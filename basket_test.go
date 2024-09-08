package main

import (
	"github.com/cucumber/godog"
)

type shopping struct{}

func (sg *shopping) thereIsASithLordLightsaverWhichCosts(arg1 int) error {
	return godog.ErrPending
}
