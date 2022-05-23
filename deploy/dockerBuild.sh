#!/bin/bash
source spin.sh

dockerProfile="mausved"
prefix="wave_"

cd ..

echo -e "\033[32m *** BUILDING $dockerProfile/$prefix$1 ***\033[36m"
docker build -t "$dockerProfile/$prefix$1" -f env/prod/k8/$1/Dockerfile .


#spin

if [ $? -eq 0 ]; then
  echo -e "\033[32m *** SUCCESS BUILD $dockerProfile/$prefix$1 ***\033[0m"
else
  echo -e "\033[31m *** ERROR DUE BUILDING $dockerProfile/$prefix$1 ***\033[0m"
  exit 127
fi
