package zing

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	// external
	"github.com/PuerkitoBio/goquery"

	// internal
	glod "github.com/sniperkit/glod/pkg"
)

const (
	ZingPrefix = "zing"
)

const (
	zingSong         string = "http://mp3.zing.vn/bai-hat/"
	zingAlbum        string = "http://mp3.zing.vn/album/"
	playList         string = "http://mp3.zing.vn/playlist/"
	linkDownloadSong string = "http://api.mp3.zing.vn/api/mobile/song/getsonginfo?requestdata=" //+ {"id":"IDBAIHAT"}
)

type Zing struct {
}

type ZingResponse struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Source struct {
		Url  string `json:"128"`
		Url2 string `json:"320"`
		Url3 string `json:"lossless"`
	} `json:"source"`
}

// function that input is a link then return an slice of url that permantly download file and error(if it has)
func (z *Zing) GetDirectLink(link string) ([]glod.Response, error) {
	if link == "" {
		return nil, errors.New("Empty Link")
	}

	var listLink []string

	if strings.Contains(link, zingSong) {
		if len(strings.Split(link, "/")) < 6 {
			return nil, errors.New("Invalid link")
		}
		listLink = append(listLink, link)
	} else if strings.Contains(link, zingAlbum) || strings.Contains(link, playList) {
		doc, err := goquery.NewDocument(link)
		if err != nil {
			return nil, err
		}
		doc.Find(".item-song").Find("a").Each(func(i int, s *goquery.Selection) {
			url, _ := s.Attr("href")
			if strings.Contains(url, "bai-hat") {
				listLink = append(listLink, url)
			}
		})
	}

	list := GetMedias(listLink)

	return list, nil
}

// GetMediaID return song id from given link
func GetMediaID(link string) string {

	temp := strings.Split(link, "/")
	_temp := strings.Split(temp[5], ".")

	ID := _temp[0]

	return ID
}

// GetMedias return list of song
func GetMedias(listLink []string) []glod.Response {

	var list []glod.Response
	var song glod.Response
	var zingResponse ZingResponse

	for i := range listLink {
		id := GetMediaID(listLink[i])
		link := linkDownloadSong + "{\"id\":\"" + id + "\"}"

		res, err := http.Get(link)
		if err != nil {
			return list
		}
		defer res.Body.Close()

		if err := json.NewDecoder(res.Body).Decode(&zingResponse); err != nil {
			log.Println(err)
		}

		song.Artist = zingResponse.Artist
		song.Title = zingResponse.Title
		song.StreamURL = zingResponse.Source.Url

		list = append(list, song)
	}
	return list
}
