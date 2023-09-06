package interf

type Interf interface {
	AcquireLock() bool
	ReleaseLock() bool
}

func Register(i Interf) Interf {
	return i
}
