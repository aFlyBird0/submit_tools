package core

import (
	"errors"
	"fmt"
	"strings"
)

type Alias []string

type ToSubmit struct {
	Alias   Alias // 能标示该人员的所有别名字段
	Evicted []int // 舍弃的别名下标
	Show    []int // 展示时使用的别名下标
}

func (to *ToSubmit) Len() int {
	return len(to.Alias)
}

// EffectiveAlias 返回有效的字段
func (to *ToSubmit) EffectiveAlias() (effective []string) {
	effective = make([]string, 0, len(to.Alias)-len(to.Evicted))
	for i, alia := range to.Alias {
		if !in(i, to.Evicted) {
			effective = append(effective, alia)
		}
	}
	return
}

func in(index int, arr []int) bool {
	for _, i := range arr {
		if i == index {
			return true
		}
	}
	return false
}

// ShowedAlias 返回展示时使用的字段
func (to *ToSubmit) ShowedAlias() (showed []string) {
	showed = make([]string, 0, len(to.Show))
	for _, showIndex := range to.Show {
		showed = append(showed, to.Alias[showIndex])
	}
	return
}

// String 展示出待提交人员的有效字段
func (to *ToSubmit) String() string {
	// show有定义，只展示show中的别名
	if len(to.Show) > 0 {
		return strings.Join(to.ShowedAlias(), "\t")
	} else {
		// show 未定义，展示未被舍弃的所有别名
		return strings.Join(to.EffectiveAlias(), "\t")
	}
}

// AliasLoader 获取 n 个人的别名数组
type AliasLoader interface {
	LoadAlias() ([]Alias, error)
}

type AliasesHandler interface {
	Handle(aliases []Alias) ([]Alias, error)
}

type Tagger interface {
	TagEvicted(tos []*ToSubmit) error
	TagShow(tos []*ToSubmit) error
}

type DefaultTagger struct {
	Evicted []int
	Show    []int
}

// TagEvicted 标记无用字段，假设toSubmit字段数量与格式一致
func (d *DefaultTagger) TagEvicted(tos []*ToSubmit) error {
	if len(tos) == 0 {
		return nil
	}
	for _, index := range d.Evicted {
		if index < 0 || index > len(tos[0].Alias) {
			return errors.New("index out of range")
		}
	}
	for _, to := range tos {
		to.Evicted = d.Evicted
	}
	return nil
}

// TagShow 标记需要显示的字段，假设toSubmit字段数量与格式一致
func (d *DefaultTagger) TagShow(tos []*ToSubmit) error {
	if len(tos) == 0 {
		return nil
	}
	for _, index := range d.Show {
		if index < 0 || index > len(tos[0].Alias) {
			return errors.New("index out of range")
		}
	}
	for _, to := range tos {
		to.Show = d.Show
	}
	return nil
}

type toSubmitGenerator struct {
	AliasLoader
	Tagger
	handlers []AliasesHandler
}

func (g *toSubmitGenerator) LoadToSubmit() (tos []*ToSubmit, err error) {
	aliases, err := g.LoadAlias()
	for _, handler := range g.handlers {
		if aliases, err = handler.Handle(aliases); err != nil {
			return nil, errors.New("alias handler error")
		}
	}
	if err != nil {
		return nil, err
	}
	for _, alias := range aliases {
		tos = append(tos, &ToSubmit{
			Alias: alias,
		})
	}

	if err := g.TagShow(tos); err != nil {
		fmt.Println(err)
		return nil, err
	}
	// 暂时忽略无效字段抛弃错误，错误了就不抛弃就是，效果一样
	if err := g.TagEvicted(tos); err != nil {
		fmt.Println(err)
	}
	return
}

func NewToSubmitGenerator(loader AliasLoader, tagger Tagger, handlers ...AliasesHandler) *toSubmitGenerator {
	g := toSubmitGenerator{}
	g.AliasLoader = loader
	g.Tagger = tagger
	g.handlers = handlers
	return &g
}
