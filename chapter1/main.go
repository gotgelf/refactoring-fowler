package main

import (
	"fmt"
)

type performance struct {
	playID        string
	audience      int
	play          play
	amount        int
	volumeCredits int
}

type play struct {
	name     string
	playType string
}

type invoice struct {
	customer     string
	performances []performance
}

type statementData struct {
	customer           string
	performances       []performance
	totalAmount        int
	totalVolumeCredits int
}

func main() {
	var plays = map[string]play{
		"hamlet": {
			name:     "Hamlet",
			playType: "tragedy",
		},
		"as-like": {
			name:     "As You Like It",
			playType: "comedy",
		},
		"othello": {
			name:     "Othello",
			playType: "tragedy",
		},
	}

	var performances = []performance{
		{
			playID:   "hamlet",
			audience: 55,
		},
		{
			playID:   "as-like",
			audience: 35,
		},
		{
			playID:   "othello",
			audience: 40,
		},
	}

	invoice := invoice{
		customer:     "BigCo",
		performances: performances,
	}
	result := statement(invoice, plays)
	fmt.Println(result)
}

func statement(inv invoice, plays map[string]play) string {
	return renderPlainText(createStatementData(inv, plays))
	//return renderHtml(createStatementData(inv, plays))
}

func renderPlainText(data statementData) string {
	result := fmt.Sprintf("Statement for %s\n", data.customer)
	for _, perf := range data.performances {
		result += fmt.Sprintf(" %s: %s (%d) seats\n", perf.play.name, usd(perf.amount), perf.audience)
	}
	result += fmt.Sprintf("Amount owed is %s\n", usd(data.totalAmount))
	result += fmt.Sprintf("You earned  %d credits\n", data.totalVolumeCredits)
	return result
}

func htmlStatement(inv invoice, plays map[string]play) string {
	return renderHtml(createStatementData(inv, plays))
}

func renderHtml(data statementData) string {
	result := "h1" + fmt.Sprintf("Statement for %s\n", data.customer) + "</h1>\n"
	result += "<table>\n"
	result += "<tr><th>play</th><th>seats</th><th>cost</th></tr>"
	for _, perf := range data.performances {
		result += " <tr><td>" + fmt.Sprintf("%s", perf.play.name) + "</td><td>" + fmt.Sprintf("%d", perf.audience) + "</td>"
		result += "<td>" + fmt.Sprintf("%s", usd(perf.amount)) + "</td></tr>\n"
	}
	result += "</table>\n"
	result += "<p>" + fmt.Sprintf("Amount owed is %s\n", usd(data.totalAmount)) + "</p>\n"
	result += "<p>" + fmt.Sprintf("You earned  %d credits\n", data.totalVolumeCredits) + "</p>\n"
	return result
}

func usd(a int) string {
	return fmt.Sprintf("$%d", a/100)
}
