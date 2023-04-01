//go:build tools

package tools

import (
	// protobuf compiler plugins
	_ "github.com/twitchtv/twirp/protoc-gen-twirp"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
