# Go Modules & Basics

## Core Concepts
- **Module**: A collection of related Go packages versioned together as a single unit
- **go.mod**: Tracks modules and their dependencies
- **go.sum**: Contains cryptographic checksums for dependency integrity
- **Semantic Versioning**: MAJOR.MINOR.PATCH format

## Essential Commands

### Initialization
```bash
go mod init github.com/username/repo-name
```
Creates a new module with the specified path (typically your repository URL)

### Maintenance
| Command | Purpose |
|---------|---------|
| `go mod tidy` | Add missing and remove unused dependencies |
| `go mod download` | Download all dependencies |
| `go mod verify` | Verify checksums against go.sum |
| `go mod vendor` | Create vendor directory with dependency copies |

### Analysis
| Command | Purpose |
|---------|---------|
| `go mod graph` | Display dependency graph |
| `go mod why <package>` | Explain why a package is needed |
| `go list -m all` | List all dependencies |

### Dependency Management
```bash
go get <module>@<version>           # Add/update specific version
go build -mod=vendor                # Build using vendor directory
go mod edit -replace old=./new      # Replace module path for local development
```

## go.mod File Structure
```go
module github.com/username/repo-name  // Module path
go 1.21                               // Go version

require github.com/pkg/errors v0.9.1  // Dependencies

replace old/module => new/module v1.0 // Module replacement (for local dev)

exclude bad/module v0.1.0             // Exclude specific version
```

## Configuration
- **Module Proxy**: Cache server for faster downloads (via `GOPROXY`)
- **Private Modules**: Use `GOPRIVATE` to bypass proxy for private repositories

---

## Hello World Example

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### Key Components
- **`package main`**: Defines an executable program (required for executables vs libraries)
- **`import "fmt"`**: Imports the fmt package for formatted I/O
- **`func main()`**: Entry point of the program (must be in package main; only one allowed)
- **`fmt.Println`**: Prints output with automatic newline

### Running & Building
```bash
go run main.go    # Run directly
go build          # Create executable for current OS/architecture
```

---

## Exercise 1: Creating and Using Modules

### Step 1: Create the Greetings Module
**File: `greetings/greetings.go`**
```go
package greetings

import "fmt"

// Hello returns a greeting message for the given name
func Hello(name string) string {
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}
```

**Key Points:**
- `:=` operator declares and assigns variables with type inference
- Only works inside functions (not at package level)
- Similar to TypeScript's `let` keyword

### Step 2: Call from Another Module
**File: `hello/main.go`**
```go
package main

import (
    "fmt"
    "example.com/greetings"
)

func main() {
    message := greetings.Hello("Vighnesh")
    fmt.Println(message)
}
```

### Step 3: Configure Local Module Path
```bash
# In hello directory
go mod init example.com/hello
go mod edit -replace example.com/greetings=../greetings
go mod tidy
```

**Output:**
```
go: found example.com/greetings in example.com/greetings v0.0.0-00010101000000-000000000000
```

### Step 4: Run
```bash
go run .
# Output: Hi, Vighnesh. Welcome!
```

**What Happened:**
1. `go mod edit -replace` configured local path for development
2. `go mod tidy` resolved and downloaded dependencies
3. The replace directive allows using local modules before publishing


#### Adding error handling to Greetings Module

[[File: 01/Exercise1/greetings/greetings.go]]
```go
package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting message for the named person.

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty Name")
	}
	message := fmt.Sprintf("Hi. %v Wellcome", name)
	// nil - no error if there is no error nil will be empty if there is an error it will contain the error message
	return message, nil
}
```
- we imported the errors package to create error messages.
- The Hello function now returns two values: a string (the greeting message) and an error.
- If the name is empty, it returns an error using errors.New
- If the name is valid, it returns the greeting message and nil (indicating no error).
- nil is used in Go to represent the absence of a value, similar to null in other languages.
- so if nil is returned it means there is no error if there is an error the error variable will contain the error message


##### adding logging in hello file 
[[File: 01/main.go]]
```go
package main

// should it be main or hello ?
// It should be main if we want to run this file directly
// It should be hello if we want to import this file as a package in another file

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("Greetings : ")

	message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
```
- We imported the log package to handle logging.
- We set a prefix for log messages using log.SetPrefix
- We called the Hello function from the greetings package and handled the returned error.
- If there is an error, we log it using log.Fatal, which prints the error message and exits the program.
- If there is no error, we print the greeting message to the console


##### Adding Randomized Greeting Formats

