package main

import (
	"fmt"
	"net/http"
	"regexp"
)

func main() {
	response, e := http.Get("https://www.uta-net.com/search/?Aselect=2&Bselect=3&Keyword=炎")
	if e != nil {
		panic(e)
	}

	html := make([]byte, 1024)
	for {
		buf := make([]byte, 1024)
		n, e := response.Body.Read(buf)
		html = append(html, buf[:n]...)

		if (n == 0) || (e != nil) {
			break
		}
	}

	html_string := string(html)
	re, e := regexp.Compile("<tr>.+?</tr>")
	if e != nil {
		panic(e)
	}
	trs := re.FindAllString(html_string, -1)

	for i := range trs {
		fmt.Println(trs[i])
	}
}
