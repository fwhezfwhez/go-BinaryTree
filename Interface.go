package BTree

type BTreeI interface {
	//if a bt contains data
	Contain(data interface{}) (bool, error)
	//get the node which has max data
	FindMax() (*BinaryNode, error)
	//get the node which has min data
	FindMin() (*BinaryNode, error)
	//insert a data
	Insert(x interface{}) (*BinaryNode, error)
	//delete a node which's data is x,if several nodes' data are all x,only delete one of them
	Remove(x interface{}) (*BinaryNode, error)
	//get a node 's depth in a bt
	GetDepth(node *BinaryNode) (int, error)
	//Get the nearest ancestor of two node
	GetAncestor(node1, node2 *BinaryNode) (*BinaryNode, error)


	//Assume node1 and node2 is contained,get the nearest distance of two node,however,it didn't support compare two nodes which have the same data,New(11) and new(11) is uncomparable
	//if node1 or node2 isn't included,nil pointer error occurred
	GetDistance(node1, node2 *BinaryNode) (int, error)
	//get max depth of a bt, returns maxDepth,leftMaxDepth,rightMaxDepth,error
	GetMaxDepth() (int, int, int, error)
	//get max two node distance of a bt
	GetMaxDistance() (int, error)
	//get a bt's node number
	 GetNodesNum(sum *int)
	//get a bt's node number escaping arg
	GetNodesNumEscapingArgs() int
	//get the maxHeight of a bt
	GetMaxHeight() int
	//get a node's height
	GetNodeHeight(node *BinaryNode) int

	//transfer a bt to a array asc sorted
	ToAscArray(rs *[]interface{}, flag *int)
	//transfer a bt to an array asc escaping args
	ToAscArrayEscapingArgs() []interface{}
	//transfer a bt to a array desc sorted
	ToDescArray(rs *[]interface{}, flag *int)
	//transfer a bt to an array desc escaping args
	ToDescArrayEscapingArgs() []interface{}
	//Assume got a node , find its root by this function
	FindRoot() *BinaryNode
	//Cached a bt in global Cache
	Cached() *BTreeCache
	//Get a bt's cache,if not existed,throws an error
	GetCache() (*BTreeCache, error)
	//Get a bt's cache,if not existed,generate one then return it
	MustGetCache() *BTreeCache
	//transfer a bt to a linked list order asc
	ToAscLinkedList() (*SortedLinkedList, error)
	//transfer a bt to a LinkedList desc
	ToDescLinkedList() (*SortedLinkedList, error)
}
