package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

    "github.com/atoy322/UtaScraper/pkg/http_functions"
    "github.com/atoy322/UtaScraper/internal/parser"
)

func main() {
	if len(os.Args) == 1 {
		return
	}
	query := strings.Join(os.Args[1:], "%20")
	query = strings.Replace(query, "　", "%20", -1)

	html, e := http_functions.Get("https://www.uta-net.com/search/?Aselect=2&Bselect=3&Keyword=" + query)
	if e != nil {
		panic(e)
	}

	html_string := string(html)
	songs := parser.ParseSongs(html_string)

	for i := range songs {
		song := songs[i]
		fmt.Printf("[%3d] : %s / %s\n", i, song.SongName, song.Name)
	}

	fmt.Print("Song id : ")
	stdin_text := make([]byte, 1024)
	n, _ := os.Stdin.Read(stdin_text)
	line := string(stdin_text[:n])
	line = strings.Replace(line, "\n", "", -1)
	line = strings.Replace(line, "\r", "", -1)
	index, e := strconv.Atoi(line)
	if e != nil {
		panic(e)
	}
	target_song := songs[index]

	html, e = http_functions.Get(target_song.URL)
	if e != nil {
		panic(e)
	}

	html_string = string(html)
	kashi := parser.ParseKashi(html_string)

	fmt.Println(kashi)

	file, e := os.Create(target_song.SongName + " - " + target_song.Name + ".txt")
	if e != nil {
		panic(e)
	}
	defer file.Close()

	file.WriteString(kashi)
}
