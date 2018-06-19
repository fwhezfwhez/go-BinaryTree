BTree realization in go

**start**

go get github.com/fwhezfwhez/go-BinaryTree

Example:
```go
package main
import (
	bt "github.com/fwhezfwhez/go-BinaryTree"
	"fmt"
)
var node,node2 *bt.BinaryNode
func init() {
		node=bt.New(7)
    	node2,_ =node.Insert(5)
    	node.Insert(9)
    	node.Insert(11)
    	node.Insert(11)
    	node.Insert(10)
    	node.Insert(8)
}
func main() {
	//1.Contain,whether a btree contains data x
	fmt.Println(node.Contain(9))
	
	//2.FindMax,find the max data in a b tree
	fmt.Println(node.FindMax())
	
	//3.FindMin,find the min data in a b tree
	fmt.Println(node.FindMin())
	
	//4.Insert,insert a data to  a b tree
	fmt.Println(node.Insert(15))
	
	//5.Remove,remove a data from a b tree
	fmt.Println(node.Remove(15))
	
	//6.Get nearest Ancestor of two node
	//make sure the nodes are contained in b tree,or nil pointer error occurred 
	fmt.Println(node.GetAncestor(bt.New(11),bt.New(9)))
	
	//7.Get a node's depth
	//make sure the node is contained in b tree,or nil pointer error occurred 
	fmt.Println(node.GetDepth(bt.New(9)))
	
	//8.Get the shortest distance of two node 
	//make sure the nodes are contained in b tree,or nil pointer error occurred 
	fmt.Println(node.GetDistance(bt.New(11),bt.New(9)))
	
	//9.Get the max depth and the left max depth,right max depth of a b tree or a node
	fmt.Println(node.GetMaxDepth())
	fmt.Println(node.Left.GetMaxDepth())
	
	//10.Get the max distance of two leaves of a b tree
	fmt.Println(node.GetMaxDistance())
}
```