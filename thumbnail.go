package picfit

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/url"
	"strconv"

	"github.com/facette/natsort"
)

// Options is the options passed to GenerateThumbnailURL() function.
type Options struct {
	Op            string
	Crop          bool
	Upscale       int
	BaseURL       string
	DefaultMethod string
	SecretKey     string
}

// NewOptions returns a new Options instance.
func NewOptions() *Options {
	return &Options{
		Op:            "thumbnail",
		Crop:          false,
		Upscale:       0,
		DefaultMethod: "display",
	}
}

// SignParams returns signature from params.
func SignParams(key string, params map[string]string) string {
	mac := hmac.New(sha1.New, []byte(key))
	var sortedKeys []string
	for k := range params {
		sortedKeys = append(sortedKeys, k)
	}
	natsort.Sort(sortedKeys)
	values := url.Values{}
	for _, k := range sortedKeys {
		values.Add(k, params[k])
	}
	mac.Write([]byte(values.Encode()))
	return hex.EncodeToString(mac.Sum(nil))
}

// GenerateThumbnailURL returns thumbnail URL without signature and query string.
func GenerateThumbnailURL(path string, geometry string, options *Options) (string, error) {
	g, err := ParseGeometry(geometry)
	if err != nil {
		return "", err
	}

	w, h, op := "", "", "thumbnail"

	if options.Crop {
		op = "resize"
	}

	if g.X != 0 {
		w = strconv.Itoa(g.X)
	}

	if g.Y != 0 {
		h = strconv.Itoa(g.Y)
	}

	params := map[string]string{
		"w":    w,
		"h":    h,
		"op":   op,
		"path": path,
	}

	if options.Upscale != 0 {
		params["upscale"] = string(options.Upscale)
	}

	u := fmt.Sprintf(
		"%s/%s/%s/%s/%sx%s/%s",
		options.BaseURL,
		options.DefaultMethod,
		SignParams(options.SecretKey, params),
		params["op"],
		params["w"],
		params["h"],
		params["path"],
	)

	if options.Upscale != 0 {
		u = fmt.Sprintf("%s?upscale=%d", u, options.Upscale)
	}

	return u, nil
}
