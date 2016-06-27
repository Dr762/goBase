#!/bin/bash


if [ "$GOENV" != "goBase" ]; then
    if [ -z "$GO_ORIG_PATH" ]; then
        export GO_ORIG_PATH=$PATH
    fi

    export GOENV=goBase
    export GOPATH="$HOME/goBase"
    export PATH="$GO_ORIG_PATH:$GOPATH/bin"
    export GO15VENDOREXPERIMENT=1
fi
