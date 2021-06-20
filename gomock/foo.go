package gomock

type Foo interface {
	Bar(x int) int
}

func SUT(f Foo) int {
	return f.Bar(99)
}
