package upload_client

import "google.golang.org/api/drive/v3"

type UploadFileOption struct {
	PCloud PCloudUploadFileOption
}
type UploadFileResponse struct {
	PCloud         PCloudUploadFileResponse
	GoogleDriveApi *drive.File
}
type GetFilePublicLinkResponse struct {
	PCloud PCloudGetFilePublicLinkResponse
	Link   string
}

//go:generate mockgen -destination=upload_client_mock.go -package=upload_client . IClient
type IClient interface {
	SetAccessToken(s string)
	CheckAuthorization() (err error)
	UploadFile(uploadPath AllowUploadLocation, fileName string, data []byte, option UploadFileOption) (resp UploadFileResponse, err error)
	GetFilePublicLink(filePath string, fileId int) (resp GetFilePublicLinkResponse, err error)
	GetPublicThumbnail(filePath string, fileId int, width int, height int) (link string, err error)
	GetFileLink(uploadPath AllowUploadLocation, fileName string, fileId int, width int, height int) (link string, err error)
}
