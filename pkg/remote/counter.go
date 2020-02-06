package remote

import (
	"sync"
)

// Counter , 计数器
type Counter struct {
	sync.Mutex
	count int
}

// Add , 加一个数
func (c *Counter) Add(n int) {
	c.Lock()
	defer c.Unlock()
	c.count = c.count + n
}

// Value , 查看计数器当前数值
func (c *Counter) Value() int {
	c.Lock()
	defer c.Unlock()
	return c.count
}
