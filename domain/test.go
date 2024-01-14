package domain

type TestUseCase interface {
	GetTest() string
	GetTestById(id int) string
	PostTest() string
	PatchTest(id int) string
	DeleteTest(id int) string
}
