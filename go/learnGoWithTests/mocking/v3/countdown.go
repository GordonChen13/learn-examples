package main

const write = "write"
const sleep= "sleep"

type CountdownOperationSpy struct {
	Calls []string
}

func (s *CountdownOperationSpy) Sleep()  {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationSpy) Write()  {
	s.Calls = append(s.Calls, write)
	return
}
