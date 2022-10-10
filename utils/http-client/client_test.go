package httpClient

import (
	"bytes"
	"encoding/json"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"os"
)

var _ = Describe("Http Client", func() {
	var _client IHttpClient
	When("real endpoint", func() {
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
		It("Get With Query Param", func() {
			_client.SetQueryParam("postId", "1")
			resp, err := _client.Get("https://jsonplaceholder.typicode.com/comments")
			var result []struct {
				PostId int    `json:"postId"`
				Id     int    `json:"id"`
				Name   string `json:"name"`
				Email  string `json:"email"`
				Body   string `json:"body"`
			}
			err = json.Unmarshal(resp.Body, &result)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(result)).Should(Equal(5))
		})
		It("Post With Query Param and auth token", func() {
			profileImgBytes, _ := ioutil.ReadFile("./wakuwaku.jpeg")
			_client.SetAuthToken(os.Getenv("PCLOUD_ACCESS_TOKEN"))
			_client.SetQueryParam("folderid", "14868593431")
			_client.SetFileReader("filename", "test_for_wakuwaku.jpeg", bytes.NewReader(profileImgBytes))
			resp, err := _client.Post("https://api.pcloud.com/uploadfile")
			Expect(err).ShouldNot(HaveOccurred())
			result := new(struct {
				Result   int `json:"result"`
				Metadata []struct {
					Name           string  `json:"name"`
					Created        string  `json:"created"`
					Thumb          bool    `json:"thumb"`
					Modified       string  `json:"modified"`
					Isfolder       bool    `json:"isfolder"`
					Height         int     `json:"height"`
					Fileid         int64   `json:"fileid"`
					Width          int     `json:"width"`
					Hash           float64 `json:"hash"`
					Path           string  `json:"path"`
					Category       int     `json:"category"`
					Id             string  `json:"id"`
					Isshared       bool    `json:"isshared"`
					Ismine         bool    `json:"ismine"`
					Size           int     `json:"size"`
					Parentfolderid int64   `json:"parentfolderid"`
					Contenttype    string  `json:"contenttype"`
					Icon           string  `json:"icon"`
				} `json:"metadata"`
				Checksums []struct {
					Sha1 string `json:"sha1"`
					Md5  string `json:"md5"`
				} `json:"checksums"`
				Fileids []int64 `json:"fileids"`
			})
			err = json.Unmarshal(resp.Body, &result)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result.Result).Should(Equal(0))
		})

	})
})
