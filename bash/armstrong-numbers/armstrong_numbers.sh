#!/usr/bin/env bash

# just going for something terse here

sum=0
for (( i = 0; i < ${#1}; i++ )); do
  ((sum+=${1:$i:1}**${#1}))
done

(($sum == $1)) && echo "true" || echo "false"
