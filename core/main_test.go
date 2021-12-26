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
	fmt.Println("未提交人员名单:")
	for _, un := range m.GetUnSubmitted() {
		fmt.Println(un.String())
	}
}
