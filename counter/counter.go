package counter

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"
)

type Counter struct {
	counts       map[string]int
	mu           sync.RWMutex
	filePath     string
	saveInterval time.Duration
}

func NewCounter(filePath string, saveInterval time.Duration) *Counter {
	c := &Counter{
		counts:       make(map[string]int),
		filePath:     filePath,
		saveInterval: saveInterval,
	}
	c.loadFromFile()
	go c.periodicSave()
	return c
}

func (c *Counter) Increment(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counts[key]++
}

func (c *Counter) Get(key string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.counts[key]
}

func (c *Counter) GetAll() map[string]int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	counts := make(map[string]int)
	for k, v := range c.counts {
		counts[k] = v
	}
	return counts
}

func (c *Counter) saveToFile() error {
	c.mu.RLock()
	data, err := json.Marshal(c.counts)
	c.mu.RUnlock()
	if err != nil {
		return err
	}
	return os.WriteFile(c.filePath, data, 0644)
}

func (c *Counter) loadFromFile() {
	data, err := os.ReadFile(c.filePath)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Println("error finding counter file")
		}
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	json.Unmarshal(data, &c.counts)
}

func (c *Counter) periodicSave() {
	ticker := time.NewTicker(c.saveInterval)
	for range ticker.C {
		if err := c.saveToFile(); err != nil {
			log.Println("error saving to counter file")
		}
	}
}
