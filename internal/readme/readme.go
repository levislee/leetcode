package readme

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/openset/leetcode/internal/base"
	"github.com/openset/leetcode/internal/leetcode"
)

var CmdReadme = &base.Command{
	Run:       runReadme,
	UsageLine: "readme",
	Short:     "build README.md file",
	Long:      "build README.md file.",
}

func runReadme(cmd *base.Command, args []string) {
	if len(args) == 1 && args[0] == "page" {
		buildCmd = "page"
		lockStr = " &hearts;"
		fileName = "index.md"
	}
	if len(args) > 1 {
		cmd.Usage()
	}
	var bf bytes.Buffer
	bf.WriteString(fmt.Sprintf(format, buildCmd, strings.Repeat(" ", 15-len(buildCmd))))
	bf.WriteString(base.AuthInfo)
	bf.WriteString(defaultStr)
	writeProblems(&bf)
	base.FilePutContents(fileName, bf.Bytes())
}

func writeProblems(bf *bytes.Buffer) {
	problems := leetcode.ProblemsAll()
	problemsSet := make(map[int]string)
	maxId := 0
	for _, problem := range problems.StatStatusPairs {
		id := problem.Stat.FrontendQuestionId
		title := strings.TrimSpace(problem.Stat.QuestionTitle)
		needPaid := ""
		if problem.PaidOnly {
			needPaid += lockStr
		}
		slug := problem.Stat.QuestionTitleSlug
		levelName := problem.Difficulty.LevelName()
		format := "| <span id=\"%d\">%d</span> | [%s](https://leetcode.com/problems/%s)%s | [Go](https://github.com/openset/leetcode/tree/master/solution/%s) | %s |\n"
		problemsSet[id] = fmt.Sprintf(format, id, id, title, slug, needPaid, slug, levelName)
		if id > maxId {
			maxId = id
		}
	}

	step, long := 50, 300
	bf.WriteString("<table><thead>\n")
	for i := 0; i < maxId; i += long {
		bf.WriteString("<tr>\n")
		for j := 0; j < long/step; j++ {
			bf.WriteString(fmt.Sprintf("<th align=\"center\"><a href=\"#%d\">[%d-%d]</a></th>\n", 1+i+j*step, 1+i+j*step, i+j*step+step))
		}
		bf.WriteString("</tr>\n")
	}
	bf.WriteString("</thead></table>\n")

	bf.WriteString("\n")
	bf.WriteString("| # | Title | Solution | Difficulty |\n")
	bf.WriteString("| :-: | - | - | :-: |\n")
	for i := 0; i <= maxId; i++ {
		if row, ok := problemsSet[i]; ok {
			bf.WriteString(row)
		}
	}
}

var buildCmd = "readme"

var lockStr = " 🔒"

var fileName = "README.md"

var format = "<!--|This file generated by command(leetcode %s); DO NOT EDIT.%s|-->"

var defaultStr = `
# [LeetCode](https://openset.github.io/leetcode)
LeetCode Problems' Solutions

[![Build Status](https://travis-ci.org/openset/leetcode.svg?branch=master)](https://travis-ci.org/openset/leetcode)
[![codecov](https://codecov.io/gh/openset/leetcode/branch/master/graph/badge.svg)](https://codecov.io/gh/openset/leetcode)
[![Go Report Card](https://goreportcard.com/badge/github.com/openset/leetcode)](https://goreportcard.com/report/github.com/openset/leetcode)
[![GitHub contributors](https://img.shields.io/github/contributors/openset/leetcode.svg)](https://github.com/openset/leetcode/graphs/contributors)
[![license](https://img.shields.io/github/license/openset/leetcode.svg)](https://github.com/openset/leetcode/blob/master/LICENSE)
[![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/openset/leetcode.svg?colorB=green)](https://github.com/openset/leetcode/archive/master.zip)

`
