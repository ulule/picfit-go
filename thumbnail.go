package picfit

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"net/url"
	"strconv"
)

// Options is the options passed to GenerateThumbnailURL() function.
type Options struct {
	Ops           []string
	Upscale       *int
	BaseURL       string
	DefaultMethod string
	SecretKey     string
}

// NewOptions returns a new Options instance.
func NewOptions() *Options {
	return &Options{
		DefaultMethod: "display",
	}
}

// SignParams returns values signature.
func SignParams(key string, values url.Values) string {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(values.Encode()))
	return hex.EncodeToString(mac.Sum(nil))
}

// GenerateThumbnailURL returns thumbnail URL without signature and query string.
func GenerateThumbnailURL(path string, geometry string, options *Options) (string, error) {
	values := make(url.Values)
	values.Set("path", path)

	g, err := ParseGeometry(geometry)
	if err != nil {
		return "", err
	}
	if g.X != 0 {
		values.Set("w", strconv.Itoa(g.X))
	}
	if g.Y != 0 {
		values.Set("h", strconv.Itoa(g.Y))
	}

	for i := range options.Ops {
		values.Add("op", options.Ops[i])
	}

	if options.Upscale != nil {
		values.Set("upscale", strconv.Itoa(*options.Upscale))
	}

	values.Set("sig", SignParams(options.SecretKey, values))

	u, err := url.Parse(options.BaseURL)
	if err != nil {
		return "", err
	}
	u.Path = options.DefaultMethod
	u.RawQuery = values.Encode()
	return u.String(), nil
}
