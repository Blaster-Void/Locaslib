📂 Locaslib

A lightweight Go library for reading and inspecting files in the working directory.

Go Reference
GitHub
🚀 Installation
bash

go get github.com/Blaster-Void/Locaslib
# intsall

```go get github.com/Blaster-Void/Locaslib```

📖 Features
1. ReadWD() ([]os.DirEntry, error)

Returns a slice of `os.DirEntry` for all files/directories in the current working directory.
```go

entries, err := Locaslib.ReadWD()
```

2. ReadWDSTR() ([]string, error)

Returns a slice of filenames (as strings) in the current directory.
```go

files, err := Locaslib.ReadWDSTR()
```
3. TypeofFile() []string

Classifies each item in the directory as |DIR: name (directory) or |File: name (file).
```go

fileTypes := Locaslib.TypeofFile()
```
📝 Example
```//‍‍‍‍‍go

package main

import (
	"fmt"
	"github.com/Blaster-Void/Locaslib"
)

func main() {
	// Get all files/dirs as []os.DirEntry
	entries, _ := Locaslib.ReadWD()
	fmt.Println("Entries:", entries)

	// Get filenames as strings
	fileNames, _ := Locaslib.ReadWDSTR()
	fmt.Println("Files:", fileNames)

	// Classify files/dirs
	classified := Locaslib.TypeofFile()
	for _, item := range classified {
		fmt.Println(item)
	}
}
```

📜 License

MIT © Blaster-Void
🔗 Links

    GitHub: github.com/Blaster-Void/Locaslib

    GoDoc: pkg.go.dev/github.com/Blaster-Void/Locaslib

🎯 Why Use Locaslib?

    Zero dependencies.

    Simple API for directory operations.

    Handy for CLI tools or file management tasks.

Feel free to contribute! 🛠️
