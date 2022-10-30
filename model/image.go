package model

import (
	"my-imgur/lib/upload_client"
	"my-imgur/lib/upload_client/google_drive"
	"net/url"
	"os"
)

type Image struct {
	uploadClient upload_client.IClient
}

func (i *Image) GetPublicThumbnailLink(fileName string, fileId int, width int, height int) (link string, err error) {
	link, err = i.uploadClient.GetFileLink(upload_client.OBSIDIAN, fileName, fileId, width, height)
	return
}

func (i *Image) UploadFile(fileName string, data []byte) (path string, err error) {
	res, err := i.uploadClient.UploadFile(upload_client.OBSIDIAN, fileName, data, upload_client.UploadFileOption{PCloud: upload_client.PCloudUploadFileOption{RenameIfExists: true}})
	if err != nil {
		return "", err
	}
	path = os.Getenv("PUBLIC_DOMAIN") + "/v1/temp-link/obsidian/" + url.PathEscape(res.GoogleDriveApi.Name)
	//path = os.Getenv("PUBLIC_DOMAIN") + "/obsidian/" + res.Metadata[0].Name
	return
}

func NewImage() IImage {
	client := google_drive.NewClient()
	return &Image{
		uploadClient: client,
	}
}

//go:generate mockgen -destination=image_mock.go -package=model . IImage
type IImage interface {
	UploadFile(fileName string, data []byte) (path string, err error)
	GetPublicThumbnailLink(fileName string, fileId int, width int, height int) (link string, err error)
}
