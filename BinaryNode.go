package BTree

import (
	"reflect"
	"errors"
	"fmt"
	"github.com/fwhezfwhez/go-queue"
)

type BinaryNode struct {
	Data   interface{}
	Parant *BinaryNode
	Left   *BinaryNode
	Right  *BinaryNode
}

type SortedLinkedList struct {
	Start *BinaryNode
}
type AVL struct {
	Root *BinaryNode
}

type RedBlack struct {
	Root *BinaryNode
}

//new a node
func New(data interface{}) *BinaryNode {
	return &BinaryNode{data, nil, nil, nil}
}

//new a node with parent
func NewWithParent(data interface{}, prt *BinaryNode) *BinaryNode {
	return &BinaryNode{data, prt, nil, nil}
}

//new a node with left child
func NewWithLeft(data interface{}, left *BinaryNode) *BinaryNode {
	return &BinaryNode{data, nil, left, nil}
}

// balance the copy of bt to an avl
func (root *BinaryNode) BalanceCopy() *BinaryNode {
	copy := *root
	return copy.Balance()
}

//balance the bt to an avl
//unfinished
func (bt *BinaryNode) Balance() *BinaryNode {
	if bt == nil {
		return nil
	}

	if bt.Left != nil && bt.Right != nil {
		bt.Left = bt.Left.Balance()
		bt.Right = bt.Right.Balance()
	} else {
		if (bt.Left.GetMaxHeight() - bt.Right.GetMaxHeight()) > 1 {
			if bt.Left.Left.GetMaxHeight() >= bt.Left.Right.GetMaxHeight() {
				return bt.RotateWithLeftChild()
			} else {
				return bt.DoubleRotateLeftChild()
			}
		} else {
			if (bt.Right.GetMaxHeight() - bt.Left.GetMaxHeight()) > 1 {
				if bt.Right.Right.GetMaxHeight() >= bt.Right.Left.GetMaxHeight() {
					return bt.RotateWithRightChild()
				} else {
					return bt.DoubleRotateRightChild()
				}
			}
		}
		return bt
	}

	return bt
}

func (bt *BinaryNode) RotateWithLeftChild() *BinaryNode {
	if bt.Left != nil && bt.Left.Right != nil {
		tmp := bt.Left
		bt.Left = tmp.Right
		tmp.Right.Parant = bt
		tmp.Right = bt
		tmp.Parant = bt.Parant

		if bt.Parant != nil {
			if bt.Parant.Left == bt {
				bt.Parant.Left = tmp
			} else {
				bt.Parant.Right = tmp
			}
		}
		bt.Parant = tmp
		return tmp
	} else {
		return bt
	}
}

func (bt *BinaryNode) DoubleRotateLeftChild() *BinaryNode {
	if bt.Left != nil {
		bt.Left = bt.Left.RotateWithRightChild()
		return bt.RotateWithLeftChild()
	} else {
		return bt
	}
}

func (bt *BinaryNode) DoubleRotateRightChild() *BinaryNode {
	if bt.Right != nil {
		bt.Right = bt.Right.RotateWithLeftChild()
		return bt.RotateWithRightChild()
	} else {
		return bt
	}

}

func (bt *BinaryNode) RotateWithRightChild() *BinaryNode {
	if bt.Right != nil && bt.Right.Left != nil {
		tmp := bt.Right
		bt.Right = tmp.Left
		tmp.Left.Parant = bt

		tmp.Left = bt
		tmp.Parant = bt.Parant

		if bt.Parant != nil {
			if bt.Parant.Left == bt {
				bt.Parant.Left = tmp
			} else {
				bt.Parant.Right = tmp
			}
		}
		bt.Parant = tmp
		return tmp
	} else {
		return bt
	}

}

