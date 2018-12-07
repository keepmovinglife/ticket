package model

import (
	"time"
)

type Market struct {
	Id          int64     `xorm:"pk autoincr BIGINT(20)"`
	No          int64     `xorm:"BIGINT(20)"`
	Symbol      string    `xorm:"not null comment('币种标识') VARCHAR(25)"`
	ChanelName  string    `xorm:"not null comment('渠道名称') VARCHAR(25)"`
	Name        string    `xorm:"not null comment('币种名称') VARCHAR(25)"`
	NameEn      string    `xorm:"not null comment('英文名称') VARCHAR(25)"`
	NameCn      string    `xorm:"comment('中文名称') VARCHAR(25)"`
	Pair        string    `xorm:"comment('交易对') VARCHAR(25)"`
	Rate        string    `xorm:"comment('当前价格') VARCHAR(25)"`
	VolA        string    `xorm:"comment('被兑换货币的交易量') VARCHAR(25)"`
	VolB        string    `xorm:"comment('兑换货币交易量') VARCHAR(25)"`
	CurrA       string    `xorm:"comment('被兑换货币') VARCHAR(25)"`
	CurrB       string    `xorm:"comment('兑换货币') VARCHAR(25)"`
	CurrSuffix  string    `xorm:"comment('货币类型后缀') VARCHAR(25)"`
	RatePercent string    `xorm:"comment('涨跌百分比') VARCHAR(25)"`
	Trend       string    `xorm:"comment(' 24小时趋势 up涨 down跌') VARCHAR(25)"`
	Supply      string    `xorm:"comment('币种供应量 ') VARCHAR(25)"`
	Marketcap   string    `xorm:"comment('总市值') VARCHAR(25)"`
	Plot        string    `xorm:"comment('趋势数据') VARCHAR(25)"`
	CreateTime  time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	Updatetime  time.Time `xorm:"comment('更新时间') TIMESTAMP"`
}
