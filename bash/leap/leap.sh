#!/usr/bin/env bash

usage_string="Usage: leap.sh <year>"

if [[ $# -ne 1 ]]; then
  echo $usage_string
  exit 1
elif ! [[ $1 =~ ^[0-9]+$ ]]; then 
  echo $usage_string
  exit 1
elif [[ $(($1 % 4)) -eq 0 ]]; then
  if [[ $(($1 % 100)) -eq 0 ]]; then
    if [[ $(($1 % 400)) -eq 0 ]]; then
      echo true
    else
      echo false
    fi
  else 
    echo true
  fi
else 
  echo false
fi

