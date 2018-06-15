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
func (bt *BinaryNode) Insert(x interface{}) (error) {
	rs, err := Compare(bt.Data, x)
	if err != nil {
		return err
	}

	if rs == 1 {
		if bt.Left != nil {
			return bt.Left.Insert(x)
		} else {
			bt.Left = NewWithParent(x, bt)
		}
	} else {
		if bt.Right != nil {
			return bt.Right.Insert(x)
		} else {
			bt.Right = NewWithParent(x, bt)
		}

	}
	return nil
}

//delete a node which's data is x,if several nodes' data are all x,only delete one of them
func (bt *BinaryNode) Remove(x interface{}) (*BinaryNode,error) {
	if bt ==nil {
		return nil,nil
	}
	rs,er:=Compare(bt.Data,x)
	if er!=nil{
		return nil,er
	}
	if rs ==-1 {
		bt.Right,er= bt.Right.Remove(x)
		if er!=nil{
			return nil,er
		}
	}else if rs==1 {
		bt.Left,er = bt.Left.Remove(x)
		if er!=nil{
			return nil,er
		}
	}else if bt.Left!=nil && bt.Right!=nil {
		tmp,_:=bt.Right.FindMin()
		bt.Right,_ = bt.Right.Remove(tmp.Data)
	}else{
		if bt.Left!=nil{
			bt = bt.Left
		}else{
			bt=bt.Right
		}
	}

	return bt,nil
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
