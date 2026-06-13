package autoload

import "fmt"

// MySQLConfig 保存 MySQL 连接设置。
type MySQLConfig struct {
	Host                    string `json:"host"`
	Port                    int    `json:"port"`
	Database                string `json:"database"`
	Username                string `json:"username"`
	Password                string `json:"password"`
	Charset                 string `json:"charset"`
	MaxOpenConns            int    `json:"max_open_conns"`
	MaxIdleConns            int    `json:"max_idle_conns"`
	ConnMaxLifetimeSeconds int    `json:"conn_max_lifetime_seconds"`
}

// DSN 返回 MySQL 数据源名称。
func (c MySQLConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local",
		c.Username, c.Password, c.Host, c.Port, c.Database, c.Charset)
}