//if a bt contains data
func (bt *BinaryNode) Contain(data interface{}) (bool, *BinaryNode, error) {
	if data == nil {
		return false, nil,nil
	}
	rs, err := Compare(bt.Data, data)
	if err != nil {
		return false,nil, err
	}

	if rs == 0 {
		return true,bt, nil
	}
	if rs == 1 {
		return bt.Left.Contain(data)
	}
	if rs == -1 {
		return bt.Right.Contain(data)
	}
	return false, nil,nil
}

//get the node which has max data
func (bt *BinaryNode) FindMax() (*BinaryNode, error) {
	if bt.Right != nil {
		return bt.Right.FindMax()
	} else {
		return bt, nil
	}
}

//get the node which has min data
func (bt *BinaryNode) FindMin() (*BinaryNode, error) {
	if bt.Left != nil {
		return bt.Left.FindMin()
	} else {
		return bt, nil
	}
}

//insert a data
func (bt *BinaryNode) Insert(x interface{}) (*BinaryNode, error) {
	rs, err := Compare(bt.Data, x)
	if err != nil {
		return nil, err
	}

	if rs == 1 {
		if bt.Left != nil {
			return bt.Left.Insert(x)
		} else {
			bt.Left = NewWithParent(x, bt)
			return bt.Left, nil
		}
	} else {
		if bt.Right != nil {
			return bt.Right.Insert(x)
		} else {
			bt.Right = NewWithParent(x, bt)
			return bt.Right, nil
		}

	}
	return nil, nil
}

//delete a node which's data is x,if several nodes' data are all x,only delete one of them
func (bt *BinaryNode) Remove(x interface{}) (*BinaryNode, error) {
	if bt == nil {
		return nil, nil
	}
	rs, er := Compare(bt.Data, x)
	if er != nil {
		return nil, er
	}
	if rs == -1 {
		bt.Right, er = bt.Right.Remove(x)
		if er != nil {
			return nil, er
		}
	} else if rs == 1 {
		bt.Left, er = bt.Left.Remove(x)
		if er != nil {
			return nil, er
		}
	} else if bt.Left != nil && bt.Right != nil {
		tmp, _ := bt.Right.FindMin()
		bt.Right.Remove(tmp.Data)
		if bt.Parant.Right == bt && bt.Parant != nil {
			bt.Parant.Right = tmp
			tmp.Parant = bt.Parant

			tmp.Left = bt.Left
			bt.Left.Parant = tmp

			if bt.Right != tmp {
				tmp.Right = bt.Right
				bt.Right.Parant = tmp
			}
			return tmp, nil
		} else if bt.Parant.Left == bt && bt.Parant != nil {
			bt.Parant.Left = tmp
			tmp.Parant = bt.Parant

			tmp.Left = bt.Left
			bt.Left.Parant = tmp

			if bt.Right != tmp {
				tmp.Right = bt.Right
				bt.Right.Parant = tmp
			}
			return tmp, nil
		}

		//bt.Right, _ = bt.Right.Remove(tmp.Data)
	} else {
		if bt.Left != nil {
			bt = bt.Left
		} else {
			bt = bt.Right
		}
	}

	return bt, nil
}

//get a node 's depth in a bt
func (bt *BinaryNode) GetDepth(node *BinaryNode) (int, error) {
	rs, er := Compare(node.Data, bt.Data)
	if er != nil {
		return -1, er
	}
	if rs == -1 {
		tmpLeft, _ := bt.Left.GetDepth(node)
		return tmpLeft + 1, nil
	} else if rs == 0 {
		return 0, nil
	} else {
		tmpRight, _ := bt.Right.GetDepth(node)
		return tmpRight + 1, nil
	}

}

//Get the nearest ancestor of two node
func (bt *BinaryNode) GetAncestor(node1, node2 *BinaryNode) (*BinaryNode, error) {
	rs1, er := Compare(node1.Data, bt.Data)
	if er != nil {
		return nil, er
	}
	rs2, er := Compare(node2.Data, bt.Data)
	if er != nil {
		return nil, er
	}
	if rs1 == 1 && rs2 == 1 {
		return bt.Right.GetAncestor(node1, node2)
	} else if rs1 == -1 && rs2 == -1 {
		return bt.Left.GetAncestor(node1, node2)
	} else {
		return bt, nil
	}
}

