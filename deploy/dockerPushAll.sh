#!/bin/bash

cd ../env/prod/k8 || exit
for dir in *
do
  if [ -d "$dir" ] && [ "$dir" != "redis" ] && [ "$dir" != "utils" ] && [ "$dir" != "secrets" ]
  then
    cd ../../../deploy/ || exit
    bash dockerPush.sh "$dir"
    cd ../env/prod/k8 || exit
  fi
done

cd ../../.. || exit
