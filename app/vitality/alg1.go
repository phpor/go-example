package vitality

import (
	"math"
	"time"
)

type Alg1 struct {
	now    time.Time
	v      *Vitality
	maxSum float64
}

func NewAlg1(v *Vitality) *Alg1 {
	maxSum := (v.MaxTimesOneDay + 1) * int(v.maxDays) / 2
	return &Alg1{
		now:    time.Now(),
		v:      v,
		maxSum: float64(maxSum),
	}
}

// 不能给最近的天数加更大的权重和曲线本身的特性综合使用
func (a *Alg1) SumOneDay(date time.Time, times float64) float64 {
	days := a.now.Sub(date).Seconds() / 86400
	v := a.v
	return ((v.maxDays - days) / v.maxDays) * (1*v.DaysWeight + times*v.TimesWeight)
}

func (a *Alg1) CalcSum(sum float64) float64 {

	return float64(1) / math.Pow(a.v.ElapseWeight, (a.maxSum-sum)/a.maxSum*10) * 100
}
