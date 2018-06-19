package BTree

import (
	"reflect"
	"errors"
)

type BinaryNode struct {
	Data   interface{}
	Parant *BinaryNode
	Left   *BinaryNode
	Right  *BinaryNode
}

type AVL struct {
	root *BinaryNode
}

type RedBlack struct {
	root *BinaryNode
}

//new a node
func New(data interface{}) *BinaryNode {
	return &BinaryNode{data, nil, nil, nil}
}
func NewWithParent(data interface{}, prt *BinaryNode) *BinaryNode {
	return &BinaryNode{data, prt, nil, nil}
}

//if a bt contains data
func (bt *BinaryNode) Contain(data interface{}) (bool, error) {
	if data == nil {
		return false, nil
	}
	rs, err := Compare(bt.Data, data)
	if err != nil {
		return false, err
	}

	if rs == 0 {
		return true, nil
	}
	if rs == 1 {
		return bt.Left.Contain(data)
	}
	if rs == -1 {
		return bt.Right.Contain(data)
	}
	return false, nil
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
		bt.Right, _ = bt.Right.Remove(tmp.Data)
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
func (bt *BinaryNode) GetMaxDepth() (int,int,int,error) {
	if bt == nil {
		return 0,0,0, nil
	}
	var leftMax, rightMax int
	var er error
	if bt.Left != nil {
		leftMax,_,_, er = bt.Left.GetMaxDepth()
		if er != nil {
			return -1,-1,-1, er
		}
		leftMax++
	}
	if bt.Right != nil {
		rightMax,_,_, er = bt.Right.GetMaxDepth()
		if er != nil {
			return -1,-1,-1, er
		}
		rightMax++
	}
	if leftMax > rightMax {
		return leftMax,leftMax,rightMax, nil
	}
	return rightMax,leftMax,rightMax, nil

}

func (bt *BinaryNode) GetMaxDistance() (int, error) {
	if bt == nil {
		return 0, nil
	}
	var lMax1 = 0
	var rMax1 = 0
	var lMax2 = 0
	var rMax2 = 0
	var er error

	_,lMax1,rMax1,er=bt.GetMaxDepth()
	if bt.Left != nil {
		lMax2, er = bt.Left.GetMaxDistance()
		if er != nil {
			return -1, er
		}
		if bt.Right!=nil{
			lMax2++
		}
	}
	if bt.Right != nil {

		rMax2, er = bt.Right.GetMaxDistance()
		if er != nil {
			return -1, er
		}
		if bt.Left!=nil{
			rMax2++
		}
	}
	return max((lMax1+rMax1), lMax2, rMax2), nil
}

//transfer a bt to a redBlack pattern
//what is a red-black tree,refer to www.baidu.com, www.google.com
func ToRedBlack() (*BinaryNode, error) {
	return nil, nil
}

//transfer a bt to an avl parttern
//what is an avl tree,refer to www.baidu.com,www.google.com
func ToAVL() (*BinaryNode, error) {
	return nil, nil
}

//rotate left
func (av *RedBlack) LeftRotate() (*AVL, error) {
	return nil, nil
}

//rotate right
func (av *RedBlack) RightRotate() (*AVL, error) {
	return nil, nil
}

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
