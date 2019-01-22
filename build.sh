#!/bin/bash

cd cmd/apcupsd_exporter
go install
cd ../..
mv "$GOPATH/bin/apcupsd_exporter"  "bin/apcupsd_exporter"

