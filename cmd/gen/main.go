package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

//go:generate go run main.go
func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../../internal/data/query",
		Mode:    gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	g.UseDB(connect())
	//autoUpdateTimeField := gen.FieldGORMTag("update_time", func(tag field.GormTag) field.GormTag {
	//	return tag.Set("autoUpdateTime", "true")
	//})
	//autoCreateTimeField := gen.FieldGORMTag("create_time", func(tag field.GormTag) field.GormTag {
	//	return tag.Set("autoCreateTime", "true")
	//})
	//softDeleteField := gen.FieldType("delete_time", "soft_delete.DeletedAt")
	//fieldOpts := []gen.ModelOpt{autoCreateTimeField, autoUpdateTimeField, softDeleteField}
	//g.ApplyBasic(g.GenerateAllTable(fieldOpts...)...)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}

func connect() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:13306)/shop?charset=utf8mb4&parseTime=True"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
