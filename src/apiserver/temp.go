package main

import (
	"fmt"
)

func main()  {
	//LABLE:
	//for i := 0; i<10; i++{
	//	fmt.Println(i)
	//	for{
	//		break LABLE
	//		fmt.Println("good")
	//	}
	//}
	//fmt.Println('a')
	//fmt.Println("te4st")
	//a := [...]int{5,2,6,8,10}
	//fmt.Println(a)
	//num := len(a)
	//for i := 0; i < num; i++ {
	//	for j := i+1; j < num ; j++  {
	//		if a[i] < a [j] {
	//			temp := a[i]
	//			a[i] = a [j]
	//			a [j] = temp
	//		}
	//	}
	//}
	//fmt.Println(a)
	//fmt.Println(&a)

	a := make(map[string]int)
    a["test"] = 3
    a["good"] = 5
    a["wo"] = 6
    fmt.Println(a)
    go func() {
		fmt.Println("processs")
	}()
	//time.Sleep(time.Millisecond * 10) // this is bad, don't do this!
    for k,v := range a {
		fmt.Println(k)
		fmt.Println(v)
	}
	fmt.Println(a)

}
