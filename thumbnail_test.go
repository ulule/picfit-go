package picfit

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetThumbnailURL(t *testing.T) {
	baseURL := "https://img.test.com"

	assert := assert.New(t)

	options := NewOptions()
	options.BaseURL = baseURL

	url, _ := GenerateThumbnailURL("bidule.jpg", "30x30", options)
	sig := strings.Split(url, "/")[4]
	expectedURL := fmt.Sprintf(
		"%s/%s/%s/thumbnail/30x30/bidule.jpg",
		options.BaseURL,
		options.DefaultMethod,
		sig,
	)

	assert.Equal(url, expectedURL)

	options = NewOptions()
	options.BaseURL = baseURL
	options.Upscale = 20
	url, _ = GenerateThumbnailURL("bidule", "30x30", options)

	assert.Contains(url, "?upscale=20")

	options = NewOptions()
	options.BaseURL = baseURL
	options.Crop = true

	url, _ = GenerateThumbnailURL("bidule.jpg", "30x30", options)
	sig = strings.Split(url, "/")[4]
	expectedURL = fmt.Sprintf(
		"%s/display/%s/resize/30x30/bidule.jpg",
		options.BaseURL,
		sig,
	)

	assert.Equal(url, expectedURL)
}
