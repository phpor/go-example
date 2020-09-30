package main

import (
	"flag"
	"fmt"
	"github.com/phpor/go-example/app/vitality"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	addr := flag.String("addr", ":8181", "listen addr")
	flag.Parse()
	http.HandleFunc("/v", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		newDayWeight, _ := strconv.ParseInt(request.Form.Get("newDayWeight"), 10, 64)
		ElapseWeight, _ := strconv.ParseFloat(request.Form.Get("ElapseWeight"), 64)
		maxDays, _ := strconv.ParseInt(request.Form.Get("maxDays"), 10, 64)
		activeDays, _ := strconv.ParseInt(request.Form.Get("activeDays"), 10, 64)
		DaysWeight, _ := strconv.ParseFloat(request.Form.Get("DaysWeight"), 64)

		if maxDays == 0 {
			maxDays = 360
		}

		v1 := vitality.NewVitality(int(maxDays), 2)
		v1.ElapseWeight = ElapseWeight // 该值越大，则历史值的作用就越小
		v1.DaysWeight = DaysWeight
		v1.Alg = vitality.NewAlg4(v1).SetNewDayWeight(int(newDayWeight))

		x := 1
		var ret []string
		for {
			if x > int(activeDays) {
				break
			}
			y := v1.Calc(vitality.MakeActions(x, x, 1, 1))
			//fmt.Printf("%d %.2f\n",x,y)
			ret = append(ret, fmt.Sprintf("%.2f", y))
			x++
		}
		writer.Write([]byte("["))
		writer.Write([]byte(strings.Join(ret, ",")))
		writer.Write([]byte("]"))
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		panic(err)
	}
}
