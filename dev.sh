#!/bin/bash
ROOT="$( cd "$(dirname "$0")" ; pwd -P)"

docker run -it --rm -v $ROOT:/opt/src --name godev julitopower/golangemacsdevelopment /bin/bash
