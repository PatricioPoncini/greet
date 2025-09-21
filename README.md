<div align="center">

# `Greet 👋`
<img src="https://i.imgur.com/riT59a4.png" alt="Greet Logo" width="200"/>

**Greet** is a simple Go package to demonstrate how to create and use structs with methods.  
This is my first Go package.

</div>

## `Features`
- Defines a `User` struct with a `Name` field.
- Provides a `SayHi` method that returns a greeting message.

## `Installation`
```bash
go get github.com/PatricioPoncini/greet@v0.0.4
```

## `Usage`
```go
package main

import (
	"fmt"

	"github.com/PatricioPoncini/greet"
)

func main() {
	user := greet.User{Name: "Patricio"}
	fmt.Println(user.SayHi()) // Output: "Hi, Patricio!"
}
```

## `Notes`
- This package is intended for learning and demonstration purposes.
- All exported types and functions start with uppercase letters, so they can be used in other projects.