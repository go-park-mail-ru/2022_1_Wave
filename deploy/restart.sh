#!/bin/bash

bash deploy.sh $1 && kubectl rollout restart deployment $1
