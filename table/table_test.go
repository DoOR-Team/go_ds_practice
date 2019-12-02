package table

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

/*
【问题背景】国际乒联现在主席沙拉拉自从上任以来就立志于推行一系列改革，以推动乒乓球运动在全球的普及。其中11分制改革引起了很大的争议，有一部分球员因为无法适应新规则只能选择退役。华华就是其中一位，他退役之后走上了乒乓球研究工作，意图弄明白11分制和21分制对选手的不同影响。在开展他的研究之前，他首先需要对他多年比赛的统计数据进行一些分析，所以需要你的帮忙。

【问题描述】华华通过以下方式进行分析，首先将比赛每个球的胜负列成一张表，然后分别计算在11分制和21分制下，双方的比赛结果（截至记录末尾）。
比如现在有这么一份记录，（其中W表示华华获得一分，L表示华华对手获得一分）：
WWWWWWWWWWWWWWWWWWWWWWLW
在11分制下，此时比赛的结果是华华第一局11比0获胜，第二局11比0获胜，正在进行第三局，当前比分1比1。而在21分制下，此时比赛结果是华华第一局21比0获胜，正在进行第二局，比分2比1。如果一局比赛刚开始，则此时比分为0比0。
你的程序就是要对于一系列比赛信息的输入（WL形式），输出正确的结果。

【输入格式】每个输入文件包含若干行字符串（每行至多20个字母），字符串有大写的W、L和E组成。其中E表示比赛信息结束，程序应该忽略E之后的所有内容。

【输出格式】输出由两部分组成，每部分有若干行，每一行对应一局比赛的比分（按比赛信息输入顺序）。其中第一部分是11分制下的结果，第二部分是21分制下的结果，两部分之间由一个空行分隔。

【输入样例】
WWWWWWWWWWWWWWWWWWWW
WWLWE

【输出样例】
11:0
11:0
1:1

21:0
2:1
 */

func test(tableFunction func(string) string) bool {
	ret := true
	dir := "./"

	for i:=0; i<10; i++ {
		fmt.Printf("case %d\n", i)
		inputfile := fmt.Sprintf("%stable/table.in%d", dir, i)
		answerfile := fmt.Sprintf("%stable/table.ou%d", dir, i)
		inf, _ := os.Open(inputfile)
		ansf, _ := os.Open(answerfile)
		buf, _ := ioutil.ReadAll(inf)
		myAnswer := tableFunction(bytes.NewBuffer(buf).String())
		buf, _ = ioutil.ReadAll(ansf)
		answer := strings.Replace(bytes.NewBuffer(buf).String(),"\r\n", "\n",  -1)

		if strings.Compare(answer, myAnswer) == 0 {
			fmt.Println("ok")
		} else {
			fmt.Println("wrong answer")
			fmt.Println("my")
			fmt.Println(myAnswer)
			fmt.Println("answer")
			fmt.Println(answer)
			ret = false
		}
	}
	return ret
}


// 通过如下方式进行测试。
// 复制这个函数，并且吧test(xxx)中的参数改为你的函数，这个函数必须是输入string，输出string
func TestJialou(t *testing.T) {
	if !test(tableJialou) {
		t.Error("some case error")
	}
}