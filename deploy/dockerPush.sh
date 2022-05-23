#!/bin/bash
source spin.sh

dockerProfile="mausved"
prefix="wave_"

echo -e "\033[32m *** PUSHING $dockerProfile/$prefix$1 ***\033[36m"
docker push "$dockerProfile/$prefix$1"


#spin

if [ $? -eq 0 ]
then
  echo -e "\033[32m *** SUCCESS PUSHED $dockerProfile/$prefix$1 ***"
else
  echo -e "\033[31m *** ERROR DUE PUSHING $dockerProfile/$prefix$1 ***"
  echo -e "\033[0m"
  exit 127
fi

echo -e "\033[0m"
