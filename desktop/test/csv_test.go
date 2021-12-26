package test

import (
	"fmt"
	"strings"
	"submit_tools/desktop/core"
	"testing"
)

func TestGetPersonsFromCsv(t *testing.T) {
	_, err := core.GetPersonsFromCsv("./static/desktop/static/list/后端ID表-全体.csv")
	if err != nil {
		fmt.Println(err)
	}
}

func TestStringSplit(t *testing.T) {
	str := ",,,,,,,,"
	fmt.Println(len(strings.Split(str, ",")))
}
