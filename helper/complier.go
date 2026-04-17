package helper

func Complier(text string) string {

	text = Case(text)
	text = Punc(text)
	text = Quote(text)
	text = Article(text)
	text = Base(text)
	text = BaseN(text)

	return text

}
