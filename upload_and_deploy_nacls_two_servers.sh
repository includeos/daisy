#!/bin/bash
# Script for uploading all the nacls created and deploying them to the starbases
set -x

# Mothership cli settings
bin=/Users/martin/go/src/github.com/includeos/mothership/mothership
username=
password=
host=
port=
notls=""
cmd="$bin --username $username --password $password --host $host --port $port $notls"

# Other settings
clean_image_tag=daisy-template
clean_image_tag_server1=daisy-template-server1
clean_image_tag_server2=daisy-template-server2
upgrade_image_tag=daisy-in-use
alias_prefix=daisy

# First get the uuids of all the instances that we are updating
ids_server1=($($cmd search --instancefilter tag $clean_image_tag_server1 -o id))
ids_server2=($($cmd search --instancefilter tag $clean_image_tag_server2 -o id))
num_ids_server1=${#ids_server1[@]}
num_ids_server2=${#ids_server2[@]}

# Find all the nacls
nacls=($(find nacls/* | sort -t "-" -k 2 -g))
num_nacls=${#nacls[@]}

# Check to see if there is a mismatch between number of instances and nacls
#if [[ $num_ids -ne $num_nacls ]]; then
#  echo There are: $num_ids instances and $num_nacls nacls, exiting
#  exit 1
#fi


# Loop over instances, performing upgrades for each of them, also give them an alias
for i in $(seq 0 $(( --num_nacls )) ); do
  server_int=$((i/2))
  if [ $((i%2)) -eq 0 ]; then
    id=${ids_server1[$server_int]}
  else
    id=${ids_server2[$server_int]}
  fi

  # Push the nacl
  naclID=$($cmd push-nacl ${nacls[$i]} -o id)

  # Set the alias
  $cmd instance-alias $id $alias_prefix$((i+1))

  # Perform the upgrade
  $cmd upgrade --nacl $naclID --imageTag $upgrade_image_tag $id
  if [ $? -ne 0 ]; then
    echo "Command: <$cmd upgrade --nacl $naclID --imageTag $upgrade_image_tag $id> failed"
  fi
done
