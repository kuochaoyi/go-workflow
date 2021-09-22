package model

// Model 其它数据结构的公共部分
type Model struct {
	ID int `gorm:"primary_key" json:"id,omitempty"`
}
