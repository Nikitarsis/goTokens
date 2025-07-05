package tokenProducer

type ITokenComponentConfig interface {
	GetKeyChannelSize() uint
	GetJtiChannelSize() uint
	GetIssuer() string
}
