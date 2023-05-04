package ghaproxy

import (
	"context"
	"io"
	"log"

	"github.com/buchgr/bazel-remote/v2/cache"

	actionscache "github.com/tonistiigi/go-actions-cache"
)

func Put(ctx context.Context, kind cache.EntryKind, hash string, logicalSize int64, sizeOnDisk int64, rc io.ReadCloser) {
}

func Get(ctx context.Context, kind cache.EntryKind, hash string) (rc io.ReadCloser, size int64, err error) {

}

func Contains(ctx context.Context, kind cache.EntryKind, hash string) (bool, int64) {

}

func New(
	cache_url string,
	auth_token string,
	storageMode string, accessLogger cache.Logger,
	errorLogger cache.Logger, numUploaders, maxQueuedUploads int,
) cache.Proxy {

	var err error
	var opt = actionscache.Opt{}
	cache, err := actionscache.New(auth_token, cache_url, opt)
	cache.
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	if storageMode != "zstd" && storageMode != "uncompressed" {
		log.Fatalf("Unsupported storage mode for the azblobproxy backend: %q, must be one of \"zstd\" or \"uncompressed\"",
			storageMode)
	}

}
