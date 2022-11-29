package main

import (
	"fmt"
	"math"
)

func createStatementData(inv invoice, plays map[string]play) statementData {
	playFor := func(perf performance) play {
		return plays[perf.playID]
	}
	enrichPerformance := func(perf performance) performance {
		calculator := createPerformanceCalculator(perf, playFor(perf))
		var result performance
		result = perf
		result.play = calculator.play()
		result.amount = calculator.amount()
		result.volumeCredits = calculator.volumesCreditsFor()

		return result
	}

	var result statementData
	result.customer = inv.customer
	for _, v := range inv.performances {
		result.performances = append(result.performances, enrichPerformance(v))
	}
	result.totalAmount = totalAmount(result)
	result.totalVolumeCredits = totalVolumeCredits(result)

	return result
}

func totalAmount(data statementData) int {
	result := 0
	for _, perf := range data.performances {
		result += perf.amount
	}
	return result
}

func totalVolumeCredits(data statementData) int {
	result := 0
	for _, perf := range data.performances {
		result += perf.volumeCredits
	}
	return result
}

type performanceCalculator interface {
	amount() int
	volumesCreditsFor() int
	play() play
}

type basicCalculator struct {
	perf performance
	pl   play
}

type tragedyCalculator basicCalculator

type comedyCalculator basicCalculator

func createPerformanceCalculator(perf performance, pl play) performanceCalculator {
	switch pl.playType {
	case "tragedy":
		return &tragedyCalculator{
			perf: perf,
			pl:   pl,
		}
	case "comedy":
		return &comedyCalculator{
			perf: perf,
			pl:   pl,
		}
	default:
		err := fmt.Sprintf("unknown type: %s", pl.playType)
		panic(err)
	}
}

func (tc tragedyCalculator) amount() int {
	result := 40000
	if tc.perf.audience > 30 {
		result += 1000 * (tc.perf.audience - 30)
	}
	return result
}

func (tc tragedyCalculator) volumesCreditsFor() int {
	result := 0
	result += int(math.Max(float64(tc.perf.audience-30), 0))
	return result
}

func (tc tragedyCalculator) play() play {
	return tc.pl
}

func (cc comedyCalculator) amount() int {
	result := 30000
	if cc.perf.audience > 20 {
		result += 10000 + 500*(cc.perf.audience-20)
	}
	result += 300 * cc.perf.audience
	return result
}

func (cc comedyCalculator) volumesCreditsFor() int {
	result := 0
	result += int(math.Max(float64(cc.perf.audience-30), 0))

	result += int(math.Floor(float64(cc.perf.audience / 5)))
	return result
}

func (cc comedyCalculator) play() play {
	return cc.pl
}
