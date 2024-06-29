package resolver

import (
	"errors"
	"regexp"
	"strconv"
)

const (
	minTitleLength          = 5
	maxTitleLength          = 100
	maxPostContentLength    = 10000
	minPostContentLength    = 5
	maxPostID               = 1 << 31
	maxCommentID            = 1 << 31
	maxCommentContentLength = 2000
	minCommentContentLength = 5
)

func cursorValidation(cursor string) error {
	if _, err := strconv.ParseUint(cursor, 10, 0); err != nil {
		return errors.New(`"after" isn't correct'`)
	}
	return nil
}

func limitValidation(limit int) error {
	if limit < 1 {
		return errors.New(`"first" isn't correct'`)
	}
	return nil
}

func postTitleValidation(title string) error {
	if len(title) < minTitleLength || len(title) > maxTitleLength {
		return errors.New(`"title" length isn't correct'`)
	}

	pattern := `^[a-zA-Z][a-zA-Z0-9 !.,?'’-]*$`
	if _, err := regexp.MatchString(pattern, title); err != nil {
		return errors.New(`"title" has an incorrect format'`)
	}
	return nil
}

func postContentValidation(content string) error {
	if len(content) < minPostContentLength || len(content) > maxPostContentLength {
		return errors.New(`"content" length isn't correct'`)
	}
	return nil
}

func commentContentValidation(content string) error {
	if len(content) < minCommentContentLength || len(content) > maxCommentContentLength {
		return errors.New(`"content" length isn't correct'`)
	}
	return nil
}
