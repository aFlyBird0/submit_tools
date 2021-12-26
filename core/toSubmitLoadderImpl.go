package core

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type RemoveHeaderHandler struct {
	EvictedLines []int
}

func (r *RemoveHeaderHandler) Handle(aliases []Alias) (res []Alias, err error) {
	res = make([]Alias, 0, len(aliases))
	for i, alias := range aliases {
		// in 函数天生忽略无效字段
		if !in(i, r.EvictedLines) {
			res = append(res, alias)
		}
	}
	return
}

type AliasLoaderFromAliasSlice struct {
	AliasSlice [][]string
}

func (loader *AliasLoaderFromAliasSlice) LoadAlias() (aliases []Alias, err error) {
	for _, alias := range loader.AliasSlice {
		if len(alias) != 0 {
			aliases = append(aliases, alias)
		}
	}
	return
}

type AliasLoaderFromStringSlice struct {
	Submits []string
	AliaSep string // 别名分隔符
}

func (loader *AliasLoaderFromStringSlice) LoadAlias() (aliases []Alias, err error) {
	for _, submit := range loader.Submits {
		aliases = append(aliases, strings.Split(submit, loader.AliaSep))
	}
	return
}

type AliasLoaderFromString struct {
	Submits string
	LineSep string // 换行符
	AliaSep string // 别名分隔符
}

func (loader AliasLoaderFromString) LoadAlias() ([]Alias, error) {
	lines := strings.Split(loader.Submits, loader.LineSep)
	aliases := make([]Alias, 0, len(lines))
	for _, line := range lines {
		aliases = append(aliases, strings.Split(line, loader.AliaSep))
	}
	return aliases, nil
}

type AliasLoaderFromIOReader struct {
	Reader  io.Reader
	LineSep string // 换行符
	AliaSep string // 别名分隔符
}

func (loader *AliasLoaderFromIOReader) LoadAlias() (aliases []Alias, err error) {
	toSubmitBytes, err := ioutil.ReadAll(loader.Reader)
	// 将全部提交信息按人分割
	submits := strings.Split(string(toSubmitBytes), loader.LineSep)
	// 将每个人的信息，分割成一个个的别名
	for _, submit := range submits {
		aliases = append(aliases, strings.Split(submit, loader.AliaSep))
	}
	return
}

type AliasLoaderFromCsv struct {
	Filename string
	AliaSep  string
}

func (loader *AliasLoaderFromCsv) LoadAlias() ([]Alias, error) {
	fs, err := os.Open(loader.Filename)
	if err != nil {
		return nil, fmt.Errorf("can not read, err is %+v", err)
	}
	defer fs.Close()

	r := csv.NewReader(fs)
	// csv文件中的分隔符 和 同一行的别名分隔符的第一个元素视为等同
	if loader.AliaSep == "" {
		loader.AliaSep = ","
	}
	r.Comma = rune(loader.AliaSep[0])
	content, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("can not readall, err is %+v", err)
	}
	loaderFromAliasSlice := AliasLoaderFromAliasSlice{AliasSlice: content}
	return loaderFromAliasSlice.LoadAlias()
}

type AliasLoaderFromFilename struct {
	Filename string
	AliaSep  string
}

func (loader *AliasLoaderFromFilename) LoadAlias() ([]Alias, error) {
	if reader, err := os.Open(loader.Filename); err == nil {
		loaderFromIOReader := AliasLoaderFromIOReader{Reader: reader, AliaSep: loader.AliaSep}
		return loaderFromIOReader.LoadAlias()
	} else {
		return nil, err
	}
}
