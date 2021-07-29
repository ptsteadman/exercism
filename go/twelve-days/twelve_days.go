package twelve

import (
	"fmt"
	"strings"
)

var ordinals = []string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}
var gifts = []string{
	"a Partridge in a Pear Tree",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

// Verse returns the nth verse of the Twelve Days of Christmas song, 1-indexed by day
func Verse(day int) string {
	var allGifts string
	// count down from current day
	for d := day; d >= 1; d-- {
		gift := gifts[d-1]
		if d == 1 {
			if day != 1 {
				gift = "and " + gift
			}
		} else {
			gift += ", "
		}
		allGifts += gift
	}

	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s.", ordinals[day-1], allGifts)
}

// Song returns the entire Twelve Days of Christmas song, with verses separated by newlines
func Song() string {
	verses := []string{}
	for i := 1; i <= 12; i++ {
		verses = append(verses, Verse(i))
	}
	return strings.Join(verses, "\n")
}
