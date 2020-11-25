package currency

import (
	. "github.com/onsi/gomega"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
)

func TestCurrency(t *testing.T) {
	g := NewGomegaWithT(t)
	matchCurrencies := getLatestCurrencies(g)
	btcCurrency := getLatestBtcCurrency(g)

	sut, err := NewCurrency()
	g.Expect(err).ShouldNot(HaveOccurred())

	t.Run("usd currency", func(t *testing.T) {
		g.Expect(gjson.Get(matchCurrencies, "rates.USD").Float()).Should(BeEquivalentTo(sut.USD()))
	})

	t.Run("eur currency", func(t *testing.T) {
		g.Expect(gjson.Get(matchCurrencies, "rates.EUR").Float()).Should(BeEquivalentTo(sut.EUR()))
	})

	t.Run("brl currency", func(t *testing.T) {
		g.Expect(gjson.Get(matchCurrencies, "rates.BRL").Float()).Should(BeEquivalentTo(sut.BRL()))
	})

	t.Run("btc currency", func(t *testing.T) {
		g.Expect(btcCurrency).Should(BeEquivalentTo(sut.BTC()))
	})

	t.Run("eth currency", func(t *testing.T) {
		g.Expect(gjson.Get(matchCurrencies, "rates.ETH").Float()).Should(BeEquivalentTo(sut.ETH()))
	})

	t.Run("cad currency", func(t *testing.T) {
		value, err := sut.Extra("CAD")

		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(gjson.Get(matchCurrencies, "rates.CAD").Float()).Should(BeEquivalentTo(value))
	})

	t.Run("cad lower case currency", func(t *testing.T) {
		value, err := sut.Extra("cad")

		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(gjson.Get(matchCurrencies, "rates.CAD").Float()).Should(BeEquivalentTo(value))
	})

	t.Run("invalid currency", func(t *testing.T) {
		value, err := sut.Extra("INVALID")

		g.Expect(err).Should(HaveOccurred())
		g.Expect(value).Should(BeEquivalentTo(0))
	})
}

func getLatestCurrencies(g *GomegaWithT) string {
	resp, err := http.Get(host)
	g.Expect(err).ShouldNot(HaveOccurred())
	g.Expect(resp.StatusCode).Should(BeEquivalentTo(http.StatusOK))
	body, err := ioutil.ReadAll(resp.Body)
	g.Expect(err).ShouldNot(HaveOccurred())
	return string(body)
}

func getLatestBtcCurrency(g *GomegaWithT) float64 {
	resp, err := http.Get(btcHost)
	g.Expect(err).ShouldNot(HaveOccurred())
	body, err := ioutil.ReadAll(resp.Body)
	g.Expect(err).ShouldNot(HaveOccurred())
	btc, err := strconv.ParseFloat(string(body), 64)
	g.Expect(err).ShouldNot(HaveOccurred())
	return btc
}
