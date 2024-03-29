package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"vitality"
)

func main() {
	addr := flag.String("addr", ":8181", "listen addr")
	root := flag.String("root", "../", "www root")
	flag.Parse()
	fs := http.FileServer(http.Dir(*root))
	http.Handle("/html/", fs)
	http.HandleFunc("/v", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		newDayWeight, _ := strconv.ParseInt(request.Form.Get("newDayWeight"), 10, 64)
		ElapseWeight, _ := strconv.ParseFloat(request.Form.Get("ElapseWeight"), 64)
		maxDays, _ := strconv.ParseInt(request.Form.Get("maxDays"), 10, 64)
		activeDays, _ := strconv.ParseInt(request.Form.Get("activeDays"), 10, 64)
		DaysWeight, _ := strconv.ParseFloat(request.Form.Get("DaysWeight"), 64)
		accumulation, _ := strconv.ParseBool(request.Form.Get("accumulation"))
		alg := request.Form.Get("alg")

		if maxDays == 0 {
			maxDays = 360
		}

		v1 := vitality.NewVitality(int(maxDays), 2)
		v1.ElapseWeight = ElapseWeight // 该值越大，则历史值的作用就越小
		v1.DaysWeight = DaysWeight
		switch alg {
		case "php":
			v1.Alg = vitality.NewAlgPhp(v1)
		default:
			v1.Alg = vitality.NewAlg4(v1).SetNewDayWeight(int(newDayWeight))
		}

		x := 1
		d := 1
		var ret []string
		for {
			if x > int(activeDays) {
				break
			}
			if accumulation {
				d = x
			}
			y := v1.Calc(vitality.MakeActions(x, d, 1, 1))
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
