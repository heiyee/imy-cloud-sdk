package cloud_basic

type InitializerFactory interface {
	CreateInitializer() Initializer
}
