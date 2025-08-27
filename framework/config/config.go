package config

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/tsfans/go/framework/logger"
	"gopkg.in/yaml.v3"
)

var (
	config     *ConfigManger
	log        = logger.Get()
	configlock sync.RWMutex
)

type ConfigManger struct {
	config         map[string]any
	path           string
	autoReload     bool
	reloadInterval int // seconds
	lastModified   *time.Time
}

func init() {
	config = &ConfigManger{}
	path := os.Getenv("CONFIG_PATH")
	if len(path) == 0 {
		path = "./conf"
	}
	config.path = path
	config.autoReload = os.Getenv("CONFIG_AUTO_RELOAD") == "true"
	if reloadInterval := os.Getenv("CONFIG_RELOAD_INTERVAL"); reloadInterval != "" {
		i, err := strconv.ParseInt(reloadInterval, 10, 64)
		if err != nil {
			log.Panicf("parse CONFIG_RELOAD_INTERVAL error: %s", err)
		}
		config.reloadInterval = int(i)
	}

	exePath, _ := os.Executable()
	exeDir := filepath.Dir(exePath)
	log.Debugf("cwd:%s,configPath:%s,autoReload:%t,reloadInterval:%d", exeDir, config.path, config.autoReload, config.reloadInterval)
	config.load()

	if config.autoReload {
		go func() {
			for range time.Tick(time.Duration(config.reloadInterval) * time.Second) {
				config.load()
			}
		}()
	}
}

func (cfg *ConfigManger) load() {
	configlock.Lock()
	defer configlock.Unlock()

	fileInfo, err := os.Stat(cfg.path)
	if err != nil {
		log.Errorf("stat config file error: %s", err)
		return
	}

	if cfg.lastModified != nil && !fileInfo.ModTime().After(*cfg.lastModified) {
		return
	}

	data, err := os.ReadFile(config.path)
	if err != nil {
		log.Errorf("read config file error: %s", err)
	}

	err = yaml.Unmarshal(data, &config.config)
	if err != nil {
		log.Errorf("unmarshal config file error: %s", err)
	}

	now := time.Now()
	config.lastModified = &now
	log.Infof("config file loaded: %s", config.path)
}

func (cfg *ConfigManger) Get(key string) any {
	configlock.RLock()
	defer configlock.RUnlock()

	keys := strings.Split(key, ".")
	current := cfg.config

	var res any
	for i, part := range keys {
		if val, exists := current[part]; exists {
			if i == len(keys)-1 {
				res = val
				break
			}
			if nextMap, ok := val.(map[string]any); ok {
				current = nextMap
			} else {
				res = nil
				break
			}
		} else {
			res = nil
			break
		}
	}
	if res == nil {
		log.Debugf("config key not found: %s", key)
	}
	return res
}

func Exists(key string) bool { return config.Get(key) != nil }

func GetString(key string, defaultVal string) string {
	if val := config.Get(key); val != nil {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return defaultVal
}

func GetInt(key string, defaultVal int) int {
	if val := config.Get(key); val != nil {
		if i, ok := val.(float64); ok {
			return int(i)
		}
	}
	return defaultVal
}

func GetFloat(key string, defaultVal float64) float64 {
	if val := config.Get(key); val != nil {
		if f, ok := val.(float64); ok {
			return f
		}
	}
	return defaultVal
}

func GetBool(key string, defaultVal bool) bool {
	if val := config.Get(key); val != nil {
		if b, ok := val.(bool); ok {
			return b
		}
	}
	return defaultVal
}

func GetArray[T any](key string, defaultVal []T) []T {
	if val := config.Get(key); val != nil {
		if arr, ok := val.([]any); ok {
			var res []T
			for _, v := range arr {
				if t, ok := v.(T); ok {
					res = append(res, t)
				}
			}
			return res
		}
	}
	return defaultVal
}

func GetValue(key string) any {
	return config.Get(key)
}
