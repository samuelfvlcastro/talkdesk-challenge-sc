// implements a rune optimized radix tree aka prefix tree
// ASCII table: https://www.cs.cmu.edu/~pattis/15-1XX/common/handouts/ascii.html
// TODO: implement the Remove method

package rtrie

import (
	"errors"
	"fmt"
)

// Tree radix tree where Nodes are stored
type Tree struct {
	// Root Tree base Node
	Root *Node
	// minCp represents low-end ascii codepoint
	minCp rune
	// maxCp represents high-end ascii codepoint
	maxCp rune
	// cLength calculated max Node Children
	cLength int
}

// Node rune information container for a single data structure
type Node struct {
	Key      rune
	IsEnd    bool
	Value    string
	Children []*Node
}

// New instantiates a new rtrie with the selected upper and lower bounds
func New(initialRune, finalRune rune) *Tree {
	cLength := int(finalRune) - int(initialRune) + 1
	return &Tree{
		Root: &Node{
			Children: make([]*Node, cLength),
		},
		minCp:   initialRune,
		maxCp:   finalRune,
		cLength: cLength,
	}
}

// Insert walks the tree according to the string to be inserted adding it if not contained
func (t *Tree) Insert(stg string) error {
	root := t.Root
	for _, cp := range stg {
		if err := t.validateRune(cp); err != nil {
			return err
		}

		idx := int(cp) - int(t.minCp)
		if root.Children[idx] == nil {
			root.Children[idx] = &Node{
				Key:      cp,
				Children: make([]*Node, t.cLength),
			}
		}
		root = root.Children[idx]
	}

	root.IsEnd = true
	root.Value = stg

	return nil
}

// Remove walks the tree according to the string to be removed
func (t *Tree) Remove(stg string) error {
	return errors.New("not implemented")
}

// Search transverses the tree trying to find the given string
func (t *Tree) Search(stg string) (string, error) {
	found := ""
	root := t.Root
	for _, cp := range stg {
		if err := t.validateRune(cp); err != nil {
			return found, err
		}

		idx := int(cp) - int(t.minCp)
		if root.Children[idx] == nil {
			break
		}

		root = root.Children[idx]
		if root.IsEnd {
			found = root.Value
		}
	}

	return found, nil
}

func (t *Tree) validateRune(cp rune) error {
	if cp < t.minCp || cp > t.maxCp {
		return fmt.Errorf("rune %q out of bounds (%d<=%d<=%d)", cp, t.minCp, cp, t.maxCp)
	}

	return nil
}
