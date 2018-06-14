package chiasenhac

import (
	glod "github.com/sniperkit/glod/pkg"
)

const (
	CsnPrefix         = "chiasenhac"
	csnDownloadSuffix = "download.html"
	csnAblum          = "chiasenhac.vn/nghe-album/"
	csnSong1          = "chiasenhac.vn/mp3"
	csnSong2          = "chiasenhac.vn/nhac-hot"

	csnStreamURL = "/128/"
)

type ChiaSeNhac struct {
}

// TODO: code chiesenhac
func (csn *ChiaSeNhac) GetDirectLink(link string) ([]glod.Response, error) {
	var list []glod.Response
	return list, nil
}
