package define_api

import (
	"fmt"
	"path/filepath"
)

func StartDefineAPI(projectSrcPath string) func() *APICounter {

	fmt.Printf("Source Directory: %s\n", projectSrcPath)

	viewsFolderPath := filepath.Join(projectSrcPath, "views")
	componentsFolderPath := filepath.Join(projectSrcPath, "components")

	viewsCounter := ScanFilesWithVueExtension(viewsFolderPath)
	componentsCounter := ScanFilesWithVueExtension(componentsFolderPath)

	viewsCounter.OptionsAPI += componentsCounter.OptionsAPI
	viewsCounter.CompositionAPI += componentsCounter.CompositionAPI
	viewsCounter.TotalFiles += componentsCounter.TotalFiles

	if viewsCounter.TotalFiles > 0 {
		viewsCounter.OptionsPercent = float64(viewsCounter.OptionsAPI) / float64(viewsCounter.TotalFiles) * 100
		viewsCounter.CompositionPercent = float64(viewsCounter.CompositionAPI) / float64(viewsCounter.TotalFiles) * 100
	}

	fmt.Printf("\nResults:\n")
	fmt.Printf("Options API Files: %d\n", viewsCounter.OptionsAPI)
	fmt.Printf("Composition API Files: %d\n", viewsCounter.CompositionAPI)
	fmt.Printf("Total Files: %d\n", viewsCounter.TotalFiles)
	fmt.Printf("Options API Percentage: %.2f%%\n", viewsCounter.OptionsPercent)
	fmt.Printf("Composition API Percentage: %.2f%%\n", viewsCounter.CompositionPercent)

	return func() *APICounter {
		return viewsCounter
	}
}
