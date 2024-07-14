package utils

import (
    "sync"
    "time"
)

type CacheItem struct {
    Value      interface{}
    Expiration int64
}

type Cache struct {
    items map[string]CacheItem
    mutex sync.RWMutex
    ttl   time.Duration
}

// NewCache creates a new cache with a given TTL (time-to-live) duration
func NewCache(ttl time.Duration) *Cache {
    return &Cache{
        items: make(map[string]CacheItem),
        ttl:   ttl,
    }
}

// Set adds a new item to the cache with the given key and value
func (c *Cache) Set(key string, value interface{}) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    c.items[key] = CacheItem{
        Value:      value,
        Expiration: time.Now().Add(c.ttl).Unix(),
    }
}

// Get retrieves an item from the cache by key
func (c *Cache) Get(key string) (interface{}, bool) {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    item, found := c.items[key]
    if !found || time.Now().Unix() > item.Expiration {
        return nil, false
    }
    return item.Value, true
}
