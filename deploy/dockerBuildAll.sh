#!/bin/bash

cd ../env/prod/k8 || exit
for dir in *
do
  if [ -d "$dir" ] && { [ "$dir" != "redis" ] || [ "$dir" != "postgres" ]; }
  then
    cd ../../../deploy || exit
    bash dockerBuild.sh "$dir"
    cd ../env/prod/k8 || exit
  fi
done
