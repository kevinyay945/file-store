package upload_client

type AllowUploadLocation int

const (
	OBSIDIAN AllowUploadLocation = iota
	OTHER
)

//func All() []AllowUploadLocation {
//	return []AllowUploadLocation{OBSIDIAN, OTHER}
//}

//func Of(fileId int) (AllowUploadLocation, error) {
//	for _, b := range All() {
//		if b.FileId() == fileId {
//			return b, nil
//		}
//	}
//	return -1, errors.New("No match")
//}
