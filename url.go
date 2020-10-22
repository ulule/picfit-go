package picfit

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Options is the options passed to BuildURL.
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

// SignParams returns signature from params.
func SignParams(key string, values url.Values) string {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(values.Encode()))
	return hex.EncodeToString(mac.Sum(nil))
}

// BuildURL builds a picfit URL.
func BuildURL(path string, geometry string, options *Options) (string, error) {
	secretValues := make(url.Values)
	extraValues := make(url.Values)
	g, err := ParseGeometry(geometry)
	if err != nil {
		return "", err
	}

	w := ""
	h := ""

	if g.X != 0 {
		w = strconv.Itoa(g.X)
		secretValues.Add("w", w)
	}

	if g.Y != 0 {
		h = strconv.Itoa(g.Y)
		secretValues.Add("h", h)
	}

	if options.Upscale != nil {
		secretValues.Add("upscale", strconv.Itoa(*options.Upscale))
		extraValues.Add("upscale", strconv.Itoa(*options.Upscale))
	}

	op := options.Ops[0]

	secretValues.Add("path", path)

	for i := range options.Ops {
		secretValues.Add("op", options.Ops[i])
		if i != 0 {
			extraValues.Add("op", options.Ops[i])
		}
	}

	u, err := url.Parse(options.BaseURL)
	if err != nil {
		return "", err
	}

	u.Path = strings.Join([]string{
		options.DefaultMethod,
		SignParams(options.SecretKey, secretValues),
		op,
		fmt.Sprintf("%sx%s", w, h),
		path,
	}, "/")

	u.RawQuery = extraValues.Encode()

	return u.String(), nil
}
