package main

type Context struct {
	Config *AppConfig
	Store  *Store
}

func NewContext() *Context {
	c := NewConfig()

	s := NewStore(*c)
	err := s.Load()
	if err != nil {
		panic(err)
	}

	return &Context{
		Config: c,
		Store:  s,
	}
}