[[File: 01/Exercise1/greetings/greetings.go]]
```go
package greetings
import (
    "errors"
    "fmt"
    "math/rand"
)
// Hello returns a greeting message for the named person.
func Hello(name string) (string, error) {
    if name == "" {
        return "", errors.New("Empty Name")
    }
    format, err := randomFormat()
    if err != nil {
        return "", err
    }
    message := fmt.Sprintf(format, name)
    // nil - no error if there is no error nil will be empty if there is an error it will contain the error message
    return message, nil
}
func randomFormat() (string, error) {
    // return a random format from a set of formats
    formats := []string{
        "Hi %v, Welcome!",
        "Hello %v, How are you?",
        "Greetings %v!",
    }
    return formats[rand.Intn(len(formats))], nil
}
```
- We imported the math/rand package to generate random numbers.
- The randomFormat function returns a   random greeting format from a predefined list.
- In the Hello function, we call randomFormat to get a random greeting format.
- We use the returned format to create the greeting message with fmt.Sprintf.


#### Hellos funtionality 
    till now we were able to return only one greeting message
    but now we will modify the Hello function to return multiple greeting messages
[[File: 01/Exercise1/greetings/greetings.go]]
```go
package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting message for the named person.

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty Name")
	}
	format, err := randomFormat()
	if err != nil {
		return "", err
	}
	message := fmt.Sprintf(format, name)
	// nil - no error if there is no error nil will be empty if there is an error it will contain the error message
	return message, nil
}

func randomFormat() (string, error) {
	// return a random format from a set of formats
	formats := []string{
		"Hi %v, Welcome!",
		"Hello %v, How are you?",
		"Greetings %v!",
	}
	return formats[rand.Intn(len(formats))], nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}
```

- We added a new function Hellos that takes a slice of names and returns a map of names to greeting messages.
- Inside Hellos, we create a map to store the messages.
- We loop through each name, call the Hello function to get the greeting message, and store it in the map.
- If there is an error while getting a greeting for any name, we return nil and the error.
- Finally, we return the map of messages and nil (indicating no error).

-- also one important lesson here is Keeping your modules backward compatible is crucial when making changes. By adding new functions like Hellos instead of modifying existing ones, you ensure that existing code depending on your module continues to work without issues.
https://go.dev/blog/module-compatibility


#### Adding Testing 
- go has build in Unit testing support 

[[File: 01/Exercise1/greetings/greetings_test.go]]
```go
package greetings

import (
	"regexp"
	"testing"
)

// t *testing.T - testing framework passes a pointer to testing.T type to the test function
func TestHello(t *testing.T) {
	//
	name := "luffy"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("luffy")
	if err != nil {
		t.Fatalf("Hello(%q) returned error: %v", name, err)
	}
	if !want.MatchString(msg) {
		t.Fatalf("Hello(%q) = %q, want match for %q", name, msg, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	_, err := Hello("")
	if err == nil {
		t.Fatal("Expected an error for empty name, but got none")
	}
}
```

- We created a new file greetings_test.go in the greetings package.
- We imported the testing package to write unit tests.
- We defined two test functions: TestHello and TestHelloEmpty.
- In TestHello, we check if the Hello function returns a greeting message containing the provided name
- In TestHelloEmpty, we check if the Hello function returns an error when given an empty name.
- We use t.Fatalf to report errors and fail the test if the conditions are not met.
- To run the tests, use the command:
```bash
go test
```
```go
PASS
ok      example.com/greetings 0.123s
```


#### compiling and installing the application 
  go run 
    - shortcut to compile and run the application in one step
    - when youre making freq changes it dosnt make sense to compile every time
  - go build
    - compiles the application and creates an executable file
    - useful when you want to distribute or deploy the application
  - go install
    - compiles and installs the application to your Go workspace's bin directory
    - makes the application globally accessible from the command line
  - go build vs go install
    - go build creates an executable in the current directory
    - go install places the executable in the Go workspace's bin directory for easy access


- one notable thing i learned here is your code and youre enviornment management 
  - i have zsh installed 
  - ussullay go is configured with bash shell
  - so when i was trying to run go install it was giving me error that go binary not found
  - so i had to add go binary path to my zsh config file
  - export PATH=$PATH:/usr/local/go/bin
  - after adding this line to my .zshrc file it worked perfectly fine
  - this is important because different shell have different config files and if you switch between shells you have to make sure that your enviornment variables are set correctly in each shell's config file
  

  