package main

import "fmt"

func main() {
	m := createMonthActivity()
	mUsers := map[int]int{} // 月活的用户集合
	uTimes := 0             // 每月的人次
	for _, v := range m {
		fmt.Printf("%d: %d\n", v.date, len(v.users))
		// 打印每天的低中高频用户
		freqNum := map[*freq]int{}
		for _, user := range v.users {
			freqNum[user.freq]++
			mUsers[user.uid]++
			uTimes++
		}
		for _, v := range freqs {
			fmt.Printf("%s: %d\t", v.name, freqNum[v])
		}
		println()
	}
	println()
	fmt.Printf("mUsers: %d   uTimes: %d\n", len(mUsers), uTimes)
}
