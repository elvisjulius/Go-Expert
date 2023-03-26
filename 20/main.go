package main

func main() {
	a := 10
	b := 2
	c := 30

	// if a > b {
	// 	println(a)
	// } else {
	// 	print(b)
	// }

	if a > b && c > a {
		println("a > b && c > a")
	}

	if a > b || c > a {
		println("a > b || c > a")
	}

	switch a {
	case 1:
		println("case a")
	case 2:
		println("case b")
	case 3:
		println("case c")
	default:
		println("case default")
	}
}
