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
	// Standard string interpolation with formatting
	fmt.Println(fstr.Eval(
		"Hello, {name}! This is {person} {age} years old. With {balance:,.2f} USD in the bank. With GPA of {gpa:.2f}.",
		map[string]interface{}{
			"name":    "World",
			"person":  "John Doe",
			"age":     23,
			"balance": 123456789.64789,
			"gpa":     3.57869,
		},
	))
	// Output:
	// Hello, World! This is John Doe 23 years old. With 123,456,789.65 USD in the bank. With GPA of 3.58.

	// Extended syntax for key-value pairing
	fmt.Println(fstr.Eval(
		"{name=} {age=} {gpa=:,.2f} {total=:,.3f}",
		map[string]interface{}{
			"name":  "Ziad Mansour",
			"age":   23,
			"gpa":   3.1495,
			"total": 123456789.9787968,
		},
	))
	// Output:
	// name=Ziad Mansour age=23 gpa=3.15 total=123,456,789.979
}
```

## Contributing
Contributions are welcome! Feel free to submit pull requests, create issues, or provide feedback.
