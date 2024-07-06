# 项目简介
tiktok是基于接口定义实现相对应功能的微服务项目
# 项目架构
- 整体借鉴了整洁架构的思想，对项目进行分层。
- 尽可能的保持了单一职责的原则来设计接口，以及通过适当的抽象来实现依赖分离，从而实现了底层实现的可替代性。

**或许你对整洁架构感到陌生,在看源代码时不知所措,晕头转向。不要担心，你可以看这个**

[tiktok中的整洁架构](https://o0e45m7p53e.feishu.cn/docx/X8yHdfa6yoqb9FxrAbCctQ56n5c)

# 项目结构
```text
├─.github                # 定义github Action
│  └─workflows
├─app                    # 项目主要实现部分
│  ├─gateway             # Gateway 模块
│  │  ├─cmd
│  │  ├─config
│  │  ├─domain           # 实体层 (领域层)
│  │  │  ├─model         # model定义被可被复用或需要抽象的接口
│  │  │  ├─repository    # 定义持久化的接口
│  │  │  └─service       # 核心业务处理
│  │  ├─interface        # 接口层，转换数据的格式，使数据可以在整个项目中流动
│  │  │  ├─handler       # 对HTTP传入的数据进行简单处理
│  │  │  ├─middleware    # HTTP层的中间件
│  │  │  ├─pack          # 对常用函数的包装
│  │  │  ├─persistence   # 持久化，实现repository中定义的接口以完成依赖反转
│  │  │  │  └─database   # db
│  │  │  ├─router
│  │  │  └─rpc           # rpc-client
│  │  └─usecase          # 用例层，衔接起interface与domain，简单的业务处理
│  │      └─pack
│  ├─interaction         # interaction 模块
│  │  ├─cmd
│  │  │  └─pack
│  │  ├─config
│  │  ├─domain           # 领域层,接口的定义或核心代码实现
│  │  │  ├─model         # 定义需要接口层实现的接口
│  │  │  ├─repository    # 定义需要persistence实现的接口
│  │  │  └─service       # 核心业务代码实现
│  │  ├─interface        # 接口层,控制数据的流动和实现领域层的接口
│  │  │  └─persistence   # 持久化,实现repository定义的接口
│  │  │      └─database
│  │  └─usecase          # 用例层,对业务的逻辑梳理
│  │      └─pack
│  ├─relation            # 同interaction模块
│  ├─user                # 同上
│  └─video               # 同上
├─deploy                 # 部署所需的相关文件
│  ├─common
│  ├─gateway
│  ├─interaction
│  ├─relation
│  ├─user
│  └─video
├─docs                   # 文档
├─pkg
│  ├─consts              # 定义一些常量
│  ├─discovery           # 服务发现
│  ├─errno               # 定义整个服务的错误类型
│  ├─idl
│  │  └─script
│  ├─inject              # 依赖注入
│  ├─kafka
│  ├─kitex_gen           # idl生成的代码
│  ├─log                 # 日志模块
│  ├─oss
│  ├─snowflake           # 生成ID
│  ├─trace
│  ├─types
│  └─utils               # 工具函数
└─scripts                # 一些脚本

```


# 文档
[项目报告](https://o0e45m7p53e.feishu.cn/docx/Mbh6d1GbBouSfUxHVzrc9QeDnyg)
