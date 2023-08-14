package DB

// DBConfig is a struct that represents the database configuration
type DBConfig struct {
	Username string `json:"username"` //账号
	Password string `json:"password"` //密码
	Host     string `json:"host"`     //数据库地址，可以是Ip或者域名
	Port     int    `json:"port"`     //数据库端口
	Dbname   string `json:"dbname"`   //数据库名
	Timeout  string `json:"timeout"`  //连接超时，10秒
}
