#!/bin/bash

bash dockerBuildAll.sh
if [ $? -eq 0 ]; then
  echo -e "\033[32m *** SUCCESS BUILD ALL  ***\033[0m"
else
  echo -e "\033[31m *** ERROR DUE BUILDING ALL ***\033[0m"
  exit 127
fi
bash dockerPushAll.sh
if [ $? -eq 0 ]; then
  echo -e "\033[32m *** SUCCESS BUILD ALL ***\033[0m"
else
  echo -e "\033[31m *** ERROR DUE BUILDING ALL ***\033[0m"
  exit 127
fi