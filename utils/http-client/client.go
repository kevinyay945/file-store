package httpClient

import "github.com/go-resty/resty/v2"

type Response struct {
	Body []byte
}

type IHttpClient interface {
	Get(url string) (resp Response, err error)
	SetAuthToken(s string) IHttpClient
}

type HttpClient struct {
	resty     *resty.Client
	authToken string
}

func (h *HttpClient) SetAuthToken(s string) IHttpClient {
	h.authToken = s
	return h
}

func NewHttpClient() IHttpClient {
	_resty := resty.New()
	return &HttpClient{
		resty: _resty,
	}
}

func (h *HttpClient) Get(url string) (resp Response, err error) {
	r := h.resty.R()
	if h.authToken != "" {
		r.SetAuthToken(h.authToken)
	}
	response, err := r.Get(url)
	resp.Body = response.Body()
	return
}
