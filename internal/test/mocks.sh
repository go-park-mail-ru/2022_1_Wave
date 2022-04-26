#!/bin/bash
echo -e "\033[32m GenerateMocks"
echo -e "\033[0m"
mockery --name=".*(Agent|Repo|UseCase$)" --dir="../" -r
echo -e "\033[32m Success\n"
echo -e "\033[0m"
