package enum

type CarType int

const (
	US    CarType = 1
	Japan CarType = 2
)

func (carType CarType) GetDesciption() string {
	switch carType {
	case US:
		return "US"
	case Japan:
		return "Japan"
	default:
		return ""
	}
}
