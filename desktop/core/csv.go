package core

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func GetPersonsFromCsv(filename string) (persons []Person, err error) {
	fs, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("can not read, err is %+v", err)
	}
	defer fs.Close()

	r := csv.NewReader(fs)
	content, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("can not readall, err is %+v", err)
	}
	for _, personSlot := range content[1:] {
		//fmt.Printf("%d, ", len(personSlot))
		personName := strings.TrimSpace(personSlot[0])
		if personName == "" {
			return nil, fmt.Errorf("人员信息表错误, 表内某行人员姓名为空")
		}
		//fmt.Println(personSlot)
		alias := make([]string, 0)
		person := Person{Name: personName}
		for _, alia := range personSlot[1:] {
			alias = append(alias, strings.TrimSpace(alia))
		}
		person.Alias = alias
		persons = append(persons, person)
	}
	return persons, nil
}
