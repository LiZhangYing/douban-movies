package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/xuri/excelize/v2"
)

type mytype []map[string]string

func ExportExecl(jsonFileName string, execlFileName string) {
	var data mytype
	file, err := ioutil.ReadFile(jsonFileName)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	f := excelize.NewFile()
	categories := map[string]string{"B1": "排名", "C1": "电影名", "D1": "首映年份", "E1": "评论人数"}
	for k, v := range categories {
		f.SetCellValue("Sheet1", k, v)
	}
	for k, m := range data {
		k = k + 2 // 从表中第 2 行开始记录数据
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", k), m["id"])
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", k), m["title"])
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", k), m["year"])
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", k), m["ratingPeople"])
	}
	if err := f.SaveAs(execlFileName); err != nil {
		fmt.Println(err)
	}
}
