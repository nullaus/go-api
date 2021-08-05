package twitter

import (
	"time"

	"github.com/allegro/bigcache"
	"github.com/eko/gocache/cache"
	"github.com/eko/gocache/store"
)

var bigcacheClient, _ = bigcache.NewBigCache(bigcache.DefaultConfig(5 * time.Minute))
var bigcacheStore = store.NewBigcache(bigcacheClient, nil) // No otions provided (as second argument)

var cacheManager = cache.New(bigcacheStore)
