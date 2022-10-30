package pcloud

import (
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"my-imgur/lib/upload_client"
	"os"
)

var _ = Describe("PCloud", func() {
	var PCloudClient upload_client.IClient

	PWhen("Real Upload", func() {
		var PCLOUD_ACCESS_TOKEN string
		BeforeEach(func() {
			err := godotenv.Load("../../.env.development")
			Expect(err).ShouldNot(HaveOccurred())
			PCLOUD_ACCESS_TOKEN = os.Getenv("PCLOUD_ACCESS_TOKEN")
			PCloudClient = NewClient()
		})

		haveAuthorization := func(b bool) {
			err := PCloudClient.CheckAuthorization()
			if b {
				Expect(err).ShouldNot(HaveOccurred())
			} else {
				Expect(err).Should(HaveOccurred())
			}
		}

		Context("Check Auth", func() {
			It("Success", func() {
				PCloudClient.SetAccessToken(PCLOUD_ACCESS_TOKEN)
				haveAuthorization(true)
			})
			It("Fail", func() {
				haveAuthorization(false)
			})
		})

		Context("File Upload", func() {
			// https://docs.pcloud.com/methods/file/uploadfile.html
			BeforeEach(func() {
				PCloudClient.SetAccessToken(PCLOUD_ACCESS_TOKEN)
				haveAuthorization(true)
			})
			It("Success", func() {
				data, _ := ioutil.ReadFile("./wakuwaku.jpeg")
				resp, err := PCloudClient.UploadFile(upload_client.OTHER, "wakuwaku.jpeg", data, upload_client.UploadFileOption{PCloud: upload_client.PCloudUploadFileOption{RenameIfExists: true}})
				Expect(resp.PCloud.Result).Should(Equal(0))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("getfilepublink ", func() {
			// https://docs.pcloud.com/methods/public_links/getfilepublink.html
			It("Success", func() {
				resp, err := PCloudClient.GetFilePublicLink("/Public Asset/obsidian/aaa.jpeg", 0)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(resp.PCloud.Metadata.Fileid).Should(Equal(int64(43946969297)))
			})
		})
		Context("getpubthumblink ", func() {
			// https://docs.pcloud.com/methods/public_links/getfilepublink.html
			It("Success", func() {
				link, err := PCloudClient.GetPublicThumbnail("/Public Asset/obsidian/aaa.jpeg", 0, 1024, 768)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(link).Should(Equal("https://api.pcloud.com/getpubthumb?code=XZaarSVZ9RtK8qBnsLjta2naRLzMQbjvCU97&size=1024x768"))
			})
		})
	})

})
