package pcloud

type AllowUploadLocation int

const (
	OBSIDIAN AllowUploadLocation = iota
	OTHER
)

type allowUploadLocation struct {
	fileId int
	path   string
}

var allowUploadLocationList = map[AllowUploadLocation]allowUploadLocation{
	OBSIDIAN: {14875353183, "/Public Asset/obsidian"},
	OTHER:    {14874889844, "/Public Asset/other"},
}

func (b AllowUploadLocation) FileId() int {
	return allowUploadLocationList[b].fileId
}

func (b AllowUploadLocation) Path() string {
	return allowUploadLocationList[b].path
}

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
