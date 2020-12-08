package node

import "fmt"
const(
	DEPOT=0		//创建节点时用来控制创建的节点类型是车场还是顾客
	CUSTOMER=1
)
type Node struct{
	id int
	x 			float64
	y 			float64
	type_       int			//类型  该节点是否为车场
	demand      float64
	readyTime 	float64		//最早开始时间
	dueDate		float64		//最迟开始时间
	serviceTime 	float64
}

func (n *Node) ID() int { //前面括号表示当前结构体简称   返回节点ID
	return n.id
}

func (n *Node) Type() int {
	return n.type_
}

func (n *Node) Position() (float64, float64) { //返回横纵坐标
	return n.x, n.y
}

func (n *Node) Demand() float64 {
	return n.demand
}

func (n *Node) ReadyTime() float64 {
	return n.readyTime
}

func (n *Node) DueDate() float64 {
	return n.dueDate
}

func (n *Node) ServiceTime() float64 {
	return n.serviceTime
}

func (n *Node) PrintInfo() {
	fmt.Println("id ", n.id, "type ", n.type_, "pos ", n.x, n.y,
		"dem ", n.demand, "ready ", n.readyTime,
		"due ", n.dueDate, "servicetime ", n.serviceTime)
}