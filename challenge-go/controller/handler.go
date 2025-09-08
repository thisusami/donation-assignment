package controller

import (
	"github.com/thisusami/donation-assignment/external"
	"github.com/thisusami/donation-assignment/service"
)

type Controller struct{}

func (c *Controller) Handler(path string) error {
	externalClient := external.NewOmiseClient()
	calculate := service.NewCalculator(path, externalClient)
	summary := service.NewSummary(calculate)
	err := summary.Summarize()
	if err != nil {
		return err
	}
	return nil
}
func NewController() *Controller {
	return &Controller{}
}
