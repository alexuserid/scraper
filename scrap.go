package main

import (
	"fmt"
	"flag"
	"net/http"
	"io/ioutil"
	"regexp"
)

var (
	link      string  = "http://google.com"
	link_num int     = 100
	field     string
)

func main() {
	// fmt.Scanf("%s", &link) //так не выходит, не меняется флаг с количеством ссылок. Почему?
	flag.StringVar(&link, "l", link, "ссылка")
	flag.IntVar(&link_num, "n", link_num, "количество необходимых ссылок")
	flag.Parse()
	new()
}

func new() {
	resp, _ := http.Get(link)//что тут вообще происходит..?
	bytes, _ := ioutil.ReadAll(resp.Body)//я еще не настолько хорошо знаю английский чтоб так хорошо понимать всю документацию((
	resp.Body.Close()//Опять же, зачем закрывать? 

	re := regexp.MustCompile("(http)+://[a-z.]+")
	resultlinks := re.FindAllString(string(bytes), link_num)
	for _, list := range resultlinks {
		if list == field {
			return
		} else {
			field = list
			fmt.Println(list)
		}
	}
}
	