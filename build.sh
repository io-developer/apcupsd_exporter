#!/bin/bash

export GOBIN="$(pwd)/bin"
export CGO_ENABLED=0

cd cmd/apcupsd_exporter
go build -tags netgo -a -v
go install
cd ../..

