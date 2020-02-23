package main

import (
	"fmt"
	"time"
)

func main() {

	for timer := 10; timer >= 0; timer-- {
		if (timer == 0) {
			fmt.Println("Exiting loop!")
			break
		}
		fmt.Println(timer)

		//sleep for 1 sec
		time.Sleep(time.Second)
	}
}