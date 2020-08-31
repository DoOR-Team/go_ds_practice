# 数据结构练习

一个文件夹一个题目，文件夹内对应test文件注释里有题目描述

在对应文件夹中添加对应的自己实现的代码，并取名（以bowlingJialou），在test文件中添加Test测试用例

创建一个属于自己的文件，例如bowling_jialou.go，并在文件中添加自己的实现函数：

```go
func bowlingJialou(input string) string {
// 实现自己的处理细节
  return result
}
```

test中添加：

```go
func TestJialou(t *testing.T) {
	if !test(bowlingJialou) {
		t.Error("some case error")
	}
}
```

文件夹内包含测试数据，测试时会告知哪个case有问题，还是通过