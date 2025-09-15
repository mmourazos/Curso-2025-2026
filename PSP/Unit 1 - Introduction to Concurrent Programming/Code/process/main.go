package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exPath := filepath.Dir(ex)

	fmt.Printf("Executable path: %s\n", exPath)

	cmd := exec.Command("pwsh", "-c", "ls", "C:\\")
	// if err := cmd.Run(); err != nil {
	// 	log.Fatal(err)
	// }

	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Output: %s\n", out)
	fmt.Println("End of main process.")
}
