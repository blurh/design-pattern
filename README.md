# Design Patterns 设计模式

## 设计原则:  
- 单一职责原则(Single Responsibility Principle, SRP):  
  类的职责应该单一, 一个方法只做一件事  
    
- 开闭原则(Open/Close Principle, OCP):  
  对扩展开放, 对修改关闭  
  程序需要扩展时, 添加类来扩展而不是修改原来的代码  
    
- 里式替换原则(Liskov Substitution Principle, LSP):  
  在基于类的语言中, 里氏替换原则通常解释为: 具有各种具体子类型的抽象基类的规范  
  子类对象能够替换程序(program)中父类对象出现的任何地方, 并且保证原来程序的逻辑行为(behavior)不变及正确性不被破坏  
  但 Go 没有类或继承, 因此无法根据抽象类层次结构实现替换  
    
- 依赖倒置原则(Dependence Inversion Principle, DIP):  
  高层次的模块不应该依赖于低层次的模块, 两者都应该依赖于抽象接口  
    
- 接口隔离原则(Interface Segregation Principle, ISP):  
  类不应该依赖它不需要的接口, 即一个类对另一个类的依赖应该建立在最小的接口上  
    
- 迪米特法则(Law of Demeter, LOD), 又称最少知识原则(Principle of Least Knowledge):  
  不该有直接依赖关系的类之间, 不要有依赖. 有依赖关系的类之间, 尽量只依赖必要的接口  
