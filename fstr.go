/*
Package fstr, akin to Python's f-strings, is a Go utility that simplifies string formatting with interpolation.
It seamlessly integrates expressions into string literals for dynamic construction.
Enhancing readability and efficiency in string formatting.
This package introduces two primary functions: `Interpolate` and `Eval`.

The `Interpolate` function replaces placeholders within a string with values from a provided map. It supports simple placeholders (e.g., {key}) and formatted placeholders (e.g., {key:.2f}) for various formatting options.

Example of `Interpolate` usage:

	var template = "Hello {name}, your balance is {balance:.2f}"
	var data = map[string]interface{}{"name": "John Doe", "balance": 123.456}
	if message, err := fstr.Interpolate(template, data); err == nil {
		fmt.Println(message) // Output: "Hello John Doe, your balance is 123.46"
	}

The `Eval` function takes this concept further by allowing more complex expressions and direct invocation within print functions. It evaluates the string with embedded expressions directly and returns the formatted string.

Example of `Eval` usage:

	fmt.Println(fstr.Eval(
		"Welcome {name}, your registration is {status}.",
		map[string]interface{}{
			"name": "Jane Doe",
			"status": "complete",
		}
	)) // Output: "Welcome Jane Doe, your registration is complete."

Both functions are invaluable for generating dynamic text where the template remains consistent, but the data changes, facilitating ease of maintenance and clarity in code involving string operations.
*/
package fstr

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"text/template"
)

// Interpolate performs string interpolation on the provided format string using the given data map.
// It replaces placeholders in the format like {key} or {key:format} with corresponding values from the data map.
//
// The function supports:
//   - Simple placeholders like {key} which are replaced by the value of 'key' from the data map.
//   - Formatted placeholders like {key:.2f} or {key:,} which are replaced with the value formatted according to the specifier.
//
// The function uses Go's text/template package for template processing and supports custom formatting through the formatNumber function.
//
// Arguments:
//   - format: The format string containing placeholders.
//   - data: A map of keys and values used to replace placeholders in the format string.
//
// Returns:
//   - The interpolated string or an error if the template parsing or execution fails.
func Interpolate(format string, data map[string]interface{}) (string, error) {
	format = preprocess(format)
	t, err := template.New("fstr").Funcs(template.FuncMap{
		"formatNumber": formatNumber,
	}).Parse(format)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}
	var output bytes.Buffer
	// convert any int value inside data to float64
	for k, v := range data {
		switch v.(type) {
		case int:
			data[k] = float64(v.(int))
		default:
			data[k] = v
		}
	}
	if err := t.Execute(&output, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}
	return output.String(), nil
}

// Eval is a convenience wrapper around Interpolate. It takes a format string and a data map,
// interpolates the format string with values from the data map, and returns the result.
// If an error occurs during interpolation, Eval panics with that error.
//
// Use Eval when you want a simple way to use string interpolation without handling errors each time.
// It's especially useful in scenarios where you're sure the format string and data won't cause errors,
// or in quick scripts or applications where error handling isn't critical. For more robust applications,
// consider using Interpolate directly and handling errors appropriately.
//
// Parameters:
//   - format: A string with placeholders to be replaced by values from the data map.
//     Placeholders are in the form {key} or {key:format}.
//   - data:   A map[string]interface{} where each key corresponds to a placeholder in the format string,
//     and the associated value is what you want to replace the placeholder with.
//
// Returns:
//   - A string with all placeholders replaced by corresponding data values.
//
// Panics:
//   - If an error occurs during interpolation, Eval panics with the error returned by Interpolate.
//
// Example usage:
//
//	result := fstr.Eval("Hello, {name}!", map[string]interface{}{"name": "Alice"})
//	fmt.Println(result) // Prints: Hello, Alice!
func Eval(format string, data map[string]interface{}) string {
	result, err := Interpolate(format, data)
	if err != nil {
		panic(err)
	}
	return result
}

// Print is a convenience wrapper around Eval. It takes a format string and a data map,
// interpolates the format string with values from the data map, and prints the result to stdout.
// If an error occurs during interpolation, Print panics with that error.
func Print(format string, data map[string]interface{}) {
	fmt.Print(Eval(format, data))
}

