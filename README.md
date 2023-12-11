# Patterns

## A Pattern Execution Application

### Overview
This application utilizes the Command Pattern to execute specified patterns. It includes a `PatternOperator` struct that manages and executes patterns based on provided names.

The purpose is to hold a variety of useful programming patterns implemented in Go as an experimental reference. These are for experimental and educational purposes that could be applied to projects for design purposes.

### Usage
To run a specific pattern, use the `-pattern` flag followed by the pattern name. For example:

### Clone the repo

```sh
git clone github.com/lkendrickd/patterns/.git
```

### Change directory to the project directory
```sh
cd github.com/lkendrickd/patterns
```

### Run the pattern application - default

You can also set the pattern to be executed via an **environment variable** called **PATTERN**. The application prioritizes environment variables over command line flags. To set a pattern via environment variable, use:

```sh
go run cmd/patterns.go -pattern foo
```
**OR**

```sh
PATTERN=foo go run cmd/patterns.go
```

**NOTE: - Variable Presidence**
The application takes environment variables as the truth source of the pattern name.
What that means is even if you pass a pattern to the -pattern flag and the environment variable PATTERN
is set it will use the value of $PATTERN



### Adding Your Own Pattern
1. Define your pattern function matching the `Patterner` interface.
2. Create an instance of your pattern using `patterner.NewPattern`.
3. Add your pattern to the `PatternOperator` using `patternOperator.AddPattern`.

#### Example - just prints bar to stdout as an example using an anonymous function or adhoc function
Adding a pattern named "bar" just give it a name and a function:
```go
patternOperator.AddPattern(patterner.NewPattern(
    "bar",
    func() error {
        // Your pattern logic here
        fmt.Println("bar executed")
        // Do something else arbitrary or pattern related.
        return nil
    },
))
```

#### Another example you can have your function outside and not be anonymous

Define the function someplace in it's own package or in cmd/patterns.go

```go

// Define the function outside of the AddPattern - not an anonymous function
// bar is a function that is called to execute the bar pattern if there was one
func bar() error {
	fmt.Println("bar")
	return nil
}

// Add the bar function as a pattern
patternOperator.AddPattern(patterner.NewPattern("bar", bar)

```