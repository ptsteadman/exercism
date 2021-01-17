#!/usr/bin/env bash

# learnings:
# - Special parameters:
#   $# is number of positional params
#   $@ expands the positional params
#   $? exit status of last command
#   $- current option flags
#  Word Splitting:
#   Word splitting is performed on the results of almost all unquoted expansions.
#   https://mywiki.wooledge.org/WordSplitting
#   IFS = "input field separator"
#   double quoting to avoid word splitting:
#     var="This is a variable"; args "$var"
#   word splitting also does not occur between [[ ]]

if [[ $# -ne 1 ]]; then
  echo "Usage: error_handling.sh <person>"
  exit 1
fi

echo "Hello, $1"
