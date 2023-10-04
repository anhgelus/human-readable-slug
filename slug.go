package main

import (
	"math/rand"
)

const minLetters = "abcdefghijklmnopqrstuvwxyz"
const numbers = "0123456789"

var DefaultOptions = CustomSlug{
	CanBe2ConsecutiveLetters: false,
	HasNumbers:               true,
	HasCapitalLetters:        true,
}

var letters []rune
var badCombination [][2]rune

func init() {
	letters = DefaultOptions.generateLetters()
	DefaultOptions.cLetters = letters
	badCombination = [][2]rune{
		{
			'I', 'l',
		},
		{
			'O', '0',
		},
		{
			'm', 'n',
		},
	}
}

func AddBadCombination(a rune, b rune) {
	badCombination = append(badCombination, [2]rune{a, b})
}

// GenerateSlug returns a random slug with the size specified
func GenerateSlug(random int64, size uint) string {
	return generateSlug(random, size, &DefaultOptions)
}

func generateSlug(random int64, size uint, options *CustomSlug) string {
	rand.NewSource(random)
	var slug string
	last := 'ø'
	for i := uint(0); i < size; i++ {
		l := randomLetter(options.cLetters)
		// handle bad readable cases
		if last != 'ø' {
			for isBadCombination(l, last, options) {
				l = randomLetter(options.cLetters)
			}
		}
		slug += string(l)
		last = l
	}
	return slug
}

// randomLetter returns a random letter
func randomLetter(letters []rune) rune {
	n := rand.Intn(len(letters))
	return letters[n]
}

// isBadCombination returns true if a and b are a bad combination of letter
func isBadCombination(a rune, b rune, options *CustomSlug) bool {
	if !options.CanBe2ConsecutiveLetters {
		if a == b {
			return true
		}
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
	return isHumanReadable(slug, &DefaultOptions)
}

func isHumanReadable(slug string, options *CustomSlug) bool {
	last := 'ø'
	for _, c := range []rune(slug) {
		if last != 'ø' {
			if c == last {
				return false
			}
			if isBadCombination(last, c, options) {
				return false
			}
		}
		last = c
	}
	return true
}
