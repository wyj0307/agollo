package agollo

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Conf ...
type Conf struct {
	AppID            string   `json:"appId,omitempty"`
	Cluster          string   `json:"cluster,omitempty"`
	NameSpaceNames   []string `json:"namespaceNames,omitempty"`
	IP               string   `json:"ip,omitempty"`
	OneNamespaceMode bool     `json:"oneNamespaceMode,omitempty"` // by xingdonghai
	TagName          string   `json:"tagname,omitempty"`          // by xingdonghai
	EnableLocalCache bool     `json:"enable_local_cache"`         // by xingdonghai

}

// NewConf create Conf from file
func NewConf(name string) (*Conf, error) {
	f, err := os.Open(name)
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}
	defer f.Close()

	var ret Conf
	if err := json.NewDecoder(f).Decode(&ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

// NewConfWithENV create Conf form system envionment variables
// by xingdonghai
func NewConfWithENV() (*Conf, error) {
	conf := &Conf{
		AppID:            os.Getenv("agollo_appid"),
		Cluster:          os.Getenv("agollo_cluster"),
		NameSpaceNames:   strings.Split(os.Getenv("agollo_namespaces"), ","),
		IP:               os.Getenv("agollo_ip"),
		TagName:          os.Getenv("agollo_tagname"),
		OneNamespaceMode: os.Getenv("agollo_onenamespacemode") == "1" || os.Getenv("agollo_onenamespacemode") == "yes",
	}
	if conf.AppID == "" || conf.IP == "" {
		return nil, fmt.Errorf(errMissENV)
	}
	if conf.Cluster == "" {
		conf.Cluster = "default"
	}

	return conf, nil
}
