package providers

type Transformer interface {
	Transform() ([]Vehicle, error)
}
