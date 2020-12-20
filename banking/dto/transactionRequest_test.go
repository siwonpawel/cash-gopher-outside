package dto

import (
	"net/http"
	"testing"
)

func Test_should_return_error_when_transaction_type_is_not_deposit_or_withdrawl(t *testing.T) {
	// Arrange
	request := TransactionRequest{
		Amount:          1,
		TransactionType: "invalid transaction type",
	}

	// Act
	appError := request.Validate()

	// Assert
	if appError.Message != "Transaction type need to be withdrawal or deposit" {
		t.Error("Invalid message while testing transaction type: " + appError.Message)
	}

	if appError.Code != http.StatusUnprocessableEntity {
		t.Error(" Invalid return code while testing transaction type")
	}
}

func Test_should_return_error_when_amount_is_less_than_zero(t *testing.T) {
	// Arrange
	request := TransactionRequest{
		Amount:          -100,
		TransactionType: "deposit",
	}

	// Act
	appError := request.Validate()

	// Assert
	if appError.Message != "Amount cannot be less or equal to 0.00" {
		t.Error("Invalid message while testing transaction amount")
	}
}
