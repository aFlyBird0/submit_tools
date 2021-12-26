package core

type Submission struct {
	Sub string //某人或所有人提交的一整个字符串
}

type SubmissionEachPerson interface {
	AliasLoader() []string
}
