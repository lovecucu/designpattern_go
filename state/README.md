## 状态模式

### 定义

- **状态模式**允许对象在内部状态改变时改变它的行为，对象看起来好像修改了它的类

### 设计原则

- 封装变化
- 针对接口编程，而不是针对实现编程
- 多用组合，少用继承
- 为交互对象之间的松耦合设计而努力
- 类应该对扩展开放，对修改关闭
- 依赖抽象，不要依赖具体类
- 最少知识原则：只和你的密友谈话
- 单一职责原则：类应该只有一个改变的理由

### 要点

- 状态模式允许一个对象基于内部状态而拥有不同的行为
- 和程序状态机（PSM）不同，状态模式用类代表状态
- Context会将行为委托给当前状态对象
- 通过将每个状态封装进一个类，我们把以后需要做的任何改变局部化了
- 状态模式和策略模式有相同的类图，但是它们的意图不同。（状态模式预定义了策略的改变）
- 策略模式通常会用行为或算法来配置Context类
- 状态模式允许Context随着状态的改变而改变行为
- 状态转换可以由State类或Context类控制
- 使用状态模式通常会导致设计中类的数目大量增加
- 状态类可以被多个Context实例共享