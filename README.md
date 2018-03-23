Package radar
==================
[![GoDoc](https://godoc.org/github.com/thedevsaddam/radar?status.svg)](https://godoc.org/github.com/thedevsaddam/radar)
[![License](https://img.shields.io/dub/l/vibe-d.svg)](https://github.com/thedevsaddam/radar/blob/dev/LICENSE.md)

Package radar help to debug nested function call and Trace current file/line

### Installation

Install the package using
```go
$ go get github.com/thedevsaddam/radar/...
```

### Usage

To use the package import it in your `*.go` code
```go
import "github.com/thedevsaddam/radar"
```
### Example

##### Beam

```go
package main

import "github.com/thedevsaddam/radar"

func main() {
	foo()
}

func foo() {
	baz()
}

func baz() {
	radar.Beam()
}
```

**Output**
```bash
➜ go run main.go
 -> /Users/thedevsaddam/Workspace/go/code/radar/main.go: main.baz(14)
```

##### Trace

```go
package main

import "github.com/thedevsaddam/radar"

func main() {
	foo()
}

func foo() {
	baz()
}

func baz() {
	radar.Trace()
}
```

**Output**
```bash
➜ go run main.go
 -> 1: /Users/thedevsaddam/Workspace/go/code/radar/main.go: main.baz(14)
 -> 2: /Users/thedevsaddam/Workspace/go/code/radar/main.go: main.foo(10)
 -> 3: /Users/thedevsaddam/Workspace/go/code/radar/main.go: main.main(6)
```

### **Contribution**
If you are interested to make the package better please send pull requests or create an issue so that others can fix. Read the [contribution guide here](CONTRIBUTING.md).

### **License**
The **radar** is an open-source software licensed under the [MIT License](LICENSE.md).
