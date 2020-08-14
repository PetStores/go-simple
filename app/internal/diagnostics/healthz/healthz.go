package healthz

type Status int

const (
	Ok Status = iota
	Warning
	Fatal
)

type Resource struct {
	Name string `json:"name"`
	Status Status `json:"status"`
	Message string `json:"msg"`
}

type Check func() []Resource

