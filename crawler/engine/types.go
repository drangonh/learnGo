package engine

type ParserFunc func(contents []byte, url string) ParseResult

type Request struct {
	Url       string
	ParseFunc ParserFunc
}

type ParseResult struct {
	Requests []Request //记录请求结果的URL列表，用于之后继续请求
	Items    []Item    //记录请求结果的名称
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}
