package dailymotion_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var d Dailymotion

func TestLinkEmptyString(t *testing.T) {
	_, err := d.GetDirectLink("")
	assert.Equal(t, "Empty Link", err.Error())
}

func TestDownloadSuccess(t *testing.T) {
	_, err := d.GetDirectLink("https://www.facebook.com/pokerorganization/videos/508013846071164/")
	assert.Nil(t, err, "We are expecting nil error here")
}
