package utils

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

func GbkToUtf8(str []byte) (b []byte, err error) {
	r := transform.NewReader(bytes.NewReader(str), simplifiedchinese.GB18030.NewDecoder())
	b, err = ioutil.ReadAll(r)
	return
}

func StrToUtf8(str string) string {
	b, _ := GbkToUtf8([]byte(str))
	return string(b)
}
