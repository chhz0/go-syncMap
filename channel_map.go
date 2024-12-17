package gosyncmap

type empty struct{}

// 使用channel实现的线程安全map
type ChannelMap struct {
	mp map[string]any

	lockChan chan empty
}

// Clear implements SyncMap.
func (cm *ChannelMap) Clear() {
	cm.lockChan <- empty{}
	defer func() {
		<-cm.lockChan
	}()

	cm.mp = make(map[string]any)
}

// Delete implements SyncMap.
func (cm *ChannelMap) Delete(key string) {
	cm.lockChan <- empty{}
	defer func() {
		<-cm.lockChan
	}()

	delete(cm.mp, key)
}

// Load implements SyncMap.
func (cm *ChannelMap) Load(key string) (value any, ok bool) {
	cm.lockChan <- empty{}
	defer func() {
		<-cm.lockChan
	}()

	value, ok = cm.mp[key]
	return value, ok
}

// Store implements SyncMap.
func (cm *ChannelMap) Store(key string, value any) {
	cm.lockChan <- empty{}
	defer func() {
		<-cm.lockChan
	}()

	cm.mp[key] = value
}

func NewChannelMap() *ChannelMap {
	return &ChannelMap{
		mp:       make(map[string]any),
		lockChan: make(chan empty, 1),
	}
}

var _ SyncMap = (*ChannelMap)(nil)
