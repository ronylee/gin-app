# gin-app
基于 gin 的 api 框架应用

### 编译

cd build && ./debug.sh

### 获取配置

fmt.Print(config.Bases.Mysql.Host)
fmt.Print(config.Get("app.debug"))


