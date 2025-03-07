package main

import (
	"fmt"
	"math"
)

type Point1 struct { // 一维空间的点
	X int
}

type Centroid1 struct {
	ID     int
	Pos    Point1
	Points []Point1
}

func Distance1(p1, p2 Point1) int {
	return int(math.Abs(float64(p1.X - p2.X))) // 使用平方差计算距离，可以避免负数的出现
}

func GeneratePoints(num int) []Point1 {
	points := make([]Point1, num)
	for i := range points {
		points[i] = Point1{X: i} // 生成随机的X坐标
	}
	return points
}

// InitializeCentroids1 初始化簇中心
func InitializeCentroids1(points []Point1, k int) []Centroid1 {
	centroids := make([]Centroid1, k)
	// 这里初始化簇中心很简单，直接取前k个点作为簇中心，就是为了证明簇中心选择不好的话，结果会很差
	for i := 1; i <= k; {
		centroids[i-1] = Centroid1{ID: i, Pos: points[i]}
		i++
	}
	return centroids
}

// AssignPointsToCentroids1 分配数据点到最近的簇中心
func AssignPointsToCentroids1(points []Point1, centroids []Centroid1) {
	for i := range points {
		minDistance := Distance1(points[i], centroids[0].Pos)
		closestCentroidID := 0
		for j := 1; j < len(centroids); j++ {
			dist := Distance1(points[i], centroids[j].Pos)
			if dist < minDistance {
				minDistance = dist
				closestCentroidID = j
			}
		}
		centroids[closestCentroidID].Points = append(centroids[closestCentroidID].Points, points[i])
	}
}

// UpdateCentroids1 更新簇中心为簇内点的平均值
func UpdateCentroids1(centroids []Centroid1) {
	for i := range centroids {
		if len(centroids[i].Points) > 0 {
			var sumX int
			for _, p := range centroids[i].Points {
				sumX += p.X
			}
			centroids[i].Pos = Point1{X: sumX / len(centroids[i].Points)}
		}
	}
}

// CentroidsChanged1 检查簇中心是否变化
func CentroidsChanged1(oldCentroids, newCentroids []Centroid1) bool {
	for i := range oldCentroids {
		if Distance1(oldCentroids[i].Pos, newCentroids[i].Pos) > 0 {
			return true
		}
	}
	return false
}

// CleanCentroids 清空簇中心的点列表以便重新分配
func CleanCentroids(centroids []Centroid1) {
	for i := range centroids {
		centroids[i].Points = []Point1{}
	}
}

// PrintClusters 打印簇信息
func PrintClusters(centroids []Centroid1) {
	for _, centroid := range centroids {
		fmt.Printf("Cluster %d (Center: %d): ", centroid.ID, centroid.Pos.X)
		for _, point := range centroid.Points {
			fmt.Printf("(%d) ", point.X)
		}
		fmt.Println()
	}
}

func main() {
	k := 10

	// 生成随机数据点
	points := GeneratePoints(100)

	// 初始化簇中心
	centroids := InitializeCentroids1(points, k)

	// 迭代直到簇中心不再变化
	for {
		oldCentroids := make([]Centroid1, len(centroids))
		copy(oldCentroids, centroids)

		// 分配点到最近的簇中心
		AssignPointsToCentroids1(points, centroids)
		PrintClusters(centroids)
		println("--------------------------")

		// 更新簇中心
		UpdateCentroids1(centroids)

		// 检查簇中心是否变化
		if !CentroidsChanged1(oldCentroids, centroids) {
			break
		}
		CleanCentroids(centroids)
	}

	// 打印最终的簇信息
	PrintClusters(centroids)
}
