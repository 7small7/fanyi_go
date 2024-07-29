# [中文文档](./README.md) | [English Document](./README_EN.md)

`fanyi_go`是一款利用Go语言编写的，聚合类的翻译命令工具，集成阿里、百度、腾讯、星火、微软、谷歌等多家翻译接口。支持Windows系统、Mac系统、Linux系统，只要支持Go语言的操作系统都可以使用。

使用`fanyi_go`可以实现轻松实现多平台翻译，实现开箱即用的效果。

# 如何使用

## 环境要求

`fanyi-go`考虑到版本向下兼容，在编写时并未使用引入Go语言的高版本语法，因此你几乎是可以在Go的所有版本中运行该项目，不过推荐使用在Go 1.8+的版本。

## 命令使用

在使用`fanyi-go`时，你可以通过如下命令，获取到帮助说明。
```shell
tutu@tutudeMac-mini ~/p/w/t/fanyi_go (main)> go run main.go -h
-c string
    the content that needs to be translated.
-l string
    the language of the content to be translated has optional values of en (English), zh (Chinese), ja (Japanese) (default "en")
-t string
    the platform used for translation has optional values of bd (Baidu), tx (Tencent), and al (Alibaba). (default "baidu")
```
- c 参数表示需要翻译的内容
- l 参数表示需要翻译内容的语言，可选值为 en (English), zh (Chinese), ja (Japanese)，需要遵循一定的规则，内容由`当前语言`+ `2` + `目标语言`组成，例如你要将中文翻译成英文，你可以使用`zh2en`，其中`zh`表示中文，`en`表示英文。每一个平台的格式可能不太一样，具体可以参考文章底部的平台翻译语言对照表。
- t 参数表示使用的平台，可选值为 bd (Baidu), tx (Tencent), and al (Alibaba)。


# 翻译对照表

1、腾讯翻译对照表。[文档链接地址](https://cloud.tencent.com/document/product/551/15619)

2、阿里翻译对照表。[文档链接地址](https://help.aliyun.com/zh/machine-translation/support/supported-languages-and-codes?spm=a2c4g.11186623.0.0.39ef4f6ftcTZOH)