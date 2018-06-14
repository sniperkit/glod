package cmd

import (
	"strings"

	// internal - core
	glod "github.com/sniperkit/glod/pkg"

	// internal - plugins
	"github.com/sniperkit/glod/pkg/plugin/provider/chiasenhac"
	"github.com/sniperkit/glod/pkg/plugin/provider/daylimotion"
	"github.com/sniperkit/glod/pkg/plugin/provider/facebook"
	"github.com/sniperkit/glod/pkg/plugin/provider/instagram"
	"github.com/sniperkit/glod/pkg/plugin/provider/nhaccuatui"
	"github.com/sniperkit/glod/pkg/plugin/provider/soundcloud"
	"github.com/sniperkit/glod/pkg/plugin/provider/twitter"
	"github.com/sniperkit/glod/pkg/plugin/provider/vimeo"
	"github.com/sniperkit/glod/pkg/plugin/provider/youtube"
	"github.com/sniperkit/glod/pkg/plugin/provider/zing"
)

func getGlod(link string) glod.Source {
	switch {

	case strings.Contains(link, initFacebook):
		// Facebook
		return &facebook.Facebook{}

	case strings.Contains(link, initInstagram):
		// Instagram
		return &instagram.Instagram{}

	case strings.Contains(link, initTwitter):
		// Twitter
		return &twitter.Twitter{}

	case strings.Contains(link, initYoutube):
		// Youtube
		return &youtube.Youtube{}

	case strings.Contains(link, initDailymotion):
		// Dailymotion
		return &dailymotion.Dailymotion{}

	case strings.Contains(link, initVimeo):
		// Vimeo
		return &vimeo.Vimeo{}

	case strings.Contains(link, initSoundCloud):
		// SoundCloud
		return &soundcloud.SoundCloud{}

	case strings.Contains(link, initNhacCuaTui):
		// nhaccuatui
		return &nhaccuatui.NhacCuaTui{}

	case strings.Contains(link, initChiaSeNhac):
		// ChiaSeNhac
		return &chiasenhac.ChiaSeNhac{}

	case strings.Contains(link, initZingMp3):
		// Zing
		return &zing.Zing{}
	}

	return nil
}
