package BTree

import "testing"
var node,node2,node3,node4,node5 *BinaryNode
func Init(){
	node=New(7)
	node.Insert(5)
	node.Insert(9)
	node.Insert(11)
	node.Insert(11)
	node.Insert(10)
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
	er:=node.Insert(15)
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