/**
 * @Author: xiaoxiao
 * @Description:
 * @File:  rotate_count
 * @Version: 1.0.0
 * @Date: 2019-08-25 02:04
 */

package models

import "github.com/xhaoxiong/util/self"

type RotateCount struct {
	self.Model
	Count int64
}

type RotateCountRecord struct {
	self.Model
	UserAgent string
	Ip        string
	Address   string
	RotateStr string `gorm:"type:text"`
}
