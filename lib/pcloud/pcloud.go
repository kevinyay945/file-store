package pcloud

import (
	"bytes"
	"encoding/json"
	"errors"
	"my-imgur/utils/http-client"
	"strconv"
)

type IClient interface {
	SetAccessToken(s string)
	CheckAuthorization() (err error)
	UploadFile(uploadPath AllowUploadLocation, fileName string, data []byte) (resp UploadFileResponse, err error)
}

type UploadFileResponse struct {
	Result   int   `json:"result"`
	Fileids  []int `json:"fileids"`
	Metadata []struct {
		Ismine         bool    `json:"ismine"`
		Id             string  `json:"id"`
		Created        string  `json:"created"`
		Modified       string  `json:"modified"`
		Hash           float64 `json:"hash"`
		Isshared       bool    `json:"isshared"`
		Isfolder       bool    `json:"isfolder"`
		Category       int     `json:"category"`
		Parentfolderid int     `json:"parentfolderid"`
		Icon           string  `json:"icon"`
		Fileid         int     `json:"fileid"`
		Height         int     `json:"height"`
		Width          int     `json:"width"`
		Path           string  `json:"path"`
		Name           string  `json:"name"`
		Contenttype    string  `json:"contenttype"`
		Size           int     `json:"size"`
		Thumb          bool    `json:"thumb"`
	} `json:"metadata"`
}

type Client struct {
	accessToken string
	httpClient  httpClient.IHttpClient
}

func (c *Client) UploadFile(uploadPath AllowUploadLocation, fileName string, data []byte) (resp UploadFileResponse, err error) {
	id := uploadPath.FileId()
	c.httpClient.SetQueryParam("folderid", strconv.Itoa(id))
	defer c.httpClient.ResetQueryParam()
	c.httpClient.SetFileReader("filename", fileName, bytes.NewReader(data))
	defer c.httpClient.ResetFileReader()
	post, err := c.httpClient.Post("https://api.pcloud.com/uploadfile")
	if err != nil {
		return
	}
	err = json.Unmarshal(post.Body, &resp)
	return
}

func (c *Client) SetAccessToken(s string) {
	c.accessToken = s
	c.httpClient.SetAuthToken(s)
}

func (c *Client) CheckAuthorization() (err error) {
	resp, err := c.httpClient.Get("https://api.pcloud.com/userinfo")
	if err != nil {
		return err
	}
	result := new(struct {
		Result int    `json:"result"`
		Error  string `json:"error"`
	})
	err = json.Unmarshal(resp.Body, &result)
	if err != nil {
		return err
	}
	if result.Result != 0 {
		return errors.New(result.Error)
	}
	return
}

func NewClient() IClient {
	iHttpClient := httpClient.NewHttpClient()
	return &Client{
		httpClient: iHttpClient,
	}
}
