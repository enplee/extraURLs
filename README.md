# extraURLs
<a title="Supported Platforms" target="_blank" href="https://github.com/panjf2000/gnet"><img src="https://img.shields.io/badge/platform-Linux%20%7C%20macOS%20%7C%20Windows-549688?style=flat-square&logo=appveyor"></a>
### 使用Go实现基于RegExp的文本RRL提取，参考xurl。
# 🚀 使用范例
```go
import "github.com/enplee/extraURLs/"

func main() {
	rxRelaxed := extraURLs.Relaxed()
	rxRelaxed.FindString("Do gophers live in golang.org?")  // "golang.org"
	rxRelaxed.FindString("This string does not have a URL") // ""

	rxStrict := extraURLs.Strict()
	rxStrict.FindAllString("must have scheme: http://foo.com/.", -1) // []string{"http://foo.com/"}
	rxStrict.FindAllString("no scheme, no match: foo.com", -1)       // []string{}
}
```

### 💡参考
#### 1. 域名命名规范 https://help.aliyun.com/document_detail/54066.html
+ 只能使用英文字母（a-z，不区分大小写）、数字（0~9）以及连接符（-）。不支持使用空格及以下字符：
!?%$等
+ 连接符（-）不能连续出现、不能单独注册，也不能放在开头和结尾。
#### 2. ipV4规范 https://zh.wikipedia.org/wiki/IPv4
+ 1.0.0.0 ~ 255.255.255.255
#### 3. ipV6规范 https://zh.wikipedia.org/wiki/IPv6
+ IPv6二进位制下为128位长度，以16位为一组，每组以冒号“:”隔开，可以分为8组，每组以4位十六进制方式表示
+ 例如：2001:0db8:86a3:08d3:1319:8a2e:0370:7344
+ 注意：IPv6IPv6在某些条件下可以省略:
    * 每项数字前导的0可以省略，省略后前导数字仍是0则继续
    * 可以用双冒号“::”表示一组0或多组连续的0，但只能出现一次
    * 如果这个地址实际上是IPv4的地址，后32位可以用10进制数表示
  
# 🌐 思路
[enplee的博客](https://enplee.github.io)