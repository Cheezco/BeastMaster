package monitoring

import "sync"

type CpuEnergyStorage struct {
	mu    sync.Mutex
	value int
}

func (c *CpuEnergyStorage) Update(value int) {
	c.mu.Lock()

	c.value = value
	c.mu.Unlock()
}

func (c *CpuEnergyStorage) Value() int {
	c.mu.Lock()

	defer c.mu.Unlock()
	return c.value
}
