package human_readable_slug

import (
	"math/rand"
	"strings"
)

const minLetters = "abcdefghijklmnopqrstuvwxyz"
const numbers = "0123456789"

var letters string
var badCombination [][2]string

func init() {
	letters = minLetters + strings.ToUpper(minLetters) + numbers
	badCombination = [][2]string{
		{
			"I", "l",
		},
		{
			"O", "0",
		},
		{
			"m", "n",
		},
	}
}

// GenerateSlug returns a random slug with the size specified
func GenerateSlug(random int64, size uint) string {
	rand.NewSource(random)
	var slug string
	last := ""
	for i := uint(0); i < size; i++ {
		l := randomLetter()
		// handle bad readable cases
		if len(last) != 0 {
			for isBadCombination(l, last) {
				l = randomLetter()
			}
		}
		slug += l
		last = l
	}
	return slug
}

// randomLetter returns a random letter
func randomLetter() string {
	n := rand.Intn(len(letters))
	return string(letters[n])
}

// isBadCombination returns true if a and b are a bad combination of letter
func isBadCombination(a string, b string) bool {
	if a == b {
		return true
	}
	for _, bad := range badCombination {
		if a == bad[0] && b == bad[1] {
			return true
		}
		if a == bad[1] && b == bad[0] {
			return true
		}
	}
	return false
}

// IsHumanReadable checks if a slug is human readable
func IsHumanReadable(slug string) bool {
	last := ""
	for _, tmp := range slug {
		c := string(tmp)
		if last != "" {
			if c == last {
				return false
			}
			if isBadCombination(last, c) {
				return false
			}
		}
		last = c
	}
	return true
}
