# 简单工厂模式

## 含义
简单工厂模式(Simple Factory Pattern)：又称为静态工厂方法(Static Factory Method)模式，属于创建型模式，可以根据参数的不同返回不同类的实例，简单工厂模式专门定义一个类来负责创建其他类的实例，被创建的实例通常都具有共同的父类。

> 由于go中不支持静态方法，所以工厂类使用方法代替


## 模式结构
[![简单工厂模式结构](https://z3.ax1x.com/2021/10/19/5035Xq.md.png)](https://imgtu.com/i/5035Xq)

## 代码
```go
package simple_factory

import (
	"fmt"
)

type animal interface {
	eat(food string) string
}

type dog struct {
	name string
	age  int
}

func (d *dog) eat(food string) string {
	return fmt.Sprintf("dog %s eating %s", d.name, food)
}

type cat struct {
	name string
	age  int
}

func (c *cat) eat(food string) string {
	return fmt.Sprintf("cat %s eating %s", c.name, food)
}

func NewAnimal(a string) animal {
	switch a {
	case "cat":
		return &cat{"miao", 2}
	case "dog":
		return &dog{"wang", 1}
	default:
		return &cat{"m", 0}
	}
}

func main() {
	dog1 := NewAnimal("dog")
	cat1 := NewAnimal("cat")

	fmt.Printf("%s", dog1.eat("b"))
	fmt.Printf("%s", cat1.eat("c"))
}

```


## 分析
### 优点
* 简单工厂模式对对象的创建进行了封装，使得创建对象过程和对象本身的业务处理进行分离没降低了系统耦合性
* 工厂类根据条件决定创建那种对象，客户端免除直接创建对象的责任，而仅仅消费对象，实现了对责任的分割
* 客户端无需知道创建的对象的类名，只需要知道对象对应的参数即可

### 缺点
* 系统扩展困难，增加新的类时需要修改工厂类的判断逻辑，与开闭原则相违背
* 对象类型较多时，可能造成工厂方法逻辑复杂，不利于系统扩展和维护


## 适用场景
* 工厂类负责创建的对象比较少：由于创建的对象较少，不会造成工厂方法中的业务逻辑太过复杂。
* 客户端只知道传入工厂类的参数，对于如何创建对象不关心：客户端既不需要关心创建细节，甚至连类名都不需要记住，只需要知道类型所对应的参数

## 总结
简单工厂模式是一种创建型模式，工厂类提供的静态方法根据参数来决定创建哪个对象，创建过程封装在了静态方法里边，客户端不关心对象的创建过程，只负责消费对象，实现了对责任的分离，降低系统耦合性