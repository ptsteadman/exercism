#!/usr/bin/env bash

reversed=''
original="$1" # needs to be quoted to prevent globbing

# for loop would be faster, could avoid modifying the original string
# could also just echo the characters as we go?
while [ ${#original} -gt 0 ]; do
  reversed=$reversed${original: -1}
  original=${original: : -1}
done

echo "$reversed" # needs to be quoted

