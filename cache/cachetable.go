package cache

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// CacheTable is a table within the cache
type CacheTable struct {
	sync.Mutex

	name string

	items map[interface{}]*CacheItem

	cleanupTimer *time.Timer

	cleanupInterval *time.Duration

	logger *log.Logger

	loadData func(key interface{}, args ...interface{}) *CacheItem

	addedItem func(item *CacheItem)

	aboutToDeleteItem func(item *CacheItem)
}

// Count returns how many items are currently stored in the cache.
func (table *CacheTable) Count() int {
	table.Lock()
	defer table.Unlock()
	return len(table.items)
}

// Foreach all items
func (table *CacheTable) Foreach(trans func(key interface{}, item *CacheItem)) {
	table.Lock()
	defer table.Unlock()
	for k, v := range table.items {
		trans(k, v)
	}
}

// SetDataLoader configures a data-loader callback, which will be called when
// trying to access a non-existing key. The key and 0...n additional arguments
// are passed to the callback function.
func (table *CacheTable) SetDataLoader(f func(interface{}, ...interface{}) *CacheItem) {
	table.Lock()
	defer table.Unlock()
	table.loadData = f
}

// SetAddedItemCallback configures a callback, which will be called every time
// a new item is added to the cache.
func (table *CacheTable) SetAddedItemCallback(f func(*CacheItem)) {
	table.Lock()
	defer table.Unlock()
	table.addedItem = f
}

// SetAboutToDeleteItemCallback configures a callback, which will be called
// every time an item is about to be removed from the cache.
func (table *CacheTable) SetAboutToDeleteItemCallback(f func(*CacheItem)) {
	table.Lock()
	defer table.Unlock()
	table.aboutToDeleteItem = f
}

// SetLogger sets the logger to be used by this cache table.
func (table *CacheTable) SetLogger(log *log.Logger) {
	table.Lock()
	defer table.Unlock()
	table.logger = log
}

// Expiration check loop, triggered by a self-adjusting timer.
func (table *CacheTable) expirationCheck() {
	table.Lock()
	defer table.Unlock()

	if table.cleanupTimer != nil {
		table.cleanupTimer.Stop()
	}
	if table.cleanupInterval > 0 {
		table.log("Experition check trigered after", table.cleanupInterval, "for table", table.name)
	} else {
		table.log("Experition check ")
	}
}

// func (table *CacheTable) addInterval(item *CacheItem) {
// 	table.log("Adding item with key", item.key, "and lifespan of", item.lifeSpan, "to table", table.name)
// 	table.items[item.key] = item

// }

func (table *CacheTable) Add(key interface{}, data interface{}, lifeSpan time.Duration) *CacheItem {
	item := NewCacheItem(key, data, lifeSpan)

}

func (table *CacheTable) log(v ...interface{}) {
	if table.logger == nil {
		fmt.Printf(v...)
	} else {
		table.logger.Println(v...)
	}
}
