#!/bin/bash
echo -e "\033[32m GenerateMocks\033[0m"
mockery --name=".*(Agent|Repo|UseCase$)" --dir="../" -r
if [ $? -eq 0 ]; then
  echo -e "\033[32m SUCCESS GENERATE MOCKS \n\033[0m\033[0m"
else
    echo -e "\033[31m *** ERROR DUE GENERATE MOCKS ***\033[0m"
    exit 127
fi