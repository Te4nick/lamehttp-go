package handle

import (
	"fmt"
	"strings"
)

type URITrie struct {
	root    map[string]*URITrie
	Methods map[string]HandlerFunc
}

func NewURITrie() *URITrie {
	return &URITrie{
		root:    make(map[string]*URITrie),
		Methods: nil,
	}
}

func (t *URITrie) Put(key string, value map[string]HandlerFunc) *URITrie {
	before, after, _ := strings.Cut(key, "/")
	newRoot := NewURITrie()
	t.root[before] = newRoot
	if after != "" {
		return newRoot.Put(after, value)
	}

	newRoot.Methods = value
	return newRoot
}

func (t *URITrie) Get(key string) *URITrie {
	before, after, _ := strings.Cut(key, "/")
	subTrie, ok := t.root[before]
	if !ok {
		return nil
	}

	if after != "" {
		return subTrie.Get(after)
	}

	return t
}

func (t *URITrie) Print() {
	t.pprint(0)
}

func (t *URITrie) pprint(level int) {
	if len(t.root) == 0 {
		return
	}

	for key, subTrie := range t.root {
		fmt.Println(strings.Repeat("\t", level), key)
		subTrie.pprint(level + 1)
	}
}
