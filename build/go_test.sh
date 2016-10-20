#!/bin/bash
set -e

PACKAGES=$(go list ./... | grep -v "/vendor")
go test ${PACKAGES}
go vet ${PACKAGES}
