package main

import "fmt"

var (
	SimpleFactoryMap = map[string]IConfigParse{
		"json": SimpleFactory("json"),
		"xml":  SimpleFactory("xml"),
	}
)

func main() {
	// 1. 简单工厂模式
	fJson := SimpleFactory("json")
	xJson := SimpleFactory("xml")
	fmt.Printf("json:%s, xml:%s\n", fJson.Load("a.json").Host, xJson.Load("a.xml").Host)

	// 1.1 简单工厂模式，将对象缓存起来,减少内存开销
	cacheJson := SimpleFactoryMap["json"]
	cacheXml := SimpleFactoryMap["xml"]
	fmt.Printf("cache json:%s, xml:%s\n", cacheJson.Load("a.json").Host, cacheXml.Load("a.xml").Host)
}
