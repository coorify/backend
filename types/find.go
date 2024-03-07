package types

type KeywordBody struct {
	Keyword string `form:"keyword"`
}

type FindBody struct {
	PageBody
	KeywordBody
}
