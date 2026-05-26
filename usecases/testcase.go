package usecases

type TestCase struct {
	Input  any
	Output any
}

type KateInput struct {
	DocsAmount         int
	TimeCollegueLeaves int
	Floors             []int
	IdCollegue         int
}
