package pcloud

import (
	"bytes"
	"encoding/json"
	"errors"
	"my-imgur/lib/upload_client"
	"my-imgur/utils/http-client"
	"net/url"
	"os"
	"strconv"
)

//go:generate mockgen -destination=pcloud_mock.go -package=upload-client . IClient
type Client struct {
	accessToken string
	httpClient  httpClient.IHttpClient
}

func (c *Client) GetFileLink(uploadPath upload_client.AllowUploadLocation, fileName string, fileId int, width int, height int) (link string, err error) {
	//TODO implement me
	panic("implement me")
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
	query.Set("code", resp.PCloud.Code)
	query.Set("size", strconv.Itoa(width)+"x"+strconv.Itoa(height))
	u.RawQuery = query.Encode()
	link = u.String()
	return
}

func (c *Client) GetFilePublicLink(filePath string, fileId int) (resp upload_client.GetFilePublicLinkResponse, err error) {
	if filePath != "" {
		c.httpClient = c.httpClient.SetQueryParam("path", filePath)
	}
	if fileId != 0 {
		c.httpClient = c.httpClient.SetQueryParam("fileid", strconv.Itoa(fileId))
	}
	response, err := c.httpClient.Get("https://api.pcloud.com/getfilepublink")
	if err != nil {
		return upload_client.GetFilePublicLinkResponse{}, err
	}
	err = json.Unmarshal(response.Body, &resp)
	if err != nil {
		return upload_client.GetFilePublicLinkResponse{}, err
	}
	return
}

func (c *Client) UploadFile(uploadPath upload_client.AllowUploadLocation, fileName string, data []byte, option upload_client.UploadFileOption) (resp upload_client.UploadFileResponse, err error) {
	id := GetFileId(uploadPath)
	c.httpClient.SetQueryParam("folderid", strconv.Itoa(id))
	if option.PCloud.RenameIfExists {
		c.httpClient.SetQueryParam("renameifexists", strconv.Itoa(1))
	}
	if option.PCloud.CTime != 0 {
		c.httpClient.SetQueryParam("ctime", strconv.Itoa(option.PCloud.CTime))
	}
	if option.PCloud.MTime != 0 {
		c.httpClient.SetQueryParam("mtime", strconv.Itoa(option.PCloud.MTime))
	}
	if option.PCloud.NoPartial {
		c.httpClient.SetQueryParam("nopartial", strconv.Itoa(1))
	}
	if option.PCloud.ProgressHash != "" {
		c.httpClient.SetQueryParam("progresshash", option.PCloud.ProgressHash)
	}
	c.httpClient.SetFileReader("filename", fileName, bytes.NewReader(data))
	defer c.httpClient.ResetQueryParam()
	defer c.httpClient.ResetFileReader()
	post, err := c.httpClient.Post("https://api.pcloud.com/uploadfile")
	if err != nil {
		return
	}
	err = json.Unmarshal(post.Body, &resp)
	if resp.PCloud.Result != 0 {
		err = errors.New(resp.PCloud.Error)
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

func NewClient() upload_client.IClient {
	iHttpClient := httpClient.NewHttpClient()
	iHttpClient.SetAuthToken(os.Getenv("PCLOUD_ACCESS_TOKEN"))
	return &Client{
		httpClient: iHttpClient,
	}
}
