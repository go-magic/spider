package mode

type Spider interface {
	Request(url string)[]UrlParse
}

type UrlParse struct {
	Method Spider
	Url string
}

func NewUrlParse(Method Spider,Url string)UrlParse{
	return UrlParse{Method:Method,Url: Url}
}


type MulSpider interface {

}
