#!/bin/bash


if [ "$GOENV" != "GoBase" ]; then
    if [ -z "$GO_ORIG_PATH" ]; then
        export GO_ORIG_PATH=$PATH
    fi

    export GOENV=GoBase
    export GOPATH="$HOME/GoBase"
    export PATH="$GO_ORIG_PATH:$GOPATH/bin"
    export GO15VENDOREXPERIMENT=1
fi
