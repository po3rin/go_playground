package port

type Error interface {
	Error() string
	StatusCode() int
}
