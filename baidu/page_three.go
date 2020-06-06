package baidu

import (
	"io"
	"io/ioutil"
	"net/http"
	"spider/mode"
	"strings"
)

type Content struct {

}

func (c Content)Request(url string)[]mode.UrlParse{
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
	return c.parse(res.Body)
}

func (c Content)parse(body io.Reader)[]mode.UrlParse{
	b,err := ioutil.ReadAll(body)
	if err != nil {
		return []mode.UrlParse{}
	}
	strBody := string(b)
	parses := make([]mode.UrlParse,0)
	if strings.Contains(strBody,"7") {
		img := Image{}
		parse := mode.NewUrlParse(img,"http://www.test1.com")
		parses = append(parses,parse)
	}
	if strings.Contains(strBody,"8") {
		img := Music{}
		parse := mode.NewUrlParse(img,"http://www.test2.com")
		parses = append(parses,parse)
	}
	if strings.Contains(strBody,"9") {
		img := Content{}
		parse := mode.NewUrlParse(img,"http://www.test3.com")
		parses = append(parses,parse)
	}
	c.Save()
	return parses
}

func (c Content)Save()  {
	
}

