package main

import "math/rand"

type activity struct {
	users []*user
	date  int
}

var days = 30

func createMonthActivity() []*activity {
	i := 1
	var m []*activity
	for {
		m = append(m, createDayActivity(i))
		i++
		if i > days {
			break
		}
	}
	return m
}

func createDayActivity(d int) *activity {
	a := &activity{date: d}
	var x int
	for _, user := range users { // 每天每人最多活跃一次，仅表示一下活跃与否而已
		x = int(rand.Int31n(30))
		if x <= user.freq.min+int(rand.Int31n(int32(user.freq.max-user.freq.min))) {
			a.users = append(a.users, user)
		}
	}
	return a
}
