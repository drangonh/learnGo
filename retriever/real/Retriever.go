package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	TimeOut  time.Duration
	UseAgent string
}

func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	result, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	return string(result)
}
