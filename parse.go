package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

type Douban struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	Year         string `json:"year"`
	RatingPeople string `json:"ratingPeople"`
}

var result []*Douban

func Parse(c *colly.Collector) []*Douban {
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 5 * time.Second, // 随机延迟
	})

	c.OnHTML("ol.grid_view", func(element *colly.HTMLElement) {
		element.DOM.Find("li").Each(func(i int, selection *goquery.Selection) {
			href, found := selection.Find("div.hd > a").Attr("href")
			if found {
				c.Visit(href) // 获取子页面电影详情
			}
		})
	})

	c.OnHTML("body", func(element *colly.HTMLElement) {
		selection := element.DOM.Find("div#content")
		id := selection.Find("div.top250 > span.top250-no").Text()
		title := selection.Find("h1 > span").First().Text()
		year := selection.Find("h1 > span.year").Text()
		ratingPeople := selection.Find("a.rating_people > span").Text()

		// 处理数据
		trimId := strings.Trim(id, "No.")

		re := regexp.MustCompile("[\u4e00-\u9fa5·\uff1a\\d]*") // 匹配中文数字和特殊符号(中文冒号和间隔号)
		regexpTitle := re.FindString(title)
		fmt.Println(regexpTitle)

		trimLeftYear := strings.TrimLeft(year, "(")
		trimYear := strings.TrimRight(trimLeftYear, ")")

		movies := &Douban{
			Id:           trimId,
			Title:        regexpTitle,
			Year:         trimYear,
			RatingPeople: ratingPeople,
		}
		result = append(result, movies)
	})

	for i := 0; i < 10; i++ {
		c.Visit("https://movie.douban.com/top250?start=" + strconv.Itoa(i*25))
	}
	return result
}
