## API of server
Base URL = <server_addr:port>/api/v1/

REST接口,返回值全部是是json

	########################################
	- 描述
	- 接口名: 请求方法(POST,GET),不写就是两者都可
		期望参数类型
	- 返回数据
	#########################################

	- 登陆(token过期时间，没想好)
	- login: POST
		username
		password
	- token(string)

	- 查看服务的有那些机器
	- list: 
		token: needed
	- [host1, host2, host3...]

	- 机器名匹配和密码获取
	- match:
		token: needed
		alias: string(eg vps2012)
		user: string(eg root)
	- [{"hostname": fullname, "password": string}, ...]

	- 查看总使用次数
	- count:
		token: needed
	- cnt(int) 使用的次数

