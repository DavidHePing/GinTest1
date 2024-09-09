package domain

type TestModel struct {
	Name   string `form:"Name" json:"Name" binding:"required"`
	Number int    `form:"Number" json:"Number" binding:"required"` //cannot be default if required
}
