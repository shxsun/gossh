     xxxxxxxxxxx                  x
          x                       x
          x        xxxx      xxxx x     xxxx
          x       x    x    x     x    x    x
          x       x    x    x     x    x    x
          x       x    x    x     x    x    x
          x        xxxx      xxxxxx     xxxx


	(###) indicates no people doing it.

	- ###:
		gossh配置文件的读取，默认的读取顺序($PWD/gossh.ini, ~/.gossh.ini, /etc/gossh.ini)
	- ###:
		gossh的命令行参数的解析
	- ###:
		gossh首次使用时的用户登录，获取密码（要求不能显示出来）
	- ###:
		gossh的自动安装包(jumbo)

	- ###:
		server的登录界面
	- skyblue:		2013-11-12 ~ 2013-11-12 finished
		server的数据存储设计(要求加密存储，数据即使恢复不了了，也不能泄漏）
		最好能包括使用量统计
	- skyblue:		2013-11-13 coding
		server的数据存取函数
	- ###
		server的数据管理页面
		* server的数据添加
		* server的数据删除
		* server的数据更新
	- skyblue:		2013-11-13 update
		server的api接口说明文档(提供给client使用）

### Usefull sources
#### server
* mysql [xorm](https://github.com/lunny/xorm) or redis [godis](https://github.com/comatosekid/godis)
* json rest api [jas](https://github.com/coocood/jas)

#### client
* command usage help: [go-flags](https://github.com/jessevdk/go-flags)
* get password [gopass](https://github.com/howeyc/gopass)
* [gcfg](https://code.google.com/p/gcfg)

