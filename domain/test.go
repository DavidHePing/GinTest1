package domain

type TestUseCase interface {
	GetTest() string
	GetTestById(id string) string
	PostTest() string
	PatchTest(id string) string
	DeleteTest(id string) string
}
