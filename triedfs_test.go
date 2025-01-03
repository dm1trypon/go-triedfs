package triedfs

import (
	"testing"
)

type testStruct struct {
	a int
	b string
	c [4]float64
}

func TestTrie_Search(t *testing.T) {
	tests := []struct {
		name      string
		trie      any
		values    []any
		findValue any
		exp       bool
	}{
		{
			name:      "SearchExistingString_ReturnsTrue",
			trie:      NewTrie[rune](),
			values:    []any{"one", "two", "three"},
			findValue: "one",
			exp:       true,
		},
		{
			name:      "SearchNonExistentString_ReturnsFalse",
			trie:      NewTrie[rune](),
			values:    []any{"one", "two", "three"},
			findValue: "four",
			exp:       false,
		},
		{
			name: "SearchWithValidPrefix_ReturnsTrue",
			trie: NewTrie[string](),
			values: []any{
				[]string{"one", "two", "three"},
				[]string{"four", "five", "six"},
				[]string{"seven", "eight", "nine"},
			},
			findValue: []string{"one", "two", "three"},
			exp:       true,
		},
		{
			name: "SearchWithInvalidCombination_ReturnsFalse",
			trie: NewTrie[string](),
			values: []any{
				[]string{"one", "two", "three"},
				[]string{"four", "five", "six"},
				[]string{"seven", "eight", "nine"},
			},
			findValue: []string{"one", "five"},
			exp:       false,
		},
		{
			name: "SearchStructKey_Exists_ReturnsTrue",
			trie: NewTrie[testStruct](),
			values: []any{
				testStruct{1, "one", [4]float64{1.00, 2.00, 3.00, 4.00}},
				testStruct{2, "two", [4]float64{2.00, 3.00, 4.00, 5.00}},
				testStruct{3, "three", [4]float64{3.00, 4.00, 5.00, 6.00}},
			},
			findValue: []testStruct{
				{2, "two", [4]float64{2.00, 3.00, 4.00, 5.00}},
			},
			exp: true,
		},
		{
			name: "SearchStructKey_ModifiedValue_ReturnsFalse",
			trie: NewTrie[testStruct](),
			values: []any{
				testStruct{1, "one", [4]float64{1.00, 2.00, 3.00, 4.00}},
				testStruct{2, "two", [4]float64{2.00, 3.00, 4.00, 5.00}},
				testStruct{3, "three", [4]float64{3.00, 4.00, 5.00, 6.00}},
			},
			findValue: []testStruct{
				{2, "two", [4]float64{2.00, 9.00, 4.00, 5.00}},
			},
			exp: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addValuesToTrie(tt.trie, tt.values)

			if result := searchTrie(tt.trie, tt.findValue); result != tt.exp {
				t.Errorf("expected %v, got %v", tt.exp, result)
			}
		})
	}
}

func addValuesToTrie(trie any, values []any) {
	for i := range values {
		switch v := trie.(type) {
		case *Trie[rune]:
			v.Add([]rune(values[i].(string)))
		case *Trie[string]:
			v.Add(values[i].([]string))
		case *Trie[int]:
			v.Add(values[i].([]int))
		case *Trie[testStruct]:
			v.Add([]testStruct{values[i].(testStruct)})
		default:
			return
		}
	}
}

func searchTrie(trie any, findValue any) bool {
	switch v := trie.(type) {
	case *Trie[rune]:
		return v.Search([]rune(findValue.(string)))
	case *Trie[string]:
		return v.Search(findValue.([]string))
	case *Trie[int]:
		return v.Search(findValue.([]int))
	case *Trie[testStruct]:
		return v.Search(findValue.([]testStruct))
	default:
		return false
	}
}

// cpu: Intel(R) Core(TM) i7-7700 CPU @ 3.60GHz
// BenchmarkAdd
// BenchmarkAdd-8            108346             10959 ns/op               1 B/op            0 allocs/op
func BenchmarkAdd(b *testing.B) {
	trie := NewTrie[int]()
	values := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		values[i] = i
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		trie.Add(values)
	}

	b.ReportAllocs()
}

// cpu: Intel(R) Core(TM) i7-7700 CPU @ 3.60GHz
// BenchmarkSearch
// BenchmarkSearch-8         111585             10934 ns/op               0 B/op           0 allocs/op
func BenchmarkSearch(b *testing.B) {
	trie := NewTrie[int]()
	values := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		values[i] = i
	}
	trie.Add(values)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		trie.Search(values)
	}

	b.ReportAllocs()
}
