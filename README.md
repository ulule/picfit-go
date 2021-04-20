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

### BuildURL()

The `BuildURL()` function returns a pre-formatted URL for picfit server.

This function takes three required parameters:

* `path` - your original image path
* `geometry` - width and height formatted like this: `widthxheight` (example: "20x30")
* `options` - picfit options

Supported options are:

* `Ops` - see [Operations](https://github.com/thoas/picfit#operations)
* `Quality` - see [General Parameters](https://github.com/thoas/picfit#general-parameters)
* `Upscale` - see [General Parameters](https://github.com/thoas/picfit#general-parameters)
* `DefaultMethod` - see [Methods](https://github.com/thoas/picfit#methods) (defaults to `display`)
* `SecretKey` - your secret key (see [Security](https://github.com/thoas/picfit#security))

Options is just an instance of `Options` struct:

```go
// Create your own instance, with your own parameters.
options := &picfit.Options{
	Op:            "thumbnail",
	BaseURL:       "https://img.yourpicfitserver.com",
	DefaultMethod: "display",
	SecretKey:     "$ecretkeyplizkeepitseeeecret"
}

// Or, use the default ones (same as above) with NewOptions()
options := picfit.NewOptions()

// And, of course, you can override everything...
options.BaseURL = "https://img.superserver.com"
options.SecretKey = "qwerty1234ohitsbad"
```

Then, generate your URL:

```go
url, err := picfit.BuildURL("image.jpg", "90x90", options)
if err != nil {
	return fmt.Errorf("couldn't build picfit URL: %w", err)
}
```

## Contributing

* Ping us on twitter [@thoas](https://twitter.com/thoas), [@oibafsellig](https://twitter.com/oibafsellig)
* Fork the [project](https://github.com/ulule/picfit-go)
* Fix [bugs](https://github.com/ulule/picfit-go/issues)
