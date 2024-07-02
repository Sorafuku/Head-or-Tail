package main

import	(
	"fmt"
	"math/rand"
)

func main(){
	headcnt := 0
	tailcnt := 0
	var name string
	headorTails := []string{
		"Heads",
		"Tails",
	}

	fmt.Println("Who are you?")
	fmt.Scan(&name)
	fmt.Println("Hello, "+name+"!")


	fmt.Println("Tossing a coin...")
	for i:=0;i<3;i++ {
		if(rand.Intn(2)==0){
			fmt.Println("Round ",i+1,": ", headorTails[0])
			headcnt++
		} else {
			fmt.Println("Round ",i+1,": ", headorTails[1])
			tailcnt++
		}
	}
	fmt.Println("Heads: ",headcnt," Tails: ",tailcnt)

	if(headcnt>tailcnt) {
		fmt.Println("You won!")
	} else {
		fmt.Println("You lost!")
	}
	
}

