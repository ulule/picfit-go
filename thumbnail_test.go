package picfit

import (
	"fmt"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildParams(t *testing.T) {
	is := assert.New(t)

	params := map[string]string{
		"w":    "20",
		"h":    "20",
		"op":   "thumbnail",
		"path": "/my/path/to/image.jpg",
	}

	expected := url.Values{
		"h":    []string{"20"},
		"op":   []string{"thumbnail"},
		"path": []string{"/my/path/to/image.jpg"},
		"w":    []string{"20"},
	}

	is.EqualValues(BuildParams(params), expected)

	params = map[string]string{
		"w":    "20",
		"h":    "",
		"op":   "thumbnail",
		"path": "/my/path/to/image.jpg",
	}

	expected = url.Values{
		"op":   []string{"thumbnail"},
		"path": []string{"/my/path/to/image.jpg"},
		"w":    []string{"20"},
	}

	is.EqualValues(BuildParams(params), expected)
}

func TestGetThumbnailURL(t *testing.T) {
	baseURL := "https://img.test.com"
	is := assert.New(t)

	options := NewOptions()
	options.BaseURL = baseURL
	url, _ := GenerateThumbnailURL("bidule.jpg", "30x30", options)
	sig := strings.Split(url, "/")[4]
	expectedURL := fmt.Sprintf(
		"%s/%s/%s/thumbnail/30x30/bidule.jpg",
		options.BaseURL,
		options.DefaultMethod,
		sig)
	is.Equal(expectedURL, url)

	options = NewOptions()
	options.BaseURL = baseURL
	options.Upscale = 20
	url, _ = GenerateThumbnailURL("bidule", "30x30", options)
	is.Contains(url, "?upscale=20")

	options = NewOptions()
	options.BaseURL = baseURL
	options.Op = "resize"
	url, _ = GenerateThumbnailURL("bidule.jpg", "30x30", options)
	sig = strings.Split(url, "/")[4]
	expectedURL = fmt.Sprintf(
		"%s/display/%s/resize/30x30/bidule.jpg",
		options.BaseURL,
		sig)
	is.Equal(expectedURL, url)

	options = NewOptions()
	options.Op = "foobar"
	_, err := GenerateThumbnailURL("bidule.jpg", "30x30", options)
	is.NotNil(err)

	options = NewOptions()
	options.BaseURL = baseURL
	options.Crop = true
	url, _ = GenerateThumbnailURL("bidule.jpg", "30x30", options)
	sig = strings.Split(url, "/")[4]
	expectedURL = fmt.Sprintf(
		"%s/display/%s/resize/30x30/bidule.jpg",
		options.BaseURL,
		sig)
	is.Equal(expectedURL, url)
}
