package logging

// from Standard Library os/error.go
type timeout interface {
	Timeout() bool
}
