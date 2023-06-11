package faker

type Faker[T any] interface {
	Create() *T
}
