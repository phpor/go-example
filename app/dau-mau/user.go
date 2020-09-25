package main

import "math/rand"

type user struct {
	uid  int
	freq *freq
}

var users []*user
var userCount = 10000

func init() {
	users = initUser(userCount)
}

type zone struct {
	from int
	to   int
}

func newZone(from, to int) *zone {
	return &zone{
		from: from,
		to:   to,
	}
}

// 月活中，每种频率的用户的占比配置，据此初始化月活
var month_dist = map[*zone]*freq{ // 根据这个配置，日活用户约40， 和实际情况差别不太大
	newZone(0, 33):   f_1_3,
	newZone(34, 56):  f_4_10,
	newZone(57, 75):  f_11_24,
	newZone(76, 84):  f_25_29,
	newZone(85, 100): f_30,
}

func initUser(sum int) []*user {
	ret := make([]*user, sum)
	i := 0
	f := f_1_3
	for {
		r := int(rand.Int31n(100))
		for k, v := range month_dist {
			if r >= k.from && r <= k.to {
				f = v
				break
			}
		}
		ret[i] = &user{
			uid:  i,
			freq: f,
		}
		i++
		if i >= sum {
			break
		}
	}
	return ret
}
