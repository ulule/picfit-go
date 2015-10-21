# picfit-go

[![Build Status](https://travis-ci.org/ulule/picfit-go.svg)](https://travis-ci.org/ulule/picfit-go)

A Go client library for generating URLs with [picfit](https://github.com/thoas/picfit)

## Installation

```bash
$ go get github.com/ulule/picfit-go
```

## Usage

First, import the package:

```go
import "github.com/ulule/picfit-go"
```

Now, you can access the package through `picfit` namespace.

## API

### GenerateThumbnailURL()

The `GenerateThumbnailURL()` function returns a pre-formatted URL for picfit server.

This function takes three required parameters:

* `path` - your original image path
* `geometry` - width and height formatted like this: `widthxheight` (example: "20x30")
* `options` - picfit options

Supported options are:

* `Op` - see [Operations](https://github.com/thoas/picfit#operations) (defaults to `thumbnail`)
* `Crop` - either crop image or not (`true` or `false`, defaults to `false`)
* `Upscale` - see [General Parameters](https://github.com/thoas/picfit#general-parameters) (defaults to `0`)
* `DefaultMethod` - see [Methods](https://github.com/thoas/picfit#methods) (defaults to `display`)
* `SecretKey` - your secret key (see [Security](https://github.com/thoas/picfit#security))

Options is just an instance of `Options` struct:

```go
// Create your own instance, with your own parameters.
options := &picfit.Options{
	Op:            "thumbnail",
	Crop:          false,
	Upscale:       0,
	BaseURL:       "https://img.yourpicfitserver.com",
	DefaultMethod: "display",
	SecretKey:     "$ecretkeyplizkeepitseeeecret"
}

// Or, use the default ones (same as above) with NewOptions()
options := picfit.NewOptions()

// And, of course, you can override everything...
options.BaseURL = "https://img.superserver.com"
options.SecretKey = "qwerty1234ohitsbad"
options.Crop = true
```

Then, generate your URL:

```go
url, err := picfit.GenerateThumbnailURL("image.jpg", "90x90", options)
if err != nil {
	fmt.Println("Oops, sorry guys")
}
```

## Contributing

* Ping us on twitter [@thoas](https://twitter.com/thoas), [@oibafsellig](https://twitter.com/oibafsellig)
* Fork the [project](https://github.com/ulule/picfit-go)
* Fix [bugs](https://github.com/ulule/picfit-go/issues)
