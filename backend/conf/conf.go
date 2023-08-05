package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	ValidEmail string
	SmtpHost   string
	SmtpEmail  string
	SmtpPass   string

	PhotoHost       string
	CoursePath      string
	AssignmentPath  string
	AssSolutionPath string
	AvatarPath      string
	CourseImgPath   string

	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string
)

func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadServer(file)
	LoadMysqlData(file)
	LoadEmail(file)
	LoadPhotoPath(file)
	LoadRedisData(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func LoadEmail(file *ini.File) {
	ValidEmail = file.Section("email").Key("ValidEmail").String()
	SmtpHost = file.Section("email").Key("SmtpHost").String()
	SmtpEmail = file.Section("email").Key("SmtpEmail").String()
	SmtpPass = file.Section("email").Key("SmtpPass").String()
}

func LoadPhotoPath(file *ini.File) {
	PhotoHost = file.Section("path").Key("Host").String()
	CoursePath = file.Section("path").Key("CoursePath").String()
	AssignmentPath = file.Section("path").Key("AssignmentPath").String()
	AssSolutionPath = file.Section("path").Key("AssSolutionPath").String()
	AvatarPath = file.Section("path").Key("AvatarPath").String()
	CourseImgPath = file.Section("path").Key("CourseImgPath").String()
}

func LoadRedisData(file *ini.File) {
	RedisDb = file.Section("redis").Key("RedisDb").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}
