#!/usr/bin/env bash

input_upper=${1^^}
score=0

for ((i=0; i<${#input_upper}; i++)); do
  case ${input_upper:$i:1} in
    A | E | I | O | U | L | N | R | S | T)
      ((score+=1));;
    D | G)
      ((score+=2));;
    B | C | M | P)
      ((score+=3));;
    F | H | V | W | Y)   
      ((score+=4));;
    K)   
      ((score+=5));;
    J | X)
      ((score+=8));;
    Q | Z)
      ((score+=10));;
   esac
done

echo $score
