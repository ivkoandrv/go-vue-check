package define_api

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ScanFilesWithVueExtension(rootPath string) *APICounter {
	counter := &APICounter{}

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error accessing path: ", err)
			return nil
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".vue") {
			fileSrc, err := DetectAPI(path)
			if err != nil {
				fmt.Printf("Error detecting API for %s: %v\n", path, err)
			} else {
				switch fileSrc {
				case "Options API", "Options API (TypeScript)":
					counter.OptionsAPI++
				case "Composition API (Script Setup)", "Composition API (TypeScript with Script Setup)":
					counter.CompositionAPI++
				}
				counter.TotalFiles++
			}
		}

		return nil
	}

	err := filepath.Walk(rootPath, walkFn)
	if err != nil {
		fmt.Println("Error walking through directories: ", err)
	}

	return counter
}

func DetectAPI(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	contentStr := strings.ToLower(string(content))

	if strings.Contains(contentStr, "<script setup>") {
		return "Composition API (Script Setup)", nil
	} else if strings.Contains(contentStr, "<script lang=\"ts\">") {
		if strings.Contains(contentStr, "<script setup>") {
			return "Composition API (TypeScript with Script Setup)", nil
		}
		return "Options API (TypeScript)", nil
	} else if strings.Contains(contentStr, "<script>") {
		return "Options API", nil
	}

	return "Unknown API", nil
}
