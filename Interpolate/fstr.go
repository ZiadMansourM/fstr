package fstr

import (
	"bytes"
	"strings"
	"text/template"
)

// Interpolate replaces placeholders with corresponding data in the format string.
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

// preprocess converts {var} to {{.var}} for Go template compatibility
func preprocess(format string) string {
	// This is a simplistic implementation. You might need a more robust parser
	// for complex cases or to avoid replacing inside strings or escaped braces.
	format = strings.Replace(format, "{", "{{.", -1)
	format = strings.Replace(format, "}", "}}", -1)
	return format

}
