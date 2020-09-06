# docker consul

## Start consul 
### 获取镜像

```sh 
docker consul
```

### 启动镜像
```sh
docker run -d --name=consul_go -p 8500:8500 \
        consul agent -server -bootstrap -ui -client 0.0.0.0
```

#### 参数说明

- -server 代表以服务端启动
- -bootstrap 指定自己为leader
- -ui 启动内置管理的web界面
- -client 0.0.0.0 指定客户端可以访问的IP地址 

注意
```shell script
sudo iptables -I INPUT -p tcp --dport 8500 -j ACCEPT
```

#### consul_go inspect 

```json
{
 "Ports": {
            "8300/tcp": null,
            "8301/tcp": null,
            "8301/udp": null,
            "8302/tcp": null,
            "8302/udp": null,
            "8500/tcp": [
                {
                    "HostIp": "0.0.0.0",
                    "HostPort": "8500"
                }
            ],
            "8600/tcp": null,
            "8600/udp": null
        },
        "SandboxKey": "/var/run/docker/netns/08b4f4133957",
        "SecondaryIPAddresses": null,
        "SecondaryIPv6Addresses": null,
        "EndpointID": "2bc9a0a3e0cee056c35522c2a94e369b86b58d6b65d835a14c640191d3982aac",
        "Gateway": "172.17.0.1",
        "GlobalIPv6Address": "",
        "GlobalIPv6PrefixLen": 0,
        "IPAddress": "172.17.0.4",
        "IPPrefixLen": 16
}
```


### 本机启动

```sh
consul agent -dev
```

## Register Service With GO

### 

