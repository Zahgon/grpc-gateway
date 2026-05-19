package utilities

// DoubleArray is a Double Array implementation of trie on sequences of strings.
type DoubleArray struct {
	// Encoding keeps an encoding from string to int
	Encoding map[string]int
	// Base is the base array of Double Array
	Base []int
	// Check is the check array of Double Array
	Check []int
}

// NewDoubleArray builds a DoubleArray from a set of sequences of strings.
func NewDoubleArray(seqs [][]string) *DoubleArray { _ = "STUB: not implemented"; return nil }

func registerTokens(da *DoubleArray, seqs [][]string) [][]int {
	_ = "STUB: not implemented"
	return nil
}

type node struct {
	row, col    int
	left, right int
}

func (n node) value(seqs [][]int) int { _ = "STUB: not implemented"; return 0 }

func (n node) children(seqs [][]int) []*node { _ = "STUB: not implemented"; return nil }

func addSeqs(da *DoubleArray, seqs [][]int, pos int, n node) { _ = "STUB: not implemented"; return }

func ensureSize(da *DoubleArray, i int) { _ = "STUB: not implemented"; return }

type byLex [][]int

func (l byLex) Len() int           { _ = "STUB: not implemented"; return 0 }
func (l byLex) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (l byLex) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// HasCommonPrefix determines if any sequence in the DoubleArray is a prefix of the given sequence.
func (da *DoubleArray) HasCommonPrefix(seq []string) bool { _ = "STUB: not implemented"; return false }
