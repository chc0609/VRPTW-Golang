package main

import ( //导入包不使用会报错
	"go_code/project1/node"
	"go_code/project1/ga"
	"bufio"
	"fmt"
	"os"
	"strings"
)
//设置参数时 要记得设置datainitialization里的行数
func main()  {
	fileName :="D:/GoProj/src/go_code/project1/dataset/C1/C101.txt"                   //数据集所在目录
	vehicleCapacityFilename :=node.GetVehicleCapacityFilename(fileName) //获取数据集所在目录的capacity文件路径
	nodes :=node.LoadData(vehicleCapacityFilename,fileName)
	//for _, node := range nodes.List() {
	//	node.PrintInfo()
	//}
	ga.CreatIndividual(nodes)

}

func VcFilename(filename string) string {    //这函数是获得capacity的文件目录  具体是通过更换输入目录的尾缀
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

