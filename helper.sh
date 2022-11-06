#!/bin/bash

build_backend () {
    echo "Building new backend..."
    [[ -f bin/backend.go ]] && rm bin/backend.go
    cd back && go build -o ../bin/backend.go
}

start_frontend () {
    cd front && npm start
}


### Main ###
if [ $# -eq 0 ]
then
  build_backend
fi  

while getopts ":bf" opt; do
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