package parser

import (
	"regexp"
	"strings"
)

type Songs struct {
	SongName string
	Name     string
	URL      string
}

func ParseSongs(html_string string) []Songs {
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
	result_map := make([]Songs, len(table_rows))

	for i := range table_rows {
		table_datas := re2.FindAllString(table_rows[i], -1)
		path := re3.FindString(table_datas[0])
		name := re4.ReplaceAllString(table_datas[0], "")
		musician := re4.ReplaceAllString(table_datas[1], "")
		songs := &Songs{SongName: name, Name: musician, URL: "https://www.uta-net.com" + path}
		result_map[i] = *songs
	}

	return result_map
}

func ParseKashi(html_string string) string {
	re1, e1 := regexp.Compile("<div id=\"kashi_area\" itemprop=\"text\">.+?</div>")
	re2, e2 := regexp.Compile("<.+?>")
	if e1 != nil {
		panic(e1)
	} else if e2 != nil {
		panic(e2)
	}

	kashi := re1.FindString(html_string)
	kashi = strings.ReplaceAll(kashi, "<br>", "\r\n")
	kashi = strings.ReplaceAll(kashi, "<br />", "\r\n")
	kashi = re2.ReplaceAllString(kashi, "")
	return kashi
}
