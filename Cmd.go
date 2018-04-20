package main

import (
	"fmt"
)

func startCMDGame() {
	setupGame()
	// for {
	// 	reader := bufio.NewReader(os.Stdin)

	// }
}

func promptPlayer() {
	fmt.Printf("It is Player %v's turn", theGame.turn)
}
