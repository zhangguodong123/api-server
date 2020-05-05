// Copyright 2020 xwi88.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Configs init api configs
package configs

import (
	"sync"

	"github.com/prometheus/common/log"
	"github.com/xwi88/config4go"
	"github.com/xwi88/kit4go/json"
)

var gConfig *APIConfig
var lock sync.RWMutex

// Init init the config
func Init(configFile *string) {
	lock.RLock()
	defer lock.RUnlock()

	// load file
	var cfg APIConfig
	if configFile != nil {
		err := config4go.LoadConfig(*configFile, &cfg)
		if err != nil {
			panic(err.Error())
		}
	} else {
		log.Fatalln("config file not set!")
	}

	gConfig = &cfg
}

// GetConfig get global config
func GetConfig() *APIConfig {
	lock.RLock()
	defer lock.RUnlock()
	return gConfig
}

// GetConfigJSONString get global config json string
func GetConfigJSONString(withIndent bool) string {
	lock.RLock()
	defer lock.RUnlock()
	if gConfig != nil {
		indent := ""
		if withIndent {
			indent = " "
		}
		gCBytes, _ := json.MarshalIndent(gConfig, "", indent)
		return string(gCBytes)
	}
	return ""
}
