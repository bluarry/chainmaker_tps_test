# 长安链介绍
“长安链·ChainMaker”具备自主可控、灵活装配、软硬一体、开源开放的突出特点，由北京微芯研究院、清华大学、北京航空航天大学、腾讯和百度等知名高校、企业共同研发。取名“长安链”，喻意“长治久安、再创辉煌、链接世界”。

长安链作为区块链开源底层软件平台，包涵区块链核心框架、丰富的组件库和工具集，致力于为用户高效、精准地解决差异化区块链实现需求，构建高性能、高可信、高安全的新型数字基础设施，同时也是国内首个自主可控区块链软硬件技术体系。

长安链的应用场景，涵盖供应链金融、碳交易、食品追溯等一系列关乎国计民生的重大领域。

长安链拥有高效并行调度算法、高性能可信安全智能合约执行引擎、流水线共识算法等国际领先的区块链底层技术，具备高并发、低延时、大规模节点组网等先进技术优势，交易吞吐能力可达10万TPS，位居全球领先水平；

# TPS验证
长安链的介绍中，提到自己的TPS达到了10w。
此项目则进行性能方面的验证
ps: 由于接触的还不够深，若有错误，请指正<br/>
测试网络为手动生成四节点网络，利用docker单机启动<br/>

# 验证过程
## 2021.07.21 第一次测试
测试方案解释：
通过GoRoutine连接多个ChainClient，多个ChainClient进行并发发起交易,计算这些交易的总用时，最终计算TPS为1,不太合理。<br/>
![TPS 为 1](./images/20210721_001.png)
## 2021.07.21 第二次测试
经过群里大佬提醒，对相关参数进行调整,TPS达到3k多
![TPS 3K](./images/2021-07-21_12-55-03.png)


# 验证结论
目前测试到的最大TPS为3k多,离10w还有很远的距离
