# embi

### 监控系统
  * 用于监控系统的访问情况

### 技术:
	* 使用mysql存储，server启动自动建表，不需要制定内部结构，即可实现任意平台的web服务监控，目前不支持定制化的日志监控
### 使用:
	* 启动服务器: ./embi
	* 启动client: 客户端使用直接初始化 github.com/et-zone/embi/client 包,调用即可将数据存储到数据仓库
	* 可视化: http://ip:6661  直接访问即可，网页比较粗糙，将就看吧
	
#### 有时间的时候，后期会做优化，