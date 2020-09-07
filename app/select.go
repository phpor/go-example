package main

// see also: https://studygolang.com/articles/28353?fr=sidebar

func main() {
	selectModelA()

}

func selectModelC() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	ch1 <- 1
	ch2 <- 1
	select { // 这里先select 谁是不确定的； 如果很的很在意顺序的话，就写成多个select, 但是这样就不能实现同时阻塞在多个case上了，似乎这个比较麻烦
	case <-ch1:
		println("ch1")
	case <-ch2:
		println("ch2")
	}

}

// 如何让select 有顺序
func selectModelA() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	ch1 <- 1
	ch2 <- 1
	select { // 这样虽然能实现，但是chan太多就很恶心了
	case <-ch1:
		println("ch1")
	case <-ch2:
		select {
		case <-ch1:
			println("ch1")
		default:
			println("ch2")
		}
	}
}

func selectModelB() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	ch1 <- 1
	ch2 <- 1
	// 这种select没法在循环中使用，非要在循环中使用的话，就得借助time.Ticker 或time.Sleep
	select {
	case <-ch1:
		println("ch1")
	default:

	}
	select {
	case <-ch2:
		println("ch2")
	default:

	}

}
