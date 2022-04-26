#!/bin/bash
mockery --name=".*(Agent|Repo|UseCase$)" --dir="../" -r
