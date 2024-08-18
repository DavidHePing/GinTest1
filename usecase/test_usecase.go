package usecase

import (
	"GinTest1/domain"
	"fmt"
)

type TestUsecase struct {
}

func NewTestUseCase() *TestUsecase {
	return &TestUsecase{}
}

func (*TestUsecase) GetTest() string {
	return "Test Get"
}
func (*TestUsecase) GetTestById(id int) string {
	return fmt.Sprint("Test Get id = ", id)
}
func (*TestUsecase) PostTest(testModel domain.TestModel) string {
	return fmt.Sprint("Test Post ", testModel)
}

func (*TestUsecase) PatchTest(id int, testModel domain.TestModel) string {
	return fmt.Sprint("Test Patch id = ", id, " Body is: ", testModel)
}

func (*TestUsecase) DeleteTest(id int) string {
	return fmt.Sprint("Test Delete id = ", id)
}
