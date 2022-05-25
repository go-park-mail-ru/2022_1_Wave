#!/bin/bash

sudo find . -regex ".*\(grpc\)*\.pb.*" -exec rm {} \;