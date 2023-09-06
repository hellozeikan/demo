package interf

type Interf interface {
	AcquireLock() bool
	ReleaseLock()
}

func Register(i Interf) Interf {
	return i
}
