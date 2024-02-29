package usecase

import (
	"GinTest1/domain"
	"fmt"
)

type testUseCase struct {
}

func NewTestUseCase() *testUseCase {
	return &testUseCase{}
}

func (*testUseCase) GetTest() string {
	return "Test Get"
}
func (*testUseCase) GetTestById(id int) string {
	return fmt.Sprint("Test Get id = ", id)
}
func (*testUseCase) PostTest(testModel domain.TestModel) string {
	return fmt.Sprint("Test Post ", testModel)
}

func (*testUseCase) PatchTest(id int, testModel domain.TestModel) string {
	return fmt.Sprint("Test Patch id = ", id, " Body is: ", testModel)
}

func (*testUseCase) DeleteTest(id int) string {
	return fmt.Sprint("Test Delete id = ", id)
}
