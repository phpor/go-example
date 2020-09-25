package main

type freq struct {
	min  int
	max  int
	name string
}

var f_1_3 = &freq{
	min:  1,
	max:  3,
	name: "f_1_3",
}

var f_4_10 = &freq{
	min:  4,
	max:  10,
	name: "f_4_10",
}

var f_11_24 = &freq{
	min:  11,
	max:  24,
	name: "f_11_24",
}
var f_25_29 = &freq{
	min:  25,
	max:  29,
	name: "f_25_29",
}
var f_30 = &freq{
	min:  30,
	max:  31,
	name: "f_30",
}

var freqs = []*freq{
	f_1_3, f_4_10, f_11_24, f_25_29, f_30,
}
