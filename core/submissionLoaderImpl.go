package core

import (
	"io"
	"io/ioutil"
	"os"
)

// SubmissionLoaderFromStringSlice 每个人的提交是一个字符串
type SubmissionLoaderFromStringSlice struct {
	Submissions []string
}

func (s *SubmissionLoaderFromStringSlice) LoadSubmission() ([]*Submission, error) {
	submissions := make([]*Submission, 0)
	for _, submission := range s.Submissions {
		submissions = append(submissions, &Submission{
			Sub: submission,
		})
	}
	return submissions, nil
}

// SubmissionLoaderFromString 所有人的提交是一个字符串
type SubmissionLoaderFromString struct {
	Submission string
}

func (s *SubmissionLoaderFromString) LoadSubmission() ([]*Submission, error) {
	return []*Submission{
		{
			Sub: s.Submission,
		},
	}, nil
}

type SubmissionLoaderSingleFromIOReader struct {
	Reader io.Reader
}

func (loader *SubmissionLoaderSingleFromIOReader) LoadSubmission() ([]*Submission, error) {
	if toSubmitBytes, err := ioutil.ReadAll(loader.Reader); err != nil {
		return nil, err
	} else {
		submissionSingle := &Submission{
			Sub: string(toSubmitBytes),
		}
		return []*Submission{submissionSingle}, nil
	}
}

type SubmissionLoaderSingleFromFilename struct {
	Filename string
}

func (loader *SubmissionLoaderSingleFromFilename) LoadSubmission() ([]*Submission, error) {
	if reader, err := os.Open(loader.Filename); err != nil {
		return nil, err
	} else {
		IOLoader := &SubmissionLoaderSingleFromIOReader{
			Reader: reader,
		}
		return IOLoader.LoadSubmission()
	}
}
