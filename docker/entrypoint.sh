#!/bin/bash

# This script creates the necessary files before starting Elysium-appd

# only create the priv_validator_state.json if it doesn't exist and the command is start
if [[ $1 == "start" && ! -f ${ELYSIUM_HOME}/data/priv_validator_state.json ]]
then
    mkdir -p ${ELYSIUM_HOME}/data
    cat <<EOF > ${ELYSIUM_HOME}/data/priv_validator_state.json
{
  "height": "0",
  "round": 0,
  "step": 0
}
EOF
fi

/bin/elysium-appd --home ${ELYSIUM_HOME} $@
