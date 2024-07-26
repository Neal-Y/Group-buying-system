package util

import (
	"sync"
)

type NotificationCache struct {
	cache sync.Map
}

func NewNotificationCache() *NotificationCache {
	return &NotificationCache{}
}

func (c *NotificationCache) Get(productID int) int {
	value, ok := c.cache.Load(productID)
	if !ok {
		return 0
	}
	return value.(int)
}

func (c *NotificationCache) Set(productID int, stock int) {
	c.cache.Store(productID, stock)
}
