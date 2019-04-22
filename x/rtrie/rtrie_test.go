package rtrie

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestNew(t *testing.T) {
	RegisterTestingT(t)

	expected := &Tree{
		Root: &Node{
			Children: make([]*Node, 26),
		},
		minCp:   97,
		maxCp:   122,
		cLength: 26,
	}

	trie := New(97, 122)
	Expect(trie).To(Equal(expected), "should return the expected trie")
}

func TestInsert(t *testing.T) {
	RegisterTestingT(t)
	cpMin := rune(97)
	cpMax := rune(122)
	trie := New(cpMin, cpMax)
	trie.Insert("cris")

	Expect(trie.Root.Children[99-cpMin].Key).To(Equal(rune(99)), "should return 99 the ascii codepoint for 'c'")
	Expect(trie.Root.Children[99-cpMin].Children[114-cpMin].Key).To(Equal(rune(114)), "should return 114 the ascii codepoint for 'r'")
	Expect(trie.Root.Children[99-cpMin].Children[114-cpMin].Children[105-cpMin].Key).To(Equal(rune(105)), "should return 105 the ascii codepoint for 'i'")
	Expect(trie.Root.Children[99-cpMin].Children[114-cpMin].Children[105-cpMin].Children[115-cpMin].Key).To(Equal(rune(115)), "should return 115 the ascii codepoint for 's'")

	Expect(trie.Root.Children[99-cpMin].Children[114-cpMin].Children[105-cpMin].Children[115-cpMin].Value).To(Equal("cris"), "should the final value")
	Expect(trie.Root.Children[99-cpMin].Children[114-cpMin].Children[105-cpMin].Children[115-cpMin].IsEnd).To(BeTrue(), "should be leaf")
}

func TestSearch(t *testing.T) {
	RegisterTestingT(t)
	lettersTrie := New(97, 122)

	lettersTrie.Insert("craft")
	lettersTrie.Insert("crafty")
	lettersTrie.Insert("crate")

	Expect(lettersTrie.Search("crafting")).To(Equal("craft"), "should find the expected prefix")

	numbersTrie := New(48, 57)
	numbersTrie.Insert("1")
	numbersTrie.Insert("120")
	numbersTrie.Insert("121")
	numbersTrie.Insert("122")
	numbersTrie.Insert("123")
	numbersTrie.Insert("1234")

	Expect(numbersTrie.Search("134859487")).To(Equal("1"), "should find the expected prefix")
	Expect(numbersTrie.Search("122847294")).To(Equal("122"), "should find the expected prefix")
	Expect(numbersTrie.Search("123474839")).To(Equal("1234"), "should find the expected prefix")
	Expect(numbersTrie.Search("224837364")).To(Equal(""), "should find the expected prefix")
}
