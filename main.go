// 爬取豆瓣电影 TOP250
package main

import (
	"bufio"
	"fmt"
	"github.com/go-crawler/douban-movie/parse"
	"os"
	"time"
)

var (
	baseUrl = "https://movie.douban.com/top250"
)

// 新增数据
func Add(movies []parse.DoubanMovie) {
	f, _ := os.Create("./data/"+time.Now().Format("2006-01-02")+".md")
	writer := bufio.NewWriter(f)
	defer f.Close()
	table:= fmt.Sprintf("|%+v|%+v|%+v|%+v|%+v|%+v|%+v|%+v|%+v|%+v|\n",
		"Year", "Title", "Subtitle",
		"Area",  "Quote", "Tag",
		"Star", "Comment", "Desc","Other")

	writer.Write([]byte(table))
	writer.Write([]byte("|----|----|----|----|----|----|----|----|----|----|\n")) //nolint:errcheck
	for _, movie := range movies {
		writer.Write([]byte(fmt.Sprintf("|%+v|%+v|%+v|%+v|%+v|%+v|%+v|%+v|%+v|%+v|\n", //nolint:errcheck
			movie.Year,movie.Title,movie.Subtitle,
			movie.Area,movie.Quote,movie.Tag,
			movie.Star,movie.Comment,movie.Desc,movie.Other)))
	}
	writer.Flush()
	file, _ := os.ReadFile("./data/" + time.Now().Format("2006-01-02") + ".md")
	last, _ := os.Create("last.md")
	defer last.Close()
	writer = bufio.NewWriter(last)
	writer.Write(file)
	writer.Flush()
}

// 开始爬取
func Start() {
	var movies []parse.DoubanMovie

	pages := parse.GetPages(baseUrl)
	for _, page := range pages {
		doc := parse.GetDoc(baseUrl + page.Url)
		movies = append(movies, parse.ParseMovies(doc)...)
	}

	Add(movies)
}

func main() {
	Start()

	//defer model.DB.Close()
}
