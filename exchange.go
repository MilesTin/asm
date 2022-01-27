package main

func exchange(a, b float64) (float64, float64) {
	tmp := a + 1
	a = b
	b = tmp
	return a, b
}
