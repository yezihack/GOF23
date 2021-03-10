package main

// 定义解析结果的结构体
type SimpleConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

// 定义配置的接口
type IConfigParse interface {
	Load(file string) *SimpleConfig
}

// JSON 解析器
type JsonConfigParse struct {
	file string
}

func (j *JsonConfigParse) Load(file string) *SimpleConfig {
	j.file = file
	return &SimpleConfig{
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		Password: "1234",
	}
}

// XML 解析器
type XmlConfigParse struct {
	file string
}

func (x *XmlConfigParse) Load(file string) *SimpleConfig {
	x.file = file
	return &SimpleConfig{
		Host:     "192.168.1.100",
		Port:     3306,
		User:     "root",
		Password: "1234",
	}
}

// 创建简单工厂模式
func SimpleFactory(t string) IConfigParse {
	switch t {
	case "json":
		return &JsonConfigParse{}
	case "xml":
		return &XmlConfigParse{}
	default:
		return nil
	}
}
