#!/usr/bin/env bash

# learned:
# arithmetic operations like incrementing also need to be inside (( )) 
# positional parameters inside ${} don't need $ prefix

if (($# != 2)); then
  echo "Usage: hamming.sh <string1> <string2>"
  exit 1
fi

string1_len=${#1}
string2_len=${#2}

if ((string1_len != string2_len)); then
  echo "left and right strands must be of equal length"
  exit 1
fi

distance=0

for (( i=0; i<string1_len; i++ )) 
do
  if [[ "${1:i:1}" != "${2:i:1}" ]]; then
    ((distance++))
  fi
done

echo $distance

