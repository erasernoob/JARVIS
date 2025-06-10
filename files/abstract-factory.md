---
title: 设计模式-创建型-抽象工厂模式
publishDate: 2025-05-10
description: 设计模式之抽象工厂模式
tags:
  - backend
  - clutter
  - design-pattern
---
## 抽象工厂模式
### 具体场景引入
开发过程中代码中有一些类，包括：
1. 一系列相关产品（**家具**）`椅子`、`沙发`、`桌子`
2. 系列产品的不同变体（家具的**风格**）`现代`、`维多利亚`、`装饰艺术`
显然，我们需要找到一次性能够创建一套风格统一的家具发送给客户的方法。**并且不希望在添加新产品或者是新风格时修改已有的代码**。
### 解决方案
>抽象模式建议为系列中的每件产品都明确的声明接口：（椅子、沙发、桌子）。**保证所有的产品的变体都继承这些产品的接口**
![](./Pasted%20image%2020250507094431.png)

与之相对应的，工厂类的声明也是**大同小异**。**抽象工厂类 (Interface)**的声明，包含了所有产品的构造函数，这些构造函数一定返回的都是**抽象**的产品类型。
![](./Pasted%20image%2020250507094928.png)

对于**产品变体的处理**，正如上图所示。不同的产品变体，实现抽象工厂类，基于该接口创建对应的不同的工厂类（对于产品变体）。**一个工厂类只能返回一个对应风格的椅子**。这样就达成了一次性产生相同风格的家具的目的。
#### 初始化说明
在抽象工厂模式中，**客户端只是接触抽象接口，创建哪个实际的工厂对象，由应用程序在初始化阶段读取客户端提供的配置文件中做出选择，设定某个具体的工厂类别。**。
#### 整体架构
![](./Pasted%20image%2020250507095420.png)
### 具体适用场景
- 如果代码需要与多个**不同系列**的**相关产品**进行交互，并且需要未来的扩展性，不希望直接根据据产品的具体类进行创建，此时可以使用抽象工厂
### 实现方式
1. 首先，根据不同的产品类型和产品变体为维度绘制矩阵。
2. 声明产品抽象接口，让所有具体的产品类实现对应的抽象接口
3. 声明工厂抽象接口，**并在接口中为每一个抽象产品提供构建方法**
4. 为每一个产品变体声明一个具体工厂类
5. 在应用程序中开发初始化代码。 该代码根据应用程序配置或当前环境， 对特定具体工厂类进行初始化。 然后将该工厂对象传递给所有需要创建产品的类。
6. **找出代码中所有对产品构造函数的直接调用， 将其替换为对工厂对象中相应构建方法的调用。**
### 代码示例
> 如果你想要购买一组运动装备， 比如一双鞋与一件衬衫这样由两种不同产品组合而成的套装。 相信你会想去购买同一品牌的商品， 这样商品之间能够互相搭配起来。

```go
package main

// 抽象产品类
type Shoe struct {
	brand string
	color string
}

type IShoe interface {
	getColor() (string, error)
	getBrand() (string, error)
}

func (s *Shoe) getBrand() (string, error) {
	return s.brand, nil
}

func (s *Shoe) getColor() (string, error) {
	return s.color, nil
}

type Shirt struct {
	brand string
	color string
}

type IShirt interface {
	getColor() (string, error)
	getBrand() (string, error)
}

func (s *Shirt) getBrand() (string, error) {
	return s.brand, nil
}

func (s *Shirt) getColor() (string, error) {
	return s.color, nil
}

type AdidasShirt struct {
	Shirt
}

type AdidasShoe struct {
	Shoe
}

type NikeShirt struct {
	Shirt
}

type NikeShoe struct {
	Shoe
}

type NikeFactory struct{}

func (*NikeFactory) createShoe() IShoe {
	// 这里NikeShirt间接实现了Shirt
	return &NikeShoe{
		Shoe: Shoe{
			brand: "Nike",
			color: "blue",
		},
	}
}

func (*NikeFactory) createShirt() IShirt {
	// 这里NikeShirt间接实现了Shirt
	return &NikeShirt{
		Shirt: Shirt{
			brand: "Nike",
			color: "blue",
		},
	}
}

type AdidasFactory struct{}

func (*AdidasFactory) createShoe() IShoe {
	// 这里AdidasShirt间接实现了Shirt
	return &AdidasShoe{
		Shoe: Shoe{
			brand: "adidas",
			color: "blue",
		},
	}
}

func (*AdidasFactory) createShirt() IShirt {
	// 这里AdidasShirt间接实现了Shirt
	return &AdidasShirt{
		Shirt: Shirt{
			brand: "adidas",
			color: "blue",
		},
	}
}

type ISportsFactory interface {
	createShoe() IShoe
	createShirt() IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &AdidasFactory{}, nil
	} else {
		return &NikeFactory{}, nil

	}

}

```

## 大数据智能团队：
1. 大数据智能团队是由西南石油大学计算机软件学院成立的学生工作室
## 其他
大数据智能与创新团队（以下简称“团队”）成立于2018年，由计算机科学学院数据科学与大数据技术专业教研室、智能与网络化系统研究中心全力打造，大数据驱动、结合统计学、数据挖掘、人工智能等技术，把大数据分析和平台开发能力、实践能力和创新能力培养作为交叉点和教学重点。在日常学习工作中，团队积极组织参加各种大数据相关学习、培训、学科竞赛，比如中国高校计算机大赛 - 人工智能创意赛/大数据挑战赛、中国大学生服务外包创新创业大赛、全国大学生电子商务三创赛等。结合社会发展需要，团队能为现代企事业单位提供大数据平台系统设计、实施与维护、大数据分析的高级数据科学人才、创新型人才和复合型人才。
获奖列表教育部产学研google创新创业基金（1项）教育部产学研阿里云创新创业联合基金（1项）全国大学生电子商务“创新、创意及创业”挑战赛国赛三等奖（1项））省赛一等奖（3项）中国大学生计算机设计大赛国家级三等奖“互联网+”大学生创新创业大赛省赛银奖（2项）铜奖（1项）四川省教育厅创业训练项目（5项）
校级开放实验（5项）
团队建设
团队定期会有老师帮大家答疑解惑，开阔视野发散思维