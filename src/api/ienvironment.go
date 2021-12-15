package api

type IEnvironment interface {
	Define(name string, obj interface{}) IRuntimeError
	Get(name IToken) (obj interface{}, err IRuntimeError)
}
