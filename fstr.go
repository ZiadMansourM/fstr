/*
Package fstr is a utility for string interpolation similar to Python's f-strings.
It allows embedding expressions inside string literals for dynamic string formatting.
*/
package fstr

import (
	"bytes"
	"strings"
	"text/template"
)

// Interpolate performs string interpolation on the format string using data from the provided map.
// The function replaces placeholders in the format like {key} with corresponding values from the data map.
// Returns the interpolated string or an error if the template parsing or execution fails.
func Interpolate(format string, data map[string]interface{}) (string, error) {
	format = preprocess(format)
	t, err := template.New("fstr").Parse(format)
	if err != nil {
		return "", err // Simplify error handling for demonstration
	}
	var output bytes.Buffer
	if err := t.Execute(&output, data); err != nil {
		return "", err
	}
	return output.String(), nil
}

// preprocess converts simple brace placeholders in the format string to Go's template syntax.
// For example, it converts "{var}" to "{{.var}}" for compatibility with text/template parsing.
// It's a simplistic implementation and assumes format is well-formed and {, } are used as placeholder delimiters.
func preprocess(format string) string {
	// This is a simplistic implementation. You might need a more robust parser
	// for complex cases or to avoid replacing inside strings or escaped braces.
	format = strings.Replace(format, "{", "{{.", -1)
	format = strings.Replace(format, "}", "}}", -1)
	return format
}
