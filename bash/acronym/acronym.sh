#!/usr/bin/env bash


read -ra words -d '' <<<"$1"

for word in "${words[@]}"; do
  stripped=${word//[^a-zA-Z]/}
  first_char=${stripped:0:1}
  if ! [[ -z $first_char ]]; then 
    printf $(tr '[a-z]' '[A-Z' <<< $first_char)
  fi
done


