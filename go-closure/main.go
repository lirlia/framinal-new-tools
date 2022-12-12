package main

import "fmt"

func main() {

	counter1 := getCounter()
	counter2 := getCounter()

	fmt.Println(counter1()) // 1
	fmt.Println(counter1()) // 2

	fmt.Println(counter2()) // 1
	fmt.Println(counter2()) // 2
	fmt.Println(counter2()) // 3

	counter3, counter4 := getCounter2()

	fmt.Println(counter3()) // 1
	fmt.Println(counter3()) // 2
	fmt.Println(counter4()) // 3
	fmt.Println(counter4()) // 4

	i := 0
	counter5:= getCounter3(&i)
	fmt.Println(counter5()) // 1
	i++
	fmt.Println(counter5()) // 3
}

func getCounter() func () int{
	count  := 0

	if count == 0 {
		return func() int {
			count++
			return count
		}
	}

	return func() int {
		count += 2
		return count
	}
}

func getCounter2() (f1 func () int, f2 func () int){
	count := 0

	f1 = func() int {
			count++
			return count
		}

	f2 = func() int {
		count++
		return count
	}

	return f1, f2
}

func getCounter3(count *int) func () int{
	return func() int {
		*count++
		return *count
	}
}
