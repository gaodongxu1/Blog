package models

import (
	_ "github.com/gaodongxu1/Blog/application/blog/routers/admin"
	_ "github.com/gaodongxu1/Blog/application/blog/routers/api"
	"time"
)
type Student struct {
 Id         uint64 `orm:pk`
 String168  string `orm:"size(168)"`
 Int64      int64
 Int32      int32
 String     string
 Bool       bool
 Time       time.Time `orm:"index"`
 Timedate   time.Time `orm:"type(date)"`
 Stringtext string    `orm:"type(text)"`
 Rune       rune
 Float32    float32
 Float64    float64
 Int        int
 Int8       int8
 Uint       uint
 Uint8      uint8
 Uint16     uint16
 Byte       byte
}
