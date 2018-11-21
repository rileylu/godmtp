package adapter

type Adapter interface {
	Adapt(content []byte)(result []byte,err error)
}
