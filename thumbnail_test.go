package picfit

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetThumbnailURL(t *testing.T) {
	baseURL := "https://img.test.com"
	is := assert.New(t)

	options := NewOptions()
	options.BaseURL = baseURL
	options.Ops = []string{"thumbnail"}
	s, err := GenerateThumbnailURL("bidule.jpg", "30x30", options)
	is.NoError(err)
	u, err := url.Parse(s)
	is.NoError(err)
	expectedURL := fmt.Sprintf("%s/%s?h=30&op=thumbnail&path=bidule.jpg&sig=%s&w=30",
		options.BaseURL,
		options.DefaultMethod,
		u.Query().Get("sig"),
	)
	is.Equal(expectedURL, s)

	options = NewOptions()
	options.BaseURL = baseURL
	options.Upscale = newint(20)
	s, err = GenerateThumbnailURL("bidule", "30x30", options)
	is.NoError(err)
	is.Contains(s, "upscale=20")

	options = NewOptions()
	options.BaseURL = baseURL
	options.Ops = []string{"resize"}
	s, err = GenerateThumbnailURL("bidule.jpg", "30x30", options)
	is.NoError(err)
	u, err = url.Parse(s)
	is.NoError(err)
	expectedURL = fmt.Sprintf("%s/%s?h=30&op=resize&path=bidule.jpg&sig=%s&w=30",
		options.BaseURL,
		options.DefaultMethod,
		u.Query().Get("sig"),
	)
	is.Equal(expectedURL, s)
}

func newint(i int) *int { return &i }
