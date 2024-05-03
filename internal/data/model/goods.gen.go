// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameGood = "goods"

// Good mapped from table <goods>
type Good struct {
	ID          int64   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Price       float64 `gorm:"column:price;comment:价格" json:"price"`             // 价格
	Deal        int32   `gorm:"column:deal;comment:销量" json:"deal"`               // 销量
	Description string  `gorm:"column:description;comment:描述" json:"description"` // 描述
	Shop        string  `gorm:"column:shop;comment:店铺" json:"shop"`               // 店铺
	Location    string  `gorm:"column:location;comment:地址" json:"location"`       // 地址
	Postfree    int32   `gorm:"column:postfree;comment:是否包邮" json:"postfree"`     // 是否包邮
	Category    string  `gorm:"column:category;comment:类别" json:"category"`       // 类别
}

// TableName Good's table name
func (*Good) TableName() string {
	return TableNameGood
}