//Assume node1 and node2 is contained,get the nearest distance of two node,however,it didn't support compare two nodes which have the same data,New(11) and new(11) is uncomparable
//if node1 or node2 isn't included,nil pointer error occurred
func (bt *BinaryNode) GetDistance(node1, node2 *BinaryNode) (int, error) {
	ancestor, er := bt.GetAncestor(node1, node2)
	if er != nil {
		return -1, er
	}
	tmp1, er := ancestor.GetDepth(node1)
	if er != nil {
		return -1, er
	}
	tmp2, er := ancestor.GetDepth(node2)
	if er != nil {
		return -1, er
	}
	return tmp1 + tmp2, nil
}

//get max depth of a bt, returns maxDepth,leftMaxDepth,rightMaxDepth,error
func (bt *BinaryNode) GetMaxDepth() (int, int, int, error) {
	if bt == nil {
		return 0, 0, 0, nil
	}
	var leftMax, rightMax int
	var er error
	if bt.Left != nil {
		leftMax, _, _, er = bt.Left.GetMaxDepth()
		if er != nil {
			return -1, -1, -1, er
		}
		leftMax++
	}
	if bt.Right != nil {
		rightMax, _, _, er = bt.Right.GetMaxDepth()
		if er != nil {
			return -1, -1, -1, er
		}
		rightMax++
	}
	if leftMax > rightMax {
		return leftMax, leftMax, rightMax, nil
	}
	return rightMax, leftMax, rightMax, nil

}

//get max two node distance of a bt
func (bt *BinaryNode) GetMaxDistance() (int, error) {
	if bt == nil {
		return 0, nil
	}
	var lMax1 = 0
	var rMax1 = 0
	var lMax2 = 0
	var rMax2 = 0
	var er error

	_, lMax1, rMax1, er = bt.GetMaxDepth()
	if bt.Left != nil {
		lMax2, er = bt.Left.GetMaxDistance()
		if er != nil {
			return -1, er
		}
		if bt.Right != nil {
			lMax2++
		}
	}
	if bt.Right != nil {

		rMax2, er = bt.Right.GetMaxDistance()
		if er != nil {
			return -1, er
		}
		if bt.Left != nil {
			rMax2++
		}
	}
	return max((lMax1 + rMax1), lMax2, rMax2), nil
}

//get a bt's node number
func (bt *BinaryNode) GetNodesNum(sum *int) {
	if bt != nil {
		*sum += 1
		if bt.Right != nil {
			bt.Right.GetNodesNum(sum)
		}
		if bt.Left != nil {
			bt.Left.GetNodesNum(sum)
		}
	}
}

//get a bt's node number escaping arg
func (bt *BinaryNode) GetNodesNumEscapingArgs() int {
	var sum int
	bt.GetNodesNum(&sum)
	return sum
}

//get the maxHeight of a bt
func (bt *BinaryNode) GetMaxHeight() int {
	if bt == nil {
		return 0
	}
	tmp, _, _, _ := bt.GetMaxDepth()
	return tmp
}

//get a node's height
func (bt *BinaryNode) GetNodeHeight(node *BinaryNode) int {
	total, _, _, _ := bt.GetMaxDepth()
	tmp, _ := bt.GetDepth(node)
	return total - tmp
}

//transfer a bt to a array asc sorted
func (bt *BinaryNode) ToAscArray(rs *[]interface{}, flag *int) {
	if bt.Left != nil {
		bt.Left.ToAscArray(rs, flag)
	}
	(*rs)[*flag] = bt.Data
	*flag++
	if bt.Right != nil {
		bt.Right.ToAscArray(rs, flag)
	}
}

//transfer a bt to an array asc escaping args
func (bt *BinaryNode) ToAscArrayEscapingArgs() []interface{} {
	var sum int
	var rs []interface{}
	var flag int
	bt.GetNodesNum(&sum)
	rs = make([]interface{}, sum)
	bt.ToAscArray(&rs, &flag)
	return rs
}

