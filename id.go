package matcha // import "gomatcha.io/matcha"

import (
	"sync"

	_ "gomatcha.io/matcha/bridge"
)

var MainLocker sync.Locker

func init() {
	MainLocker = &sync.Mutex{}
}
