#!/usr/bin/env bash

# learnings:
# IFS can be more than one character
# IFS prepended to command applies only to that command
# <<< redirection operator is called a "here string" does not perform filename
# expansion or word splitting
# read command is used to retrieve data from stdin aka file descriptor 0 
# read foo reads data from stdin to variable foo
# read -a flag reads lines into an array
# read -d flag specifies the "line" delimiter
# for in array loop syntax

IFS='- ' read -ra words -d '' <<< "$1"

for word in "${words[@]}"; do
  stripped=${word//[^a-zA-Z]/}
  first_char=${stripped:0:1}
  printf "%s" $(tr '[a-z]' '[A-Z]' <<< "$first_char")
done


