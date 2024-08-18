package connectToMysql

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

type TransactionStudy struct {
	Id      int64  `mapstructure:"id" gorm:"primarykey"`
	Name    string `mapstructure:"name"`
	Age     int32  `mapstructure:"age"`
	Address string `mapstructure:"address"`
}

func NewTransactionStudy(id int64, name string, age int32, address string) *TransactionStudy {
	return &TransactionStudy{Id: id, Name: name, Age: age, Address: address}
}

type Server struct {
	Port int32 `yaml:"port" mapstructure:"port"`
}
type DataBases struct {
	Mysql MySQLConfig `yaml:"mysql" mapstructure:"mysql"`
}

type Config struct {
	DataBases DataBases `yaml:"databases" mapstructure:"databases"`
	Server    Server    `yaml:"server" mapstructure:"server"`
}

type MySQLConfig struct {
	Username string `yaml:"username" mapstructure:"username"`
	Password string `yaml:"password" mapstructure:"password"`
	Ip       string `yaml:"ip" mapstructure:"ip"`
	Port     int32  `yaml:"port" mapstructure:"port"`
	Database string `yaml:"database" mapstructure:"database"`
}

// TableName 指定表名
func (TransactionStudy) TableName() string {
	return "transaction_study"
}

// MysqlClient is mysql client
var MysqlClient *gorm.DB

const mysqlConnectStr string = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"

// ConnectToMysql // 连接数据库
func ConnectToMysql(nacosClient config_client.IConfigClient) {
	mysqlConfigYaml, err := nacosClient.GetConfig(vo.ConfigParam{
		DataId: "paper-upload-project",
		Group:  "go",
	})
	//fmt.Println(mysqlConfigYaml)

	if err != nil {
		log.Fatalf("Parse config.mysql segment error: %s\n", err)
	}
	var config Config
	yaml.Unmarshal([]byte(mysqlConfigYaml), &config)

	//dsn格式 user:pass@tcp(ip:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf(mysqlConnectStr,
		config.DataBases.Mysql.Username,
		config.DataBases.Mysql.Password,
		config.DataBases.Mysql.Ip,
		config.DataBases.Mysql.Port,
		config.DataBases.Mysql.Database)

	mysqlClient, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",   // 表前缀
			SingularTable: true, // 禁用表名复数
		}})
	if err != nil {
		panic(err)
	}

	mysqlClientDB, _ := mysqlClient.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	mysqlClientDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	mysqlClientDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	mysqlClientDB.SetConnMaxLifetime(time.Hour)

	MysqlClient = mysqlClient
}
