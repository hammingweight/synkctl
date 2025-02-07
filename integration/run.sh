#!/usr/bin/env bash
export TEST_USER=$(grep user $HOME/.synk/config | cut -d: -f2 | tr -d ' ')
export TEST_PASSWORD=$(grep password $HOME/.synk/config | cut -d: -f2 | tr -d ' ')
export TEST_INVERTER_SN=$(grep default_inverter_sn $HOME/.synk/config | cut -d: -f2 | tr -d ' ')
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
go test -v -count=1 -tags integration $SCRIPT_DIR
