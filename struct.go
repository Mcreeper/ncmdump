package ncmdump

import (
	"encoding/json"
	"strconv"
)

type Album struct {
	Id       interface{} `json:"albumId"`
	Name     string  `json:"album"`
	CoverUrl string  `json:"albumPic"`
}


// @Mcreeper: change for -- Error information:          interface conversion: interface {} is string, not float64
type Artist_Old struct {
	Name string
	Id   float64
}
// @see https://stackoverflow.com/questions/42377989/unmarshal-json-array-of-arrays-in-go
func (a *Artist) UnmarshalJSON_Old(data []byte) error {
	var v []interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	a.Name = v[0].(string)
	a.Id = v[1].(float64)
	return nil
}

// @Mcreeper changes starting here

type Artist struct {
	Name string
	Id   interface{}  // 改为 interface{}
}

func (a *Artist) UnmarshalJSON(data []byte) error {
	var v []interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	// 安全获取 Name
	if len(v) > 0 {
		if name, ok := v[0].(string); ok {
			a.Name = name
		}
	}

	// 安全获取 Id，支持数字和字符串
	if len(v) > 1 {
		switch id := v[1].(type) {
		case float64:
			a.Id = id
		case string:
			// 尝试转换为 float64
			if f, err := strconv.ParseFloat(id, 64); err == nil {
				a.Id = f
			} else {
				a.Id = id  // 保留字符串
			}
		}
	}
	return nil
}

// @Mcreeper changes ending here

// @ref https://music.163.com/#/song?id={id}
type Meta struct {
	Id       interface{} `json:"musicId"`
	Name     string  `json:"musicName"`
	*Album   `json:",inline"`
	Artists  []Artist `json:"artist"`
	BitRate  float64  `json:"bitrate"`
	Duration float64  `json:"duration"`
	Format   string   `json:"format"`
	Comment  string   `json:"-"`
}
