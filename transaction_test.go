package goplay

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaymentFromTransactions(t *testing.T) {
	ref := randomIDStringWithPrefix("trans")
	custRef := randomIDStringWithPrefix("trans")
	externalReference := RandomIDString(8)
	firstPaymentMethodCode := RandomIDString(8)
	firstAmount := rand.Intn(100)
	secondPaymentMethodCode := RandomIDString(8)
	secondAmount := rand.Intn(100)
	description := RandomIDString(8)
	refundAmount := rand.Intn(secondAmount)

	t1 := transaction{
		id:                0,
		orderReference:    ref,
		customerReference: custRef,
		status:            created,
		transactionType:   paymentType,
		externalReference: externalReference,
		paymentMethodCode: firstPaymentMethodCode,
		amount:            firstAmount,
		description:       description,
		version:           1,
	}

	t2 := transaction{
		id:                0,
		orderReference:    ref,
		customerReference: custRef,
		status:            created,
		transactionType:   paymentType,
		externalReference: externalReference,
		paymentMethodCode: secondPaymentMethodCode,
		amount:            secondAmount,
		description:       description,
		version:           2,
	}

	t3 := transaction{
		id:                0,
		orderReference:    ref,
		customerReference: custRef,
		status:            confirmed,
		transactionType:   paymentType,
		externalReference: externalReference,
		paymentMethodCode: firstPaymentMethodCode,
		amount:            firstAmount,
		description:       description,
		version:           3,
	}

	t4 := transaction{
		id:                0,
		orderReference:    ref,
		customerReference: custRef,
		status:            confirmed,
		transactionType:   paymentType,
		externalReference: externalReference,
		paymentMethodCode: secondPaymentMethodCode,
		amount:            secondAmount,
		description:       description,
		version:           4,
	}

	t5 := transaction{
		id:                0,
		orderReference:    ref,
		customerReference: custRef,
		status:            captured,
		transactionType:   paymentType,
		externalReference: externalReference,
		paymentMethodCode: secondPaymentMethodCode,
		amount:            secondAmount,
		description:       description,
		version:           5,
	}

	t6 := transaction{
		id:                0,
		orderReference:    ref,
		customerReference: custRef,
		status:            refundToSource,
		transactionType:   paymentType,
		externalReference: externalReference,
		paymentMethodCode: secondPaymentMethodCode,
		amount:            refundAmount,
		description:       description,
		version:           5,
	}

	expected := payment{
		orderReference:    ref,
		customerReference: custRef,
		status:            captured,
		paymentBreakdown: []paymentBreakdownItem{
			{paymentMethodCode: firstPaymentMethodCode, amount: firstAmount},
			{paymentMethodCode: secondPaymentMethodCode, amount: secondAmount - refundAmount},
		},
		description: description,
		amount:      firstAmount,
	}
	payment := paymentFromTransactions([]transaction{t1, t2, t3, t4, t5, t6})
	assert.Equal(t, expected, payment)
}
