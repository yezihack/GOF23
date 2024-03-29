# 工厂方法模式

定义一个用于创建对象的接口，让子类决定实例化哪一个类。工厂方法使一个类的实例化延迟到其子类

工厂方法实现过程：
1. 首先定义一个抽象产品 product，该产品负责定义产品的共性。
2. 再定义一个工厂，creator， 创建具体产品由具体工厂来完成。


类图如下
1. Product 是抽象产品
2. ConcreteProduct 是具体产品的实现，实现 Product 抽象产品
3. Creator 是抽象工厂
4. ConcreteCreator 是具体工厂的实现，实现了 Creator 抽象工厂
5. ConcreteCreator 具体工厂依赖具体产品，如巧妇难为无米之炊。
![](https://cdn.jsdelivr.net/gh/yezihack/assets/b/202111291805261.png)

# 优缺点
## 优点
1. 你可以避免创建者和具体产品之间的紧密耦合。
1. 单一职责原则。 你可以将产品创建代码放在程序的单一位置， 从而使得代码更容易维护。
1. 开闭原则。 无需更改现有客户端代码， 你就可以在程序中引入新的产品类型。

## 缺点
1. 应用工厂方法模式需要引入许多新的子类， 代码可能会因此变得更复杂。 最好的情况是将该模式引入创建者类的现有层次结构中。

# 应用
1. 日志记录器，记录可能记录到本地，数据库，远程服务器等。用户可以选择记录到什么地方。
1. 数据库的访问。可能是MySQL, PG, MongoDB等。 
1. 缓存方式。可能使用 Redis，Memcache, 本地文件等


# 总结
1. 定义一个抽象产品
2. 定义一个抽象工厂
3. 具体产品实现抽象产品
4. 具体工厂实现抽象工厂
5. 具体工厂依赖具体产品