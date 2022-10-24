#!/bin/bash


get_datetime() {
    # Get current date and time into variable
    datetime=$(date +%Y-%m-%d_%H-%M-%S)
}

build_backend () {
    [[ -f bin/backend.go ]] && rm bin/backend.go
    cd back && go build -o ../bin/backend.go
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
      command=$OPTARG
      if [ $command == "build" ]
      then
        build_backend
      elif [ $command == "start" ]
      then
        start_backend
      fi
      ;;
    f)
      start_frontend
      ;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      ;;
  esac
done