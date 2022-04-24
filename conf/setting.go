package conf

import (
	"io/ioutil"
	"sams_helper/notice"

	"gopkg.in/yaml.v2"
)

type ProxySet struct {
	IsEnabled bool   `yaml:"isEnabled"`
	ProxyUrl  string `yaml:"proxyUrl"`
}

type AutoFixPurchaseLimitSet struct {
	IsEnabled  bool `yaml:"isEnabled"`
	FixOffline bool `yaml:"fixOffline"`
	FixOnline  bool `yaml:"fixOnline"`
}

type GuaranteedSupplyGoodsOpts struct {
	DescKeywords    []string `yaml:"descKeywords"`
	DescFullfillCnt int16    `yaml:"descFullfillCount"`
	DescBlackList   []string `yaml:"descBlackList"`
}

type Setting struct {
	AuthToken               string                    `yaml:"authToken"`
	RunMode                 int                       `yaml:"runMode"`
	BruteCapacity           bool                      `yaml:"bruteCapacity"`
	DeviceType              int64                     `yaml:"deviceType"`
	DeliveryType            int64                     `yaml:"deliveryType"`
	StoreType               int64                     `yaml:"storeType"`
	FloorId                 int64                     `yaml:"floorId"`
	IgnoreInvalid           bool                      `yaml:"ignoreInvalid"`
	AutoFixPurchaseLimitSet AutoFixPurchaseLimitSet   `yaml:"autoFixPurchaseLimit"`
	PerDateLen              int                       `yaml:"perDateLen"`
	SassId                  string                    `yaml:"sassId"`
	ProxySet                ProxySet                  `yaml:"proxySet"`
	NoticeSet               notice.NoticerSet         `yaml:"noticeSet"`
	GSGoodsOption           GuaranteedSupplyGoodsOpts `yaml:"gsGoodsOptions"`
	GoodSpuId               string
}

func InitSetting() (error, Setting) {
	setting := Setting{}
	file, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return err, setting
	}

	err = yaml.Unmarshal(file, &setting)
	if err != nil {
		return err, setting
	}
	return nil, setting
}
