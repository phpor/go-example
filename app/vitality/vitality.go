package vitality

import (
	"fmt"
	"math"
	"time"
)

type Algorithm interface {
	SumOneDay(date time.Time, times float64) float64
	CalcSum(sum float64) float64
}

type Vitality struct {
	maxDays        float64       // 最多算多少天的
	earliestTime   time.Time     // 该时间之前的都不计算在内
	duration       time.Duration // 该时间之前的都不计算在内
	TimesWeight    float64       // 次数权重
	DaysWeight     float64       // 天数权重
	ElapseWeight   float64       // 距离现在远近对活跃度影响的因子, 该值越小，拖尾的贡献度越大，活跃度综合越能接近 100% ，定义域 ( 0 , 100 )
	MaxTimesOneDay int           // 每天计算的最大的次数
	Debug          bool
	Alg            Algorithm
}

func NewVitality(maxDays int, MaxTimesOneDay int) *Vitality {
	return &Vitality{
		maxDays:        float64(maxDays),
		earliestTime:   time.Now().Add(-time.Duration(maxDays*86400) * time.Second),
		duration:       time.Duration(maxDays*86400) * time.Second,
		TimesWeight:    1,
		DaysWeight:     1,
		ElapseWeight:   2,
		MaxTimesOneDay: MaxTimesOneDay,
	}
}

type action struct {
	date  time.Time // 活跃日期，按当日零点计算
	times int       // 当日活跃次数
}

func MakeActions(daysAgo int, days int, times int, interval int) []action {
	start := time.Now().Add(-time.Duration(daysAgo) * 24 * time.Hour)
	var a []action
	for {
		if days <= 0 {
			break
		}
		a = append(a, action{
			date:  start.AddDate(0, 0, interval),
			times: times,
		})
		start = start.AddDate(0, 0, interval)
		days--
	}
	return a
}

func (v Vitality) Calc(actions []action) float64 {
	sum := float64(0)
	now := time.Now()
	for _, action := range actions {
		if action.date.After(now) { // 活跃时间在将来的，就不要了
			continue
		}
		if action.date.Before(v.earliestTime) { // 时间太久的就不要了
			continue
		}
		times := math.Min(float64(action.times), float64(v.MaxTimesOneDay))

		// 关键问题是，sum的值域难以确定
		sum += v.Alg.SumOneDay(action.date, times)
	}
	if v.Debug {
		fmt.Printf("sum: %.2f\t", sum)
	}
	return v.Alg.CalcSum(sum)
}
