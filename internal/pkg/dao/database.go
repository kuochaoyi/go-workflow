package dao

import (
	"log"
	"strconv"

	config "github.com/kuochaoyi/go-workflow/workflow-config"
	"github.com/kuochaoyi/go-workflow/workflow-engine/model"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Model 其它数据结构的公共部分
type Model struct {
	ID int `gorm:"primary_key" json:"id,omitempty"`
}

// 配置
var conf = *config.Config

// Setup 初始化一个db连接
func Setup() {
	//var err error
	log.Println("数据库初始化！！")
	//db, err = gorm.Open(conf.DbType, fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.DbUser, conf.DbPassword, conf.DbHost, conf.DbPort, conf.DbName))
	//if err != nil {
	//	log.Fatalf("数据库连接失败 err: %v", err)
	//}
	// 启用Logger，显示详细日志
	db, _ := New()

	//mode, _ := strconv.ParseBool(conf.DbLogMode)
	//
	//db.LogMode(mode)

	//db.SingularTable(true) //全局设置表名不可以为复数形式

	// db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	//idle, err := strconv.Atoi(conf.DbMaxIdleConns)
	//if err != nil {
	//	panic(err)
	//}
	//db.DB().SetMaxIdleConns(idle)
	//open, err := strconv.Atoi(conf.DbMaxOpenConns)
	//if err != nil {
	//	panic(err)
	//}
	//db.DB().SetMaxOpenConns(open)

	//db.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").AutoMigrate(&model.Procdef{}).
	db.AutoMigrate(&model.Procdef{}).
		AutoMigrate(&model.Execution{}).AutoMigrate(&model.Task{}).
		AutoMigrate(&model.ProcInst{}).AutoMigrate(&model.Identitylink{}).
		AutoMigrate(&model.ExecutionHistory{}).
		AutoMigrate(&model.IdentitylinkHistory{}).
		AutoMigrate(&model.ProcInstHistory{}).
		AutoMigrate(&model.TaskHistory{}).
		AutoMigrate(&model.ProcdefHistory{})
	db.Model(&model.Procdef{}).AddIndex("idx_id", "id")
	db.Model(&model.ProcInst{}).AddIndex("idx_id", "id")
	db.Model(&model.Execution{}).AddForeignKey("proc_inst_id", "proc_inst(id)", "CASCADE", "RESTRICT").AddIndex("idx_id", "id")
	db.Model(&model.Identitylink{}).AddForeignKey("proc_inst_id", "proc_inst(id)", "CASCADE", "RESTRICT").AddIndex("idx_id", "id")
	db.Model(&model.Task{}).AddForeignKey("proc_inst_id", "proc_inst(id)", "CASCADE", "RESTRICT").AddIndex("idx_id", "id")
	//---------------------历史纪录------------------------------
	db.Model(&model.ProcInstHistory{}).AddIndex("idx_id", "id")
	db.Model(&model.ExecutionHistory{}).AddForeignKey("proc_inst_id", "proc_inst_history(id)", "CASCADE", "RESTRICT").AddIndex("idx_id", "id")
	db.Model(&model.IdentitylinkHistory{}).AddForeignKey("proc_inst_id", "proc_inst_history(id)", "CASCADE", "RESTRICT").AddIndex("idx_id", "id")
	db.Model(&model.TaskHistory{}).
		//  AddForeignKey("proc_inst_id", "proc_inst_history(id)", "CASCADE", "RESTRICT").
		AddIndex("idx_id", "id")
	// db.Model(&Comment{}).AddForeignKey("proc_inst_id", "proc_inst(id)", "CASCADE", "RESTRICT")
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer db.Close()
}

// GetDB getdb
func GetDB() *gorm.DB {
	return db
}

// GetTx GetTx
func GetTx() *gorm.DB {
	return db.Begin()
}
