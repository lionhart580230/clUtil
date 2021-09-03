package clSuperMap

import "sync"

// 超级Map
type clSuperMap struct {
	data map[string] string
	locker sync.RWMutex
}