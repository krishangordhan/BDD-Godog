package main

import (
	"github.com/cucumber/godog"
)

type shopping struct {
	shelf *Shelf
}

type shoppingBasket struct {
	shelf *Shelf
}

func (sb *shoppingBasket) addProduct(productName string, price float64) (err error) {
	sb.shelf.AddProduct(productName, price)
	return
}

func (sh *shopping) thereIsASithLordLightsaverWhichCosts(arg1 int) error {
	return godog.ErrPending
}

func (sh *shopping) thereIsAJediLightsaberWhichCosts(arg1 int) error {
	return godog.ErrPending
}

func (sh *shopping) iAddThetoTheBasket(arg1 int) error {
	return godog.ErrPending
}

func (sh *shopping) iShouldHaveProductsInTheBasket(arg1 int) error {
	return godog.ErrPending
}

func (sh *shopping) theOverallBasketPriceShouldBe(arg1 int) error {
	return godog.ErrPending
}

func featureContext(ctx *godog.ScenarioContext) {
	ctx.Before(func(sc *godog.Scenario) (Shelf, error) {
		return NewShelf(), nil
	})

	sh := &shopping{}
	sb := shoppingBasket{}

	ctx.Step(`^there is a "([a-zA-Z\s]+)", which costs £(\d+)$`, sb.addProduct)
	ctx.Step(`^there is a "([a-zA-Z\s]+)", which costs £(\d+)$`, sb.addProduct)
	ctx.Step(`^I add (\d+) to the basket$`, sh.iAddThetoTheBasket)
	ctx.Step(`^I should have (\d+) products in the basket$`, sh.iShouldHaveProductsInTheBasket)
	ctx.Step(`^the overall basket price should be (\d+)$`, sh.theOverallBasketPriceShouldBe)
}
