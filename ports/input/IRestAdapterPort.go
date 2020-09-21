package input

type IRestAdapterPort interface {
	Start(port string) error
}
