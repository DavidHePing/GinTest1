package usecase

import (
	"fmt"
)

type testUseCase struct {
}

func NewTestUseCase() *testUseCase {
	return &testUseCase{}
}

// @Schemes
// @Tags test
// @Accept json
// @Produce json
// @Success 200 {string} Test Get
// @Router /test [get]
func (*testUseCase) GetTest() string {
	return "Test Get"
}

// @Schemes
// @Tags test
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Success 200 {string} Test Get
// @Router /test/{id} [get]
func (*testUseCase) GetTestById(id string) string {
	return fmt.Sprint("Test Get id = ", id)
}

// @Schemes
// @Tags test
// @Accept json
// @Produce json
// @Success 200 {string} Test Get
// @Router /test [post]
func (*testUseCase) PostTest() string {
	return "Test Post"
}

// @Schemes
// @Tags test
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Success 200 {string} Test Get
// @Router /test/{id} [patch]
func (*testUseCase) PatchTest(id string) string {
	return fmt.Sprint("Test Patch id = ", id)
}

// @Schemes
// @Tags test
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Success 200 {string} Test Get
// @Router /test/{id} [delete]
func (*testUseCase) DeleteTest(id string) string {
	return fmt.Sprint("Test Delete id = ", id)
}
