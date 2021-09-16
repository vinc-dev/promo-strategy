package config

import (
	"fmt"
	"os"
	"reflect"
	"time"
)

// Get get config by key from env then base config
func Get(key string) string {
	r := os.Getenv(key)
	if r != "" {
		return r
	}
	if configValue, ok := base[key]; ok {
		return configValue
	}
	return ""
}

// GetIfDuration get config as duration
func GetIfDuration(key string) time.Duration {
	var t time.Duration
	return getIf(key, t).(time.Duration)
}

func getIf(key string, p interface{}) interface{} {
	t := reflect.TypeOf(p)
	if _, ok := baseInterface[key]; !ok {
		panic(fmt.Sprintf(`config with key "%s" not found`, key))
	}
	if keyType := reflect.TypeOf(baseInterface[key]); t != keyType {
		panic(fmt.Sprintf(`different type of config with key "%s" got %s expected %s`, key, keyType, t))
	}
	return baseInterface[key]
}

func mergeConfig(configs ...map[string]string) map[string]string {
	result := map[string]string{}
	for _, configMap := range configs {
		for key, configValue := range configMap {
			if _, ok := result[key]; ok {
				panic(fmt.Sprintf(`duplicate config key "%s" detected`, key))
			}
			result[key] = configValue
		}
	}
	return result
}

func mergeConfigInterface(configInterfaces ...map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for _, configInterfaceMap := range configInterfaces {
		for key, configValue := range configInterfaceMap {
			if _, ok := result[key]; ok {
				panic(fmt.Sprintf(`duplicate config interface key "%s" detected`, key))
			}
			result[key] = configValue
		}
	}
	return result
}
