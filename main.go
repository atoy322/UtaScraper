package main

import (
	"fmt"
	"net/http"
	"regexp"
)

func main() {
	response, e := http.Get("https://www.uta-net.com/search/?Aselect=2&Bselect=3&Keyword=F.O.O.L")
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
	re1, e1 := regexp.Compile("<tr>.+?</tr>")
    re2, e2 := regexp.Compile("<td class=\".*?td[1-2]\">.+?</td>")
    re3, e3 := regexp.Compile("/song/[0-9]+")
    re4, e4 := regexp.Compile("<.+?>")
	if e1 != nil {
		panic(e1)
	} else if e2 != nil {
        panic(e2)
    } else if e3 != nil {
        panic(e3)
    } else if e4 != nil {
        panic(e4)
    }
	table_rows := re1.FindAllString(html_string, -1)

	for i := range table_rows {
        table_datas := re2.FindAllString(table_rows[i], -1)
        path := re3.FindString(table_datas[0])
        name := re4.ReplaceAllString(table_datas[0], "")
        musician := re4.ReplaceAllString(table_datas[1], "")
        fmt.Printf("u=%s n=%s m=%s\n", path, name, musician)
	}
}
