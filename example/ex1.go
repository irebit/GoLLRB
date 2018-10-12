package main

import (
	"fmt"
	"github.com/petar/GoLLRB/llrb"
)

func showItem(item llrb.Item) bool {
	fmt.Print(item)
	return true
}

func main() {
	tree := llrb.New()
	tree.ReplaceOrInsert(llrb.Int(0))
	tree.ReplaceOrInsert(llrb.Int(1))
	tree.ReplaceOrInsert(llrb.Int(2))
	tree.ReplaceOrInsert(llrb.Int(3))
	tree.ReplaceOrInsert(llrb.Int(4))
	tree.ReplaceOrInsert(llrb.Int(5))
	tree.ReplaceOrInsert(llrb.Int(6))
	tree.ReplaceOrInsert(llrb.Int(7))
	tree.ReplaceOrInsert(llrb.Int(8))
	tree.ReplaceOrInsert(llrb.Int(9))
	max := tree.Max()
	fmt.Print(max)
	min := tree.DeleteMin()
	fmt.Print(min)
	item := tree.Delete(llrb.Int(4))
	fmt.Print(item)
	tree.AscendRange(llrb.Int(1), llrb.Int(10), showItem)
}
