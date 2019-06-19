package main

func main() {
	// 1 - use 2 nested for loops
	for i := 1; i <= 25; i++ {
		for j := 1; j <= i; j++ {
			print("G")
		}
		println()
	}
	// 2 -  use only one for loop and string concatenation
	str := "G"
	for i := 1; i <= 25; i++ {
		println(str)
		str += "G"
	}
	//3 - use one for loop and strings.Repeat
	for i:=1; i<=25; i++{
		str2 := strings.Repeat("G",i)
		fmt.Println(str2)
	}
}
