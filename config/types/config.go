package types

type Server struct {
	WebDomain string
	Port      string
	Mode      string
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

type Redis struct {
	Network  string
	Addr     string // Redis 服务器地址
	Password string // Redis 访问密码，如果没有设置密码则为空字符串
	DB       int    // 选择使用的数据库，默认为0
}

type Email struct {
	SMTPServer     string
	SMTPPort       int
	From           string
	Password       string
	DefaultSubject string
}
