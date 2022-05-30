package recursion

var m = map[int]int{}

func Stage(n int) int {
	if n == 2 {
		return 2
	}
	if n == 1 {
		return 1
	}

	var res int
	if v, ok := m[n]; ok {
		res = v
	} else {
		res = Stage(n-1) + Stage(n-2)
		m[n] = res
	}
	return res
}

func Stage2(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	fn1 := 2
	fn2 := 1
	res := 0

	for i := 3; i <= n; i++ {
		res = fn1 + fn2
		fn2 = fn1
		fn1 = res
	}
	return res
}
