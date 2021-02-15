#!/usr/bin/env bash

if [[ "$1" == "total" ]]; then 
  echo $(bc <<< "2^64 - 1")
  exit 0
fi

if (($1>64)) || (($1<1)); then 
  echo "Error: invalid input"
  exit 1
fi

echo $(bc <<< "2^($1-1)")
