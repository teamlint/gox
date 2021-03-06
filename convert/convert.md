

# convert
`import "."`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package convert 类型转化




## <a name="pkg-index">Index</a>
* [func ToBool(i interface{}) bool](#ToBool)
* [func ToBytes(i interface{}) []byte](#ToBytes)
* [func ToFloat32(i interface{}) float32](#ToFloat32)
* [func ToFloat64(i interface{}) float64](#ToFloat64)
* [func ToInt(i interface{}) int](#ToInt)
* [func ToInt16(i interface{}) int16](#ToInt16)
* [func ToInt32(i interface{}) int32](#ToInt32)
* [func ToInt64(i interface{}) int64](#ToInt64)
* [func ToInt8(i interface{}) int8](#ToInt8)
* [func ToString(i interface{}) string](#ToString)
* [func ToStringSlice(i interface{}) []string](#ToStringSlice)
* [func ToTime(i interface{}, format ...string) time.Time](#ToTime)
* [func ToTimeDuration(i interface{}) time.Duration](#ToTimeDuration)
* [func ToType(i interface{}, t string) interface{}](#ToType)
* [func ToUint(i interface{}) uint](#ToUint)
* [func ToUint16(i interface{}) uint16](#ToUint16)
* [func ToUint32(i interface{}) uint32](#ToUint32)
* [func ToUint64(i interface{}) uint64](#ToUint64)
* [func ToUint8(i interface{}) uint8](#ToUint8)


#### <a name="pkg-files">Package files</a>
[convert.go](/src/target/convert.go) 





## <a name="ToBool">func</a> [ToBool](/src/target/convert.go?s=3751:3782#L172)
``` go
func ToBool(i interface{}) bool
```
ToBool false: "", 0, false, off



## <a name="ToBytes">func</a> [ToBytes](/src/target/convert.go?s=2257:2291#L103)
``` go
func ToBytes(i interface{}) []byte
```
ToBytes 转化为[]byte



## <a name="ToFloat32">func</a> [ToFloat32](/src/target/convert.go?s=6581:6618#L356)
``` go
func ToFloat32(i interface{}) float32
```
ToFloat32 转化为float32



## <a name="ToFloat64">func</a> [ToFloat64](/src/target/convert.go?s=6792:6829#L368)
``` go
func ToFloat64(i interface{}) float64
```
ToFloat64 转化为float64



## <a name="ToInt">func</a> [ToInt](/src/target/convert.go?s=3990:4019#L186)
``` go
func ToInt(i interface{}) int
```
ToInt 转化为int



## <a name="ToInt16">func</a> [ToInt16](/src/target/convert.go?s=4773:4806#L238)
``` go
func ToInt16(i interface{}) int16
```
ToInt16 转化为int16



## <a name="ToInt32">func</a> [ToInt32](/src/target/convert.go?s=4931:4964#L249)
``` go
func ToInt32(i interface{}) int32
```
ToInt32 转化int32



## <a name="ToInt64">func</a> [ToInt64](/src/target/convert.go?s=5092:5125#L260)
``` go
func ToInt64(i interface{}) int64
```
ToInt64 转化为int64



## <a name="ToInt8">func</a> [ToInt8](/src/target/convert.go?s=4616:4647#L227)
``` go
func ToInt8(i interface{}) int8
```
ToInt8 转化为int8



## <a name="ToString">func</a> [ToString](/src/target/convert.go?s=2441:2476#L114)
``` go
func ToString(i interface{}) string
```
ToString 基础的字符串类型转换



## <a name="ToStringSlice">func</a> [ToStringSlice](/src/target/convert.go?s=3408:3450#L155)
``` go
func ToStringSlice(i interface{}) []string
```
ToStringSlice 转化为字符串切片



## <a name="ToTime">func</a> [ToTime](/src/target/convert.go?s=1450:1504#L67)
``` go
func ToTime(i interface{}, format ...string) time.Time
```
ToTime 转化为time.Time



## <a name="ToTimeDuration">func</a> [ToTimeDuration](/src/target/convert.go?s=2142:2190#L98)
``` go
func ToTimeDuration(i interface{}) time.Duration
```
ToTimeDuration 将变量i转换为time.Duration类型



## <a name="ToType">func</a> [ToType](/src/target/convert.go?s=711:759#L25)
``` go
func ToType(i interface{}, t string) interface{}
```
ToType 将变量i转换为字符串指定的类型t



## <a name="ToUint">func</a> [ToUint](/src/target/convert.go?s=5251:5282#L271)
``` go
func ToUint(i interface{}) uint
```
ToUint 转化为unit



## <a name="ToUint16">func</a> [ToUint16](/src/target/convert.go?s=6075:6110#L323)
``` go
func ToUint16(i interface{}) uint16
```
ToUint16 转化为uint16



## <a name="ToUint32">func</a> [ToUint32](/src/target/convert.go?s=6243:6278#L334)
``` go
func ToUint32(i interface{}) uint32
```
ToUint32 转化为uint32



## <a name="ToUint64">func</a> [ToUint64](/src/target/convert.go?s=6411:6446#L345)
``` go
func ToUint64(i interface{}) uint64
```
ToUint64 转化为uint64



## <a name="ToUint8">func</a> [ToUint8](/src/target/convert.go?s=5911:5944#L312)
``` go
func ToUint8(i interface{}) uint8
```
ToUint8 转化为uint8








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
