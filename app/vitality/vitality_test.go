package vitality

import (
	"fmt"
	"testing"
)

type expect struct {
	min float64
	max float64
}

type testData struct {
	from     int
	days     int
	times    int
	interval int
	expect   expect
}

func newVitality() *Vitality {
	v1 := NewVitality(360, 2)
	v1.ElapseWeight = 2 // 该值越大，则历史值的作用就越小
	v1.DaysWeight = 2
	v1.Alg = NewAlg4(v1).SetNewDayWeight(30)
	return v1
}

func Test_calc(t *testing.T) {
	test(newVitality(), t)
}

func Test_adjust_args(t *testing.T) {
	v := newVitality()
	//v.Debug = true
	testAdjuestArgs(v, t)
}

// 这个必须有测试用例
func testAdjuestArgs(v *Vitality, t *testing.T) {
	// 通过该测试用例，用来调整参数，观察结果值
	data := []testData{
		{300, 1, 1, 1, expect{0, 0.1}},
		{270, 1, 1, 1, expect{0, 0.1}},
		{240, 1, 1, 1, expect{0, 0.1}},
		{210, 1, 1, 1, expect{0, 0.2}},
		{180, 1, 1, 1, expect{0.1, 0.2}},
		{150, 1, 1, 1, expect{0.1, 0.3}},
		{120, 1, 1, 1, expect{0.2, 1}},
		{90, 1, 1, 1, expect{0.6, 1.2}},
		{60, 1, 1, 1, expect{1, 2}},
		{30, 1, 1, 1, expect{2, 3.3}},
		{1, 1, 1, 1, expect{2, 5}},
		//30天内，隔一天活跃一次和每天都活跃的活跃度不是一倍的差距，但是还是需要有一定的差距
		{32, 30, 1, 1, expect{40, 60}},
		{32, 15, 1, 2, expect{30, 45}},
	}
	format := "%d days ago, vitality: %.2f\n"
	for _, d := range data {
		v1 := v.Calc(MakeActions(d.from, d.days, d.times, d.interval))
		fmt.Printf(format, d.from, v1)
		if v1 < d.expect.min || v1 > d.expect.max {
			t.Errorf("fail: %+v actual: %.2f", d, v1)
		}
	}
	//
}

// 这个必须有测试用例, 完善一下这个测试用例，使得能一直可测试
func test(v *Vitality, t *testing.T) {

	data := []testData{
		{390, 300, 2, 1, expect{20, 40}},
		{60, 30, 2, 1, expect{20, 45}},
		{30, 20, 2, 1, expect{30, 50}},
		{5, 4, 2, 1, expect{10, 20}},
		{420, 360, 2, 1, expect{40, 60}},
		{30, 10, 2, 1, expect{25, 35}},
		{12, 11, 2, 1, expect{30, 40}},
		{90, 80, 2, 1, expect{60, 80}},
	}
	for _, d := range data {
		v1 := v.Calc(MakeActions(d.from, d.days, d.times, d.interval))
		fmt.Printf("%+v  actual: %.2f\n", d, v1)
		if v1 < d.expect.min || v1 > d.expect.max {
			t.Errorf("fail: %+v actual: %.2f", d, v1)
		}
	}
}

// 基准测试数据
func TestBench(t *testing.T) {
	v := newVitality()
	x := 1
	for {
		if x > 360 {
			break
		}
		y := v.Calc(MakeActions(x, x, 1, 1))
		fmt.Printf("%d %.2f\n", x, y)
		x++
	}
}

func Test_weight(t *testing.T) {
	n := 360
	for j := 1; j < 360; j++ {
		fmt.Printf("%d: %.2f\n", j, testWeight(j, n))
	}

}
func testWeight(j int, n int) float64 {
	sum := float64(0)
	for i := 0; i < j; i++ {
		sum += float64(n-i) / float64(n)
	}
	return sum
}
