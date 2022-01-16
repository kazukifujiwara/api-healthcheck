package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
)

const wait_seconds int64 = 3

type ApiEndpoint struct {
	Id          string `csv:"id", json:"id"`
	Name        string `csv:"name", json::"name"`
	Description string `csv:"description", json:"description"`
	Uri         string `csv:"uri", json:"uri"`
	Status      int    `default:0, json:"status"`
}

func checkApiStatus() []*ApiEndpoint {

	file, err := os.OpenFile("./apitest/apiendpoints.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	getStatusCode := func(apiEndpoint *ApiEndpoint) {
		url := apiEndpoint.Uri
		req, _ := http.NewRequest("GET", url, nil)
		// req.Header.Set("Authorization", "Bearer access-token")
		client := new(http.Client)
		resp, _ := client.Do(req)
		apiEndpoint.Status = resp.StatusCode

		jsonData, _ := json.Marshal(apiEndpoint)
		fmt.Println(string(jsonData))
	}

	apiEndpoints := []*ApiEndpoint{}
	if err := gocsv.UnmarshalFile(file, &apiEndpoints); err != nil {
		panic(err)
	}

	fmt.Println(len(apiEndpoints))

	for _, apiEndpoint := range apiEndpoints {
		go getStatusCode(apiEndpoint)
	}

	time.Sleep(1 * time.Second)

	return apiEndpoints
}

func main() {

	r := gin.Default()

	// ここからCorsの設定
	r.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"*",
		},
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできない)
		AllowMethods: []string{
			"GET",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: false,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))

	r.GET("/apiStatusCheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": checkApiStatus(),
		})
	})
	r.Run()
}
