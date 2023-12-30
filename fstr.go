/*
Package fstr is a utility for string interpolation similar to Python's f-strings.
It allows embedding expressions inside string literals for dynamic string formatting.
*/
package fstr

import (
	"bytes"
	"fmt"
	"regexp"
	"text/template"
)

// Interpolate performs string interpolation on the provided format string using the given data map.
// It supports both simple placeholders (e.g., {key}) and formatted placeholders (e.g., {key:.2f}).
//
// Placeholders should be in the form:
//   - {key}: Replaced with the value of 'key' from the data map.
//   - {key:format}: Replaced with the formatted value according to the format specifier (like .2f for float).
//
// Usage example:
//
//	data := map[string]interface{}{"name": "John Doe", "balance": 123.456}
//	result, err := Interpolate("Hello {name}, your balance is {balance:.2f}", data)
//	// result: "Hello John Doe, your balance is 123.46"
//
// Returns the interpolated string or an error if the template parsing or execution fails.
func Interpolate(format string, data map[string]interface{}) (string, error) {
	format = preprocess(format)
	t, err := template.New("fstr").Parse(format)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}
	var output bytes.Buffer
	if err := t.Execute(&output, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}
	return output.String(), nil
}

// preprocess prepares the format string for use with the Go text/template package.
// It converts custom placeholders into a format compatible with text/template syntax.
// Specifically:
//   - Converts {key} to {{.key}}
//   - Converts {key:format} to a printf directive (e.g., {{printf "%.2f" .key}})
//
// Assumes that format is well-formed and uses curly braces exclusively for defining placeholders.
func preprocess(format string) string {
	re := regexp.MustCompile(`{([a-zA-Z0-9_]+)(?::([.#0-9]*)f)?}`)
	return re.ReplaceAllStringFunc(format, func(m string) string {
		matches := re.FindStringSubmatch(m)
		if matches[2] != "" {
			return fmt.Sprintf("{{printf \"%%%sf\" .%s}}", matches[2], matches[1])
		}
		return fmt.Sprintf("{{.%s}}", matches[1])
	})
}
