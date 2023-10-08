package singleton

import "sync"

type Config struct {
	Name string
}

var instance *Config

//func init() {
//	instance = &Config{}
//}

func GetInstance() *Config {
	return instance
}

// using sync once
// the `Do` method ensure execute only one time to create new instance, make it thread-safe when we perform multi-thread
var once sync.Once

func GetInstanceBySyncOne() *Config {
	once.Do(func() {
		instance = &Config{}
	})
	return instance
}

func GetInstanceByFuncInit() *Config {
	if instance == nil {
		instance = &Config{}
	}
	return instance
}

var mu sync.Mutex

func GetInstanceByMutex() *Config {
	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		instance = &Config{}
	}
	return instance
}
