package BTree

import (
	"testing"
	"github.com/fwhezfwhez/go-queue"
)

var node, node2, node3, node4, node5, nodei *BinaryNode

func Init() {
	node = New(7)
	node2, _ = node.Insert(5)
	node3, _ = node.Insert(9)
	//node4,_ =node.Insert(11)
	//node5,_ =node.Insert(11)
	//node.Insert(10)
	//node.Insert(8)
	//nodei,_=node.Insert(1)
	//node.Insert(6)
	//node.Insert(0)
	node.Insert(3)
	node.Insert(6)
	node.Insert(1)
	node.Insert(4)
	node.Insert(0)
	node.Insert(2)
	node.Insert(11)
	node.Insert(13)
	node.Insert(10)
	node.Insert(8)
	//node2=New(5)
	//node3=New(9)
	//node4=New(11)
	//node5=New(11)
	//node.Left=node2
	//node.Right=node3
	//node3.Right=node4
	//node4.Left =node5
}
func Test_Compare(t *testing.T) {
	var k interface{} = 1
	t.Log(Compare(k, 1))
}

func TestBinaryNode_Contain(t *testing.T) {
	Init()
	t.Log(node.Contain(9))
}

func TestBinaryNode_FindMax(t *testing.T) {
	Init()
	t.Log(node.FindMax())
}

func TestBinaryNode_FindMin(t *testing.T) {
	Init()
	t.Log(node.FindMin())
}

func TestBinaryNode_Insert(t *testing.T) {
	Init()
	node.Insert(1)
	_, er := node.Insert(15)
	if er != nil {
		t.Fatal(er)
	}
	t.Log(node.FindMax())

}
func TestBinaryNode_Remove(t *testing.T) {
	Init()
	node, er := node.Remove(11)
	node, er = node.Remove(11)
	//er=node.Remove(11)
	//er=node.Remove(11)
	if er != nil {
		t.Fatal(er.Error())
	}
	t.Log(node.FindMax())
}
func TestBinaryNode_GetAncestor(t *testing.T) {
	Init()
	n1 := New(8)
	n2 := New(11)
	t.Log(node.GetAncestor(n1, n2))
}
func TestBinaryNode_GetDepth(t *testing.T) {
	Init()
	t.Log(node.GetDepth(node4))
}

func TestBinaryNode_GetDistance(t *testing.T) {
	Init()
	t.Log(node.GetDistance(New(1), New(11)))
}

func TestBinaryNode_GetMaxDepth(t *testing.T) {
	Init()
	t.Log(node.GetMaxDepth())
}
func TestBinaryNode_GetMaxDistance(t *testing.T) {
	Init()
	t.Log(node.GetMaxDistance())
}
func TestMmax(t *testing.T) {
	t.Log(max(6, 3, 4))
}

func TestBinaryNode_NodeNum(t *testing.T) {
	Init()
	var sum int
	node.GetNodesNum(&sum)
	t.Log(sum)
}

func TestBinaryNode_GetMaxHeight(t *testing.T) {
	Init()
	t.Log(node.GetMaxHeight())
	t.Log(nodei.Right)
}

func TestBinaryNode_GetNodeHeight(t *testing.T) {
	Init()
	t.Log(node.GetNodeHeight(node3))
}

func TestBinaryNode_ToSortLinkedListStruct(t *testing.T) {
	Init()
	rs, er := node.ToAscLinkedList()
	if er != nil {
		t.Fatal(er)
	}

	for ; ; {
		t.Log(rs.Start.Data)
		rs.Start = rs.Start.Right
		if rs.Start == nil {
			break
		}
	}
}

func TestBinaryNode_ToDescLinkedList(t *testing.T) {
	Init()
	rs, er := node.ToDescLinkedList()
	if er != nil {
		t.Fatal(er)
	}

	for ; ; {
		t.Log(rs.Start.Data)
		rs.Start = rs.Start.Right
		if rs.Start == nil {
			break
		}
	}
}

func TestBinaryNode_ToAscArray(t *testing.T) {
	Init()
	var sum = 0
	node.GetNodesNum(&sum)
	var rs = make([]interface{}, sum)
	var tmp = 0
	node.ToAscArray(&rs, &tmp)
	t.Log(rs)
	t.Log(tmp)
}

