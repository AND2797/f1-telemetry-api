package server

type server interface {
	Listen()
	Write()
}
