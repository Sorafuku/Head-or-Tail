package main

import	(
	"fmt"
	"math/rand"
)

func main(){
	headcnt := 0
	tailcnt := 0
	fmt.Println("Tossing a coin...")
	headorTails := []string{
		"Heads",
		"Tails",
	}
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
	
}

