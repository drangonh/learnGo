/**
 * @Author: dragon
 * @Description:
 * @File:  itemsaver_test
 * @Version: 1.0.0
 * @Date: 2020/3/30 下午4:13
 */

package persist

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
	"gomodtest/crawler/engine"
	"gomodtest/crawler/model"
	"testing"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Id:   "15184356789",
		Type: "zhenai",
		Url:  "1",
		Payload: model.Profile{
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
		}}

	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	index := "dating_test"
	err = save(client, expected, index)

	if err != nil {
		panic(err)
	}

	resp, err := client.Get().Index("dating_profile").Type(expected.Type).Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", resp.Source)

	//反序列化
	var actual engine.Item
	err = json.Unmarshal([]byte(*resp.Source), &actual)
	if err != nil {
		panic(err)
	}

	actualPayload, err := model.FromJsonObj(actual.Payload)
	if err != nil {
		panic(err)
	}
	actual.Payload = actualPayload
	if actual != expected {
		t.Errorf("%v,%v", actual, expected)
	}
}
