package main

import (
	"a4m.dev/godev/chapter4/treesort"
)

func main() {
	tree := treesort.NewTree("Lagos")
	treesort.Add(tree, "Abuja")
	treesort.Add(tree, "Kano")
	treesort.Add(tree, "Enugu")
	treesort.Add(tree, "Ibadan")
	treesort.Add(tree, "Calabar")
	treesort.PrintTree(tree, 0)
}
