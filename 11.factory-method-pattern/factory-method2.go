package factory_method_pattern

// 实现存储系统，包括本地，远程

// 第一步：定义一个抽象产品
type SaveSystemProduct interface {
	PutData(key string, val interface{}) // 存储
	GetData(key string) interface{}      // 获取
}

// 第二步：定义一个抽象工厂
type ISaveSystemCreator interface {
	Create(s string) SaveSystemProduct
}

// 第三步：具体产品1，即实现抽象产品的特征
type LocalSave struct {
	data map[string]interface{}
}

func (s *LocalSave) PutData(key string, val interface{}) {
	s.data[key] = val
}
func (s *LocalSave) GetData(key string) interface{} {
	return s.data[key]
}

func NewLocalSave() SaveSystemProduct {
	return &LocalSave{
		data: make(map[string]interface{}),
	}
}

// 第三步：具体产品2，即实现抽象产品的特征
type RemoteSave struct {
	data map[string]interface{}
}

func (s *RemoteSave) PutData(key string, val interface{}) {
	s.data[key] = val
}
func (s *RemoteSave) GetData(key string) interface{} {
	return s.data[key]
}

func NewRemoteSave() SaveSystemProduct {
	return &RemoteSave{
		data: make(map[string]interface{}),
	}
}

// 第四步：具体工厂，即实现抽象工厂
type SaveSystemCreator struct {
}

func (s *SaveSystemCreator) Create(t string) SaveSystemProduct {
	switch t {
	case "local":
		return NewLocalSave()
	case "remote":
		return NewRemoteSave()
	default:
		return nil
	}
}
func NewSaveSystemCreator() ISaveSystemCreator {
	return &SaveSystemCreator{}
}
