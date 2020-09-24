
私网环境
1.ontology节点（部署相关合约）
- OToken合约
    - 来自跨链合约代码
- FlashPool相关合约（王成）
- 治理相关合约（吕品）

2.后台程序（荣怡可以帮忙）
- wingServer（梦航）
    - 对应有postgreSQL数据库
    - 节点重置需要删除库
- OntBackEnd（李捷）
    - onto扫码
    - 前端数据展示（清算页面数据）
        - 大部分来自wingServer
    - 需要配置comptroller合约
    
3.前端（泽宇）
- 需要Ftoken、Itoken以及otoken合约地址 
- comptroller、wingGov、WingToken地址
- approve合约
- ontBackEnd
- wingServer
- supply 合约（帅帅）


目前私网维护的两套环境（荣怡、李捷目前配合后台服务的更新）
公用AutoTestGo的config_prv3.json配置文件、修改IP地址调节
- 229环境提供给贝宝在联调清算
环境部署配置：
    - wingServer（172.168.3.232机器）
    目录：/data/wingserver
    - ontBackEnd（172.168.3.230机器）
    目录：/data/lijie
    - tools工具（229机器上）
    目录：/data/gopath 
    - ontology程序
    目录：/data/node

- 226环境未正式投入使用
环境部署配置：
    - wingServer（172.168.3.226机器）
    目录：/data/wingserver
    - ontBackEnd（172.168.3.226机器）
    目录：/data/ontbackend
    - tools工具（172.168.3.226机器上）
    目录：/data/tools
     - ontology程序
    目录：/data/node

备注：
- 目前的Chain_ok备份的是已有构造好的清算数据（配置相对完善、其它Chain备份不建议使用）
- 如果需要从干净的环境开始可以使用新部署flash pool然后注册进治理合约
- 给地址分发测试币	
oToken.TransferAllTestToken(cfg, account, sdk, "AbhqQjSw9QHNoZyvnUhzKSXyzQEUTkTtNp")
- 增发测试
oToken.DelegateToProxyAllTestToken(cfg, account, sdk)

- wing-test/test下是治理和配置相关调用
- wing-test/compound/ftoken/下有借贷相关调用

```sh
<!-- 初始价格 -->
./tools oracle put-price --market-name renBTC --amount 1000000000000000
./tools oracle put-price --market-name WBTC --amount 1039000000000000

<!-- 触发清算 -->
./tools oracle put-price --market-name WBTC --amount 1039000000000000
./tools oracle put-price --market-name ETH     --amount 2250000000000000

```
flash pool合约更新操作流程：
- 第一步升级Flash Pool（王成提供合约编译好的wasm.str文件、或协助更新） ：
./tools migrate --update-cfg --comptroller --market
- 第二步记录更新后的comptroller合约以及ftoken合约地址
- 通知前后端更新相关配置（荣怡、李捷、张梦航、泽宇）
- 更新治理合约comptroller的地址（吕品帮忙）


```plantuml
@startuml
actor User  as run #blue
participant webUI
collections Otoken
box "Flash Pool" #LightBlue

collections "Market (ftokens-借贷池) " as mk
collections "Market (Itoken-保险池) " as it

control comptroller as cm
participant oracle
participant approve
participant IntersRate
end box
box "Wing Server" #Light

database postgreSQL
participant wingServer
participant ontBackEnd
participant supply

end box

box "Wing Governance" #Light

control Governance
participant WingToken
participant GlobalParam

end box

autonumber 0 1 "<b>[00]"


@enduml
```

