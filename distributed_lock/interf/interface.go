package interf

type Interf interface {
	AcquireLock(lockKey string) bool
	ReleaseLock(lockKey string) bool
}

func Register(i Interf) Interf {
	return i
}
