# 分布式系统全局唯一ID生成器 
### 全局唯一ID使用场景

分布式系统设计时，数据分片场景下，通常需要一个全局唯一id；

在消息系统中需要消息唯一ID标识来防止消息重复；

多系统打通需要一个全局唯一标识 （如集团各业务线面对不同用户，需要一个全局用户id）。

如何生成一个全局唯一id？或者说设计一个ID发号器呢？

常用如下几种方式：


### 1、UUID

Universally Unique Identifier 是自由软件基金会组织制定的唯一辨识ID生成标准，大多数系统已实现，如微软的GUID实现。

生成格式如：3d422567-f034-4ab4-b98f-a34fd263d0de

### 2、sequence表 

使用DB统一维护一张（N张）发号表， 使用主键自增值生成唯一ID。

生成格式如：1,2,3,4,5....（递增数字）

### 3、SnowFlake 雪花算法

Twitter实现的算法，使用时间戳+机器分配标识+自增序列组成64位数字ID。

生成格式如：1292755860950487050


### 阅读全文链接
[分布式系统全局唯一ID生成器](https://mp.weixin.qq.com/s?__biz=MzIyMzMxNjYwNw==&mid=2247483661&idx=1&sn=9cb88ab5b322efaf7b6b63f7d2cc836c&chksm=e8215e1ddf56d70b1c83858a91fd404046038590b868490bc9eddeaab58cbbcac2b95c5a4ecd&token=1061787983&lang=zh_CN#rd)

扫码关注微信订阅号支持：

![技术岁月](https://raw.githubusercontent.com/skyhackvip/ratelimit/master/techyears.jpg)
