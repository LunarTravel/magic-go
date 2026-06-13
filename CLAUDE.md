# magic-go 项目规范

> 本文件为 magic-go 项目的强制规范。每次对本项目（尤其 `magic-service/` 下的 Go 代码）做修改前，必须先遵循本规范。

## 1. 项目与技术栈

- **语言**: Go 1.25；**Web 框架**: Gin；**日志**: zap；**配置**: YAML(`configs/config.yaml`) + 环境变量(`.env`)
- **存储**: MySQL（`database/sql`）；**缓存**: Redis（`go-redis`）
- **架构**: DDD（领域驱动设计），模块名 `magic-service`
- **仓库定位**: `magic-go` 为 monorepo，`magic-service/` 是 Go 后端，后续 `python-agent/`、`frontend/` 为同级目录

## 2. 仓库结构

```
magic-go/                        # monorepo 根
├── CLAUDE.md                    # 本规范
├── .gitignore                   # monorepo 级忽略
└── magic-service/               # Go 后端服务
    ├── main.go / wire.go        # 入口与依赖组装
    ├── go.mod                   # module magic-service
    ├── configs/config.yaml      # 配置（${VAR:-默认值} 占位符）
    ├── .env / .env.example      # 本地私密配置 / 模板
    ├── Dockerfile
    └── internal/
        ├── domain/{module}/     # 领域层
        ├── application/{module}/# 应用层
        ├── interfaces/          # 接口层（http/）
        ├── infrastructure/      # 基础设施层
        ├── config/ constants/ di/ pkg/
```

## 3. 分层与依赖

四层位于 `magic-service/internal/`：

```
domain → application → interfaces
infrastructure（技术底座，各层可用）
```

- **业务流向**: `Interface → Application → Domain`
- **依赖方向**: 外层依赖内层；**`domain` 不依赖任何外层包**
- **仓储**: 接口定义在 `domain/{module}/repository`，实现在 `infrastructure/persistence`（Go 惯例：domain 定义 interface，infrastructure 实现）
- **禁止**: 跨领域直接调用对方仓储；领域间优先用事件解耦

## 4. 目录树指引

```
magic-service/internal/
├── domain/{module}/                 # {module} 为限界上下文：chat/agent/contact/flow/knowledge/file...
│   ├── model/                       # 实体 Entity、值对象 ValueObject
│   ├── repository/                  # 仓储接口（Go interface）
│   └── service/                     # 领域服务 *DomainService
├── application/{module}/
│   ├── service/                     # 应用服务 *AppService
│   ├── dto/                         # Request/Response DTO
│   └── assembler/                   # DTO ↔ Entity 转换
├── interfaces/
│   ├── http/
│   │   ├── handler/                 # HTTP 处理器
│   │   ├── router/                  # 路由注册（仅在此注册，禁止注解式路由）
│   │   └── middleware/             # CORS / RequestID / 日志
│   └── dto/                         # 统一响应（APIResponse）
├── infrastructure/
│   ├── persistence/{mysql,redis}/   # 仓储实现 + 客户端；PO 仅存于此
│   ├── logging/                     # zap 封装
│   └── appruntime/                  # 时区 / 优雅关闭
├── config/                          # 配置加载 + autoload 结构体
├── constants/ di/ pkg/              # 常量 / 依赖注入 Provider / 公共包
```

新增能力：**先在 `domain/{module}` 定模型与仓储契约**，再补 `application` 与 `interfaces`，基础设施实现跟外部依赖走 `infrastructure`。

## 5. Go 工程约定

- 所有代码经 `gofmt` / `goimports` 格式化；遵循 **Effective Go**
- **包名**: 小写单词、单数（如 `model` 非 `models`、`service` 非 `services`）
- **命名**: 类型 `PascalCase`（导出）并加文档注释；未导出 `camelCase`；常量 `PascalCase` 或 `UPPER_SNAKE_CASE`
- **错误**: 显式 `error` 返回值；包装用 `fmt.Errorf("...: %w", err)`；业务中不 `panic`
- **`context.Context`** 作为方法首个参数
- **路由**: 仅在 `interfaces/http/router/` 注册，禁止注解式路由

## 6. 参数传递（强类型链路）

| 方向 | 路径 |
|------|------|
| 入站 | handler(Request) → **Request DTO** → Assembler → **Domain Entity/VO** → Repository |
| 出站 | Repository → **Domain Entity** → Assembler → **Response DTO** → handler |

各层**对外签名**允许的主要业务载体：

