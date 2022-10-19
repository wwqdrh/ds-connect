1、shadow: 开启socks5代理，接收客户端的流量
2、shadow-controller: shadow的管理服务，监听当前有哪些网络，保证shadow能够加入到所有的网络中
3、客户端: socks5代理，本地dns服务，tun设备建立，到集群的流量转发到这个tun设备然后发送到集群服务中
