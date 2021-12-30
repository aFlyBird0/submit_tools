package main

import (
	"fmt"
	"github.com/aFlyBird0/submit_tools/desktop/core"
)

func main() {
	// 记得自己建文件夹
	// 人员信息根目录在 /desktop/static/list
	personInfoFilename := "3班名单.csv"
	// 提交信息根目录在 /desktop/static/submission
	submissionFilename := "3班婚育证明统计.txt"

	//fmt.Println("下面这些人快上传帅照哦")
	//fmt.Println("还没实名或还没进群的小猪都在这里啦！")
	//fmt.Println("下面这些人还没投票")
	fmt.Println("未交婚姻表或未填写特殊情况表的人统计如下")

	// 根据人员信息与提交信息获取所有人员信息
	persons, err := core.GetNotSubmitPersons(personInfoFilename, submissionFilename)
	if err != nil {
		panic(err)
	}
	for _, person := range persons {
		// 输出未提交人的名字
		fmt.Println(person.Name)

		// 可以根据实际别名情况输出额外信息，如第一个别名是 QQ
		//fmt.Printf("名字：%s\tQQ:%v\n", person.Name, person.Alias[0])
	}

}
