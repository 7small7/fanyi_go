# [English Document](./README_EN.md) | [中文文档](./README.md)

`Fanyi_go ` is a translation command tool written in Go language, which aggregates various translation interfaces such as Alibaba, Baidu, Tencent, Spark, Microsoft, and Google. Supports Windows, Mac, Linux systems, and any operating system that supports Go language can be used.

Using `fanyi_go` can easily achieve multi platform translation and achieve out of the box results.

# How to use

## Environmental requirements

`Considering version backward compatibility, fanyi go did not use the high-level syntax introduced in Go language when writing, so you can almost run the project in all versions of Go, but it is recommended to use Go 1.8+version.

## Command usage

When using 'fanyi go', you can obtain help instructions through the following command.
```shell
tutu@tutudeMac-mini  ~/p/w/t/fanyi_go (main)> go run main.go -h
-c string
the content that needs to be translated.
-l string
the language of the content to be translated has optional values of en (English), zh (Chinese), ja (Japanese) (default "en")
-t string
the platform used for translation has optional values of bd (Baidu), tx (Tencent), and al (Alibaba).  (default "baidu")
```
- The parameter 'c' represents the content that needs to be translated,How to wrap the content in single or double quotes if it is too long or contains punctuation marks, such as' - c '' What did you have for dinner tonight? ''.
- The 'l' parameter represents the language in which the content needs to be translated, with optional values of 'en' (English), 'zh' (Chinese), and 'ja' (Japanese). It must follow certain rules, and the content consists of 'current language'+'2'+'target language'. For example, if you want to translate Chinese into English, you can use 'zh2en', where 'zh' represents Chinese and 'en' represents English. The format of each platform may be different, please refer to the platform translation language comparison table at the bottom of the article for details.
- The t parameter represents the platform used, with optional values of bd (Baidu), tx (Tencent), and al (Alibaba).

** The translation result is automatically copied to the clipboard, and you can directly paste the translated content. **


# Effect Preview

```shell
#Translate Chinese into English
Go run main.go-t tx-c "Hello! My name is Tom, may I ask what your name is? "- l zh2en
Hello! My name is Tom. What's your name, please

# Translate English into Chinese
go run main.go -t tx -c "Good evening,everyone" -l en2zh
Good evening,everyone
```

# Translation Comparison Table

1. Tencent Translation Comparison Table. [Document link address](https://cloud.tencent.com/document/product/551/15619)

2. Alibaba Translation Comparison Table. [Document link address](https://help.aliyun.com/zh/machine-translation/support/supported-languages-and-codes?spm=a2c4g.11186623.0.0.39ef4f6ftcTZOH)