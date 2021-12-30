package test

import (
	"fmt"
	"github.com/aFlyBird0/submit_tools/desktop/core"
	"strings"
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
