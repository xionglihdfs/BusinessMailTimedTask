## 定时任务系统
#### 场景:

##### 1.根据MySQL查询语句, 查询结果集之后生成csv文件和Excel文件, 然后发送邮件;
##### 2.定时执行上述任务, 发送给需要以邮件方式通知的业务方; 
##### 3.大批量数据暂不支持, 根据smtp附件大小要求, 如果有限制, 后续添加压缩功能, 提高发送效率.  
​	
#### 使用说明
##### 1.修改配置文件 config/config.properties 相关配置项, 需要修改MySQL和smtp地址及用户;
##### 2.将SQL文件放入sql/, 支持中文, 这个SQL是需要发送的数据, 并且将SQL结果集交于业务方确认之后; 
##### 3.编写main.go文件, 添加邮件主题和发送邮件组信息, 然后编写定时任务代码;
##### 4.执行windows部署文件deploy.bat, 在项目目录下生产main文件;
##### 5.将main文件传到Linux下, 依赖的文件夹, config/, result/, sql/ 一并放在main目录下, 这里需要编辑 config/config.properties 为真实数据库配置, smtp也能使用即可, 授予main执行权限, 然后 ./main 运行项目.

##### 目的: 为DBA开发一套能定时发送报表邮件或者监控数据的系统, 后期还会不断扩充功能, 有问题请加微信沟通, 或者直接留言.

![image](https://github.com/xionglihdfs/BusinessMailTimedTask/blob/master/doc/%E7%86%8A%E7%86%8A.png)

