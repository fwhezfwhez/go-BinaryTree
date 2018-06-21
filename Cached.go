package BTree

//a bt cached indexed by  root
//one bt only refer to one btreeCache
type BTreeCache struct{
	Root *BinaryNode
	Height int
	MinNode *BinaryNode
	MaxNode *BinaryNode
	MaxDistance int
	AscArray []interface{}
	DescArray []interface{}
	AscLinkedList *SortedLinkedList
	DescLinkedList *SortedLinkedList
}

//save different btCache
var Cache = make(map[*BinaryNode]*BTreeCache,0)

//Refresh a btCache
func (bc *BTreeCache) Refresh()error{
	var er error
	bc.Height = bc.Root.GetMaxHeight()
	bc.MinNode,er = bc.Root.FindMin()
	if er!=nil {
		return er
	}
	bc.MaxNode,er= bc.Root.FindMax()
	if er!=nil {
		return er
	}
	bc.MaxDistance,er = bc.Root.GetMaxDistance()
	if er!=nil {
		return er
	}

	var sum int
	var flag int
	bc.Root.GetNodesNum(&sum)
	arrTmp := make([]interface{},sum)
	bc.Root.ToAscArray(&arrTmp,&flag)
	bc.AscArray = arrTmp

	var sum2 int
	var flag2 int
	bc.Root.GetNodesNum(&sum2)
	arrTmp2 := make([]interface{},sum2)
	bc.Root.ToDescArray(&arrTmp2,&flag2)
	bc.DescArray = arrTmp2

	bc.AscLinkedList,er = bc.Root.ToAscLinkedList()
	if er!=nil {
		return er
	}
	bc.DescLinkedList,er= bc.Root.ToDescLinkedList()
	if er!=nil {
		return er
	}
	return  nil
}
