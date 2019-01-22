#!/bin/bash

export GOBIN="$(pwd)/bin"

cd cmd/apcupsd_exporter
go install
cd ../..