func TestBinaryNode_ToAscArrayEscapingArgs(t *testing.T) {
	Init()
	node.Remove(5)
	t.Log(node.ToAscArrayEscapingArgs())
}

func TestBinaryNode_ToDescArrayEscapingArgs(t *testing.T) {
	Init()
	node.Remove(9)
	t.Log(node.ToDescArrayEscapingArgs())
}

func TestBinaryNode_ToDescArray(t *testing.T) {
	Init()
	var sum = 0
	node.GetNodesNum(&sum)
	var rs = make([]interface{}, sum)
	var tmp = 0
	node.ToDescArray(&rs, &tmp)
	t.Log(rs)
	t.Log(tmp)
}

func TestBinaryNode_FindRoot(t *testing.T) {
	Init()
	t.Log(node3.FindRoot())
}

func TestBinaryNode_SmartPrint(t *testing.T) {
	type User struct {
		Name string
		Id   int
		P    string
	}
	SmartPrint(User{"ft", 5, ""})
}

func TestBinaryNode_Cached(t *testing.T) {
	Init()
	node.Cached()
	rs, er := node.GetCache()
	if er != nil {
		t.Fatal(er.Error())
	}
	SmartPrint(*rs)
}
func TestBinaryNode_Abs(t *testing.T) {
	t.Log(abs(2))
	t.Log(abs(-2))
}

func TestBinaryNode_RotateWithLeftChild(t *testing.T) {
	Init()
	t.Log(node.RotateWithLeftChild())
}

func TestBinaryNode_RotateWithRightChild(t *testing.T) {
	Init()
	t.Log(node.RotateWithLeftChild().RotateWithRightChild())
}

func TestBinaryNode_DoubleRotateLeftChild(t *testing.T) {
	Init()
	t.Log(node.DoubleRotateLeftChild())
}

func TestBinaryNode_Balance(t *testing.T) {
	Init()
	rs := node.Balance()
	t.Log(rs)
	t.Log(rs.Left)
	//t.Log(node.Balance().ToDescArrayEscapingArgs())
}
func TestBinaryNode_CopyBalance(t *testing.T) {
	Init()
	//rs := node.BalanceCopy()
	//t.Log(rs.ToDescArrayEscapingArgs())
	t.Log(node.BalanceCopy().ToDescArrayEscapingArgs())
}

func TestBinaryNode_ToAVL(t *testing.T) {
	Init()
	t.Log(node.ToAVL().Root)
}

func TestBinaryNode_GetLevelData(t *testing.T) {
	Init()
	q := Queue.NewEmpty()
	node.LevelDataToQueue(3, q)
	q.Print()
}

func TestBinaryNode_LevelVisitToQueue(t *testing.T) {
	Init()
	q:=node.LevelVisitToQueue()
	q.Print()
}

func TestBinaryNode_LevelVisitToArray(t *testing.T) {
	Init()
	q:=node.LevelVisitToArray()
	t.Log(q)
}

func TestBinaryNode_LevelDataToArray(t *testing.T) {
	Init()
	arr:=node.LevelDataToArray(4)
	t.Log(arr)
}

func TestBinaryNode_GetLevelNodeNumMap(t *testing.T) {
	Init()
	rs:=node.GetLevelNodeNumMap()
	t.Log(rs)
}

func TestBinaryNode_GetLevelNodeNumArr(t *testing.T) {
	Init()
	rs:=node.GetLevelNodeNumArr()
	t.Log(rs)
}

func TestBinaryNode_Paint(t *testing.T) {
	Init()
	node.Paint()
}

func TestBinaryNode_LevelNodeToQueue(t *testing.T) {
	Init()
	q:= Queue.NewEmpty()
	node.LevelNodeToQueue(0,q)
	q.Print()
}

func TestBinaryNode_LevelVisitNodeToQueue(t *testing.T) {
	Init()
	q:=node.LevelVisitNodeToQueue()
	q.Print()
}

func TestBinaryNode_LevelVisitNodeToArray(t *testing.T) {
	Init()
	arr:=node.LevelVisitNodeToArray()
	t.Log(arr)
}