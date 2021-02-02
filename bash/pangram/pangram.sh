#!/usr/bin/env bash

# learned:
# generate range with curly braces
# string contains can be done with equality operator and *"$var"* or regex
# to lower case ${var,,} bash 4+
# to upper case ${var^^} bash 4+

input_lower=${1,,}

for l in {a..z}; do
  if [[ "$input_lower" != *"$l"* ]]; then
    echo "false"
    exit 0
  fi
done

echo "true"