// Println is a convenience wrapper around Eval. It takes a format string and a data map,
// interpolates the format string with values from the data map, and prints the result to stdout.
// If an error occurs during interpolation, Println panics with that error.
func Println(format string, data map[string]interface{}) {
	fmt.Println(Eval(format, data))
}

// preprocess converts placeholders in the format string into a syntax compatible with Go's text/template package.
// It identifies and converts simple placeholders (e.g., {key}) and formatted placeholders (e.g., {key:.2f}).
func preprocess(format string) string {
	re := regexp.MustCompile(`{([a-zA-Z0-9_]+)(=)?(?::(,|\.([0-9]+)f|,\.([0-9]+)f))?}`)
	return re.ReplaceAllStringFunc(format, func(m string) string {
		matches := re.FindStringSubmatch(m)
		switch {
		case matches[2] == "=":
			/*
				Handle all four cases:
				- {name=} => name={{.name}}
				- {balance=:,} => balance={{formatNumber .balance ","}}
				- {total:.3f} => total={{formatNumber .total ".3"}}
				- {balance=:,.2f} => balance={{formatNumber .balance ",.2"}}
			*/
			if matches[3] == "," {
				// example format: {balance:,} and balance is 123456789.111 => 123,456,789
				return fmt.Sprintf("%[1]s={{formatNumber .%[1]s \",\"}}", matches[1])
			} else if matches[4] != "" {
				// example format: {total:.3f} and total is 123456789.9787968 => 123,456,789.979
				return fmt.Sprintf("%[1]s={{formatNumber .%[1]s \".%s\"}}", matches[1], matches[4])
			} else if matches[5] != "" {
				// example format: {total:,.3f} and total is 123456789.9787968 => 123,456,789.979
				return fmt.Sprintf("%[1]s={{formatNumber .%[1]s \",.%s\"}}", matches[1], matches[5])
			} else {
				return fmt.Sprintf("%[1]s={{.%[1]s}}", matches[1])
			}
		case matches[3] == ",":
			// example format: {balance:,} and balance is 123456789.111 => 123,456,789
			return fmt.Sprintf("{{formatNumber .%s \",\"}}", matches[1])
		case matches[4] != "":
			// example format: {total:.3f} and total is 123456789.9787968 => 123,456,789.979
			return fmt.Sprintf("{{formatNumber .%s \".%s\"}}", matches[1], matches[4])
		case matches[5] != "":
			// example format: {total:,.3f} and total is 123456789.9787968 => 123,456,789.979
			return fmt.Sprintf("{{formatNumber .%s \",.%s\"}}", matches[1], matches[5])
		default:
			return fmt.Sprintf("{{.%s}}", matches[1])
		}
	})
}

// formatNumber is a helper function that formats a number according to the given format specifier.
// It supports formatting for thousands separators and decimal precision.
func formatNumber(value float64, format string) string {
	// Split the format string to identify thousands and decimal parts.
	formatParts := strings.Split(format, ".")
	if strings.Contains(formatParts[0], ",") && len(formatParts) == 1 {
		intPart := fmt.Sprintf("%.0f", value) // Get the integer part
		for i := len(intPart) - 3; i > 0; i -= 3 {
			intPart = intPart[:i] + "," + intPart[i:]
		}
		return intPart
	} else if strings.Contains(formatParts[0], ",") && len(formatParts) == 2 {
		// example format: {total:,.3f} and total is 123456789.9787968 => 123,456,789.979
		strNumber := fmt.Sprintf("%."+formatParts[1]+"f", value)
		parts := strings.Split(strNumber, ".")
		decimalPart := parts[1]
		intPart := parts[0]
		for i := len(intPart) - 3; i > 0; i -= 3 {
			intPart = intPart[:i] + "," + intPart[i:]
		}
		return intPart + "." + decimalPart
	} else if !strings.Contains(formatParts[0], ",") && len(formatParts) == 2 {
		// example format: {gpa:.4f} and gpa is 3.165789 => 3.1658
		return fmt.Sprintf("%."+formatParts[1]+"f", value)
	} else {
		panic("Invalid format")
	}
}
