package httpClient

import (
	"github.com/go-resty/resty/v2"
	"io"
	"net/url"
)

type Response struct {
	Body []byte
}

type IHttpClient interface {
	Get(url string) (resp Response, err error)
	SetAuthToken(s string) IHttpClient
	Post(s string) (resp Response, err error)
	SetQueryParam(key string, value string) IHttpClient
	SetFileReader(param string, fileName string, reader io.Reader) IHttpClient
}

type MultipartFile struct {
	Param    string
	FileName string
	Reader   io.Reader
}

type HttpClient struct {
	resty          *resty.Client
	authToken      string
	queryParam     url.Values
	multipartFiles []MultipartFile
}

func (h *HttpClient) SetFileReader(param string, fileName string, reader io.Reader) IHttpClient {
	h.multipartFiles = append(h.multipartFiles, MultipartFile{
		Param:    param,
		FileName: fileName,
		Reader:   reader,
	})
	return h
}

func (h *HttpClient) SetQueryParam(key string, value string) IHttpClient {
	h.queryParam.Set(key, value)
	return h
}

func (h *HttpClient) Post(s string) (resp Response, err error) {
	r := h.resty.R()
	if h.authToken != "" {
		r.SetAuthToken(h.authToken)
	}
	r.SetQueryParamsFromValues(h.queryParam)
	for _, file := range h.multipartFiles {
		r = r.SetFileReader(file.Param, file.FileName, file.Reader)
	}
	post, err := r.Post(s)
	if err != nil {
		return
	}
	resp = toResponse(post)
	return
}

func (h *HttpClient) SetAuthToken(s string) IHttpClient {
	h.authToken = s
	return h
}

func NewHttpClient() IHttpClient {
	_resty := resty.New()
	return &HttpClient{
		resty:          _resty,
		queryParam:     url.Values{},
		multipartFiles: []MultipartFile{},
	}
}

func (h *HttpClient) Get(url string) (resp Response, err error) {
	r := h.resty.R()
	if h.authToken != "" {
		r.SetAuthToken(h.authToken)
	}
	r = r.SetQueryParamsFromValues(h.queryParam)
	response, err := r.Get(url)
	resp = toResponse(response)
	return
}

func toResponse(response *resty.Response) (resp Response) {
	resp.Body = response.Body()
	return
}
