package main

// "github.com/heeropunjabi/go-handson/utils"

	
func abc(x interface{}, y interface{}) {

	z:= x.(int) + y.(int)
	println(z)

	// if x.(int) > y.(int) {
	// 	fmt.Println("x is greater than y")
	// } else {
	// 	fmt.Println("y is greater than x")
	// }
	
}


func main() {
	abc(10, 20)

	

}