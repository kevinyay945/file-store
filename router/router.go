package router

import (
	"bytes"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"mime/multipart"
	"my-imgur/lib/pcloud"
	"my-imgur/model"
	"net/http"
)

type MyResponse struct {
	Data struct {
		Link string `json:"link"`
	} `json:"data"`
}

type Router struct {
	pCloudClient pcloud.IClient
	imageModel   model.IImage
}

func NewRouter() *Router {
	client := pcloud.NewClient()
	image := model.NewImage()
	return &Router{
		pCloudClient: client,
		imageModel:   image,
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
	path, err := r.imageModel.UploadFile(file.Filename, fileBytes.Bytes())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	response := MyResponse{}

	response.Data.Link = path
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, response)
}
func (r *Router) GetPublicThumbnailLink(c echo.Context) error {
	param := c.Param("fileName")
	link, err := r.imageModel.GetPublicThumbnailLink(param, 0, 1024, 768)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.Redirect(http.StatusFound, link)
}
