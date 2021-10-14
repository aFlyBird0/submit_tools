package core

import (
	"strings"
)

type personInfos struct {
	Filename string
	Persons  []Person
}

type Person struct {
	Name   string   //标准名
	Alias  []string //别名，可以是邮箱、QQ号等
	Submit bool     // 是否已提交
}
type submitted string

func InitPersonInfo(filename string) (p personInfos, err error) {
	var persons []Person
	if persons, err = GetPersonsFromCsv(filename); err != nil {
		return personInfos{}, err
	}
	return personInfos{Filename: filename, Persons: persons}, nil
}

func (p *personInfos) GetPersonNotSubmit(submission string) {
	for index, person := range p.Persons {
		if strings.Contains(submission, person.Name) {
			person.Submit = true
			p.Persons[index] = person
			//fmt.Printf("找到了!%s\n", person.Name)
			continue
		}
		for _, alia := range person.Alias{
			if alia != "" && strings.Contains(submission, alia) {
				person.Submit = true
				p.Persons[index] = person
				//fmt.Printf("找到了!%s\n", person.Name)
				break
			}
		}
	}
}

func (p personInfos) Statistic(submission string) {
	p.GetPersonNotSubmit(submission)
}

func (p *Person) MarkSubmit() {
	p.Submit = true
}
