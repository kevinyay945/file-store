package pcloud

import (
	"bytes"
	"encoding/json"
	"errors"
	"my-imgur/utils/http-client"
	"net/url"
	"os"
	"strconv"
)

//go:generate mockgen -destination=pcloud_mock.go -package=pcloud . IClient
type IClient interface {
	SetAccessToken(s string)
	CheckAuthorization() (err error)
	UploadFile(uploadPath AllowUploadLocation, fileName string, data []byte, option UploadFileOption) (resp UploadFileResponse, err error)
	GetFilePublicLink(filePath string, fileId int) (resp GetFilePublicLinkResponse, err error)
	GetPublicThumbnail(filePath string, fileId int, width int, height int) (link string, err error)
}
type UploadFileOption struct {
	NoPartial      bool
	ProgressHash   string
	RenameIfExists bool
	MTime          int
	CTime          int
}
type GetFilePublicLinkResponse struct {
	Code            string `json:"code"`
	Created         string `json:"created"`
	Downloadenabled bool   `json:"downloadenabled"`
	Type            int    `json:"type"`
	Modified        string `json:"modified"`
	Downloads       int    `json:"downloads"`
	Link            string `json:"link"`
	Result          int    `json:"result"`
	Linkid          int    `json:"linkid"`
	Haspassword     bool   `json:"haspassword"`
	Traffic         int    `json:"traffic"`
	Views           int    `json:"views"`
	Metadata        struct {
		Name           string  `json:"name"`
		Created        string  `json:"created"`
		Thumb          bool    `json:"thumb"`
		Modified       string  `json:"modified"`
		Isfolder       bool    `json:"isfolder"`
		Height         int     `json:"height"`
		Fileid         int64   `json:"fileid"`
		Width          int     `json:"width"`
		Hash           float64 `json:"hash"`
		Comments       int     `json:"comments"`
		Category       int     `json:"category"`
		Id             string  `json:"id"`
		Isshared       bool    `json:"isshared"`
		Ismine         bool    `json:"ismine"`
		Size           int     `json:"size"`
		Parentfolderid int64   `json:"parentfolderid"`
		Contenttype    string  `json:"contenttype"`
		Icon           string  `json:"icon"`
	} `json:"metadata"`
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
	Error string `json:"error"`
}

type Client struct {
	accessToken string
	httpClient  httpClient.IHttpClient
}

func (c *Client) GetPublicThumbnail(filePath string, fileId int, width int, height int) (link string, err error) {
	resp, err := c.GetFilePublicLink(filePath, fileId)
	if err != nil {
		return "", err
	}
	u := url.URL{
		Scheme: "https",
		Host:   "api.pcloud.com",
		Path:   "getpubthumb",
	}
	query := u.Query()
	query.Set("code", resp.Code)
	query.Set("size", strconv.Itoa(width)+"x"+strconv.Itoa(height))
	u.RawQuery = query.Encode()
	link = u.String()
	return
}

func (c *Client) GetFilePublicLink(filePath string, fileId int) (resp GetFilePublicLinkResponse, err error) {
	if filePath != "" {
		c.httpClient = c.httpClient.SetQueryParam("path", filePath)
	}
	if fileId != 0 {
		c.httpClient = c.httpClient.SetQueryParam("fileid", strconv.Itoa(fileId))
	}
	response, err := c.httpClient.Get("https://api.pcloud.com/getfilepublink")
	if err != nil {
		return GetFilePublicLinkResponse{}, err
	}
	err = json.Unmarshal(response.Body, &resp)
	if err != nil {
		return GetFilePublicLinkResponse{}, err
	}
	return
}

func (c *Client) UploadFile(uploadPath AllowUploadLocation, fileName string, data []byte, option UploadFileOption) (resp UploadFileResponse, err error) {
	id := uploadPath.FileId()
	c.httpClient.SetQueryParam("folderid", strconv.Itoa(id))
	if option.RenameIfExists {
		c.httpClient.SetQueryParam("renameifexists", strconv.Itoa(1))
	}
	if option.CTime != 0 {
		c.httpClient.SetQueryParam("ctime", strconv.Itoa(option.CTime))
	}
	if option.MTime != 0 {
		c.httpClient.SetQueryParam("mtime", strconv.Itoa(option.MTime))
	}
	if option.NoPartial {
		c.httpClient.SetQueryParam("nopartial", strconv.Itoa(1))
	}
	if option.ProgressHash != "" {
		c.httpClient.SetQueryParam("progresshash", option.ProgressHash)
	}
	c.httpClient.SetFileReader("filename", fileName, bytes.NewReader(data))
	defer c.httpClient.ResetQueryParam()
	defer c.httpClient.ResetFileReader()
	post, err := c.httpClient.Post("https://api.pcloud.com/uploadfile")
	if err != nil {
		return
	}
	err = json.Unmarshal(post.Body, &resp)
	if resp.Result != 0 {
		err = errors.New(resp.Error)
		return
	}
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
	iHttpClient.SetAuthToken(os.Getenv("PCLOUD_ACCESS_TOKEN"))
	return &Client{
		httpClient: iHttpClient,
	}
}
