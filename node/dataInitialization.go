package node

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
const NODENUM=101
func GetVehicleCapacityFilename(filename string) string { //这函数是获得capacity的文件目录  具体是通过更换输入目录的尾缀
	paths := strings.Split(filename, "/")
	paths[len(paths)-1] = "vehicle_capacity.txt"
	vcFilename := strings.Join(paths, "/")
	return vcFilename
}

func readFile(filename string, capacity int) []string {  //这里的capacity是切片的容量    把文件中的行都读过来
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File %s could not read: %v\n", filename, err)
		os.Exit(1)
	}

	defer f.Close()

	lines := make([]string, 0, capacity)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func trim(elements []string) []string {  //把数组里的空字符串去掉，只把数字拿出来
	values := make([]string, 0, 7)
	for _, element := range elements {
		if element != "" {
			values = append(values, element)
		}
	}
	return values
}
func creatNode(values []string,type_ int) *Node{   //创建节点
	//对输入数组读取每一位的属性
	id,_ :=strconv.Atoi(values[0])
	x,_ :=strconv.ParseFloat(values[1],64)
	y,_ :=strconv.ParseFloat(values[2],64)
	demand,_ :=strconv.ParseFloat(values[3],64)
	readyTime,_ :=strconv.ParseFloat(values[4],64)
	dueDate,_ :=strconv.ParseFloat(values[5],64)
	serviceTime,_ :=strconv.ParseFloat(values[6],64)

	node :=&Node{
		id:		id,
		x:		x,
		y:		y,
		type_:	type_,
		demand: demand,
		readyTime: readyTime,
		dueDate: dueDate,
		serviceTime: serviceTime,
	}
	return node
}
func LoadData(vehicleCapacityFilename string, nodeFilename string) *NodeList {
	vcapLines := readFile(vehicleCapacityFilename, 1)
	vehicleCapacity, _ := strconv.ParseFloat(vcapLines[0], 64)	//将字符串数组转化为数字

	nodes := &NodeList{   //相当于初始化所有节点集合
		list:            make([]*Node, 0, NODENUM),
		capacity:        vehicleCapacity,
		customersIDList: make([]int, 0, NODENUM-1),
	}
	nodeLines := readFile(nodeFilename, 150)
	// Skip header i = 0   该行为表头  所以跳过
	// 读取Depot  i=1 是仓库   i的含义是行索引
	line := nodeLines[1]
	elements := strings.Split(line, " ")   //分割后还存在这空字符串和数字
	values := trim(elements)
	depot :=creatNode(values, DEPOT)
	nodes.list =append(nodes.list,depot)
	nodes.depot = depot
	// Customers
	for i :=2;i<len(nodeLines);i++{
		line :=nodeLines[i]
		elements :=strings.Split(line," ")
		values :=trim(elements)
		customer :=creatNode(values,CUSTOMER)
		nodes.list=append(nodes.list,customer)
	}

	nodes.customers=nodes.list[1:]  //切片  返回从索引为1的元素返回索引到len(list)-1的元素
	for _,customer :=range nodes.customers{   //range 返回2个值 前一个是索引 后一个是存储的对象   如果返回1个值  那返回的就是索引
		id :=customer.ID()
		nodes.customersIDList=append(nodes.customersIDList,id)
	}
	return nodes
}