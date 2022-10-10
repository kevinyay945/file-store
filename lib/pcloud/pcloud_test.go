package pcloud

import (
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"os"
)

var _ = Describe("PCloud", func() {
	var PCloudClient IClient

	When("Real Upload", func() {
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
				resp, err := PCloudClient.UploadFile(OTHER, "wakuwaku.jpeg", data)
				Expect(resp.Result).Should(Equal(0))
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
	})

})
