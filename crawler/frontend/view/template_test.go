/**
 * @Author: dragon
 * @Description:
 * @File:  template_test
 * @Version: 1.0.0
 * @Date: 2020/4/1 上午10:13
 */

package view

import (
	"gomodtest/crawler/engine"
	"gomodtest/crawler/frontend/model"
	common "gomodtest/crawler/model"
	"html/template"
	"os"
	"testing"
)

func TestTemplate(t *testing.T) {
	template := template.Must(template.ParseFiles("template.html"))
	page := model.SearchResult{
		Hits:     123,
		Start:    0,
		Query:    "",
		PrevFrom: 0,
		NextFrom: 10,
		Items: []engine.Item{
			{
				Id:   "15184356789",
				Type: "zhenai",
				Url:  "1",
				Payload: common.Profile{
					Age:        34,
					Height:     162,
					Weight:     57,
					Income:     "3001-5000元",
					Gender:     "女",
					Name:       "安静的雪",
					Xinzuo:     "牡羊座",
					Occupation: "人事/行政",
					Marriage:   "离异",
					House:      "已购房",
					Hokou:      "山东菏泽",
					Education:  "大学本科",
					Car:        "未购车",
				}}},
	}
	out, err := os.Create("template.test.html")
	if err != nil {
		panic(err)
	}
	err = template.Execute(out, page)
	if err != nil {
		panic(err)
	}
}
