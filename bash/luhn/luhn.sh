#!/usr/bin/env bash
# remove spaces
input=${1// /}
input_len=${#input}
digit_re="^[[:digit:]]{2,}$"

if ! [[ $input =~ $digit_re ]]; then
	echo "false"
	exit 0
fi

sum=0

for ((i = 1; i <= input_len; i++)); do
	char=${input:input_len-i:1}
	if ((i % 2 != 0)); then
		((sum += char))
	else
		((doubled = char * 2))
		((sum += doubled >= 10 ? doubled - 9 : doubled))
	fi
done

if ((sum % 10 == 0)); then
	echo "true"
else
	echo "false"
fi
