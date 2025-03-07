package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Point 表示一个二维空间中的点
type Point struct {
	X, Y float64
}

// Centroid 表示一个簇的中心
type Centroid struct {
	ID     int
	Pos    Point
	Points []Point
}

// Distance 计算两个点之间的欧氏距离
func Distance(p1, p2 Point) float64 {
	return (p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y)
}

// GenerateRandomPoints 生成随机数据点
func GenerateRandomPoints(num int, min, max float64) []Point {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	points := make([]Point, num)
	for i := range points {
		x := min + r.Float64()*(max-min)
		y := min + r.Float64()*(max-min)
		points[i] = Point{X: x, Y: y}
	}
	return points
}

// InitializeCentroids 初始化簇中心
func InitializeCentroids(points []Point, k int) []Centroid {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	selectedIndices := make(map[int]bool)
	centroids := make([]Centroid, k)
	for i := 0; i < k; {
		index := r.Intn(len(points))
		if !selectedIndices[index] {
			selectedIndices[index] = true
			centroids[i] = Centroid{ID: i, Pos: points[index]}
			i++
		}
	}
	return centroids
}

// AssignPointsToCentroids 分配数据点到最近的簇中心
func AssignPointsToCentroids(points []Point, centroids []Centroid) {
	for i := range points {
		minDistance := Distance(points[i], centroids[0].Pos)
		closestCentroidID := 0
		for j := 1; j < len(centroids); j++ {
			dist := Distance(points[i], centroids[j].Pos)
			if dist < minDistance {
				minDistance = dist
				closestCentroidID = j
			}
		}
		centroids[closestCentroidID].Points = append(centroids[closestCentroidID].Points, points[i])
	}
}

// UpdateCentroids 更新簇中心为簇内点的平均值
func UpdateCentroids(centroids []Centroid) {
	for i := range centroids {
		if len(centroids[i].Points) > 0 {
			var sumX, sumY float64
			for _, p := range centroids[i].Points {
				sumX += p.X
				sumY += p.Y
			}
			centroids[i].Pos = Point{X: sumX / float64(len(centroids[i].Points)), Y: sumY / float64(len(centroids[i].Points))}
		}
	}
}

// CleanCentroids 清空簇中心的点列表以便重新分配
func CleanCentroids(centroids []Centroid) {
	for i := range centroids {
		centroids[i].Points = []Point{}
	}
}

// CentroidsChanged 检查簇中心是否变化
func CentroidsChanged(oldCentroids, newCentroids []Centroid) bool {
	for i := range oldCentroids {
		if Distance(oldCentroids[i].Pos, newCentroids[i].Pos) > 1e-6 {
			return true
		}
	}
	return false
}

// PrintClusters 打印簇信息
func PrintClusters(centroids []Centroid) {
	for _, centroid := range centroids {
		fmt.Printf("Cluster %d (Center: %.2f, %.2f): ", centroid.ID, centroid.Pos.X, centroid.Pos.Y)
		for _, point := range centroid.Points {
			fmt.Printf("(%.2f, %.2f) ", point.X, point.Y)
		}
		fmt.Println()
	}
}

func main() {
	numPoints := 50
	k := 3
	minValue := 0.0
	maxValue := 10.0

	// 生成随机数据点
	points := GenerateRandomPoints(numPoints, minValue, maxValue)

	// 初始化簇中心
	centroids := InitializeCentroids(points, k)

	// 迭代直到簇中心不再变化
	for {
		oldCentroids := make([]Centroid, len(centroids))
		copy(oldCentroids, centroids)

		// 分配点到最近的簇中心
		AssignPointsToCentroids(points, centroids)

		// 更新簇中心
		UpdateCentroids(centroids)

		// 检查簇中心是否变化
		if !CentroidsChanged(oldCentroids, centroids) {
			break
		}
		CleanCentroids(centroids)
	}

	// 打印最终的簇信息
	PrintClusters(centroids)
}
