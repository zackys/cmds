package encode

import (
	"code.google.com/p/go.text/encoding/japanese"
	"code.google.com/p/go.text/transform"
	"io/ioutil"
	"strings"
)

// UTF-8 から ShiftJIS
func UTF8_to_SJIS(str string) (string, error) {
        ret, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(str), japanese.ShiftJIS.NewEncoder()))
        if err != nil {
                return "", err
        }
        return string(ret), err
}


// ShiftJIS から UTF-8
func SJIS_to_UTF8(str string) (string, error) {
        ret, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(str), japanese.ShiftJIS.NewDecoder()))
        if err != nil {
                return "", err
        }
        return string(ret), err
}