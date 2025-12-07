# The Go Programming Language - Learning Repository

Code examples and exercises from **The Go Programming Language** by Alan A. A. Donovan and Brian W. Kernighan.

## Structure

### Chapter 1: Tutorial
- **helloworld** - Basic Hello World with command-line arguments
- **echocmd** - Command-line argument echo implementations
- **dup1/dup2/dup3** - Duplicate line detection from stdin/files
- **lissajous** - Animated GIF generation with HTTP server
- **ch1_fetch** - HTTP client for fetching URLs
- **fetchall** - Concurrent URL fetching
- **webserver1/2/3** - HTTP server examples

### Chapter 2: Program Structure
- **tempconv** - Temperature conversion package (Celsius/Fahrenheit)
- **cf** - Command-line temperature converter

### Chapter 3: Basic Data Types
- **datatypes** - Exploration of Go's basic data types

### Chapter 4: Composite Types
- **arrays** - Array fundamentals and currency symbols example
- **slices** - Slice operations with months/quarters
- **append** - Slice append operations
- **reverse** - In-place slice reversal
- **nonempty** - Filtering empty strings from slices
- **sha256** - SHA256 hash computation
- **maps/learn** - Map operations with user data and sorting
- **maps/dedup** - Deduplication using maps
- **stackds** - Stack data structure implementation

### Exercises
- **ex4.5** - Adjacent duplicate elimination (generic implementation)
- **ex4.6** - Unicode whitespace squashing in byte slices

## Running Examples

```bash
# Navigate to any chapter/example directory
cd chapter1/helloworld
go run hello.go

# Run with tests
cd chapter4/nonempty
go test

# Run web server examples
cd chapter1/lissajous
go run main.go
# Visit http://localhost:8000?cycle=5
```

## Dependencies

Some examples use external packages:
```bash
go get github.com/google/uuid
```

## Notes

- Each subdirectory contains its own `go.mod` for module management
- Binary executables are gitignored
- Examples follow the book's progression and coding style
