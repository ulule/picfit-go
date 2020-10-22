package picfit

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildURL(t *testing.T) {
	baseURL := "https://img.test.com"
	is := assert.New(t)

	options := NewOptions()
	options.BaseURL = baseURL
	options.Ops = []string{"thumbnail"}
	url, _ := BuildURL("bidule.jpg", "30x30", options)
	sig := strings.Split(url, "/")[4]
	expectedURL := fmt.Sprintf(
		"%s/%s/%s/thumbnail/30x30/bidule.jpg",
		options.BaseURL,
		options.DefaultMethod,
		sig)
	is.Equal(expectedURL, url)

	options = NewOptions()
	options.BaseURL = baseURL
	options.Upscale = newint(20)
	options.Ops = []string{"thumbnail"}
	url, _ = BuildURL("bidule", "30x30", options)
	is.Contains(url, "?upscale=20")

	options = NewOptions()
	options.BaseURL = baseURL
	options.Ops = []string{"resize"}
	url, _ = BuildURL("bidule.jpg", "30x30", options)
	sig = strings.Split(url, "/")[4]
	expectedURL = fmt.Sprintf(
		"%s/display/%s/resize/30x30/bidule.jpg",
		options.BaseURL,
		sig)
	is.Equal(expectedURL, url)
}

func newint(i int) *int { return &i }
