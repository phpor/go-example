package main

import (
	"fmt"
	"math"
	"sort"
)

// Point 表示一个二维数据点及其类别
type Point struct {
	X, Y     float64
	Category string
}

// EuclideanDistance 计算两个点之间的欧几里得距离
func EuclideanDistance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p1.X-p2.X, 2) + math.Pow(p1.Y-p2.Y, 2))
}

// KNNClassifier 实现KNN分类器
type KNNClassifier struct {
	Data []Point
	K    int
}

// Predict 预测新样本的类别
func (knn *KNNClassifier) Predict(newPoint Point) string {
	// 计算所有点到新样本的距离
	distances := make([]struct {
		Point    Point
		Distance float64
	}, len(knn.Data))

	for i, p := range knn.Data {
		distances[i] = struct {
			Point    Point
			Distance float64
		}{
			Point:    p,
			Distance: EuclideanDistance(p, newPoint),
		}
	}

	// 按距离排序
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Distance < distances[j].Distance
	})

	// 计算前 K 个邻居的类别频率
	categoryCount := make(map[string]int)
	for i := 0; i < knn.K; i++ {
		categoryCount[distances[i].Point.Category]++
	}

	// 找到频率最高的类别
	var predictedCategory string
	maxCount := 0
	for category, count := range categoryCount {
		if count > maxCount {
			maxCount = count
			predictedCategory = category
		}
	}

	return predictedCategory
}

func main() {
	// 训练数据集
	data := []Point{
		{1, 2, "A"},
		{2, 3, "A"},
		{5, 4, "B"},
		{6, 7, "B"},
	}

	// 创建KNN分类器，设置K=3
	knn := KNNClassifier{
		Data: data,
		K:    4,
	}

	// 新样本数据
	newPoint := Point{X: 4, Y: 5}

	// 预测类别
	predictedCategory := knn.Predict(newPoint)
	fmt.Printf("新样本 (%.2f, %.2f) 的预测类别是: %s\n", newPoint.X, newPoint.Y, predictedCategory)
}
