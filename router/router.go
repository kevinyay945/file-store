package router

import (
	"bytes"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"mime/multipart"
	"my-imgur/lib/pcloud"
	"net/http"
	"net/url"
	"os"
)

type MyResponse struct {
	Data struct {
		Link string `json:"link"`
	} `json:"data"`
}

type Router struct {
	pCloudClient pcloud.IClient
}

func NewRouter() *Router {
	client := pcloud.NewClient()
	return &Router{
		pCloudClient: client,
	}
}

func (r *Router) UploadPicture(c echo.Context) error {
	file, err := c.FormFile("image")
	open, err := file.Open()
	defer func(open multipart.File) {
		err := open.Close()
		if err != nil {
			fmt.Println("don't close the file open")
		}
	}(open)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	fileBytes := bytes.NewBuffer(nil)
	_, err = io.Copy(fileBytes, open)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	res, err := r.pCloudClient.UploadFile(pcloud.OBSIDIAN, file.Filename, fileBytes.Bytes(), pcloud.UploadFileOption{RenameIfExists: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	response := MyResponse{}
	response.Data.Link = os.Getenv("PUBLIC_DOMAIN") + "/obsidian/" + url.PathEscape(res.Metadata[0].Name)
	return c.JSON(http.StatusOK, response)
}
func (r *Router) GetPublicThumbnailLink(c echo.Context) error {
	param := c.Param("path")
	link, err := r.pCloudClient.GetPublicThumbnail("/Public Asset/obsidian/"+param, 0, 1024, 768)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.Redirect(http.StatusFound, link)
}
