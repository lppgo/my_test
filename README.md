# my_test

点点滴滴，积累实例

## 心得体会 （有些是自己的，有些是身边人的案例）

### 1. 团队沟通很重要。

（1）尤其是一个人遇到问题较长时间不能解决，不及时寻求帮助和和领导反馈困难。  
 （2）会把问题积累并放大。拖延项目进度。甚至会影响团队士气  
 （3）个人会陷入思维误区，你遇到的问题很可能就是前人刚刚踩过的坑  
 （4）每个人水平，代码习惯等有好有差，也会有各种 Bug,团队可以相互帮助。出现问题也是只针对问题，不针对人员

（5）修改公共的包或者方法,库(util,common)，尤其要评估对现有的项目影响范围，对涉及的人员做好及时沟通。
不要自己偷偷改了东西窃喜，坑了别人甚至引发灾难。

### 2.借鉴优秀的源码。主要是用法和写法

（1）B 站 错误处理，context  
 （2）业务代码和非业务代理分离，中间件，公共方法  
 （3）请求追踪可以用 requestID,context, \*http.Request  
 （4）当需要多个 for 循环的时候，试着能不能把嵌套的 for 改成前后执行。这样遍历次数会比笛卡尔积少很多。

### 3.不要为了设计而设计（过度设计）

（1）优秀的代码是将复杂的业务简单化，代码简洁明了。

### 4.新项目做之前，先要充分理解需求，业务场景，目标

（1）进行思考，明确需求  
 （2）写代码之前，先做代码规范。数据库，go 框架，涉及的技术解决什么问题，命名规范，项目结构规范  
 （3）数据库规范，索引，非空值等  
 （4）以简单实用为主基调  
 （5）项目开始之初，编码也应该做好统一，比如 Go 默认是 utf-8。每个人编码不同，文档这些每个人打开都可能是乱码，每次都得转码。但转码之后又影响别人，而且都不想改自己的编码。影响沟通和配合  
 （6）静态代码检查，越早越好。我们在开发的后期才意识到引入静态代码检查，其实最好的做法是在项目开始时就及时使用，并以较严格的标准保证主分支的代码质量。在开发后期才引入的问题是，已经有太多代码不符合标准。所以我们不得不短期内忽略了很多检查项。来自《知乎社区核心业务 Golang 化实践https://zhuanlan.zhihu.com/p/48039838》

### 5.每一步都不得随意

写代码时，一定要对每一个变量，做到精确的控制，走到哪一步可能有什么结果，心里都要有数，模棱两可的时候，就是大概率会埋坑的时侯

`回来补充5点，rust设计就严格控制了变量的所属权`

### 6. CLI Args 和配置文件各有什么特点

- 命令行参数与配置文件各有什么优劣，特点，进行思考

### 7. 技术选型

- 要做好技术选型调研，出一份调研报告，并对几个产品的各自特点，是否适合目前的业务，是否适合当前的团队，最好能和进行讨论评审

### 8. 做好架构设计

- 架构设计等同于数据设计，梳理清楚数据的走向和逻辑。尽量避免环形依赖，数据重复请求等
