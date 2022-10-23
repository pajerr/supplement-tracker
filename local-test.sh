#!/bin/bash

get_datetime() {
    # Get current date and time into variable
    datetime=$(date +%Y-%m-%d_%H-%M-%S)
}

build_backend () {
    get_datetime
    cd back && go build -o ../bin/backend-$datetime.go
}

start_frontend () {
    cd front && npm start
}

start_backend () {
    go 
}

### Main ###
#if no argumetns given print help
if [ $# -eq 0 ]
then
  build_backend
fi  

while getopts ":b:f" opt; do
  case $opt in
    b)
      build_backend
      ;;
    f)
      start_frontend
      ;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      ;;
  esac
done