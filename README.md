# printj
Tiny library to print a pretty JSON from objects you pass in.

I missed `console.log()` of JavaScript, that's why I need this.

## Install
```sh
$ go get github.com/sigit-prayoga/printj
```

### Usage
```go
package main

import "github.com/sigit-prayoga/printj"

type person struct {
	// Make sure use upper case for first char to have a public access.
	Firstname, Lastname string
	Age                 int8
}

func main() {
	john := person{"John", "Doe", 28}
	// to have a pretty json (4 indents), pass true in second args.
	printj.Print(john, true, "Print person")

	// it works with array too
	jane := person{"Jane", "Doe", 25}
	var people []person
	people = append(people, john, jane)

	printj.Print(people, true, "Print array")
}
```

**Output**
```sh
** Print person **
 {
    "Firstname": "John",
    "Lastname": "Doe",
    "Age": 28
}

** Print array **
 [
    {
        "Firstname": "John",
        "Lastname": "Doe",
        "Age": 28
    },
    {
        "Firstname": "Jane",
        "Lastname": "Doe",
        "Age": 25
    }
]
```