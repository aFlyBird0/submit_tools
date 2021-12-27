package core

import (
	"fmt"
	"testing"
)

func TestFirstVersion(t *testing.T) {
	toLoader := AliasLoaderFromStringSlice{}
	toLoader.Submits = []string{"1 李鹤鹏 abc 123", "2 张三 3w4 123"}
	toLoader.AliaSep = " "

	tagger := DefaultTagger{Show: []int{0, 1}, Evicted: []int{3, 0}}
	g1 := NewToSubmitGenerator(&toLoader, &tagger)

	subsLoaders := SubmissionLoaderFromStringSlice{Submissions: []string{"fdjf李鹤鹏dfs123"}}

	m := NewManager([]ToSubmitLoader{g1}, []SubmissionLoader{&subsLoaders}, DefaultMatch)
	unSubmitted, confused, errs := m.GetUnSubmitted()
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
