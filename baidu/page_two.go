package baidu

import (
	"io"
	"io/ioutil"
	"net/http"
	"spider/mode"
	"strings"
)

type Music struct {

}

func (m Music)Request(url string)[]mode.UrlParse{
	var (
		err error
		req *http.Request
		res *http.Response
		client http.Client
	)
	req,err = http.NewRequest("GET",url,nil)
	if err != nil {
		return []mode.UrlParse{}
	}
	res,err = client.Do(req)
	if err != nil {
		return []mode.UrlParse{}
	}
	return m.parse(res.Body)
}

func (m Music)parse(body io.Reader)[]mode.UrlParse{
	b,err := ioutil.ReadAll(body)
	if err != nil {
		return []mode.UrlParse{}
	}
	strBody := string(b)
	parses := make([]mode.UrlParse,0)
	if strings.Contains(strBody,"4") {
		img := Image{}
		parse := mode.NewUrlParse(img,"http://www.test4.com")
		parses = append(parses,parse)
	}
	if strings.Contains(strBody,"5") {
		img := Music{}
		parse := mode.NewUrlParse(img,"http://www.test5.com")
		parses = append(parses,parse)
	}
	if strings.Contains(strBody,"6") {
		img := Content{}
		parse := mode.NewUrlParse(img,"http://www.test6.com")
		parses = append(parses,parse)
	}
	m.Save()
	return parses
}

func (m Music)Save()  {

}