package pcloud

import "my-imgur/lib/upload_client"

type allowUploadLocation struct {
	fileId int
	path   string
}

var allowUploadLocationList = map[upload_client.AllowUploadLocation]allowUploadLocation{
	upload_client.OBSIDIAN: {14875353183, "/Public Asset/obsidian"},
	upload_client.OTHER:    {14874889844, "/Public Asset/other"},
}

func GetFileId(b upload_client.AllowUploadLocation) int {
	return allowUploadLocationList[b].fileId
}

func GetPath(b upload_client.AllowUploadLocation) string {
	return allowUploadLocationList[b].path
}
