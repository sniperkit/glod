package cmd

import (
	"strings"

	// internal - core
	glod "github.com/sniperkit/glod/pkg"

	// internal - plugins
	"github.com/sniperkit/glod/pkg/plugin/provider/chiasenhac"  // TODO: code chiesenhac
	"github.com/sniperkit/glod/pkg/plugin/provider/daylimotion" // TODO: code dailymotion
	"github.com/sniperkit/glod/pkg/plugin/provider/facebook"    // TODO: code dailymotion
	"github.com/sniperkit/glod/pkg/plugin/provider/instagram"   // TODO: code dailymotion
	"github.com/sniperkit/glod/pkg/plugin/provider/nhaccuatui"  // TODO: code nhaccuatui
	"github.com/sniperkit/glod/pkg/plugin/provider/soundcloud"  // TODO: code soundcloud
	"github.com/sniperkit/glod/pkg/plugin/provider/twitter"     // TODO: code dailymotion
	"github.com/sniperkit/glod/pkg/plugin/provider/vimeo"       // TODO: code vimeo
	"github.com/sniperkit/glod/pkg/plugin/provider/youtube"     // TODO: code dailymotion
	"github.com/sniperkit/glod/pkg/plugin/provider/zing"        // TODO: code dailymotion
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
