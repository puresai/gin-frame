## 说明

gin的基础框架


## 部署与管理

- 本地编译

```
# 进入项目目录，编译linux平台代码
./manager build
```

- 配置

```
# supervisor配置文件rzc-game.ini

[program:]
autostart=true
startsecs=15
command=./rzc-game -c ./conf/config.yaml
autorestart=true
startretries=3
redirect_stderr=true
stdout_logfile_backups=20
stdout_logfile=/data/logs/rzc-game-supervisor.log
stdout_logfile_maxbytes=10MB
stderr_logfile=/data/logs/rzc-game-supervisor-err.log
stderr_logfile_maxbytes=5MB
stopwaitsecs=5               ; max num secs to wait b4 SIGKILL (default 10)
```

- 启动

> supervisorctl status
> supervisorctl update
> supervisorctl start rzc-game

- 管理

```
# 查看运行状态
./manager status

# 重启
./manager restart

# 启动
./manager start

# 停止
./manager stop
```
