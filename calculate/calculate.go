package calculate

import (
	"math"
)

type Rect struct {
	X1, Y1, X2, Y2 int32
}

//min 返回两数最小值
func min(x, y int32) int32 {
	if x < y {
		return x
	}
	return y
}

//max 返回两数最大值
func max(x, y int32) int32 {
	if x > y {
		return x
	}
	return y
}

//CompareDistance 计算两个框中心点的距离
func CompareDistance(R1 *Rect, R2 *Rect) (ret float64) {
	X1 := (R1.X1 + R1.X2) / 2
	Y1 := (R1.Y1 + R1.Y2) / 2
	X2 := (R2.X1 + R2.X2) / 2
	Y2 := (R2.Y1 + R2.Y2) / 2
	dis2 := math.Pow(float64(X1-X2), 2) + math.Pow(float64(Y1-Y2), 2)
	return math.Abs(math.Sqrt(float64(dis2)))
}

//CompareIoU 计算两个框重合比例
func CompareIoU(R1 *Rect, R2 *Rect) (ret float64) {
	var iou float64
	X1 := max(R1.X1, R2.X1)
	X2 := min(R1.X2, R2.X2)
	Y1 := max(R1.Y1, R2.Y1)
	Y2 := min(R1.Y2, R2.Y2)
	area := max((X2-X1+1), 0) * max(0, (Y2-Y1+1))
	reta := ((R1.X2-R1.X1+1)*(R1.Y2-R1.Y1+1) + (R2.X2-R2.X1+1)*(R2.Y2-R2.Y1+1) - area)
	iou = float64(area) / float64(reta)
	return iou
}
