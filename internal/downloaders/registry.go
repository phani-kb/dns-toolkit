package downloaders

import (
	"errors"
	"sync"
)

type registry struct {
	mu          sync.RWMutex
	downloaders map[string]Downloader
}

type Register interface {
	RegisterDownloader(downloader Downloader) error
	GetDownloader(name string) (Downloader, bool)
}

var downloaderRegistry = &registry{
	downloaders: make(map[string]Downloader),
}

func RegisterDownloader(downloader Downloader) error {
	name := downloader.Name()
	downloaderRegistry.mu.Lock()
	defer downloaderRegistry.mu.Unlock()
	if _, exists := downloaderRegistry.downloaders[name]; exists {
		return errors.New("downloader with the same name already exists")
	}
	downloaderRegistry.downloaders[name] = downloader
	return nil
}

func GetDownloader(name string) (Downloader, bool) {
	downloaderRegistry.mu.RLock()
	defer downloaderRegistry.mu.RUnlock()
	downloader, exists := downloaderRegistry.downloaders[name]
	return downloader, exists
}
