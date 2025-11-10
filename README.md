# 🧰 Go X

A collection of reusable **Go (Golang) helper functions and utilities** designed to simplify development across multiple projects. This repository serves as a central toolkit for common patterns, operations, and reusable logic — so you can write cleaner, faster, and more maintainable Go code.

---

## 🚀 Features

* 🧩 **Utility Functions** — common helpers for string manipulation, time, error handling, etc.
* ⚙️ **Reusable Components** — build once, use everywhere.
* 📦 **Lightweight & Dependency-free** — designed to be simple and fast.
* 🧪 **Tested** — each helper includes unit tests for reliability.

---

## 📁 Project Structure

```
gox/
├── docs/                   # Documentation helper functions
├── ... 
├── pkg/
│   ├── xconfig/        # Config helper functions, load some configuration file
│   └── ...
├── ...                  
├── go.mod
├── go.sum
├── .gitignore
├── LICENSE
└── README.md
```
---

## ⚙️ Installation

To use this helper package in your Go project, simply run:

```bash
go get github.com/agussyahrilmubarok/gox
```

Then import it in your code:

```go
import "github.com/agussyahrilmubarok/gox/xstringutil"
```

Example usage:

```go
package main

import (
    "fmt"
    "github.com/agussyahrilmubarok/gox/xstringutil"
)

func main() {
    result := xstringutil.ToCamelCase("hello_world")
    fmt.Println(result) // Output: HelloWorld
}
```

---

## 🧪 Running Tests

You can run all tests with:

```bash
go test ./...
```

---

## 🧾 License

This project is licensed under the **MIT License** — see the [LICENSE](LICENSE) file for details.

---

## 📚 References

* [The Go Programming Language](https://golang.org/doc/)
* [Effective Go](https://golang.org/doc/effective_go)