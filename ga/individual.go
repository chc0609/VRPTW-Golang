package ga

import (
	"fmt"
	"go_code/project1/node"
)

type Individual struct{
	Chromosome [][]int   //应该是用第一维表示车辆   第二维表示路径
	Distance   float64
	Fitness    float64
}

//***********************//
// Methods of Individual //
//***********************//
//都是公共的，外部可以调用
func (indv *Individual) GetVehicleNum() int{
 	vehicleNum :=len(indv.Chromosome)
 	return vehicleNum
}

func CreatIndividual(nodes *node.NodeList) *Individual{
	flattench := make([]int, len(nodes.CusotmersIDList()))  //扁平化
	chromosome := make([][]int, 0)		//二维染色体
	customersIDList := nodes.CusotmersIDList()
	copy(flattench, customersIDList)
	shuffle(flattench)		//打乱顾客顺序
	chromosome = shapeFlatToVehicles(nodes, flattench)    //生成根据一段顾客序列 转化为二维的初始染色体  第一维表示车辆  第二维表示执行路径
	fmt.Println(chromosome)
	indv := &Individual{Chromosome: make([][]int, 0)}
	return nil
}

