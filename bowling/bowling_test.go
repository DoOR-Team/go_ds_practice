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
打保龄球

源程序名 　　　bowling.??? (pas,c,cpp)
可执行文件名   bowling.exe
输入文件名　　 bowling.in
输出文件名     bowling.out

    打保龄球是用一个滚球去打击十个站立的柱，将柱击倒。一局分十轮，每轮可滚球一次或多次，以击倒的柱数为依据计分。一局得分为十轮得分之和，而每轮的得分不仅与本轮滚球情况有关，还可能与后续一两轮的滚球情况有关。即某轮某次滚球击倒的柱数不仅要计入本轮得分，还可能会计入前一两轮得分。具体的滚球击柱规则和计分方法如下：
   （1）若某一轮的第一次滚球就击倒全部十个柱，则本轮不再滚球（若是第十轮则还需
另加两次滚球，不妨称其为第十一轮和第十二轮，并不是所有的情况都需要滚第十一轮和第十二轮球）。该轮得分为本次击倒柱数10与以后两次滚球所击倒柱数之和。
   （2）若某一轮的第一次滚球未击倒十个柱，则可对剩下未倒的柱再滚球一次。如果这
两次滚球击倒全部十个柱，则本轮不再滚球（若是第十轮则还需另加一次滚球），该轮得分
为这两次共击倒柱数10与以后一次滚球所击倒柱数之和。
   （3）若某一轮两次滚球未击倒全部十个柱，则本轮不再继续滚球，该轮得分为这两次
滚球击倒的柱数之和。
    总之，若—轮中一次滚球或两次滚球击倒十个柱，则本轮得分是本轮首次滚球开始的
连续三次滚球击倒柱数之和（其中有一次或两次不是本轮滚球）。若一轮内二次滚球击倒柱
数不足十个，则本轮得分即为这两次击倒柱数之和。下面以实例说明如下(字符“/”表示击倒当前球道上的全部的柱)：
    轮          1    2   3    4    5    6    7    8    9   10   11   12
    击球情况    /    /    /    72   9/   81   8/    /    9/    /    8/
    各轮得分   30   27   19   9   18    9   20   20   20   20
累计总分   30   57  76   85  103   112  132  152  172  192
现在请你编写一个保龄球实时计分程序，用来计算和显示某轮结束后的得分情况。若某轮的得分暂时无法算出，则该轮得分不显示。

输入：
输入数据用文件bowling.in，文件内容仅有一行，为前若干轮滚球的情况，每轮滚球用一到两个字符表示，每一个字符表示一次击球，字符“/”表示击倒当前球道上的全部的柱，否则用一个数字字符表示本次滚球击倒的当前球道上的柱的数目，两轮滚球之间用一个空格字符隔开。
如上例对应的输入文件内容为：/  /  /  72  9/  81  8/  /  9/  /  8/

输出：
输出到文件bowling.out，共两行，第一行为每轮得分，第二行为到当前轮为止的总得分。每个得分之间用一个空格隔开。

样例输入：
     /    /    /   72    9/   81   8/    /    9/    /   8/

样例输出：
    30   27  19    9   18    9   20   20   20   20
30   57  76   85  103  112  132  152  172  192
 */

func test(bowlingFunction func(string) string) bool {
	ret := true
	dir := "./"

	for i:=1; i<=6; i++ {
		fmt.Printf("case %d\n", i)
		inputfile := fmt.Sprintf("%sbowling/bowling%d.in", dir, i)
		answerfile := fmt.Sprintf("%sbowling/bowling%d.out", dir, i)
		inf, _ := os.Open(inputfile)
		ansf, _ := os.Open(answerfile)
		buf, _ := ioutil.ReadAll(inf)
		myAnswer := bowlingFunction(bytes.NewBuffer(buf).String())
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
	if !test(bowlingJialou) {
		t.Error("some case error")
	}
}