| 层级 | 入参 | 出参 |
|------|------|------|
| **Interface**（HTTP handler） | `gin.Context` + 组装好的 **Request DTO** | **Response DTO** 或统一 `APIResponse`；**禁止**把 Entity/PO 直接作 API 出参 |
| **Application**（`*AppService`） | **DTO**（及租户/会话等横切上下文） | **DTO** 或标量/枚举；**禁止**返回 Entity/PO |
| **Domain**（`*DomainService`） | **Entity**、**ValueObject**、领域枚举 | **Entity**、**ValueObject**；列表用 `[]Entity`；**禁止**以 PO、Application DTO 作签名 |
| **Repository**（domain 接口 / infra 实现） | 写：**Entity**；读：主键/业务键 **标量** 或查询 **VO** | domain 接口返回 **Entity**/`[]Entity`；**PO 仅在 infrastructure 仓储实现内部存在**，经 PO↔Entity 转换后不外泄 |

- **Assembler** 负责 `DTO ↔ Entity`；**仓储实现** 负责 `PO ↔ Entity`（PO 定义在 `infrastructure/persistence`，不进 domain）
- **禁止**: AppService 向 Domain 传裸 `map`/`[]any`；Domain 依赖 Application 的 DTO；API 跳过 AppService 直调 Domain；domain 层出现 PO 类型。

## 7. 批量与性能

- **禁止**在 `for` 循环里逐条查/写数据库（避免 N+1）；优先**批量接口**（如 `FindByCodes`、`BatchInsert`）
- 非批量不可行时，注释说明原因并控制范围

## 8. 职责与异常

- **领域层**: 核心业务规则；返回领域错误；不写控制器式逻辑
- **应用层**: 用例编排、事务边界；不堆领域规则
- **接口层**: 入参出参与 HTTP；不写业务规则
- **错误**: 分层处理；接口层用 `internal/pkg/errors` 的错误码统一响应形态（`interfaces/dto.APIResponse`）

## 9. 最小实现原则

只写**当前需求**需要的方法与成员；不自动生成完整 CRUD、未要求的查询、成组 getter/setter。

## 10. 测试

**本项目不编写测试文件。** 功能正确性通过运行服务 + 手动验证（`curl` / 接口调用）确认。提交时不要附带 `*_test.go`。

## 11. 配置与安全

- 敏感配置（密码、密钥）只放 `.env`（已 gitignore，不入库）；`config.yaml` 用 `${VAR:-默认值}` 占位符引用环境变量
- `.env.example` 为模板，提交入库；`.env` 为本地真实值，禁止提交
- 外部输入须校验；防注入与越权

## 12. 项目边界（重要）

- **`magic` 项目（`/Users/macbook/Developer/Work/magic`）为只读参考**，只能阅读其代码作为迁移参考，**严禁修改其中任何文件**
- **`magic-go` 与 `magic` 的 git 完全独立**，互不关联、互不提交
- 提交前确保 `go build ./...`（在 `magic-service/` 下）通过

## 13. 新功能落地顺序

1. `domain`（实体/值对象/仓储接口/领域服务）
2. `application`（DTO、Assembler、AppService）
3. `infrastructure` / 持久化实现（若涉及）
4. `interfaces`（handler、路由、响应）

---

# Git 提交规范

## 格式

```
<类型>(<范围>): <主题>

<正文，可选>

<Footer，可选：Closes #123>
```

- **类型**（小写英文）: `feat` `fix` `refactor` `perf` `style` `docs` `chore` `build` `ci`
- **范围**: 模块或层级的中文简述，如 `聊天`、`Domain/代理`、`Interface/API`、`配置`
- **主题**: 中文、祈使语气（如「增加」「修复」）、≤50 字、**句末不加句号**

## 语言

- **范围**、**主题**、**正文**、Footer 说明性文字**必须中文**
- **类型**用约定英文小写词（`feat`/`fix` 等）；Footer 的 `Closes`/`Fixes`/`#编号` 按既有写法保留

## 正文与关联

- 与主题空一行；说明**动机、做法、影响**；单行 ≤72 字
- 需要时末尾：`Closes #n` / `Fixes #n` / `Related to #n`

## 禁止出现的内容

提交标题或正文（含注释、尾注）**不得**包含：

- 任何标明 AI / 自动化工具代写的标签或文案（如「Made with …」「Generated by …」）
- 以工具或模型名为作者的 `Co-authored-by` 等**元数据**

提交信息须像**人类开发者**书写，简洁可审计。**即：不要加 `Co-Authored-By: Claude` 之类署名。**

## 习惯

- 一提交一意图，避免巨型混杂提交
- 提交前自测通过，不带调试代码与无关临时文件
- 遵循 `.gitignore`：`.env`、编译产物、IDE 配置不提交

## 示例

```
feat(聊天): 支持会话创建

应用层组装 DTO，领域层定义会话实体与仓储契约。

Closes #12
```

```
fix(配置): 修复 .env 占位符未展开

config.yaml 改用标准 ${VAR:-默认值} 语法。
```