//transfer a bt to a array desc sorted
func (bt *BinaryNode) ToDescArray(rs *[]interface{}, flag *int) {
	if bt.Right != nil {
		bt.Right.ToDescArray(rs, flag)
	}

	(*rs)[*flag] = bt.Data
	*flag++

	if bt.Left != nil {
		bt.Left.ToDescArray(rs, flag)
	}
}

//transfer a bt to an array desc escaping args
func (bt *BinaryNode) ToDescArrayEscapingArgs() []interface{} {
	var sum int
	var rs []interface{}
	var flag int
	bt.GetNodesNum(&sum)
	rs = make([]interface{}, sum)
	bt.ToDescArray(&rs, &flag)
	return rs
}

//Assume got a node , find its root by this function
func (bt *BinaryNode) FindRoot() *BinaryNode {
	if bt.Parant != nil {
		return bt.Parant.FindRoot()
	} else {
		return bt
	}
}

//Cached a bt in global Cache
func (bt *BinaryNode) Cached() *BTreeCache {
	rootTmp := bt.FindRoot()
	for k := range Cache {
		if Cache[k].Root == bt.FindRoot() {
			Cache[k].Refresh()
			return Cache[k]
		}
	}
	Cache[rootTmp] = &BTreeCache{Root: bt.FindRoot()}
	Cache[rootTmp].Refresh()
	return Cache[rootTmp]
}

//Get a bt's cache,if not existed,throws an error
func (bt *BinaryNode) GetCache() (*BTreeCache, error) {
	for k, _ := range Cache {
		if k == bt.FindRoot() {
			return Cache[k], nil
		}
	}
	return nil, errors.New("this binary tree has not been cached yet,use bt.Cached() first,or use bt.MustGetCache()")
}

//Get a bt's cache,if not existed,generate one then return it
func (bt *BinaryNode) MustGetCache() *BTreeCache {
	for k := range Cache {
		if k == bt.FindRoot() {
			return Cache[k]
		}
	}
	return bt.Cached()
}

//transfer a bt to a redBlack pattern
//what is a red-black tree,refer to www.baidu.com, www.google.com
func (bt *BinaryNode) ToRedBlack() error {
	return nil
}

//transfer a bt to an avl parttern
//what is an avl tree,refer to www.baidu.com,www.google.com
func (bt *BinaryNode) ToAVL() *AVL{
	tmp := bt.BalanceCopy()
	return &AVL{Root: tmp}
}

//transfer a bt to a linked list order asc
func (bt *BinaryNode) ToAscLinkedList() (*SortedLinkedList, error) {
	var sum int
	bt.GetNodesNum(&sum)
	var rs = make([]interface{}, sum)
	var flag = 0
	bt.ToAscArray(&rs, &flag)

	head := New(rs[0])
	tail := head
	var tmp *BinaryNode
	for i := 1; i < len(rs); i++ {
		tmp = New(rs[i])
		tail.Right = tmp
		tmp.Left = tail
		tail = tmp
	}
	return &SortedLinkedList{head}, nil
}

//transfer a bt to a LinkedList desc
func (bt *BinaryNode) ToDescLinkedList() (*SortedLinkedList, error) {
	var sum int
	bt.GetNodesNum(&sum)
	var rs = make([]interface{}, sum)
	var flag = 0
	bt.ToDescArray(&rs, &flag)

	head := New(rs[0])
	tail := head
	var tmp *BinaryNode
	for i := 1; i < len(rs); i++ {
		tmp = New(rs[i])
		tail.Right = tmp
		tmp.Left = tail
		tail = tmp
	}
	return &SortedLinkedList{head}, nil
}

