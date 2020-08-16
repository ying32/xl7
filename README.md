# xl7
xl7 sdk for Go， 迅雷7 Go语言SDK

[迅雷7官方API文档](http://xldoc.xl7.xunlei.com/0000000026/index.html) 

对比原[迅雷下载引擎](https://github.com/ying32/xldl) ，这个版本的引擎文件更少，更小，是之前的一个简化版本，只支持HTTP协议，ThunderPlatform_SDK.zip包含官方的完整SDK和使用，传在这里是因为官方网站已经无法下载了。

这个版本的还会出现托盘图标。。。

因迅雷官方已经关闭了这个sdk，所以出现无法使用，具体原因是无法访问`client.stat.xunlei.com`这个网站，所以崩溃了，在hosts中将这个指向`127.0.0.1`即可正常使用。