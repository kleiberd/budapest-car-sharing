package providers

type Transformer interface {
	Transform(interface{}) (Vehicle, error)
}
