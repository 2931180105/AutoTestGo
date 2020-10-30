# AutoTestGo
go语言自动化测试（集成性能测试）

* 项目初始化
``sh
go mod init
go build -o main main.go
``


 **ont-test** 
 -  刷块交易(测试利息分配的时候可调用)
 - ontology性能测试

**wing-test**

目前是将项目“wing-contract-tools”中的交易构造方法copy过来进行一下适配的修改，增加测试调用的灵活性

flash pool相关的测试代码在“wing-test/compound”目录下

私网已铺地了部分数据，后续将继续增加铺地数据量：
 -  数据库前3000个地址是Supply，后1千个Borrow
 -  代码测试时可以直接使用Utils.GetAccounts方法取到需要类型的account
 -  每个account保证每种测试币100个左右






**node节点还原（帅帅和王成可以支持）：**

- 节点地址
  - 172.168.3.226   ubuntu/Onchain@2019
  - /data/ontology

- 重置节点:
  -  killall ontology
  -  删除Chain和Log： rm -rf Chain Log
- git log 选择你需还原的分支
  - git checkout {分支 id} 
  - 目前主网运行的合约分支："32eaf8150a81585e5113b624c16a5863eed6cbb8"

- 启动节点： 
  - nohup ./start.sh &



**wingServer维护(张梦航)**

- 服务地址：
  - 172.168.3.226   ubuntu/Onchain@2019
  - /data/wingserver
- config.json文件修改，最新的fToken地址和comptroller地址
  - /data/ontology/config-wingserver.json 
- 重启服务



**李婕的ont-backend**

- 修改配置文件重启
- application-dev.properties



**前端**

- 泽宇的前端，提供一下最新的config.json文件即可
- 前端地址：http://172.168.3.245:31126/




