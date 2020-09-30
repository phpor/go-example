package vitality

import (
	"fmt"
	"math"
	"time"
)

type Alg4 struct {
	now          time.Time
	v            *Vitality
	newDayWeight int // 该值是一个大于1的值, 该值越大，最近的值的权重越小
}

// 该算法仅仅能保证结果值是一个0~100之间的数值，也能体现
func NewAlg4(v *Vitality) *Alg4 {
	return &Alg4{
		now:          time.Now(),
		v:            v,
		newDayWeight: 99,
	}
}

func (a *Alg4) SetNewDayWeight(w int) *Alg4 {
	if w >= 100 {
		w = 100
	}
	if w < 0 {
		w = 0
	}
	a.newDayWeight = (100 - w) + 1
	return a
}

func (a *Alg4) SumOneDay(date time.Time, times float64) float64 {
	return a.weight(date) * (1*a.v.DaysWeight + times*a.v.TimesWeight)
}

func (a *Alg4) CalcSum(sum float64) float64 {
	return (sum / (sum + float64(a.newDayWeight))) * 100
}

func (a *Alg4) weight(t time.Time) float64 { // 距离现在越远，则该值越小，这是一个 0-1之间的值
	// 这个算法使得时间对活跃的影响是匀速变化的（没有加速度）
	days := a.now.Sub(t).Seconds() / 86400 // days 基本在0~10 之间
	weight := 1 / (math.Pow(1+a.v.ElapseWeight/10000, days))
	if a.v.Debug {
		fmt.Printf("weight: %s %d days ago %.2f\n", t.Format("2006-01-02 15:04:05"), int(days), weight)
	}
	return weight
}
