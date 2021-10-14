package main

import (
	"fmt"
	"io/ioutil"
	"submit_tools/desktop/core"
)

func main()  {
	personInfoRootDir := "./desktop/static/list/"
	submissionRootDir := "./desktop/static/submission/"
	personToSubmit, err := core.InitPersonInfo(personInfoRootDir + "后端ID表-全体.csv")
	if err != nil {
		fmt.Println(err)
	}
	submission, err :=  ioutil.ReadFile(submissionRootDir + "上交了照片的名单.txt")
	if err != nil {
		panic("文件打开错误" + err.Error())
	}
	personToSubmit.Statistic(string(submission))
	fmt.Println("下面这些人快上传帅照哦")
	//fmt.Println("还没实名或还没进群的小猪都在这里啦！")
	for _, person := range personToSubmit.Persons {
		if !person.Submit {
			fmt.Printf("名字：%s\tQQ:%v\n", person.Name, person.Alias[0])
			//fmt.Println(person.Name)
		}
	}

}