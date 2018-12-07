package model

import (
	"time"
)

type Ticket struct {
	Id            int64     `xorm:"pk autoincr BIGINT(20)"`
	ChanelName    string    `xorm:"comment('渠道名称') VARCHAR(25)"`
	Pair          string    `xorm:"not null pk comment('交易对') VARCHAR(25)"`
	Result        string    `xorm:"not null comment('获取结果') VARCHAR(25)"`
	Last          string    `xorm:"not null comment('最新成交价格') VARCHAR(25)"`
	Lowestask     string    `xorm:"not null comment('卖方最低价') VARCHAR(25)"`
	Highestbid    string    `xorm:"comment('买方最高价') VARCHAR(25)"`
	Percentchange string    `xorm:"comment('涨跌百分比') VARCHAR(25)"`
	Basevolume    string    `xorm:"comment('交易量') VARCHAR(25)"`
	Quotevolume   string    `xorm:"comment('兑换货币交易量') VARCHAR(255)"`
	High24hr      string    `xorm:"comment('24小时最高价') VARCHAR(25)"`
	Low24hr       string    `xorm:"comment('24小时最低价') VARCHAR(25)"`
	CreateTime    time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	Updatetime    time.Time `xorm:"comment('更新时间') TIMESTAMP"`
}
