package multasgt

import (
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

// Ticket represents the information related to the ticket.
type Ticket struct {
	ID       string `json:"id"`
	Entity   string `json:"entity"`
	Date     string `json:"date"`
	Ammount  string `json:"ammount"`
	Discount string `json:"discount"`
	Total    string `json:"total"`
	Location string `json:"location"`
	Info     string `json:"info"`
	Photo    string `json:"photo"`
}

// TicketChecker is the interface that all checkers must implement.
type TicketChecker interface {
	Check(plateType, plateNumber string) ([]Ticket, error)
}

func getAttribute(attrName string, n *html.Node) string {
	if n == nil {
		return ""
	}
	for i, a := range n.Attr {
		if a.Key == attrName {
			return n.Attr[i].Val
		}
	}
	return ""
}

func cleanStrings(s string) string {
	clean := regexp.MustCompile(`[\t\n\r]`)
	return clean.ReplaceAllString(strings.TrimSpace(s), "")
}
