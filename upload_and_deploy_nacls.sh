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
clean_image_tag=daisy-template
upgrade_image_tag=daisy-in-use
alias_prefix=daisy

# First get the uuids of all the instances that we are updating
ids=($($cmd search --instancefilter tag $clean_image_tag -o id))
num_ids=${#ids[@]}

# Find all the nacls
nacls=($(find nacls/* | sort -t "-" -k 2 -g))
num_nacls=${#nacls[@]}

# Check to see if there is a mismatch between number of instances and nacls
if [[ $num_ids -ne $num_nacls ]]; then
  echo There are: $num_ids instances and $num_nacls nacls, exiting
  exit 1
fi


# Loop over instances, performing upgrades for each of them, also give them an alias
for i in $(seq 0 $(( --num_nacls )) ); do
  # Push the nacl
  naclID=$($cmd push-nacl ${nacls[$i]} -o id)

  # Set the alias
  $cmd instance-alias ${ids[$i]} $alias_prefix$((i+1))

  # Perform the upgrade
  $cmd upgrade --nacl $naclID --imageTag $upgrade_image_tag ${ids[$i]}
  if [ $? -ne 0 ]; then
    echo "Command: <$cmd upgrade --nacl $naclID --imageTag $upgrade_image_tag ${ids[$i]}> failed"
  fi
done
