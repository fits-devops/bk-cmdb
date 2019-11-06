package main
// fmt 的别名
import io "fmt"

// 常量
const PI,CITY = 3.14,"广州"
// 全局变量 双引号
var name = "超付"

func main()  {
	a := 900
	// 打印函数 可见性 P 大写表示 public
	io.Println("test 你好 世界")
	io.Println(PI)
	io.Println(name)
	io.Println(a)
	io.Println(CITY)
}