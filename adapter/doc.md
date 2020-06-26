# 适配器模式
## 含义
将一个类的接口转换成客户希望的另一种接口。使得原本接口不兼容的类可以一起工作。
## 场景
* 调用者-Client；待适配者-Adaptee；适配器-Adapter
* Client 和 Adaptee，各自可完成自己的任务，但接口不匹配、无法协同工作。比如：使用第三方组件，不能改组件接口或调用方的实现。
* 通过适配器屏蔽不同Adaptee的数据类型、格式差异，Client无需关心底层的数据结构和组织。
* Adaptee只实现了Client需要的部分接口，使用 Adapter封装后，再补全其它接口，或封装多个Adaptee提供给Client用。
## 实现
1. 类适配器：Client、与Adaptee接口不同，Client希望调用Target形式的接口，适配器（Adapter）多重继承自Target、Adaptee，在实现目标接口时，调用Adaptee的某些接口。
2. 对象适配器：Client、与Adaptee接口不同，Client希望调用Target形式的接口，适配器（Adapter）继承自Target，内部创建Adaptee对象，封装接口调用。
