package main


import (  //导入包不使用会报错
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main()  {
	zz()
}

func packges(){  //随机数案例
	rand.Seed(time.Now().Unix())
	for i := 0; i < 3; i++  {
		fmt.Println(rand.Intn(100))
	}
}
func imports(){
	fmt.Printf("Now tou have %g problems.\n",math.Sqrt(9))
}

func add(x int ,y int ) int{  //类型声明在变量之后
	return x+y
}
func add2(x,y int) int{   //多参数声明
	return x+y
}

func swap(x,y string) (string,string){   //多返回值
	return y,x
}

func split(sum int) (x,y int){  //命名返回值
	//Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。
	x=sum*4/9
	y=sum-x
	return
}
var c,python,java bool
func variable(){
	var i int
	fmt.Println(i,c,python,java)
}
var i,j int=1,2
func intitializer()  {
	//变量声明可以包含初始值，每个变量对应一个。
	//如果初始化值已存在，则可以省略类型；变量会从初始值中获得类型。
	var c,python,java=true,false,"no!"
	fmt.Println(i,j,c,python,java)

}
 func short_declaration(){
 	//在函数中，简洁赋值语句 := 可在类型明确的地方代替 var 声明。
	 //函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用。
 	var i,j int=1,2
 	k :=3
 	c,python,java :=true,false,"No!"
 	fmt.Println(i,j,k,c,python,java)
 }

 func type_conversion(){  //类型转换   go不支持自动转换
 	var x,y int=3,4
 	var f =math.Sqrt(float64(x*x+y*y))
 	var z uint=uint(f);
 	fmt.Println(x,y,z)
 }

 func type_inference(){   //类型推导
	 v :=0.876+0.5i
 	fmt.Printf("v is type of %T\n",v)
 }

 const Pi=3.14
func constants(){
	//常量的声明与变量类似，只不过是使用 const 关键字。
	//常量可以是字符、字符串、布尔值或数值。
	//常量不能用 := 语法声明。
	const world ="世界"
	fmt.Println("Hello",world)
	fmt.Println("Happy",Pi,"Day")

	const Truth=true
	fmt.Println("Go rules?",Truth)
}

 func forxh(){  //for循环  三项没有要求都有
 	sum :=0
 	for i:=0;i<10;i++{
 		sum+=i
	}
	fmt.Println(sum)
 }

 func forxh2(){  //go中的while循环可以用for来代替
 	sum :=1
 	for sum<1000{
 		sum+=sum
	}
	fmt.Println(sum)
 }

 func sqrt_if(x float64) string{   //if语句
 	if x< 0{
 		return sqrt_if(-x)+"i"
	}
	return fmt.Sprint(math.Sqrt(x))
 }

 func switch_order(){
 	//没有条件的 switch
	 //没有条件的 switch 同 switch true 一样。
	 //这种形式能将一长串 if-then-else 写得更加清晰。
	 t := time.Now()
	 switch {
	 case t.Hour() < 12:
		 fmt.Println("Good morning!")
	 case t.Hour() < 17:
		 fmt.Println("Good afternoon.")
	 default:
		 fmt.Println("Good evening.")
	 }
 }

 func zz(){  //指针
 	var i int =10
 	//i的地址是什么
 	fmt.Println("i的地址",&i)
 	var ptr *int = &i
 	//1.ptr是一个指针变量   用*获取指针指向的值   &获取变量的地址
 	//2.ptr 的类型是*int
 	//3.ptr 本身的的值&i
 	fmt.Println("ptr=%v\n",ptr)
 	fmt.Println("ptr的地址=%v",&ptr)
 	fmt.Println("ptr指向的值=%v",*ptr)
 }
