package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"time"
)

type BaseDetailResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Ocr struct {
	Tips    string `json:"tips"`
	Content string `json:"content"`
	URL     string `json:"url"`
}

func main() {
	http.HandleFunc("/ocr", hello)
	http.ListenAndServe(":8000", nil)
}
func hello(w http.ResponseWriter, req *http.Request) {
	vars := req.URL.Query()

	var field Ocr
	field.URL = vars["url"][0]

	filepath := download(field.URL)

	cmd := exec.Command("tesseract", filepath, "stdout", "-l", "chi_sim")
	// out, err := cmd.CombinedOutput()
	// if err != nil {
	// 	log.Fatalf("cmd.Run() failed with %s\n", err)
	// }

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	field.Tips = errStr
	field.Content = outStr

	var sendData BaseDetailResponse
	sendData.Code = 200
	sendData.Msg = "success"
	sendData.Data = field
	res, _ := json.Marshal(sendData)
	fmt.Fprintln(w, string(res))
}
func download(imgUrl string) string {
	imgPath := "/data/images/"
	year := time.Now().Format("2006")
	month := time.Now().Format("01")
	day := time.Now().Format("02")
	imgPath += year + "-" + month + "-" + day + "/"
	if !IsDir(imgPath) {
		fmt.Println("com in")
		Mkdir(imgPath)
	}

	fileName := path.Base(imgUrl)
	resp, err := http.Get(imgUrl)
	if err != nil {
		fmt.Fprint(os.Stderr, "get url error", err)
	}

	defer resp.Body.Close()

	out, err := os.Create(imgPath + fileName)
	wt := bufio.NewWriter(out)

	defer out.Close()

	n, err := io.Copy(wt, resp.Body)
	fmt.Println("write", n)
	if err != nil {
		panic(err)
	}
	wt.Flush()
	return imgPath + fileName
}

/**
* 判断目录是否存在
 */
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}
func Mkdir(path string) bool {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return false
	}
	return true
}
