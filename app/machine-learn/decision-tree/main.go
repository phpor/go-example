package main

import (
	"fmt"
	"math"
)

// Sample 表示一个样本及其标签
type Sample struct {
	Features []int
	Label    string
}

// TreeNode 表示决策树的节点
type TreeNode struct {
	FeatureIndex int       // 分割特征索引
	Threshold    float64   // 阈值（对于二叉树通常为0.5）
	LeftChild    *TreeNode // 左子节点
	RightChild   *TreeNode // 右子节点
	Label        string    // 叶节点的标签
}

// DecisionTree 表示决策树
type DecisionTree struct {
	Root *TreeNode
}

// Entropy 计算熵
func Entropy(labels []string) float64 {
	counts := make(map[string]int)
	for _, label := range labels {
		counts[label]++
	}
	var entropy float64
	total := len(labels)
	for _, count := range counts {
		p := float64(count) / float64(total)
		entropy -= p * math.Log2(p)
	}
	return entropy
}

// InformationGain 计算信息增益
func InformationGain(samples []Sample, featureIndex int) float64 {
	totalEntropy := Entropy(getLabels(samples))

	leftSamples, rightSamples := splitSamples(samples, featureIndex)

	leftEntropy := Entropy(getLabels(leftSamples))
	rightEntropy := Entropy(getLabels(rightSamples))

	pLeft := float64(len(leftSamples)) / float64(len(samples))
	pRight := float64(len(rightSamples)) / float64(len(samples))

	return totalEntropy - (pLeft*leftEntropy + pRight*rightEntropy)
}

// getLabels 提取所有样本的标签
func getLabels(samples []Sample) []string {
	labels := make([]string, len(samples))
	for i, sample := range samples {
		labels[i] = sample.Label
	}
	return labels
}

// splitSamples 根据特征索引分割样本
func splitSamples(samples []Sample, featureIndex int) ([]Sample, []Sample) {
	var leftSamples, rightSamples []Sample
	for _, sample := range samples {
		if sample.Features[featureIndex] == 0 {
			leftSamples = append(leftSamples, sample)
		} else {
			rightSamples = append(rightSamples, sample)
		}
	}
	return leftSamples, rightSamples
}

// BuildTree 构建决策树
func (dt *DecisionTree) BuildTree(samples []Sample, maxDepth int, depth int) {
	if len(samples) == 0 || depth >= maxDepth {
		return
	}

	bestFeatureIndex := -1
	maxGain := -1.0

	for i := range samples[0].Features {
		gain := InformationGain(samples, i)
		if gain > maxGain {
			maxGain = gain
			bestFeatureIndex = i
		}
	}

	if bestFeatureIndex == -1 {
		return
	}

	leftSamples, rightSamples := splitSamples(samples, bestFeatureIndex)

	node := &TreeNode{
		FeatureIndex: bestFeatureIndex,
		Threshold:    0.5,
	}

	dt.Root = node

	if len(leftSamples) > 0 {
		node.LeftChild = &TreeNode{}
		dt.BuildTree(leftSamples, maxDepth, depth+1)
	} else {
		node.Label = mostCommonLabel(getLabels(samples))
	}

	if len(rightSamples) > 0 {
		node.RightChild = &TreeNode{}
		dt.BuildTree(rightSamples, maxDepth, depth+1)
	} else {
		node.Label = mostCommonLabel(getLabels(samples))
	}
}

// Predict 预测新样本的类别
func (dt *DecisionTree) Predict(sample Sample) string {
	node := dt.Root
	for node.LeftChild != nil && node.RightChild != nil {
		if sample.Features[node.FeatureIndex] == 0 {
			node = node.LeftChild
		} else {
			node = node.RightChild
		}
	}
	return node.Label
}

// mostCommonLabel 找到最常见的标签
func mostCommonLabel(labels []string) string {
	counts := make(map[string]int)
	for _, label := range labels {
		counts[label]++
	}
	var maxCount int
	var maxLabel string
	for label, count := range counts {
		if count > maxCount {
			maxCount = count
			maxLabel = label
		}
	}
	return maxLabel
}

func main() {
	// 训练数据集
	samples := []Sample{
		{[]int{0, 0}, "A"},
		{[]int{1, 0}, "B"},
		{[]int{0, 1}, "B"},
		{[]int{1, 1}, "A"},
	}

	// 创建决策树分类器
	dt := DecisionTree{}

	// 构建决策树，最大深度为3
	dt.BuildTree(samples, 3, 0)

	// 新样本数据
	newSample := Sample{Features: []int{0, 1}}

	// 预测类别
	predictedCategory := dt.Predict(newSample)
	fmt.Printf("新样本 %v 的预测类别是: %s\n", newSample.Features, predictedCategory)
}
