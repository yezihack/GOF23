package factory_method_pattern

// 实现存储系统，包括本地，远程

// 第一步：定义一个抽象产品
type SaveSystemProduct interface {
	PutData(key string, val interface{}) // 存储
	GetData(key string) interface{}      // 获取
}

// 第二步：定义一个抽象工厂
type SaveSystemCreator interface {
	Create(s string) SaveSystemProduct
}

// 第三步：具体产品，即实现抽象产品的特征
type SaveSystem struct {
	data map[string]interface{}
}

func (s *SaveSystem) PutData(key string, val interface{}) {
	s.data[key] = val
}
func (s *SaveSystem) GetData(key string) interface{} {
	return s.data[key]
}

type LocalSave struct {
	SaveSystem
}

func NewLocalSave() SaveSystemProduct {
	return &LocalSave{}
}
