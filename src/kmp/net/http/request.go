package http
import "net/http"
type Request struct{
	req *http.Request
}
type Response struct{
	resp *http.Response
}

func NewRequestFromOrigin(o_req *http.Request)(req *Request){
	return &Request{req:o_req}
}
