package models

import (
	"time"
	"bnwUdp/share"
)

type UdpDeviceChange struct {
	Id 		  int
	OldDevice string
	NewDevice string
	CreateAt  string
}

// 表名
func (model *UdpDeviceChange) TableName() string {
	return "ts_udp_device_change"
}

// 插入记录
func (model *UdpDeviceChange) Insert() error {
	str_insert := "insert into "+model.TableName()+" ( `old_device`, `new_device`, `create_at`) values ( '"+model.OldDevice+"', '"+model.NewDevice+"', '"+time.Now().Format("2006-01-02 15:04:05")+"');"
	_, err := share.ShareDb.Exec(str_insert)
	return err
}
