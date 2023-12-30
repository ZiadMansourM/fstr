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
```
package main

import (
    "fmt"
    "github.com/ZiadMansourM/fstr"
)

func main() {
    data := map[string]interface{}{
        "name": "John Doe",
        "age":  30,
    }
    result, err := fstr.Interpolate("Hello {name}, you are {age} years old.", data)
    if err != nil {
        // Handle error
    }
    fmt.Println(result)
}
```

## Contributing
Contributions are welcome! Feel free to submit pull requests, create issues, or provide feedback.
