package example

import (
	"fmt"
	"submit_tools/core"
	"testing"
)

func TestAsMain(t *testing.T) {
	// 加载待提交列表
	csvLoader := core.AliasLoaderFromCsv{Filename: "./secret/总名单.csv"}
	// 去除标题行
	removeHandler := core.RemoveHeaderHandler{EvictedLines: []int{0}}
	// 第4列不参与匹配，展示结果采用第0和第1列别名
	tagger := core.DefaultTagger{Evicted: []int{4}, Show: []int{0, 1, 5}}

	// 提交信息读取构造器
	g := core.NewToSubmitGenerator(&csvLoader, &tagger, &removeHandler)

	// 加载已提交文本
	submission := &core.SubmissionLoaderSingleFromFilename{Filename: "./secret/已提交.txt"}

	// 总管
	manager := core.NewEmptyManager()
	manager.SetToSubmit(g)
	manager.SetSubmission(submission)
	manager.SetMatch(core.DefaultMatch)
	unSubmitted, confused, errs := manager.GetUnSubmitted()
	if len(errs) != 0 {
		t.Errorf("errs: %v", errs)
	}
	fmt.Println("未提交人员名单:")
	if len(unSubmitted) == 0 {
		fmt.Println("所有人都提交啦！")
	} else {
		for _, un := range unSubmitted {
			fmt.Println(un.String())
		}
	}
	if len(confused) > 0 {
		fmt.Println("以下人员提交情况不确定：")
		for _, c := range confused {
			fmt.Println(c.String())
		}
	}
}
