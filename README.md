# fstr

![DALL·E 2023-12-30 13 16 30 - Create a logo with the text 'fstr' in all lowercase letters, featuring a sleek, modern font on a solid black background  The design should be minimali](https://github.com/ZiadMansourM/fstr/assets/64917739/10fbd823-a649-4657-865f-1eb6fca9781b)

A simple Go package for string interpolation, inspired by Python's f-strings.

## Overview

The `fstr` package provides a straightforward way to interpolate variables into strings. It's designed to be an intuitive tool for Go developers looking to embed dynamic expressions inside string literals.

## Features

- Easy to use with a simple API.
- Supports dynamic string interpolation similar to Python's f-strings.

## Installation

To install `fstr`, simply run:

```bash
go get github.com/ZiadMansourM/fstr
```

## Usage
```Go
package main

import (
	"fmt"
	"github.com/ZiadMansourM/fstr"
)

func main() {
	// Data map containing various types of values.
	data := map[string]interface{}{
		"name":    "John Doe",
		"age":     30,                  // integer value
		"gpa":     3.7498,              // floating-point value
		"balance": 123456.789,          // floating-point value for testing thousands separator
		"total":   987654321.123456,    // floating-point value for testing combined thousands separator and precision
	}

	// Format string with various types of placeholders.
	format := "Hello {name}, you are {age} years old. Your GPA is {gpa:.2f}. Your balance is {balance:,}. Your total assets are {total:,.2f}."

	// Perform string interpolation using the fstr.Interpolate function.
	result, err := fstr.Interpolate(format, data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the result.
	fmt.Println(result)
	// Expected output:
	// Hello John Doe, you are 30 years old. Your GPA is 3.75. Your balance is 123,456. Your total assets are 987,654,321.12.
}
```

## Contributing
Contributions are welcome! Feel free to submit pull requests, create issues, or provide feedback.
