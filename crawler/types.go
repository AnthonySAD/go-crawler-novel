package crawler

type Request struct {
	Url 		string
	ParseFunc 	func([]byte, map[string]string) ParseResult
	Info        map[string]string
}

type ParseResult struct {
	Requests 	[]Request
	HandleFunc  func(interface{})
	Content 	interface{}
}
