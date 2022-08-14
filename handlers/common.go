package handlers

import "github.com/lonelyevil/khl"

func RegisterHandlers(s *khl.Session) {
	s.AddHandler(commandHandler)
}
