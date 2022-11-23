package main

import (
	"encoding/json"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	var jsonFileName = "export/douban.json"
	var execlFileName = "export/douban.xlsx"
	var chartsBarFileName = "export/douban.html"

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36"),
	)

	movies := Parse(c)

	// 生成 JSON 文件
	file, _ := os.Create(jsonFileName)
	encoder := json.NewEncoder(file)
	encoder.SetIndent(" ", "  ")
	encoder.Encode(movies)

	// 导出 Execl 表
	ExportExecl(jsonFileName, execlFileName)

	// 导出图表
	ExportChartsBar(execlFileName, chartsBarFileName)

}
