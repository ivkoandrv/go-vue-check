package reports

import (
	"fmt"
	"go-vue-check/common"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"sort"
	"text/template"
)

type ReportData struct {
	Counter    common.APICounter
	ReportList []common.ReportRow
}

var reportFilePath string = "report.html"

func GenerateHTMLReport(Stats common.APICounter) {
	fmt.Printf("Yeah, we can generate report\n")
	if len(common.ReportList) > 0 {

		templateFile, err := ioutil.ReadFile("templates/report_api_template.html")
		if err != nil {
			log.Fatal(err)
			return
		}

		tpl, err := template.New("report_api").Parse(string(templateFile))
		if err != nil {
			log.Fatal(err)
			return
		}

		reportFile, err := os.Create(reportFilePath)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer reportFile.Close()

		sort.Slice(common.ReportList, func(i, j int) bool {
			return common.ReportList[i].APIType < common.ReportList[j].APIType
		})

		data := ReportData{
			Counter:    Stats,
			ReportList: common.ReportList,
		}

		err = tpl.Execute(reportFile, data)
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println("Report has been generated in report.html")
	} else {
		fmt.Println("ReportList is empty.")
	}
}

func OpenReport() {
	_, err := os.Stat(reportFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Report file not found. Try to launch with flag -generate")
			return
		}
		fmt.Println("Error: ", err)
		return
	}

	cmd := exec.Command("open", reportFilePath)

	err = cmd.Start()

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Printf("Open %s in default browser\n", reportFilePath)

	err = cmd.Wait()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}
