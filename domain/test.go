package domain

type TestModel struct {
	Name   string `form:"Name" json:"Name" binding:"required"`
	Number int    `form:"Number" json:"Number" binding:"required"`
}

type TestUseCase interface {
	GetTest() string
	GetTestById(id int) string
	PostTest(testModel TestModel) string
	PatchTest(id int) string
	DeleteTest(id int) string
}
