#!/usr/bin/env bash

# learned:
# echo -n echos without a newline, but is not consistent across shells
# printf has no newline by default

if [[ $# -ne 1 ]] || ! [[ $1 =~ ^[0-9]+$ ]]; then
  echo 'Usage: raindrops <int>'
  exit 1
fi

has_factor=false

if (($1 % 3 == 0)); then
  printf 'Pling'
  has_factor=true
fi
if (($1 % 5 == 0)); then
  printf 'Plang'
  has_factor=true
fi
if (($1 % 7 == 0)); then
  printf 'Plong'
  has_factor=true
fi

$has_factor || echo "$1"
