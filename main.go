package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("go", "version")
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stdout

	// err := cmd.Run()
	
	// if err != nil {
	// 	fmt.Print("Error", err)
	// }

	output, err := cmd.Output()

	if err != nil {
		fmt.Print("Error", err)
	} else {
		fmt.Printf("Output: %s\n", output)
		fmt.Println(string(output))
	}
}