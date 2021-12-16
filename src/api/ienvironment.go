package api

type IEnvironment interface {
	Assign(name IToken, obj interface{}) IRuntimeError
	Define(name string, obj interface{}) IRuntimeError
	Get(name IToken) (obj interface{}, err IRuntimeError)
}
