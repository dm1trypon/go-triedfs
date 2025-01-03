# triedfs

`triedfs` is a Go package that implements a Trie data structure with capabilities for depth-first search. It allows users to efficiently add and search for sequences of comparable values.

## Features

- **Add Sequences**: Insert sequences of values into the trie.
- **Search Sequences**: Check if a given sequence of values exists in the trie.

## Installation

To install the package, run:

```bash
go get github.com/dm1trypon/triedfs
```

## Usage

### Creating a Trie

You can create a new instance of the Trie like this:

```go
import "github.com/dm1trypon/triedfs"

trie := triedfs.NewTrie[int]() // Replace "int" with any comparable type
```

### Adding Values

To add a sequence of values to the trie, use the `Add` method:

```go
trie.Add([]int{1, 2, 3})
```

### Searching for Values

To check if a specific sequence of values exists in the trie, use the `Search` method:

```go
found := trie.Search([]int{1, 2, 3}) // returns true
notFound := trie.Search([]int{4, 5, 6}) // returns false
```

## Example

Here’s a complete example that demonstrates how to use the `triedfs` package:

```go
package main

import (
    "fmt"
	
    "github.com/dm1trypon/triedfs"
)

func main() {
    trie := triedfs.NewTrie[int]()
    
    // Add sequences to the trie
    trie.Add([]int{1, 2, 3})
    trie.Add([]int{4, 5, 6})
    
    // Search for existing and non-existing sequences
    fmt.Println(trie.Search([]int{1, 2, 3})) // Output: true
    fmt.Println(trie.Search([]int{4, 5, 7})) // Output: false
}
```
