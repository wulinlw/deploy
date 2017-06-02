GRPC实现的服务器管理工具, 需与web页面配合使用, web端是内部项目, 还不能公开

主要功能：
- 命令执行
- 文件传输
- 存活检测

服务端口
server: 50051
proxy: 80,8000

部署:
gbopsProxy web请求代理，只需部署一个
gbopsServer 服务端，部署在需要管理的服务器上

API：
#### 命令执行
url: http://127.0.0.1/index
参数：
                    
参数名  | 说明
------------- | -------------
apiName  | 默认值 ComplexCommand , 接口名
Dir | 默认值 / , 执行路径
Command |  执行的命令
ip |  目标服务器ip:port
uniqueId |  //随机字符串
                    

命令执行后，除了直接返回结果，也会通过websocket返回，在管理页面可以方便的显示，适合长时命令

#### 文件传输
url: http://127.0.0.1/index
参数：
                    
参数名  | 说明
------------- | -------------
apiName  | 默认值 SendFile ,接口名
ip  |  目标服务器ip:port
RelativePath  |  上传的相对路径，相对于/usr/local/gbops
FileContent |  文件内容
                     
仅用于脚本类型，未测试二进制文件

#### 存活检测
url: http://127.0.0.1/live
参数：
无参数，若服务器存活则返回"ok"

调用方式：
```php
$params = [
	'apiName' => 'ComplexCommand',
	'Dir' => '/',
	'Command' => $command,
	'ip' => $ip.':50051',
	'uniqueId' => $uniqueId
];
$curl = new Curl();
$body = $curl->post('http://127.0.0.1/index', $params);

