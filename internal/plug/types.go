package plug

import (
	"fmt"
	"sync"
)

type BlockingValue func() (any, error)

type ReadOnlyProperties interface {
	Get(name string) (any, bool)
	GetBlocking(name string) (any, error)
	GetString(name string) string
	GetBool(name string) bool
	GetInt(name string) int64
}

type Properties struct {
	mu  *sync.RWMutex
	ps  map[string]any
	psb map[string]BlockingValue
}

func NewProperties() *Properties {
	return &Properties{
		mu:  &sync.RWMutex{},
		ps:  map[string]any{},
		psb: map[string]BlockingValue{},
	}
}

func (p *Properties) Set(name string, value any) {
	p.mu.Lock()
	p.ps[name] = value
	p.mu.Unlock()
}

func (p *Properties) SetBlocking(name string, value BlockingValue) {
	p.mu.Lock()
	p.psb[name] = value
	p.mu.Unlock()
}

func (p *Properties) Get(name string) (any, bool) {
	p.mu.RLock()
	v, ok := p.ps[name]
	p.mu.RUnlock()
	return v, ok
}

func (p *Properties) GetBlocking(name string) (any, error) {
	p.mu.RLock()
	v, ok := p.psb[name]
	p.mu.RUnlock()
	if !ok {
		// TODO:
		return nil, nil
	}
	return v()
}

func (p *Properties) GetString(name string) string {
	v, ok := p.Get(name)
	if !ok {
		return ""
	}
	if sv, ok := v.(string); ok {
		return sv
	}
	return fmt.Sprintf("%v", v)
}

func (p *Properties) GetBool(name string) bool {
	v, ok := p.Get(name)
	if ok {
		if bv, ok := v.(bool); ok {
			return bv
		}
	}
	return false
}

func (p *Properties) GetInt(name string) int64 {
	v, ok := p.Get(name)
	if ok {
		if bv, ok := v.(int64); ok {
			return bv
		}
	}
	return 0

}

type RegistryItem[T any] struct {
	Name string
	Item T
}

type RegistryItems[T any] []RegistryItem[T]

func (ri RegistryItems[T]) Map() map[string]T {
	m := make(map[string]T, len(ri))
	for _, x := range ri {
		m[x.Name] = x.Item
	}
	return m
}
