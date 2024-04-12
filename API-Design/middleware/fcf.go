package main

type Fn func(int, int) int

func cal(sn Fn) int {
	return sn(5, 4)
}

func sum(a int, b int) int {
	return a + b
}

// func main() {
// 	fn := sum
// 	r1 := fn(1, 2)
// 	fmt.Println(r1)

// 	r2 := cal(fn)
// 	fmt.Println(r2)

// 	r3 := cal(sum)
// 	fmt.Println(r3)
// }
