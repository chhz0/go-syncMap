package gosyncmap

type SyncMap interface {
	Clear()

	Delete(key string)

	Load(key string) (value any, ok bool)

	Store(key string, value any)
}
