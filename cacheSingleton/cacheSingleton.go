package cacheSingleton

import (
	"fmt"

	"github.com/astaxie/beego/cache"
)

var globalCache cache.Cache = nil

func Get() cache.Cache {
	// not thread safe
	if globalCache == nil {
		var err error
		globalCache, err = cache.NewCache("memory", `{"interval": 60}`)
		if err != nil {
			fmt.Println("error get cache instance")
		}
	}
	return globalCache
}
