package main

import (
	"fmt"

	"github.com/JackSliwen/go_UtilityFunc/calculate"
	"github.com/JackSliwen/go_UtilityFunc/file"
)

func main() {
	// 输出hello world
	file.PathExists("")
	r1 := &calculate.Rect{
		X1: 55,
		Y1: 200,
		X2: 200,
		Y2: 300,
	}
	r2 := &calculate.Rect{
		X1: 100,
		Y1: 100,
		X2: 200,
		Y2: 300,
	}
	//image.DrawRectOnImageAndSave()
	fmt.Println(calculate.CompareDistance(r1, r2))
	fmt.Println(calculate.CompareIoU(r1, r2))
}
