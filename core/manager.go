package core

import (
	"errors"
	"fmt"
	"strings"
)

type ToSubmitLoader interface {
	LoadToSubmit() ([]*ToSubmit, error)
}

type SubmissionLoader interface {
	LoadSubmission() ([]*Submission, error)
}

type Matcher func(wants []*ToSubmit, givens []*Submission) ([]int, []int)

// DefaultMatch 匹配待提交列表和已提交列表，返回已提交列表的下标数组和可能有歧义的下标数组
var DefaultMatch = func(wants []*ToSubmit, givens []*Submission) (submittedIndex, confusedIndex []int) {
	if len(givens) == 0 {
		return
	}
	submittedIndex = make([]int, 0, len(givens))
	// 如果提交列表只有1，那么说明所有的提交信息全部揉在一起了
	if len(givens) == 1 {
		for i, want := range wants {
			for _, alia := range want.EffectiveAlias() {
				alia := strings.TrimSpace(alia)
				if alia == "" {
					continue
				}
				if strings.Contains(givens[0].Sub, alia) {
					submittedIndex = append(submittedIndex, i)
					break
				}
			}
		}
	}

	// 如果提交列表大于1，那么说明每行是一个提交
	matched := make([]bool, len(givens))
	for toSubmitIndex, want := range wants {
		match := false
		confuse := false
		for _, alia := range want.EffectiveAlias() {
			alia := strings.TrimSpace(alia)
			if alia == "" {
				continue
			}
			for j, giv := range givens {
				if strings.Contains(giv.Sub, alia) {
					if matched[j] {
						confuse = true
					}
					matched[j] = true
					match = true
					submittedIndex = append(submittedIndex, toSubmitIndex)
					break
				}
			}
			if match {
				if confuse {
					confusedIndex = append(confusedIndex, toSubmitIndex)
				}
				break
			}
		}
	}
	return
}

type manager struct {
	ToSubmitLoaders  []ToSubmitLoader
	SubmissionLoader []SubmissionLoader
	Match            Matcher
}

func (m *manager) combineTos() (tos []*ToSubmit, err error) {
	for _, loader := range m.ToSubmitLoaders {
		if to, err := loader.LoadToSubmit(); err == nil {
			tos = append(tos, to...)
		} else {
			return nil, errors.New("Error loading to submit:" + err.Error())
		}
	}
	return
}

func (m *manager) combineSubmissions() (subs []*Submission, err error) {
	for _, loader := range m.SubmissionLoader {
		if sub, err := loader.LoadSubmission(); err == nil {
			subs = append(subs, sub...)
		} else {
			return nil, errors.New("Error loading submission:" + err.Error())
		}
	}
	return
}

func (m *manager) GetUnSubmitted() (unSubmitted []*ToSubmit, confused []*ToSubmit, errs []error) {

	// 获取待提交列表和已提交信息
	tos, err1 := m.combineTos()
	if err1 != nil {
		fmt.Println(err1)
		errs = append(errs, err1)
	}
	sbs, err2 := m.combineSubmissions()
	if err2 != nil {
		fmt.Println(err2)
		errs = append(errs, err2)
	}
	if len(tos) == 0 {
		errs = append(errs, errors.New("no to submit"))
	}
	if len(sbs) == 0 {
		errs = append(errs, errors.New("no submission"))
	}
	if len(errs) > 0 {
		return
	}

	submitted, confusedIndex := m.Match(tos, sbs)
	for i, to := range tos {
		if !in(i, submitted) {
			unSubmitted = append(unSubmitted, to)
		}
	}
	for i, to := range tos {
		if in(i, confusedIndex) {
			confused = append(confused, to)
		}
	}
	return
}

func NewEmptyManager() *manager {
	return &manager{}
}

func NewManager(toLoaders []ToSubmitLoader, subsLoaders []SubmissionLoader, matcher Matcher) *manager {
	return &manager{
		ToSubmitLoaders:  toLoaders,
		SubmissionLoader: subsLoaders,
		Match:            matcher,
	}
}

func (m *manager) SetToSubmit(tos ...ToSubmitLoader) {
	m.ToSubmitLoaders = tos
}

func (m *manager) SetSubmission(subs ...SubmissionLoader) {
	m.SubmissionLoader = subs
}

func (m *manager) SetMatch(matcher Matcher) {
	m.Match = matcher
}
