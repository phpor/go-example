package vitality

import (
	"math"
	"time"
)

type Alg2 struct {
	now time.Time
	v   *Vitality
}

func NewAlg2(v *Vitality) *Alg2 {
	return &Alg2{
		now: time.Now(),
		v:   v,
	}
}
func (a *Alg2) SumOneDay(date time.Time, times float64) float64 {
	return a.weight2(a.now, date) * (1*a.v.DaysWeight + times*a.v.TimesWeight) / 100
}

func (a *Alg2) CalcSum(sum float64) float64 {
	if sum < 0.1 {
		return 0
	}
	return (1 - 1/(sum*10)) * 100
}

func (a *Alg2) weight(now time.Time, t time.Time) float64 { // 距离现在越远，则该值约小
	// 这个算法使得时间对活跃的影响是匀速变化的（没有加速度）
	days := now.Sub(t).Seconds() / 86400 // 这个days是 0 - v.maxDays 之间的值
	return 100 / (days + 1)              // 这里的 + 1 是为了避免出现除零错误
}

func (a *Alg2) weight2(now time.Time, t time.Time) float64 { // 距离现在越远，则该值越小，这是一个 0-1之间的值
	// 这个算法使得时间对活跃的影响是匀速变化的（没有加速度）
	days := now.Sub(t).Seconds() / 86400 / (a.v.maxDays / 10) // days 基本在0~10 之间
	return 1 / (math.Pow(2, days))                            // 这里的 + 1 是为了避免出现除零错误
}
