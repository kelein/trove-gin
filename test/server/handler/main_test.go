package handler

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"

	"github.com/kelein/trove-gin/internal/handler"
	"github.com/kelein/trove-gin/internal/middleware"
	"github.com/kelein/trove-gin/pkg/config"
	jwt2 "github.com/kelein/trove-gin/pkg/jwt"
	"github.com/kelein/trove-gin/pkg/log"
)

var (
	userId = "xxx"
)

var logger *log.Logger

var hdl *handler.Handler

var jwt *jwt2.JWT

var router *gin.Engine

func TestMain(m *testing.M) {
	fmt.Println("begin")
	err := os.Setenv("APP_CONF", "../../../config/local.yml")
	if err != nil {
		fmt.Println("Setenv error", err)
	}
	var envConf = flag.String("conf", "config/local.yml", "config path, eg: -conf ./config/local.yml")
	flag.Parse()
	conf := config.NewConfig(*envConf)

	// modify log directory
	logPath := filepath.Join("../../../", conf.GetString("log.log_file"))
	conf.Set("log.log_file", logPath)

	logger = log.NewLog(conf)
	hdl = handler.NewHandler(logger)

	jwt = jwt2.NewJwt(conf)
	gin.SetMode(gin.TestMode)
	router = gin.Default()
	router.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		//middleware.SignMiddleware(log),
	)

	code := m.Run()
	fmt.Println("test end")

	os.Exit(code)
}

func performRequest(r http.Handler, method, path string, body *bytes.Buffer) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	return resp
}

func genToken(t *testing.T) string {
	token, err := jwt.GenToken(userId, time.Now().Add(time.Hour*24*90))
	if err != nil {
		t.Error(err)
		return token
	}
	return token
}

func newHttpExcept(t *testing.T, router *gin.Engine) *httpexpect.Expect {
	return httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(router),
			Jar:       httpexpect.NewCookieJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			// httpexpect.NewDebugPrinter(t, true),
		},
	})
}
