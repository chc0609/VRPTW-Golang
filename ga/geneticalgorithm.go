package ga

import (
	"go_code/project1/node"
	"math"
	"math/rand"
	"time"
)

func shuffle(data []int) {   //基因打乱
	rand.Seed(time.Now().UnixNano())
	for i := len(data)-1;i>0;i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}
func IsFeasible(nodes *node.NodeList,route []int) bool {		//判断一个解是否可行
	// Capacity check
	var amount float64 = 0.0
	for _, nodeID :=range route{
		node :=nodes.GetNode(nodeID)
		demand :=node.Demand()
		amount +=demand
		if amount> nodes.GetCapacity(){
			return false
		}
	}
	//Time Window check
	startTime :=nodes.Depot().ReadyTime()
	var t float64 = startTime   //初始为仓库的开始工作时间
	for _,nodeID :=range route{
		node :=nodes.GetNode(nodeID)
		readyTime :=node.ReadyTime()
		dueDate :=node.DueDate()
		if t>dueDate {
			return false
		}
		t=math.Max(t,readyTime)+node.ServiceTime()	//这里是把路径时间和服务时间都加在一起了
		//计算下一点的到达时间
	}
	depot :=nodes.Depot()
	if t>depot.DueDate() {
		return false
	}
	return true
}
func shapeFlatToVehicles(nodes *node.NodeList, flattench []int) [][]int {    //生成初始解  就是对传进来的flattench按是否可行进行切片
	chromosome := make([][]int, 0)
	size := len(flattench)
	var cut1, cut2 int = 0, 0
	for cut1 < size {
		breakFlag := false
		route := make([]int, 0, size)
		for cut2 = cut1; cut2 < size+1; cut2++ {
			route = flattench[cut1:cut2]
			if !IsFeasible(nodes,route) {
				cut1 = cut2 - 1
				route = route[:len(route)-1]
				breakFlag = true
				break
			}
		}
		if !breakFlag {
			cut1 = cut2
		}
		chromosome = append(chromosome, route)
	}
	return chromosome
}