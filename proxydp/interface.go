package proxydp

// IGit git interface
type IGit interface {
	Clone(url string) error
	Push(url string) error
}
