package main

import (
	"os"

	"github.com/furyaxyz/elysium-app/app"
	"github.com/furyaxyz/elysium-app/cmd/elysium-appd/cmd"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, cmd.EnvPrefix, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
