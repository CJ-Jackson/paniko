package shared

const (
	UserDepName = "user-e0d6c3c9f7fd79ab13e96a6c9b7ff666"
	UserDepData = "user-c5962ba014317e92666f1cf4f6c6416c"
)

type User interface {
	Username() string
	CheckIfUser()
	CheckIfGuest()
}
