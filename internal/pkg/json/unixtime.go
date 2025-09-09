package json

import "time"

// UnixTime 自定义时间类型（解析时间戳）
type UnixTime time.Time

// UnmarshalJSON 实现自定义 JSON 解析器
func (t *UnixTime) UnmarshalJSON(b []byte) error {
	// 解析时间戳
	var ts int64
	if err := json.Unmarshal(b, &ts); err != nil {
		return err
	}
	*t = UnixTime(time.Unix(ts, 0)) // 解析 Unix 时间戳
	return nil
}

// MarshalJSON 用于反序列化
func (t UnixTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Unix()) // 以 Unix 时间戳形式输出
}
