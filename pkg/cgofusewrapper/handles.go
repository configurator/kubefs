package cgofusewrapper

import (
	"log"
	"sync"
)

type CountedHandle struct {
	count uint
	value interface{}
}

type Handles struct {
	mutex   sync.Mutex
	handles map[string]*CountedHandle
}

func (h *Handles) CreateOrIncrement(name string, newValue func() (interface{}, error)) error {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	if handle, ok := h.handles[name]; ok {
		handle.count++
		return nil
	}

	value, err := newValue()
	if err != nil {
		return err
	}

	if h.handles == nil {
		h.handles = make(map[string]*CountedHandle)
	}

	h.handles[name] = &CountedHandle{
		count: 1,
		value: value,
	}
	return nil
}

func (h *Handles) Get(name string) interface{} {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	if handle, ok := h.handles[name]; ok {
		return handle.value
	}
	return nil
}

func (h *Handles) Release(name string) uint {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	handle, ok := h.handles[name]
	if !ok {
		log.Printf("Released non-existant handle %s\n", name)
		return 0
	}
	handle.count--
	if handle.count == 0 {
		delete(h.handles, name)
	}
	return handle.count
}
