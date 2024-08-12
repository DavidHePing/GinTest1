package usecase

import (
	"GinTest1/domain"
	"fmt"
)

type testUsecase struct {
}

// Ensure that testUsecase implements domain.TestUseCase
var _ domain.TestUseCase = &testUsecase{}

func NewTestUseCase() *testUsecase {
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
