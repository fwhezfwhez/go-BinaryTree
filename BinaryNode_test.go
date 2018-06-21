package BTree

import "testing"
var node,node2,node3,node4,node5,nodei *BinaryNode
func Init(){
	node=New(7)
	node2,_ =node.Insert(5)
	node3,_ =node.Insert(9)
	node4,_ =node.Insert(11)
	node5,_ =node.Insert(11)
	node.Insert(10)
	node.Insert(8)
	nodei,_=node.Insert(1)
	node.Insert(6)
	//node2=New(5)
	//node3=New(9)
	//node4=New(11)
	//node5=New(11)
	//node.Left=node2
	//node.Right=node3
	//node3.Right=node4
	//node4.Left =node5
}
func TestCompare(t *testing.T) {
	var k interface{} = 1
	t.Log(Compare(k,1))
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
	_,er:=node.Insert(15)
	if er!=nil{
		t.Fatal(er)
	}
	t.Log(node.FindMax())

}
func TestBinaryNode_Remove(t *testing.T) {
	Init()
	node,er:=node.Remove(11)
	node,er=node.Remove(11)
	//er=node.Remove(11)
	//er=node.Remove(11)
	if er!=nil{
		t.Fatal(er.Error())
	}
	t.Log(node.FindMax())
}
func TestGetAncestor(t *testing.T) {
	Init()
	n1 := New(8)
	n2 :=New(11)
	t.Log(node.GetAncestor(n1,n2))
}
func TestBinaryNode_GetDepth(t *testing.T) {
	Init()
	t.Log(node.GetDepth(node4))
}

func TestBinaryNode_GetDistance(t *testing.T) {
	Init()
	t.Log(node.GetDistance(New(1),New(11)))
}

func TestBinaryNode_GetMaxDepth(t *testing.T) {
	Init()
	t.Log(node.GetMaxDepth())
}
func TestBinaryNode_GetMaxDistance(t *testing.T) {
	Init()
	t.Log(node.GetMaxDistance())
}
func TestMmax(t *testing.T){
	t.Log(max(6,3,4))
}

func TestNodeNum(t *testing.T){
	Init()
	var sum int
	t.Log(node.GetNodesNum(&sum))
}

func TestMaxHeight(t *testing.T){
	Init()
	t.Log(node.GetMaxHeight())
	t.Log(nodei.Right)
}

func TestBinaryNode_GetNodeHeight(t *testing.T) {
	Init()
	t.Log(node.GetNodeHeight(node3))
}

func TestToSortLinkedListStruct(t *testing.T) {
	Init()
	rs,er:=node.ToAscLinkedList()
	if er!=nil {
		t.Fatal(er)
	}

	for ;;{
		t.Log(rs.Start.Data)
		rs.Start = rs.Start.Right
		if rs.Start ==nil{
			break
		}
	}
}

func TestBinaryNode_ToDescLinkedList(t *testing.T) {
	Init()
	rs,er:=node.ToDescLinkedList()
	if er!=nil {
		t.Fatal(er)
	}

	for ;;{
		t.Log(rs.Start.Data)
		rs.Start = rs.Start.Right
		if rs.Start ==nil{
			break
		}
	}
}

func TestBinaryNode_ToAscArray(t *testing.T) {
	Init()
	var sum = 0
	var rs = make([]interface{},node.GetNodesNum(&sum))
	var tmp =0
	node.ToAscArray(&rs,&tmp)
	t.Log(rs)
	t.Log(tmp)
}

func TestBinaryNode_ToAscArrayEscapingArgs(t *testing.T) {
	Init()
	t.Log(node.ToAscArrayEscapingArgs())
}

func TestBinaryNode_ToDescArrayEscapingArgs(t *testing.T) {
	Init()
	t.Log(node.ToDescArrayEscapingArgs())
}

func TestBinaryNode_ToDescArray(t *testing.T) {
	Init()
	var sum = 0
	var rs = make([]interface{},node.GetNodesNum(&sum))
	var tmp =0
	node.ToDescArray(&rs,&tmp)
	t.Log(rs)
	t.Log(tmp)
}

func TestBinaryNode_FindRoot(t *testing.T) {
	Init()
	t.Log(node3.FindRoot())
}

func TestSmartPrint(t *testing.T) {
	type User struct {
		Name string
		Id int
		P string
	}
	SmartPrint(User{"ft",5,""})
}

func TestBinaryNode_Cached(t *testing.T) {
	Init()
	node.Cached()
	rs,er:=node.GetCache()
	if er!=nil {
		t.Fatal(er.Error())
	}
	SmartPrint(*rs)
}