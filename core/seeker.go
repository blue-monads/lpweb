package core

type Seeker interface {
	Set(hash, addr string) error
	Get(hash string) (string, error)
}
