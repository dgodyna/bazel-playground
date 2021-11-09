package repository

import (
	"github.com/dgodyna/bazel-playground/pkg/entity"
	"sync"
)

type EntityRepository interface {
	Create(entity *entity.Entity)
	Get(name string) *entity.Entity
	Update(*entity.Entity) (*entity.Entity, error)
	Delete(name string)
	List() []*entity.Entity
}

var _ EntityRepository = (*entityRepository)(nil)

type entityRepository struct {
	mux   sync.RWMutex
	cache map[string]*entity.Entity
}

func (e *entityRepository) Create(entity *entity.Entity) {
	e.mux.Lock()
	e.cache[entity.Name] = entity
	e.mux.Unlock()
}

func (e *entityRepository) Get(name string) *entity.Entity {
	e.mux.RLock()
	defer e.mux.RUnlock()
	return e.cache[name]
}

func (e *entityRepository) Update(e2 *entity.Entity) (*entity.Entity, error) {
	e.mux.Lock()
	defer e.mux.Unlock()
	e.cache[e2.Name] = e2

	return e2, nil
}

func (e *entityRepository) Delete(name string) {
	e.mux.Lock()
	defer e.mux.Unlock()
	delete(e.cache, name)
}

func (e *entityRepository) List() []*entity.Entity {
	e.mux.RLock()
	defer e.mux.RUnlock()

	res := make([]*entity.Entity, len(e.cache))
	i := 0
	for _, v := range e.cache {
		res[i] = v
		i++
	}
	return res
}

func NewEntityRepository() EntityRepository {
	return &entityRepository{
		cache: make(map[string]*entity.Entity),
	}
}
