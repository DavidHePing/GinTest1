package usecase

import (
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

func (*testUseCase) GetTestById(id string) string {
	return fmt.Sprint("Test Get id = ", id)
}

func (*testUseCase) PostTest() string {
	return "Test Post"
}

func (*testUseCase) PatchTest(id string) string {
	return fmt.Sprint("Test Patch id = ", id)
}

func (*testUseCase) DeleteTest(id string) string {
	return fmt.Sprint("Test Delete id = ", id)
}
