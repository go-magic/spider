package mode

type Engine struct {
	Parses []UrlParse
	End chan bool
}

func NewEngine(Parses []UrlParse,End chan bool)Engine{
	return Engine{Parses:Parses,End: End}
}

func (e Engine)Start(){
	if len(e.Parses) == 0 {
		return
	}
	parses := make([]UrlParse,0)
	parses = append(parses,e.Parses...)
	for len(parses) > 0 {
		p := parses[0]
		parses = parses[1:]
		select {
		case <-e.End:
			return
		default:
			ps := p.Method.Request(p.Url)
			for _,v := range ps {
				if v.Url != "" {
					parses = append(parses,v)
				}
			}
		}
	}
}
