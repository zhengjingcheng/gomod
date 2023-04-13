package main

import "fmt"

type People struct {
	id   int
	name string
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}          // 定义数组
	s := a[1:3]                         // 定义切片
	fmt.Println("before modify:", s, a) // before modify: [2 3] [1 2 3 4 5]
	s[0] = 6                            // 修改切片
	fmt.Println("after modify:", s, a)  // after modify: [6 3] [1 6 3 4 5]
	b := [...]int{1, 2, 3, 4, 5}        // 定义数组
	t := b                              // 定义切片
	fmt.Println("before modify:", t, b) // before modify: [2 3] [1 2 3 4 5]
	t[0] = 6                            // 尝试修改切片
	fmt.Println("after modify:", t, b)  // after modify: [6 3] [1 2 3 4 5]
	//数组是值类型，切片是引用类型
}

//var wg sync.WaitGroup
//wg.Add(2)
//ch := make(chan int)
//go func() {
//	for i := 1; i < 100; i += 2 {
//		fmt.Println(i)
//		ch <- i
//		<-ch
//	}
//	wg.Done()
//
//
//go func() {
//	for i := 2; i < 100; i += 2 {
//		fmt.Println(i)
//		<-ch
//		ch <- i
//	}
//	wg.Done()
//}()
//wg.Wait()
