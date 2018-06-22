BTree realization in go

**if you find any bug or want to add new requirement,call me at 1728565484@qq.com,or submit your issur above**

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
	
	//11.Get the node num of a bt
	fmt.Println(node.GetNodesNumEscapingArgs())
	
	//12.Get bt's max height and the height of a specific node
	fmt.Println(node.GetMaxHeight())
	fmt.Println(node.GetNodeHeight(node2))
	
	//13.Transfer a bt to an array asc or desc
	fmt.Println(node.ToAscArrayEscapingArgs())
	fmt.Println(node.ToDescArrayEscapingArgs())
	
	//14.Find a node' root
	fmt.Println(node2.FindRoot())
	
	//15.Cached a  bt
	node.Cached()
	rs :=node.MustGetCache() 
	bt.SmartPrint(rs)
	
	//16. transfer a bt to an sortedDoubleWayLinkList
	fmt.Println(node.ToDescLinkedList())
	fmt.Println(node.ToAscLinkedList())
	
	//17. Balance a bt'copy ,no modifying to itself
	fmt.Print(node.BalanceCopy().ToDescArrayEscapintArgs())
	//18. Balance itself
	fmt.Println(node.Balance().ToDescArrayEscapingArgs())
	//19. transfer a normal bt to an AVL
	fmt.Println(node.ToAVL().Root)
	/*
	  	//Balance a bt'copy ,no modifying to itself
          BalanceCopy() BinaryNode
          //Balance a bt to avl 
          Balance()*BinaryNode
          //Pacakge a balanced bt to the specific new struct AVL,which is extendable and now contains a root *Binary
          ToAVL()*AVL
	 */
}
```