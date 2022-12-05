package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello,World")

	/** 画出空心金字塔
	var totalLevel int = 20
	for i:=1;i<=totalLevel;i++ {
		for k := 0; k < totalLevel - i; k++ {
			fmt.Print(" ")
		}
		for j:=1;j<=2*i-1;j++ {
			if (j == 1 || j == 2 * i-1 || i == totalLevel) {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
			// fmt.Print("*")
		}
		fmt.Println()
	}
	**/

	/** 画出九九乘法表
	for i:=1;i<=9;i++ {
		for j:=1;j<=i;j++ {
			fmt.Printf("%d * %d = %d\t",i,j,i*j)
		}
		fmt.Println()
	}
	**/

	
	var count int = 0
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		// fmt.Println(n)
		count ++
		if (n == 99) {
			break
		}
	}
	fmt.Println("生成 99 共循环了",count,"次")
}