//smartPaint a bt
func (bt *BinaryNode) Paint() {
	if bt == nil {
		fmt.Println("nil")
		return
	}

	levelNum := bt.GetLevelNodeNumArr()
	queue := bt.LevelVisitToQueue()
	for i:=0;i<len(levelNum);i++{
		for j:=0;j<levelNum[i];j++{
			fmt.Print(queue.Pop())
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

//Get a bt's children
func (bt *BinaryNode) GetChildren() []*BinaryNode {
	var rs = make([]*BinaryNode, 0)
	if bt.Left != nil {
		rs = append(rs, bt.Left)
	}
	if bt.Right != nil {
		rs = append(rs, bt.Right)
	}
	return rs
}

////rotate left
//func (av *RedBlack) LeftRotate() (*AVL, error) {
//	return nil, nil
//}
//
////rotate right
//func (av *RedBlack) RightRotate() (*AVL, error) {
//	return nil, nil
//}

//compare two interface{} value ,only supports for int,int64,int32,int16,int8,string,float32,float64
func Compare(v1 interface{}, v2 interface{}) (int, error) {
	if reflect.TypeOf(v1).Kind().String() != reflect.TypeOf(v2).Kind().String() {
		return -1, errors.New("二者类型不匹配,对一个bt树链来说，不推荐存放不同数据类型的对象")
	}

	switch v1.(type) {
	case int:
		v1tmp := v1.(int)
		v2tmp := v2.(int)
		if v1tmp > v2tmp {
			return 1, nil
		} else if v1tmp == v2tmp {
			return 0, nil
		} else {
			return -1, nil
		}
	case int64:
		v1tmp := v1.(int64)
		v2tmp := v2.(int64)
		if v1tmp > v2tmp {
			return 1, nil
		} else if v1tmp == v2tmp {
			return 0, nil
		} else {
			return -1, nil
		}
	case int32:
		v1tmp := v1.(int32)
		v2tmp := v2.(int32)
		if v1tmp > v2tmp {
			return 1, nil
		} else if v1tmp == v2tmp {
			return 0, nil
		} else {
			return -1, nil
		}
	case int16:
		v1tmp := v1.(int16)
		v2tmp := v2.(int16)
		if v1tmp > v2tmp {
			return 1, nil
		} else if v1tmp == v2tmp {
			return 0, nil
		} else {
			return -1, nil
		}
	case int8:
		v1tmp := v1.(int8)
		v2tmp := v2.(int8)
		if v1tmp > v2tmp {
			return 1, nil
		} else if v1tmp == v2tmp {
			return 0, nil
		} else {
			return -1, nil
		}
	case string:
		v1tmp := v1.(string)
		v2tmp := v2.(string)
		if v1tmp > v2tmp {
			return 1, nil
		} else if v1tmp == v2tmp {
			return 0, nil
		} else {
			return -1, nil
		}
	case float32:
		v1tmp := v1.(float32)
		v2tmp := v2.(float32)
		if v1tmp > v2tmp {
			return 1, nil
		} else if v1tmp == v2tmp {
			return 0, nil
		} else {
			return -1, nil
		}
	case float64:
		v1tmp := v1.(float64)
		v2tmp := v2.(float64)
		if v1tmp > v2tmp {
			return 1, nil
		} else if v1tmp == v2tmp {
			return 0, nil
		} else {
			return -1, nil
		}
	}
	return -1, errors.New("only support compare int,int8,int16,int32,int64,float32,float64,string")
}

func max(args ... int) int {
	var maxTmp = args[0]
	for i := 0; i < len(args)-1; i++ {
		for j := i + 1; j < len(args); j++ {
			if maxTmp < args[j] {
				maxTmp = args[j]
			}
		}
	}
	return maxTmp
}

func abs(arg int) int {
	if arg > 0 {
		return arg
	} else {
		return -arg
	}
}
func SmartPrint(i interface{}) {
	var kv = make(map[string]interface{})
	vValue := reflect.ValueOf(i)
	vType := reflect.TypeOf(i)
	for i := 0; i < vValue.NumField(); i++ {
		kv[vType.Field(i).Name] = vValue.Field(i)
	}
	fmt.Println("获取到数据:")
	for k, v := range kv {
		fmt.Print(k)
		fmt.Print(":")
		fmt.Print(v)
		fmt.Println()
	}
}

func IfZero(arg interface{}) bool {
	if arg == nil {
		return true
	}
	switch v := arg.(type) {
	case int, float64, int32, int16, int64, float32:
		if v == 0 {
			return true
		}
	case string:
		if v == "" || v == "%%" || v == "%" {
			return true
		}
	case *string, *int, *int64, *int32, *int16, *int8, *float32, *float64:
		if v == nil {
			return true
		}
	default:
		return false
	}
	return false
}

//put a single level datas into a queue
func (bt *BinaryNode) LevelDataToQueue(level int,queue *Queue.Queue) int{
  if bt==nil || level<0{
  	return 0
  }
  if level==0{
  	queue.Push(bt.Data)
  	return 1
  }
  return bt.Left.LevelDataToQueue(level-1,queue)+bt.Right.LevelDataToQueue(level-1,queue)
}

//put a single level nodes into a queue
func (bt *BinaryNode) LevelNodeToQueue(level int,queue *Queue.Queue)int{
	if bt==nil || level<0{
		return 0
	}
	if level==0{
		queue.Push(bt)
		return 1
	}
	return bt.Left.LevelNodeToQueue(level-1,queue)+bt.Right.LevelNodeToQueue(level-1,queue)
}

//put all datas into a queue order by level/layer/depth
func  (bt *BinaryNode) LevelVisitToQueue() *Queue.Queue{
	q := Queue.NewEmpty()
	i := 0
	for i = 0; ; i++ {
		if bt.FindRoot().LevelDataToQueue(i, q) == 0 {
			break
		}
	}
	return q
}

//put all nodes into a queue order by level/layer/depth
func  (bt *BinaryNode) LevelVisitNodeToQueue() *Queue.Queue{
	q := Queue.NewEmpty()
	i := 0
	for i = 0; ; i++ {
		if bt.FindRoot().LevelNodeToQueue(i, q) == 0 {
			break
		}
	}
	return q
}

//put all nodes into a array  order by level/layer/depth
func (bt *BinaryNode) LevelVisitNodeToArray()[]*BinaryNode{
	tmp :=bt.LevelVisitNodeToQueue().Data
	rs := make([]*BinaryNode,len(tmp))
	for i,v:=range tmp{
		rs[i] = v.(*BinaryNode)
	}
	return rs
}
//put all datas into a array order by level/layer/depth
func  (bt *BinaryNode) LevelVisitToArray() []interface{}{
	return bt.LevelVisitToQueue().Data
}

//put a single level's datas into a array
func  (bt *BinaryNode) LevelDataToArray(i int) []interface{}{
	q:=Queue.NewEmpty()
	bt.LevelDataToQueue(i,q)
	return q.Data
}

//node number of each level,rs[3]=5 represents  level 3 has 5 nodes
func (bt *BinaryNode) GetLevelNodeNumMap()map[int]int{
	var rs = make(map[int]int,0)
	depth,_,_,_:= bt.FindRoot().GetMaxDepth()
	arrTmp := make([]interface{},0)
	for i:=0;i<depth+1;i++{
		arrTmp = bt.LevelDataToArray(i)
		rs[i] = len(arrTmp)
	}
	return rs
}

//node number of all levels,for example [1,2,4,4,2] represents level 0 has 1 node,level 1 has 2 nodes,level 3 has 4 nodes...
func (bt *BinaryNode) GetLevelNodeNumArr()[]int{
	if bt==nil{
		return nil
	}
	depth,_,_,_:= bt.FindRoot().GetMaxDepth()
	var rs = make([]int,depth+1)
	arrTmp := make([]interface{},0)
	for i:=0;i<depth+1;i++{
		arrTmp = bt.LevelDataToArray(i)
		rs[i] = len(arrTmp)
	}
	return rs
}