package cache

import (
	"sync"
	"time"
)

// CacheItem is an individual cache item
// Parameter data contains the user-set value in the cache.
type CacheItem struct {
	sync.Mutex

	key interface{}

	data interface{}

	lifeSpan time.Duration

	createdOn time.Time

	accessedOn time.Time

	accessCount uint64

	aboutToExpire func(interface{})
}

// NewCacheItem returns a newly created CacheItem.
// Parameter key is the item's cache-key.
// Parameter lifeSpan determines after which time period without an access the item
// will get removed from the cache.
// Parameter data is the item's value.
func NewCacheItem(key interface{}, data interface{}, lifeSpan time.Duration) *CacheItem {
	t := time.Now()
	return &CacheItem{
		key:           key,
		data:          data,
		lifeSpan:      lifeSpan,
		createdOn:     t,
		accessedOn:    t,
		accessCount:   0,
		aboutToExpire: nil,
	}
}

// KeepAlive marks an item to be kept for another expireDuration period.
func (item *CacheItem) KeepAlive() {
	item.Lock()
	defer item.Unlock()
	t := time.Now()
	item.accessedOn = t
	item.accessCount++
}

// LifeSpan returns this item's expiration duration.
func (item *CacheItem) LifeSpan() time.Duration {
	return item.lifeSpan
}

// CreatedOn returns when this item was added to the cache.
func (item *CacheItem) CreatedOn() time.Time {
	return item.createdOn
}

// AccessedOn returns when this item was last accessed.
func (item *CacheItem) AccessedOn() time.Time {
	item.Lock()
	defer item.Unlock()
	return item.accessedOn
}

// AccessedCount returns when this item was last accessed.
func (item *CacheItem) AccessedCount() uint64 {
	item.Lock()
	defer item.Unlock()
	return item.accessCount
}

// Key returns the key of this cached item.
func (item *CacheItem) Key() interface{} {
	return item.key
}

// Data returns the value of this cached item.
func (item *CacheItem) Data() interface{} {
	return item.data
}

// SetAboutToExpireCallback configures a callback, which will be called right
// before the item is about to be removed from the cache.
func (item *CacheItem) SetAboutToExpireCallback(f func(key interface{})) {
	item.Lock()
	defer item.Unlock()
	item.aboutToExpire = f
}
