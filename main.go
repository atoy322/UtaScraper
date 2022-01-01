package main

import (
	"fmt"
	"os"
	"strconv"
    "strings"
)

func main() {
	if len(os.Args) == 1 {
		return
	}
	query := os.Args[1]
	html, e := Get("https://www.uta-net.com/search/?Aselect=2&Bselect=3&Keyword=" + query)
	if e != nil {
		panic(e)
	}

	html_string := string(html)
	songs := ParseSongs(html_string)

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

	html, e = Get(target_song.URL)
	if e != nil {
		panic(e)
	}

	html_string = string(html)
	kashi := ParseKashi(html_string)

	fmt.Println(kashi)

	file, e := os.Create(target_song.SongName + " - " + target_song.Name + ".txt")
	if e != nil {
		panic(e)
	}
	defer file.Close()

	file.WriteString(kashi)
}
