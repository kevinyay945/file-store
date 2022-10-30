package upload_client

type PCloudUploadFileOption struct {
	NoPartial      bool
	ProgressHash   string
	RenameIfExists bool
	MTime          int
	CTime          int
}
type PCloudGetFilePublicLinkResponse struct {
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
type PCloudUploadFileResponse struct {
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
