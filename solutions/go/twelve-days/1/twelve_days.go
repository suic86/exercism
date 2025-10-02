package twelve

import (
	"fmt"
	"strings"
)

// Verse returns nth verse of the 'The Twelve Days of Christmas'.
func Verse(verse int) string {
	var line strings.Builder
	for i := verse - 1; i >= 0; i-- {
		line.WriteString(giftsAndDays[i])
	}
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s", nthDay[verse-1], line.String())
}

// Song return the lyrics of 'The Twelve Days of Christmas'.
func Song() string {
	song := []string{}
	for i := 1; i <= 12; i++ {
		song = append(song, Verse(i))
	}
	return strings.Join(song, "\n")
}

var nthDay []string = []string{
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

var giftsAndDays []string = []string{
	"a Partridge in a Pear Tree.",
	"two Turtle Doves, and ",
	"three French Hens, ",
	"four Calling Birds, ",
	"five Gold Rings, ",
	"six Geese-a-Laying, ",
	"seven Swans-a-Swimming, ",
	"eight Maids-a-Milking, ",
	"nine Ladies Dancing, ",
	"ten Lords-a-Leaping, ",
	"eleven Pipers Piping, ",
	"twelve Drummers Drumming, ",
}
