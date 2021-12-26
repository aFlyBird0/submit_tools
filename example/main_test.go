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
	if unSubs, err := manager.GetUnSubmitted(); err == nil {
		fmt.Printf("共有%d人未提交\n", len(unSubs))
		for _, unSub := range unSubs {
			fmt.Println(unSub)
		}
	} else {
		t.Error(err)
	}
}
