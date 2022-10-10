package router

import (
	"github.com/gavv/httpexpect/v2"
	"github.com/golang/mock/gomock"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"my-imgur/model"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Router", func() {
	var (
		e          *echo.Echo
		server     *httptest.Server
		testServer *httpexpect.Expect
	)
	PWhen("real upload", func() {
		BeforeEach(func() {
			err := godotenv.Load("../.env.development")
			Expect(err).ShouldNot(HaveOccurred())
			e = echo.New()
			server = httptest.NewServer(e)
			testServer = httpexpect.WithConfig(httpexpect.Config{
				BaseURL:  server.URL,
				Reporter: httpexpect.NewAssertReporter(GinkgoT()),
				Printers: []httpexpect.Printer{
					httpexpect.NewDebugPrinter(GinkgoT(), true),
				},
			})
		})
		AfterEach(func() {
			defer server.Close()
		})
		It("Upload file and get path", func() {
			router := NewRouter()
			data, _ := ioutil.ReadFile("./wakuwaku.jpeg")
			e.POST("/", router.UploadPicture)
			testServer.POST("/").
				WithMultipart().
				WithFileBytes("image", "wakuwaku.jpeg", data).
				Expect().Status(http.StatusOK).JSON().Object().
				Value("data").Object().Value("link").
				Equal("test")
		})

	})
	When("Mock", func() {
		var mockCtrl *gomock.Controller

		BeforeEach(func() {
			e = echo.New()
			server = httptest.NewServer(e)
			testServer = httpexpect.WithConfig(httpexpect.Config{
				BaseURL:  server.URL,
				Reporter: httpexpect.NewAssertReporter(GinkgoT()),
				Printers: []httpexpect.Printer{
					httpexpect.NewDebugPrinter(GinkgoT(), true),
				},
			})
		})
		AfterEach(func() {
			defer server.Close()
			defer mockCtrl.Finish()
		})
		It("Upload file and get path", func() {
			mockCtrl = gomock.NewController(GinkgoT())
			mockImage := model.NewMockIImage(mockCtrl)
			mockImage.EXPECT().UploadFile("wakuwaku.jpeg", gomock.Any()).
				Return("valid path", nil)
			router := Router{
				imageModel: mockImage,
			}

			data, _ := ioutil.ReadFile("./wakuwaku.jpeg")
			e.POST("/", router.UploadPicture)

			testServer.POST("/").
				WithMultipart().
				WithFileBytes("image", "wakuwaku.jpeg", data).
				Expect().Status(http.StatusOK).JSON().Object().
				Value("data").Object().Value("link").
				Equal("valid path")
		})
	})
})
