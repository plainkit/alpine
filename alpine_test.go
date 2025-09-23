package alpine

import (
	"strings"
	"testing"
)

func TestJavaScript(t *testing.T) {
	js := JavaScript()

	if len(js) == 0 {
		t.Fatal("JavaScript() returned empty content")
	}

	jsString := string(js)

	// Check for Alpine-specific content
	if !strings.Contains(jsString, "Alpine") {
		t.Error("JavaScript content does not appear to contain Alpine.js library")
	}

	// Should be minified (no excessive whitespace)
	lines := strings.Split(jsString, "\n")
	if len(lines) > 10 {
		t.Error("JavaScript appears to not be minified (too many lines)")
	}
}

func TestJavaScriptSize(t *testing.T) {
	js := JavaScript()

	// Alpine.js minified is typically around 30-40KB
	if len(js) < 25000 {
		t.Error("JavaScript content seems too small to be complete Alpine.js library")
	}

	if len(js) > 60000 {
		t.Error("JavaScript content seems too large for minified Alpine.js library")
	}
}
