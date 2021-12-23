package api

type InterruptType int64

const (
	// since iota starts with 0, the first value
	// defined here will be the default
	INTERRUPT_UNKNOWN InterruptType = iota

	INTERRUPT_BREAK
	INTERRUPT_CONTINUE
)

type IRuntimeError interface {
	Token() IToken
	Message() string
	String() string

	// "break", "continue" statements
	Interrupt() InterruptType
	SetInterrupt(InterruptType)
	ClearInterrupt()
}

func (t InterruptType) String() string {
	switch t {
	case INTERRUPT_BREAK:
		return "break"
	case INTERRUPT_CONTINUE:
		return "continue"
	default:
		return "unknown interrupt"
	}
}
