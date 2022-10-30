package model

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"my-imgur/lib/upload_client"
)

var _ = Describe("Image", func() {
	When("PCloud UploadFile", func() {
		It("Should escape url encoding", func() {
			mockCtrl := gomock.NewController(GinkgoT())
			client := upload_client.NewMockIClient(mockCtrl)
			image := Image{
				uploadClient: client,
			}
			client.EXPECT().UploadFile(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return(upload_client.UploadFileResponse{
					PCloud: upload_client.PCloudUploadFileResponse{
						Metadata: []struct {
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
						}{struct {
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
						}{Name: "( ).jpeg"}},
					},
				}, nil)
			path, err := image.UploadFile("( ).jpeg", []byte{})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(path).Should(ContainSubstring("%28%20%29.jpeg"))
		})
	})
})
