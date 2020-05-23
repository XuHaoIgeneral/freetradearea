# freetradearea
这就是我的大学毕业设计。 

# 背景： 
- 自由选题，docker容器的跨主机调度和使用，即docker集群。

# 设计思路
- master-worker 模式。 一个master控制多个worker，master负责任务分发、节点监控、网络控制(新加子网)、docker启动解析和销毁发起等。
  worker 主要负责docker操作，接收到master请求后执行相应指令；同时也不断向master反馈当前节点cpu、disk、ram情况，方便master打分计算。
