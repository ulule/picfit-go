package picfit

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/url"
	"strconv"
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
		Upscale:       0,
		DefaultMethod: "display",
	}
}

// BuildParams builds params for SignParams.
func BuildParams(params map[string]string) url.Values {
	values := url.Values{}
	for k, v := range params {
		if v != "" {
			values.Add(k, v)
		}
	}
	return values
}

// SignParams returns signature from params.
func SignParams(key string, params map[string]string) string {
	mac := hmac.New(sha1.New, []byte(key))
	values := BuildParams(params)
	mac.Write([]byte(values.Encode()))
	return hex.EncodeToString(mac.Sum(nil))
}

// GenerateThumbnailURL returns thumbnail URL without signature and query string.
func GenerateThumbnailURL(path string, geometry string, options *Options) (string, error) {
	g, err := ParseGeometry(geometry)
	if err != nil {
		return "", err
	}

	supportedOps := map[string]bool{
		"thumbnail": true,
		"resize":    true,
		"flip":      true,
		"rotate":    true,
	}

	if options.Op == "" {
		if options.Crop {
			options.Op = "resize"
		} else {
			options.Op = "thumbnail"
		}
	} else {
		if _, ok := supportedOps[options.Op]; !ok {
			return "", fmt.Errorf("operation %s is not supported", options.Op)
		}
	}

	w := ""
	h := ""

	if g.X != 0 {
		w = strconv.Itoa(g.X)
	}

	if g.Y != 0 {
		h = strconv.Itoa(g.Y)
	}

	params := map[string]string{
		"w":    w,
		"h":    h,
		"op":   options.Op,
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
