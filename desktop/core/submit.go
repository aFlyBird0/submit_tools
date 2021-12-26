package core

import (
	"fmt"
	"io/ioutil"
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

func (p *personInfos) StatisticPersonNotSubmit(submission string) {
	for index, person := range p.Persons {
		if strings.Contains(submission, person.Name) {
			person.Submit = true
			p.Persons[index] = person
			//fmt.Printf("找到了!%s\n", person.Name)
			continue
		}
		for _, alia := range person.Alias {
			if alia != "" && strings.Contains(submission, alia) {
				person.Submit = true
				p.Persons[index] = person
				//fmt.Printf("找到了!%s\n", person.Name)
				break
			}
		}
	}
}

const (
	personInfoRootDir = "./desktop/static/list/"
	submissionRootDir = "./desktop/static/submission/"
)

func GetNotSubmitPersons(personInfoFilename, submissionFilename string) (persons []Person, err error) {
	personToSubmit, err := InitPersonInfo(personInfoRootDir + personInfoFilename)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	submission, err := ioutil.ReadFile(submissionRootDir + submissionFilename)
	if err != nil {
		return nil, fmt.Errorf("文件打开错误" + err.Error())
	}
	personToSubmit.StatisticPersonNotSubmit(string(submission))
	for _, person := range personToSubmit.Persons {
		if !person.Submit {
			persons = append(persons, person)
		}
	}
	return persons, nil
}

func (p *Person) MarkSubmit() {
	p.Submit = true
}
