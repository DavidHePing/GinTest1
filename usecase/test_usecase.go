package usecase

import (
	"GinTest1/domain"
	"fmt"
)

type TestUseCase interface {
	GetTest() string
	GetTestById(id int) string
	PostTest(testModel domain.TestModel) string
	PatchTest(id int, testModel domain.TestModel) string
	DeleteTest(id int) string
}

type testUsecase struct {
}

func NewTestUseCase() TestUseCase {
	return &testUsecase{}
}

func (*testUsecase) GetTest() string {
	return "Test Get"
}
func (*testUsecase) GetTestById(id int) string {
	return fmt.Sprint("Test Get id = ", id)
}
func (*testUsecase) PostTest(testModel domain.TestModel) string {
	return fmt.Sprint("Test Post ", testModel)
}

func (*testUsecase) PatchTest(id int, testModel domain.TestModel) string {
	return fmt.Sprint("Test Patch id = ", id, " Body is: ", testModel)
}

func (*testUsecase) DeleteTest(id int) string {
	return fmt.Sprint("Test Delete id = ", id)
}
