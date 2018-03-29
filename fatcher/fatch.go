package fatcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"log"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var rateLimiter = time.Tick(30 * time.Millisecond)

func Fatch(url string) ([]byte, error) {
	<-rateLimiter
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong http code: %d", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)

	e := determineEncoding(bodyReader)
	utf8NewReader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8NewReader)
}

// 判断网页编码格式.  test ide to push
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("fatcher error：%v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
