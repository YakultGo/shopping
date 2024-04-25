package main

import "gorm.io/gen"

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "../../internal/data/query",
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})
	g.UseDB(connectDB(bc.Data.Database))
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}
