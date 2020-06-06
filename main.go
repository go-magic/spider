package main

import (
	"os"
	"os/signal"
	"spider/baidu"
	"spider/mode"
)

func Exit(encChan chan bool){
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	encChan <- true
}

func main(){
	i := baidu.Image{}
	url := "http://www.baidu.com"
	parse := mode.NewUrlParse(i,url)
	parses := make([]mode.UrlParse,0)
	parses = append(parses,parse)
	encChan := make(chan bool,1)
	go Exit(encChan)
	e := mode.NewEngine(parses,encChan)
	e.Start()
}