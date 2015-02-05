package eliminate

type Stack interface {
	Push(o struct{})
	Pop() struct{}
}
