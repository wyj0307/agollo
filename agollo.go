package agollo

import (
	"reflect"
)

var (
	defaultClient *Client

)

// Start agollo
// by xingdonghai modify
func Start() error {
	if conf, err := NewConfWithENV(); err == nil {
		return StartWithConf(conf)
	} else {
		return StartWithConfFile(defaultConfName)
	}
}

// StartAndUnmarshal start agollo and unmarshal map to struct
// by xingdonghai
func StartAndUnmarshal(v interface{}) error {
	if err := Start(); err != nil {
		return err
	}
	return Unmarshal(v)
}

// StartAndUnmarshalOnChange start agollo and unmarshal map to struct, listen event of configuration change after that.
// by xingdonghai
func StartAndUnmarshalOnChange(v interface{}, run func(*ChangeEvent, error)) error {
	if err := StartAndUnmarshal(v); err != nil {
		return err
	}
	OnConfigChange(func(e *ChangeEvent) {
		tmp := reflect.New(reflect.TypeOf(v).Elem()).Interface()
		err := Unmarshal(tmp)
		if err == nil {
			reflect.ValueOf(v).Elem().Set(reflect.ValueOf(tmp).Elem())
		}
		run(e, err)
	})
	return nil
}

// StartWithENV run agollo with system envionment variables
// by xingdonghai
func StartWithENV() error {
	conf, err := NewConfWithENV()
	if err != nil {
		return err
	}
	return StartWithConf(conf)
}

// StartWithConfFile run agollo with conf file
func StartWithConfFile(name string) error {
	conf, err := NewConf(name)
	if err != nil {
		return err
	}
	return StartWithConf(conf)
}

// StartWithConf run agollo with Conf
// by xingdonghai modify
func StartWithConf(conf *Conf) error {
	if conf.TagName == "" {
		conf.TagName = defaultTagName
	}
	defaultClient = NewClient(conf)

	return defaultClient.Start()
}

// Stop sync config
func Stop() error {
	return defaultClient.Stop()
}

// WatchUpdate get all updates
func WatchUpdate() <-chan *ChangeEvent {
	return defaultClient.WatchUpdate()
}

// OnConfigChange when config changed, user code would be called
func OnConfigChange(run func(*ChangeEvent)) {
	defaultClient.OnConfigChange(run)
}

// GetStringValueWithNameSpace get value from given namespace
func GetStringValueWithNameSpace(namespace, key, defaultValue string) string {
	return defaultClient.GetStringValueWithNameSpace(namespace, key, defaultValue)
}

// GetStringValue from default namespace
func GetStringValue(key, defaultValue string) string {
	return GetStringValueWithNameSpace(defaultNamespace, key, defaultValue)
}

// GetNameSpaceContent get contents of namespace
func GetNameSpaceContent(namespace, defaultValue string) string {
	return defaultClient.GetNameSpaceContent(namespace, defaultValue)
}

// GetAllKeys return all config keys in given namespace
func GetAllKeys(namespace string) []string {
	return defaultClient.GetAllKeys(namespace)
}

// Unmarshal unmarshals the config into a struct. Make sure that the tags
// on the fields of the structure are properly set.
func Unmarshal(model interface{}) error {
	return defaultClient.Unmarshal(model)
}
