package config

type Mysql struct {
	Path     string // 服务器地址
	Port     string // 端口
	Config   string // 高级配置
	Dbname   string // 数据库名
	Username string // 数据库用户名
	Password string // 数据库密码
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}
