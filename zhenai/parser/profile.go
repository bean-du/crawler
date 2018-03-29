package parser

import (
	"regexp"

	"strconv"

	"github.com/bean-du/crawler/engine"
	"github.com/bean-du/crawler/model"
)

var (
	ageRe       = regexp.MustCompile(`<td><span class="label">年龄：</span>([0-9]+)岁</td>`)
	marriageRe  = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
	nameRe      = regexp.MustCompile(`<a class="name fs24">([^<]+)</a>`)
	ganderRe    = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
	heightRe    = regexp.MustCompile(`<td><span class="label">身高：</span><span field="">([0-9]+)CM</span></td>`)
	weightRe    = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([0-9]+)KG</span></td>`)
	incomeRe    = regexp.MustCompile(`<td><span class="label">月收入：</span>([0-9]+-[0-9]+元)</td>`)
	educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
	occupation  = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
	hukouRe     = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
	xingzuoRe   = regexp.MustCompile(`<td><span class="label">星座：</span>([^<]+)</td>`)
	houseRe     = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
	carRe       = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
	idUrlRe     = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)
)

func ParseProfile(contents []byte, url string) engine.ParseResult {
	profile := model.Profile{}

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}
	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}
	profile.Marriage = extractString(contents, marriageRe)
	profile.Name = extractString(contents, nameRe)
	profile.Gander = extractString(contents, ganderRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Education = extractString(contents, educationRe)
	profile.Occupation = extractString(contents, occupation)
	profile.Hukou = extractString(contents, hukouRe)
	profile.Xingzuo = extractString(contents, xingzuoRe)
	profile.House = extractString(contents, houseRe)
	profile.Car = extractString(contents, carRe)
	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      extractString([]byte(url), idUrlRe),
				Payload: profile,
			},
		},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
