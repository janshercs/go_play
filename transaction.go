package goplay

import "time"

const (
	first uint = 1
)

type status string
type transactionType string

type transaction struct {
	id                int
	orderReference    string
	customerReference string
	status            status
	transactionType   transactionType
	externalReference string
	paymentMethodCode string
	amount            int
	description       string
	version           uint
	createdAt         time.Time
	updatedAt         time.Time
}

type paymentBreakdown []paymentBreakdownItem
type paymentBreakdownItem struct {
	paymentMethodCode string
	amount            int
}
type payment struct {
	orderReference    string
	customerReference string
	status            status
	paymentBreakdown  paymentBreakdown
	description       string
	amount            int
}

func paymentFromTransactions(transactions []transaction) (p payment) {
	for _, t := range transactions {
		p = t.Process(p)
	}
	return
}

func (t transaction) Process(p payment) payment {
	switch t.status {
	case created:
		if t.version == first {
			p = payment{
				orderReference:    t.orderReference,
				customerReference: t.customerReference,
				status:            t.status,
				description:       t.description,
				amount:            t.amount,
			}
		}

		p.paymentBreakdown = append(p.paymentBreakdown, paymentBreakdownItem{t.paymentMethodCode, t.amount})

	case confirmed, captured:
		p.status = t.status

	case refundToSource:
		for i := range p.paymentBreakdown {
			if p.paymentBreakdown[i].paymentMethodCode == t.paymentMethodCode {
				p.paymentBreakdown[i].amount -= t.amount
			}
		}
	}

	return p
}
