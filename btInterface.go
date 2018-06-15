package BTree
type BTreeMehodsI interface{
	//Where data is contained in a bt
	Contains(data interface{}) (bool)
	//
	FindMin() (interface{},error)
	FindMax() (interface{},error)
	Insert(data interface{}) error
	Remove(data interface{}) error
	PrintTree()error
}