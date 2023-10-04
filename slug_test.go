package human_readable_slug

import (
	"testing"
	"time"
)

func TestGenerateSlug(t *testing.T) {
	for i := 0; i < 10; i++ {
		slug := GenerateSlug(time.Now().Unix(), 7)
		t.Log(slug)
		if !IsHumanReadable(slug) {
			t.Error("slug", slug, "is not human readable")
		}
	}
}
