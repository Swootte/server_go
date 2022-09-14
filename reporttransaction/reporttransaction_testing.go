package reporttransaction

import "github.com/99designs/gqlgen/client"

func ReportTransactionTest(c *client.Client, options client.Option, idToken string, transaction_id string, message string) (struct{ ReportTransaction bool }, error) {
	var _response struct{ ReportTransaction bool }
	err := c.Post(
		`
		mutation ReportTransaction($transaction_id: String!, $message: String!) {
			reportTransaction(transaction_id: $transaction_id, message: $message)
		}
	`,
		&_response,
		client.Var("transaction_id", transaction_id),
		client.Var("message", message),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)
	return _response, err
}
