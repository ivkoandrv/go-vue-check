package main

import (
	"flag"
	"fmt"
	"go-vue-check/define_api"
	"go-vue-check/utils"
	"os"
	"path/filepath"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory", err)
		return
	}

	projectPath := flag.String("p", currentDir, "Path to project")
	sourceDir := flag.String("s", "src", "Source directory")

	flag.Parse()

	projectPathSrc := filepath.Join(*projectPath, *sourceDir)

	elapsedTime := utils.ElapseTimeMemory(func() {
		api := define_api.StartDefineAPI(projectPathSrc)
		if api == nil {
			fmt.Println("Error: StartDefineAPI returned nil")
			return
		}
	})

	if elapsedTime == 0 {
		fmt.Println("Error: Elapsed time is zero")
		return
	}
	fmt.Printf("Elapsed time: %.2f seconds\n", elapsedTime)

}
