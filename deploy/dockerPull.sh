#!/bin/bash
source spin.sh

dockerProfile="mausved"
prefix="wave_"

cd ..

echo -e "\033[32m *** PULLING $dockerProfile/$prefix$1 ***\033[36m"
docker pull "$dockerProfile/$prefix$1"


#spin

if [ $? -eq 0 ]; then
  echo -e "\033[32m *** SUCCESS PULLED $dockerProfile/$prefix$1 ***\033[0m"
else
  echo -e "\033[31m *** ERROR DUE PULLING $dockerProfile/$prefix$1 ***\033[0m"
  exit 127
fi
