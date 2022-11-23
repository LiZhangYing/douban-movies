package main

import (
	"fmt"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/xuri/excelize/v2"
)

func ExportChartsBar(execlFileName string, chartsBarFileName string) {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Width:  "1200px",
			Height: "4000px",
		}),
		charts.WithTitleOpts(opts.Title{
			Title: "电影评价人数",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "人数",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "片名",
		}),
	)

	f, _ := excelize.OpenFile(execlFileName)

	bar.SetXAxis(generateBarItem("c", f)).
		AddSeries("Category A", generateBarItem("e", f))
	bar.XYReversal()

	c, _ := os.Create(chartsBarFileName)
	bar.Render(c)
}

func generateBarItem(columns string, f *excelize.File) []opts.BarData {
	itemCnt := 257
	items := make([]opts.BarData, 0)
	for i := 2; i < itemCnt; i++ {
		cell, _ := f.GetCellValue("Sheet1", fmt.Sprintf("%s%d", columns, i))
		items = append(items, opts.BarData{Value: cell})
	}
	return items
}
