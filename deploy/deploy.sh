#!/bin/bash

bash dockerBuild.sh $1
bash dockerPush.sh $1
