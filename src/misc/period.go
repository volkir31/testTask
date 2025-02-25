package misc

const (
	DAY   = "day"
	MONTH = "month"
	YEAR  = "year"
)

type Period string

func (p Period) String() string {
	return string(p)
}
