package js

import _ "embed"

// AlpineMinJS contains the minified Alpine.js JavaScript library (v3.x).
// This can be served directly or embedded in HTML pages.
//
//go:embed alpine.min.js
var AlpineMinJS []byte