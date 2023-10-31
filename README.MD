## 技术栈

- 语言：Golang、Typescript；
- 后端：GoFrame、Casbin；
- 前端：Vue3、Vite、Element-Plus、TypeScript、Pinia 等；

- 管理员：`admin`，密码：`qq125346 or devops.zze`；
- 测试账号：`test`，密码：`qq125346 or devops.zze`；


## 项目运行
1、在 MySQL 中执行 `manifest/db/mysql.sql` 创建好数据库以及初始化数据；

2、然后修改 `manifest/config/config.yaml` 中的数据库连接地址（`database.default.link`），格式如下：
```sql
mysql:<用户名>:<密码>@tcp(<数据库地址>)/<库名>?loc=Local&parseTime=true
-- 例：mysql:root:root@tcp(127.0.0.1:3306)/DatabaseName?loc=Local&parseTime=true
```
3、直接运行项目根目录的 `main.go` 就可以跑起来啦~