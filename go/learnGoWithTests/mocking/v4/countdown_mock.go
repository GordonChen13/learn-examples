package main

type CountdownOperationSpy struct {
	Calls []string
}

func (s *CountdownOperationSpy) Sleep()  {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}
