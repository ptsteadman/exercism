#!/usr/bin/env bash

# wanted to experiment with using functions as conditionals
# return 1 for false, 0 for true (bc it's like an exit code)
# why to store regex as quoted variable:
# https://unix.stackexchange.com/questions/382054/how-does-storing-the-regular-expression-in-a-shell-variable-avoid-problems-with

function is_yelling() {
  has_lowercase_re="[a-z]+"
  has_uppercase_re="[A-Z]+"
  if ! [[ $1 =~ $has_lowercase_re ]] && [[ $1 =~ $has_uppercase_re ]]; then
    return 0
  fi
  return 1
}

whitespace_re="^\s*$"
if [[ $1 =~ $whitespace_re ]]; then
  echo "Fine. Be that way!"
  exit 0
fi

question_re="\?\s*$"
if [[ $1 =~ $question_re ]]; then
  if is_yelling "$1"; then
    echo "Calm down, I know what I'm doing!" 
    exit 0
  fi
  echo "Sure."
  exit 0
fi

if is_yelling "$1"; then
  echo "Whoa, chill out!"
  exit 0 
fi

echo "Whatever."

