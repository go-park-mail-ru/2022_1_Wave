#!/bin/bash
source spin.sh

dockerProfile="mausved"
prefix="wave_"

echo -e "\033[32m *** PUSHING $dockerProfile/$prefix$1 ***\033[36m"
docker push -q "$dockerProfile/$prefix$1" &


#spin

if [ $? -eq 0 ]
then
  echo -e "\033[32m *** SUCCESS PUSHED mausved/wave_$1 ***"
else
  echo -e "\033[31m *** ERROR DUE PUSHING mausved/wave_$1 ***"
fi

echo -e "\033[0m"
