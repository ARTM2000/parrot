package core

type server struct {
	Config *config
}

func RunServer(c *config) error {
	s := server{
		Config: c,
	}
	return s.run()
}

func (s *server) run() error {
	ch := make(chan bool)

	<-ch
	return nil
}
