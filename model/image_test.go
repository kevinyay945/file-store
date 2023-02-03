package model

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/api/drive/v3"
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
			file := drive.File{
				Name: "( ).jpeg",
			}
			client.EXPECT().UploadFile(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return(upload_client.UploadFileResponse{
					GoogleDriveApi: &file,
				}, nil)
			path, err := image.UploadFile("( ).jpeg", []byte{})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(path).Should(ContainSubstring("%28%20%29.jpeg"))
		})
	})
})
