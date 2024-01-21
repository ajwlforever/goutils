package main

import "sync"

type IService interface {
	getDb() string
	closeDb()
}
type Service struct {
	name string
	db   string
}

func (s *Service) getDb() string {
	println(s.name + ":" + s.db)
	return s.db
}
func (s *Service) closeDb() {
	println(s.db + ":Service:closeDb")
}

func (s *Service) String() string {
	return s.name + ":" + s.db
}
func New(t int) Service {
	if t == 0 {
		return Service{
			name: "forever",
			db:   "mysql",
		}
	} else {
		return Service{
			name: "forever",
			db:   "redis",
		}
	}

}

type IServiceFactory interface {
	New(t int) (*Service, error)
}
type ServiceFactory struct {
}

func (f *ServiceFactory) New(t int) (*Service, error) {
	if t == 0 {
		return &Service{
			name: "forever",
			db:   "mysql",
		}, nil
	} else {
		return &Service{
			name: "forever",
			db:   "redis",
		}, nil
	}
}

var factory *ServiceFactory
var one sync.Once

func GetFactory() *ServiceFactory {
	one.Do(func() {
		factory = &ServiceFactory{}
	})
	return factory
}

func main() {
	s := New(0)
	var is IService
	is = &s
	is.closeDb()

}
