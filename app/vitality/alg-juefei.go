package vitality

import (
	"math"
	"time"
)

/**
  该算法的本质是： 随着x的变大，一个很大的数加上一个很小的数和减去一个很小的数差别越来越小，但是，加上一个数总比减去一个数要大，
	使得结果总在 0~1 之间。
	还可以有其它的变种： 1-(e^(x)-e^(-x))/(e^x+e^(-x))   ; 这里的 e 可以视为一个可以调的参数，大于1就行
*/
type AlgPhp struct {
	now          time.Time
	v            *Vitality
	days         int
	newDayWeight int // 该值是一个大于1的值, 该值越大，最近的值的权重越小
}

func NewAlgPhp(v *Vitality) *AlgPhp {
	return &AlgPhp{
		now:          time.Now(),
		v:            v,
		newDayWeight: 99,
	}
}

func (a *AlgPhp) SumOneDay(date time.Time, times float64) float64 {
	a.days++
	num := int(a.now.Sub(date).Seconds() / 86400)
	return a.getTimeFactor(num) * (math.Pow(float64(times), 0.045) - 0.2) * 100
}

func (a *AlgPhp) CalcSum(sum float64) float64 {
	fm := float64(0)
	for i := 0; i <= a.days; i++ {
		fm += a.getTimeFactor(i)
	}
	vitality := sum / fm
	if vitality > 100 {
		vitality = 100
	}
	return vitality
}

func (a *AlgPhp) getTimeFactor(i int) float64 {
	x := (float64(4)/180)*float64(i) - 2
	mx := -x
	return 100 * ((math.Exp(mx)-math.Exp(x))/(math.Exp(x)+math.Exp(mx)) + 1.2)
}
