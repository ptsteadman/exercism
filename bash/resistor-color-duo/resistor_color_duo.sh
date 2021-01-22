#!/usr/bin/env bash

band_color_to_number () {
  case $1 in
    'black')
      number=0;;
    'brown')
      number=1;;
    'red')
      number=2;;
    'orange')
      number=3;;
    'yellow')
      number=4;;
    'green')
      number=5;;
    'blue')
      number=6;;
    'violet')
      number=7;;
    'grey')
      number=8;;
    'white')
      number=9;;
    *)
    esac
    echo "$number"
}

num1="$(band_color_to_number $1)"
num2="$(band_color_to_number $2)"

if [[ -z $num1 ]] || [[ -z $num2 ]]; then 
  echo 'invalid color'
  exit 1
else
  echo "$num1$num2"
fi
