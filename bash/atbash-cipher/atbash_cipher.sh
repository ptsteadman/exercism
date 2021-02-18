#!/usr/bin/env bash

print_opposite_lowercase_char() {
	# get ascii code point for letter
	printf -v s '%d' "'$1"
	# if it's in the range of a-z, get the opposite
	if ((s >= 97)) && ((s <= 122)); then
		((opp = 97 + (122 - s)))
		# convert to hex
		printf -v opp '%x' "$opp"
		# convert backslashed hex notation back to string
		printf "%b" "\x$opp"
	else
		# else, print out the original character
		printf "%s" "$1"
	fi
}

input_lower=${2,,}
# note to self: need DOUBLE SLASH for global replacement
input_stripped=${input_lower//[^[:alnum:]]/}
input_len=${#input_stripped}
for ((i = 0; i < input_len; i++)); do
	print_opposite_lowercase_char "${input_stripped:i:1}"
	if [[ "$1" == "encode" ]] && ((i + 1 != input_len)) && (((i + 1) % 5 == 0)); then
    # if encoding, at the obfuscating spaces
		printf "%s" " "
	fi
done
