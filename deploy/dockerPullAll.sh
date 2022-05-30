#!/bin/bash

bash dockerPull.sh
if [ $? -eq 0 ]; then
  echo -e "\033[32m *** SUCCESS PULLED ALL  ***\033[0m"
else
  echo -e "\033[31m *** ERROR DUE PULLING ALL ***\033[0m"
  exit 127
fi