package human_readable_slug

import "strings"

type CustomSlug struct {
	// Do you want "ABC..."?
	HasCapitalLetters bool
	// Do you want "0123..."
	HasNumbers bool
	// Do you want "aa" or "DD"...?
	CanBe2ConsecutiveLetters bool
	cLetters                 []rune
}

// GenerateSlug returns a random slug with the size specified
func (options *CustomSlug) GenerateSlug(random int64, size uint) string {
	options.cLetters = options.generateLetters()
	return generateSlug(random, size, options)
}

// IsHumanReadable checks if a slug is human readable
func (options *CustomSlug) IsHumanReadable(slug string) bool {
	return isHumanReadable(slug, options)
}

func (options *CustomSlug) generateLetters() []rune {
	raw := minLetters
	if DefaultOptions.HasNumbers {
		raw += numbers
	}
	if DefaultOptions.HasCapitalLetters {
		raw += strings.ToUpper(minLetters)
	}
	return []rune(raw)
}
