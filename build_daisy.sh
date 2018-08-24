#!/bin/bash

docker build -t daisy .
docker create --name day daisy
docker cp day:/go/bin/daisy daisy
docker rm day
