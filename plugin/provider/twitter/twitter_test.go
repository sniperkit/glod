package twitter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tw Twitter

func TestLinkEmptyString(t *testing.T) {
	_, err := tw.GetDirectLink("")
	assert.Equal(t, "Empty Link", err.Error())
}

func TestDownloadSuccess(t *testing.T) {
	_, err := tw.GetDirectLink("https://www.tiwtter.com/videos/508013846071164/")
	assert.Nil(t, err, "We are expecting nil error here")
}
