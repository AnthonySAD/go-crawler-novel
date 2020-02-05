package httpclient

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// 网页内容抓取函数
func Get(url string) ([]byte, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong state code: %d", resp.StatusCode)
	}

	return getContent(resp.Body), nil
}

//获取内容，并把内容转换成urf-8
func getContent(body io.ReadCloser) []byte{
	bodyReader := bufio.NewReader(body)
	bytes, err := bodyReader.Peek(1024)
	var e encoding.Encoding
	if err != nil {
		e = unicode.UTF8
	}else{
		e, _, _ = charset.DetermineEncoding(bytes, "")
	}
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	content, _ := ioutil.ReadAll(utf8Reader)
	return content
}

func Download(url string, filePath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
