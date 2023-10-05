package define_api

import (
	"fmt"
	"go-vue-check/common"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ScanFilesWithVueExtension(rootPath string, isGenerate bool) *common.APICounter {
	counter := &common.APICounter{}
	//reportList := reports.ReportList

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

					if isGenerate {

						optionsRow := common.ReportRow{
							FileName: path,
							APIType:  "Options API",
						}

						//fmt.Println(optionsRow)
						common.ReportList = append(common.ReportList, optionsRow)

					}
				case "Composition API (Script Setup)", "Composition API (TypeScript with Script Setup)":
					counter.CompositionAPI++

					if isGenerate {
						compositionRow := common.ReportRow{
							FileName: path,
							APIType:  "Composition API",
						}

						//fmt.Println(compositionRow)
						common.ReportList = append(common.ReportList, compositionRow)

					}
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

	if isGenerate {

		fmt.Print("REPORT ARR LENGTH")
		fmt.Print(len(common.ReportList))
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
