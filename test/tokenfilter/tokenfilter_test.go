package tokenfilter

import (
	"testing"
	"time"

	"github.com/furyaxyz/elysium-app/app"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	ibctesting "github.com/cosmos/ibc-go/v6/testing"
	"github.com/stretchr/testify/suite"
)

type TokenFilterTestSuite struct {
	suite.Suite

	coordinator *ibctesting.Coordinator

	// Elysium app including the tokenfilter middleware
	elysiumChain *ibctesting.TestChain

	// Default IBC Simapp
	otherChain *ibctesting.TestChain
}

func (suite *TokenFilterTestSuite) SetupTest() {
	chains := make(map[string]*ibctesting.TestChain)
	suite.coordinator = &ibctesting.Coordinator{
		T:           suite.T(),
		CurrentTime: time.Now(),
		Chains:      chains,
	}
	suite.elysiumChain = NewTestChain(suite.T(), suite.coordinator, ibctesting.GetChainID(1))
	suite.otherChain = ibctesting.NewTestChain(suite.T(), suite.coordinator, ibctesting.GetChainID(2))

	suite.coordinator.Chains[ibctesting.GetChainID(1)] = suite.elysiumChain
	suite.coordinator.Chains[ibctesting.GetChainID(2)] = suite.otherChain
}

func NewTransferPath(elysiumChain, otherChain *ibctesting.TestChain) *ibctesting.Path {
	path := ibctesting.NewPath(elysiumChain, otherChain)
	path.EndpointA.ChannelConfig.PortID = ibctesting.TransferPort
	path.EndpointB.ChannelConfig.PortID = ibctesting.TransferPort
	path.EndpointA.ChannelConfig.Version = types.Version
	path.EndpointB.ChannelConfig.Version = types.Version

	return path
}

// TestHandleOutboundTransfer asserts that native tokens on a elysium based chain can be transferred to
// another chain and can then return to the original elysium chain
func (suite *TokenFilterTestSuite) TestHandleOutboundTransfer() {
	// setup between elysiumChain and otherChain
	path := NewTransferPath(suite.elysiumChain, suite.otherChain)
	suite.coordinator.Setup(path)

	elysiumApp := suite.elysiumChain.App.(*app.App)
	originalBalance := elysiumApp.BankKeeper.GetBalance(suite.elysiumChain.GetContext(), suite.elysiumChain.SenderAccount.GetAddress(), sdk.DefaultBondDenom)
	// take half of the original balance
	amount := originalBalance.Amount.QuoRaw(2)
	timeoutHeight := clienttypes.NewHeight(1, 110)
	coinToSendToB := sdk.NewCoin(sdk.DefaultBondDenom, amount)

	// send half the users balance from elysiumChain to otherChain
	msg := types.NewMsgTransfer(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, coinToSendToB, suite.elysiumChain.SenderAccount.GetAddress().String(), suite.otherChain.SenderAccount.GetAddress().String(), timeoutHeight, 0, "")
	res, err := suite.elysiumChain.SendMsgs(msg)
	suite.Require().NoError(err) // message committed

	packet, err := ibctesting.ParsePacketFromEvents(res.GetEvents())
	suite.Require().NoError(err)

	// relay send
	err = path.RelayPacket(packet)
	suite.Require().NoError(err) // relay committed

	// check that the token exists on chain B
	voucherDenomTrace := types.ParseDenomTrace(types.GetPrefixedDenom(packet.GetDestPort(), packet.GetDestChannel(), sdk.DefaultBondDenom))
	balance := suite.otherChain.GetSimApp().BankKeeper.GetBalance(suite.otherChain.GetContext(), suite.otherChain.SenderAccount.GetAddress(), voucherDenomTrace.IBCDenom())
	coinSentFromAToB := types.GetTransferCoin(path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, sdk.DefaultBondDenom, amount)
	suite.Require().Equal(coinSentFromAToB, balance)

	// check that the account on elysiumChain has "amount" less tokens than before
	intermediateBalance := elysiumApp.BankKeeper.GetBalance(suite.elysiumChain.GetContext(), suite.elysiumChain.SenderAccount.GetAddress(), sdk.DefaultBondDenom)
	want := originalBalance.Amount.Sub(coinToSendToB.Amount)
	suite.Require().Equal(want, intermediateBalance.Amount)

	// Send the native elysiumChain token on otherChain back to elysiumChain
	msg = types.NewMsgTransfer(path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, coinSentFromAToB, suite.otherChain.SenderAccount.GetAddress().String(), suite.elysiumChain.SenderAccount.GetAddress().String(), timeoutHeight, 0, "")
	res, err = suite.otherChain.SendMsgs(msg)
	suite.Require().NoError(err) // message committed

	packet, err = ibctesting.ParsePacketFromEvents(res.GetEvents())
	suite.Require().NoError(err)

	err = path.RelayPacket(packet)
	suite.Require().NoError(err) // relay committed

	// check that the token was sent back i.e. the new balance is equal to the original balance
	newBalance := elysiumApp.BankKeeper.GetBalance(suite.elysiumChain.GetContext(), suite.elysiumChain.SenderAccount.GetAddress(), sdk.DefaultBondDenom)
	suite.Require().Equal(originalBalance, newBalance)
}

// TestHandleInboundTransfer asserts that inbound transfers to a elysium chain are rejected when they do not contain
// the elysium native token
func (suite *TokenFilterTestSuite) TestHandleInboundTransfer() {
	// setup between elysiumChain and otherChain
	path := NewTransferPath(suite.elysiumChain, suite.otherChain)
	suite.coordinator.Setup(path)

	amount, ok := sdk.NewIntFromString("1000")
	suite.Require().True(ok)
	timeoutHeight := clienttypes.NewHeight(1, 110)
	coinToSendToA := sdk.NewCoin(sdk.DefaultBondDenom, amount)

	// send from otherChain to elysiumChain
	msg := types.NewMsgTransfer(path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, coinToSendToA, suite.otherChain.SenderAccount.GetAddress().String(), suite.elysiumChain.SenderAccount.GetAddress().String(), timeoutHeight, 0, "")
	res, err := suite.otherChain.SendMsgs(msg)
	suite.Require().NoError(err) // message committed

	packet, err := ibctesting.ParsePacketFromEvents(res.GetEvents())
	suite.Require().NoError(err)

	// relay send
	err = path.RelayPacket(packet)
	suite.Require().NoError(err) // relay committed

	// check that the token does not exist on chain A (was rejected)
	voucherDenomTrace := types.ParseDenomTrace(types.GetPrefixedDenom(packet.GetDestPort(), packet.GetDestChannel(), sdk.DefaultBondDenom))
	balance := suite.otherChain.GetSimApp().BankKeeper.GetBalance(suite.otherChain.GetContext(), suite.otherChain.SenderAccount.GetAddress(), voucherDenomTrace.IBCDenom())
	emptyCoin := sdk.NewInt64Coin(voucherDenomTrace.IBCDenom(), 0)
	suite.Require().Equal(emptyCoin, balance)
}

func TestTokenFilterTestSuite(t *testing.T) {
	suite.Run(t, new(TokenFilterTestSuite))
}
