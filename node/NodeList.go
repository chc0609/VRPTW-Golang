package node

import "fmt"

type NodeList struct {    //集合
	list            []*Node //所有节点数组
	capacity        float64
	depot           *Node   //车场
	customers       []*Node
	customersIDList []int		//觉得没有存在的必要  暂时留着    顾客ID集合
}

func (n *NodeList) Depot() *Node { //返回起点
	return n.depot
}

func (n *NodeList) GetNode(targetID int) *Node { //输入ID 返回对应节点
	for i, node := range n.list {
		id := node.ID()
		if id == targetID {
			return n.list[i]
		}
	}
	return nil
}

func (n *NodeList) List() []*Node {   //使得外部不能直接取到list数组，只能通过该函数来得到
	return n.list
}

func (n *NodeList) Cusotmers() []*Node {		//得到顾客集合
	return n.customers
}

func (n *NodeList) CusotmersIDList() []int {
	return n.customersIDList
}

func (n *NodeList) Position(targetID int) (float64, float64) {
	for _, node := range n.list {
		id := node.ID()
		if id == targetID {
			return node.Position()
		}
	}
	return -1.0, -1.0
}
func (n *NodeList) GetCapacity() float64 { //得到
	return n.capacity
}

func (n *NodeList) PrintInfo() {
	fmt.Println("========================================")
	fmt.Println("List")
	for _, node := range n.list {
		node.PrintInfo()
	}
	fmt.Println("========================================")
	fmt.Println("Depot")
	n.Depot().PrintInfo()
	fmt.Println("========================================")
	fmt.Println("Customers")
	for _, node := range n.Cusotmers() {
		node.PrintInfo()
	}
	fmt.Println("========================================")
	fmt.Println("Customers ID List")
	fmt.Println(n.CusotmersIDList())
}