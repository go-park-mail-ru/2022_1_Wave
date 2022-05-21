#!/bin/bash

cd internal/test || return
bash mocks.sh
cd ../..

go list ./...
CVPKG=$(go list ./... | grep -Ev 'mocks|*.(P|p)roto' | tr '\n' ',')
echo -e "\033[32m go tests\033[0m"
go test -coverpkg "$CVPKG" -coverprofile cover.out ./...
go tool cover -func cover.out | grep total
if [ $? -eq 0 ]; then
  echo -e "\033[32m success\033[0m"
else
  echo -e "\033[31m success\033[0m\033[0m"
  exit 127

#go tool cover -html=cover.out