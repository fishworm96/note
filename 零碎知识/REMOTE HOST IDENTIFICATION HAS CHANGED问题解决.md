# 背景
ssh连接至[云服务器](https://cloud.tencent.com/product/cvm?from=10680)时，提示以下错误： 
![](https://cdn.nlark.com/yuque/0/2023/png/12511308/1675583640362-393f9373-ddf1-4d6a-9681-47920fb3f960.png#averageHue=%23161616&clientId=u1ef354d8-8d74-4&from=paste&id=u911d4978&originHeight=595&originWidth=1620&originalType=url&ratio=1&rotation=0&showTitle=false&status=done&style=none&taskId=u4ce29096-5f26-462a-86bb-ef25069d929&title=)
原因是第一次使用SSH连接时，会生成一个认证，储存在客户端的known_hosts中。
可使用以下指令查看：

- ssh-keygen -l -f ~/.ssh/known_hosts

由于[服务器](https://cloud.tencent.com/product/cvm?from=10680)重新安装系统了，所以会出现以上错误。
# 解决方案
ssh-keygen -R 服务器端的ip地址 
![](https://cdn.nlark.com/yuque/0/2023/png/12511308/1675583640359-95d62869-115c-477b-b05f-aea83dbc638d.png#averageHue=%23121212&clientId=u1ef354d8-8d74-4&from=paste&id=u4e59ca0e&originHeight=215&originWidth=1419&originalType=url&ratio=1&rotation=0&showTitle=false&status=done&style=none&taskId=ua4beb58d-07f6-4f37-8e0f-7c3674ab33f&title=)
重新连线,出现以下提示： 
![](https://cdn.nlark.com/yuque/0/2023/png/12511308/1675583640424-91fb78a0-3b91-4ca7-acf6-7cc854743221.png#averageHue=%23161616&clientId=u1ef354d8-8d74-4&from=paste&id=uf1872552&originHeight=266&originWidth=1620&originalType=url&ratio=1&rotation=0&showTitle=false&status=done&style=none&taskId=u18a7af9a-6eb6-4c6a-b919-e2cce25e2ea&title=)
输入yes确认即可连线成功。
