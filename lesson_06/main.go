package main

func main() {
	i := 0
	i = i + 1
	i += 1
	i++
	//fmt.Println(i)

	i = 10
	i = i - 1
	i -= 1
	i--
	//fmt.Println(i)

	for i := 0; i < 10; i++ {
		//fmt.Println(i)
	}

	var j int
	for j = 0; j < 10; j++ {
		//fmt.Println(j)
	}
	sum := 1
	for sum < 20 {
		sum += sum
		//fmt.Println(sum)
	}

	for sum <= 100 {
		sum += 1
		//fmt.Println(sum)
	}

	//for {
	//	fmt.Println("Stop me please")
	//}

}
