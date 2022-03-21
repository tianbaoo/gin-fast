package models

import (
	"fmt"
	"gin-fast/conf"
	"gin-fast/pkg/util"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
	//"github.com/go-redis/redis/v8"
)

var DB *gorm.DB

func SetUp(isOrmDebug bool) {
	conUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		                  conf.DataBase.User,
		                  conf.DataBase.Password,
		                  conf.DataBase.Host,
		                  conf.DataBase.Port,
		                  conf.DataBase.DB,
		                  conf.DataBase.Charset)

	db, err := gorm.Open(conf.DataBase.Type, conUri)
	if err != nil {
		panic(err)
	}
	DB = db
	DB.LogMode(isOrmDebug)
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return conf.DataBase.Prefix + defaultTableName
	}

	DB.AutoMigrate(&Account{})
	DB.AutoMigrate(&FileModel{})

}

type BaseModel struct {
	ID        uint64 `gorm:"primary_key'" json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

// BeforeCreate 生成全局唯一ID
func (m *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	if m.ID == 0 {
		m.ID = util.GenSonyFlakeId()
	}
	return nil
}

// Rdb 声明一个全局的rdb变量
var Rdb *redis.Client

// InitRedisClient redis旧版 非V8 https://www.liwenzhou.com/posts/Go/go_redis/
func InitRedisClient() (err error) {
	var Addr string
	Addr = fmt.Sprintf("%s:%s", conf.RedisCfg.Host, conf.RedisCfg.Port)
	Rdb = redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: conf.RedisCfg.Password, // no password set
		DB:       conf.RedisCfg.DB,  // use default DB
	})

	_, err = Rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func RedisExample() {
	err := Rdb.Set("score", 100, 5*time.Second).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}

	val, err := Rdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed, err:%v\n", err)
		return
	}
	fmt.Println("score", val)
	time.Sleep(10 * time.Second)
	val, err = Rdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed, err:%v\n", err)
		//return
	}
	fmt.Println("score", val)

	val2, err := Rdb.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	} else {
		fmt.Println("name", val2)
	}
}
