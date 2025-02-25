package request

type Request interface {
	Url() string
	ToFormData() map[string][]string
}
