package typings

type Config struct {
	Server   Server
	Postgres Postgres
}

type Server struct {
	Port string
	Mode string
}

type Postgres struct {
	Host     string
	Port     string
	Config   string
	DBName   string
	Username string
	Password string
	LogMode  string
	SSLMode  string
	TimeZone string
}

type Zap struct {
	Level        string // 级别
	Prefix       string // 日志前缀
	Format       string // 输出
	Directory    string // 日志文件夹
	MaxAge       int    // 日志留存时间
	ShowLine     bool   // 显示行
	LogInConsole bool   // 输出控制台
}
