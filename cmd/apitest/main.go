package main

import (
    "encoding/json"
    "fmt"
    "os"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/gocarina/gocsv"
)

const wait_seconds int64 = 1

type ApiEndpoint struct {
    Id string `csv:"id", json:"id"`
    Name string `csv:"name", json::"name"`
    Description string `csv:"description", json:"description"`
    Uri string `csv:"uri", json:"uri"`
    Status int `default:0, json:"status"`
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
    r.GET("/apiStatusCheck", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "data": checkApiStatus(),
        })
    })
    r.Run()
}