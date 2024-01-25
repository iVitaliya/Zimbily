package main

import (
	"fmt"
	"os"

	"github.com/iVitaliya/Zimbily/utils"
	"github.com/iVitaliya/colors-go"
)

func main() {
	var SU = utils.String()

	length := len(os.Args)
	args := os.Args[1:length]

	if length <= 0 {
		
	}

	fmt.Print(args)
}