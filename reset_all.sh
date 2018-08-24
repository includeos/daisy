#!/bin/bash
# Script for uploading all the nacls created and deploying them to the starbases

# Mothership cli settings
bin=/Users/martin/go/src/github.com/includeos/mothership/mothership
username=
password=
host=
port=
notls=""
cmd="$bin --username $username --password $password --host $host --port $port $notls"

# Other settings
image_to_clean=daisy-in-use
image_to_push=image-name

ids=($($cmd search --instancefilter tag $image_to_clean -o id))
num_ids=${#ids[@]}

for i in $(seq 0 $(( --num_ids )) ); do
  # reset all alias
  $cmd instance-alias ${ids[$i]} ""
  # Redeploy a base image
  $cmd deploy ${ids[$i]} $image_to_push
done
