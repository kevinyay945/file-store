package httpClient

import (
	"encoding/json"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("Http Client", func() {
	var _client IHttpClient
	BeforeEach(func() {
		_client = NewHttpClient()
		err := godotenv.Load("../../.env.development")
		Expect(err).ShouldNot(HaveOccurred())
	})
	It("Get", func() {
		resp, err := _client.Get("https://jsonplaceholder.typicode.com/todos/1")
		Expect(err).ShouldNot(HaveOccurred())
		f := new(struct {
			UserId    int    `json:"userId"`
			Id        int    `json:"id"`
			Title     string `json:"title"`
			Completed bool   `json:"completed"`
		})
		err = json.Unmarshal(resp.Body, f)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(*f).Should(Equal(struct {
			UserId    int    `json:"userId"`
			Id        int    `json:"id"`
			Title     string `json:"title"`
			Completed bool   `json:"completed"`
		}{
			UserId:    1,
			Id:        1,
			Title:     "delectus aut autem",
			Completed: false,
		}))
	})
	It("Get With Auth Token", func() {
		_client = _client.SetAuthToken(os.Getenv("PCLOUD_ACCESS_TOKEN"))
		resp, err := _client.Get("https://api.pcloud.com/userinfo")
		Expect(err).ShouldNot(HaveOccurred())
		result := new(struct {
			Result int `json:"result"`
		})
		err = json.Unmarshal(resp.Body, &result)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(result.Result).Should(Equal(0))

	})
})
