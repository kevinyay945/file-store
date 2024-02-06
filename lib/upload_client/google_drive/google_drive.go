package google_drive

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"io"
	"log"
	"my-imgur/lib/upload_client"
	"os"
	"strings"
)

type Client struct {
	driveService *drive.Service
}

type allowUploadLocation struct {
	folderId string
	path     string
}

var allowUploadLocationList = map[upload_client.AllowUploadLocation]allowUploadLocation{
	upload_client.OBSIDIAN: {"12o01NrAVom6FSrQgKfQVsR2G5H0dVsd3", "/public-asset"},
	//upload_client.OTHER:    {14874889844, "/Public Asset/other"},
}

func GetFolderID(b upload_client.AllowUploadLocation) string {
	return allowUploadLocationList[b].folderId
}

func GetPath(b upload_client.AllowUploadLocation) string {
	return allowUploadLocationList[b].path
}

func NewClient() upload_client.IClient {
	token := os.Getenv("GOOGLE_DRIVE_API_TOKEN")
	client := Client{}
	client.SetAccessToken(token)
	return &client
}

func (c *Client) GetFileLink(uploadPath upload_client.AllowUploadLocation, fileName string, fileId int, width int, height int) (link string, err error) {
	query := fmt.Sprintf("name = '%s'", fileName)
	r, err := c.driveService.Files.List().PageSize(10).Q(query).
		Fields("nextPageToken, files(id, name,thumbnailLink)").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}
	fmt.Println("Files:")
	if len(r.Files) == 0 {
		fmt.Println("No files found.")
	} else {
		for _, i := range r.Files {
			//fmt.Printf("%s (%s)\n", i.Name, i.Id)
			//fmt.Printf("webview link: %v\n", i.WebViewLink)
			//fmt.Printf("webcontent link: %v\n", i.WebContentLink)
			//fmt.Printf("thumbnail link: %v\n", i.ThumbnailLink)
			//fmt.Printf("exportLink link: %v\n", i.ExportLinks)
			if err != nil {
				fmt.Println("err", err.Error())
			}
			link = strings.Split(i.ThumbnailLink, "=")[0]
			//link = fmt.Sprintf("https://lh3.google.com/u/0/d/%s\n", i.Id)
			//link = fmt.Sprintf("https://drive.google.com/uc?export=view&id=%s\n", i.Id)
			//fmt.Printf("share link is %s\n", link)
		}
	}
	return
}
func (c *Client) SetAccessToken(s string) {
	bytes := []byte(s)
	c.driveService, _ = drive.NewService(context.Background(), option.WithCredentialsJSON(bytes))
}

func (c *Client) CheckAuthorization() (err error) {
	_, err = c.driveService.Files.List().PageSize(10).Fields("nextPageToken, files(id, name)").Do()
	return
}

func (c *Client) UploadFile(uploadPath upload_client.AllowUploadLocation, fileName string, data []byte, option upload_client.UploadFileOption) (resp upload_client.UploadFileResponse, err error) {
	folderId := GetFolderID(uploadPath)
	buffer := bytes.NewBuffer(data)
	resultOfFileCreation, err := createFile(c.driveService, fileName, "image/png", buffer, folderId)
	if err != nil {
		return
	}
	fmt.Println("create file result:", resultOfFileCreation)
	resp.GoogleDriveApi = resultOfFileCreation
	permission := drive.Permission{
		Type: "anyone",
		Role: "reader",
	}
	_, err = c.driveService.Permissions.Create(resultOfFileCreation.Id, &permission).Do()
	if err != nil {
		return
	}
	return
}

func (c *Client) GetFilePublicLink(filePath string, fileId int) (resp upload_client.GetFilePublicLinkResponse, err error) {
	r, err := c.driveService.Files.List().PageSize(10).Q("name = 'eac06f60-3830-11ed-bc4d-d2e490ac16f8.png'").
		Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}
	fmt.Println("Files:")
	if len(r.Files) == 0 {
		fmt.Println("No files found.")
	} else {
		for _, i := range r.Files {
			fmt.Printf("%s (%s)\n", i.Name, i.Id)
			if err != nil {
				fmt.Println("err", err.Error())
			}
			resp.Link = fmt.Sprintf("share link is https://drive.google.com/uc?export=view&id=%s\n", i.Id)
			//fmt.Printf("share link is https://drive.google.com/uc?export=view&id=%s\n", i.Id)
		}
	}
	return
}

func (c *Client) GetPublicThumbnail(filePath string, fileId int, width int, height int) (link string, err error) {
	//TODO implement me
	panic("not support")
}
func createFile(service *drive.Service, name string, mimeType string, content io.Reader, parentID string) (*drive.File, error) {
	parents := []string{}
	if parentID != "" {
		parents = []string{parentID}
	}
	f := &drive.File{
		MimeType: mimeType,
		Name:     name,
		Parents:  parents,
	}
	file, err := service.Files.Create(f).Media(content).Do()

	if err != nil {
		log.Println("Could not create file: " + err.Error())
		return nil, err
	}

	return file, nil
}
