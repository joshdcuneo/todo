package main

type Context struct {
	Config *AppConfig
	Store  *Store
	Args   *Args
}

func NewContext() *Context {
	c := NewConfig()

	s := NewStore(*c)
	err := s.Load()
	if err != nil {
		panic(err)
	}

	a := ParseArgs()

	return &Context{
		Config: c,
		Store:  s,
		Args:   a,
	}
}
