package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contexts, err := ioutil.ReadFile("citylist_test_data.html")

	if err != nil {
		panic(err)
	}

	results := ParseCityList(contexts, "")

	const resultsSize = 494

	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	if len(results.Requests) != resultsSize {
		t.Errorf("result should have %d "+
			"requests; but had %d",
			resultsSize, len(results.Requests))
	}

	for i, url := range expectedUrls {
		if results.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but "+
				"was %s",
				i, url, results.Requests[i].Url)
		}
	}
}
