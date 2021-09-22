package main

import "fmt"

func greetings(name string) string {
	message := fmt.Sprintf("Hi %v. welcome", name)
	asd := "asddfff"

	return message + asd
}

func main() {
	result := greetings("sunny")
	fmt.Println(result + "lalwani")

}

// func main() {
// 	fmt.Println("hello main")
// 	TryMe();
// }

// func TryMe(){
// fmt.Println("hello from Tryme")
// }
