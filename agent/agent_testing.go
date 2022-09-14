package agent

import (
	"server/graph/model"

	"github.com/99designs/gqlgen/client"
)

func GrantRoleTest(c *client.Client, options client.Option, idToken string, userId string, role string, pinCode string, token string) (struct{ AssignRole bool }, error) {
	var _response struct{ AssignRole bool }
	err := c.Post(
		`
		mutation AssignRole($userId: String!, $role: String!, $pinCode: String!, $token: String!){
			assignRole(userId: $userId, role: $role, pinCode: $pinCode, token: $token)
		}
		`,
		&_response,
		client.Var("userId", userId),
		client.Var("role", role),
		client.Var("pinCode", pinCode),
		client.Var("token", token),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)
	return _response, err
}

func RemoveRoleTest(c *client.Client, options client.Option, idToken string, userId string, role string, pinCode string, token string) (struct{ UnassignRole bool }, error) {
	var _response struct{ UnassignRole bool }
	err := c.Post(
		`
	mutation UnAssignRole($userId: String!, $role: String!, $pinCode: String!, $token: String!){
		unassignRole(userId: $userId, role: $role, pinCode: $pinCode, token: $token)
	}
	`,
		&_response,
		client.Var("userId", userId),
		client.Var("role", role),
		client.Var("pinCode", pinCode),
		client.Var("token", token),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func CancelWithDrawTest(c *client.Client, options client.Option, idToken string, transaction_id string, _type string, pinCode string) (struct{ CancelTransactionAgent bool }, error) {
	var _response struct{ CancelTransactionAgent bool }
	err := c.Post(
		`
        mutation CancelTransactionAgent($transaction_id: String, $type: PaymentType, $pinCode: String!) {
            cancelTransactionAgent(transaction_id: $transaction_id, type: $type, pinCode: $pinCode)
        }
	`,
		&_response,
		client.Var("transaction_id", transaction_id),
		client.Var("type", _type),
		client.Var("pinCode", pinCode),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err

}

func CancelTopUpTest(c *client.Client, options client.Option, idToken string, transaction_id string, _type string, pinCode string) (struct{ CancelTransactionAgent bool }, error) {
	var _response struct{ CancelTransactionAgent bool }
	err := c.Post(
		`
			mutation CancelTransactionAgent($transaction_id: String, $type: PaymentType, $pinCode: String!) {
				cancelTransactionAgent(transaction_id: $transaction_id, type: $type, pinCode: $pinCode)
			}
	`,
		&_response,
		client.Var("transaction_id", transaction_id),
		client.Var("type", _type),
		client.Var("pinCode", pinCode),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func ConfirmTopUpTest(c *client.Client, options client.Option, idToken string, transaction_id string, _type model.PaymentType, pinCode string, token string) (struct{ ConfirmTransactionAgent bool }, error) {
	var _response struct{ ConfirmTransactionAgent bool }
	err := c.Post(
		`
        mutation ConfirmTransactionAgent($transaction_id: String!, $type: PaymentType!, $token: String!, $pinCode: String!) {
			confirmTransactionAgent(transaction_id: $transaction_id, type: $type, token: $token, pinCode: $pinCode)
	   }
	`,
		&_response,
		client.Var("transaction_id", transaction_id),
		client.Var("type", _type),
		client.Var("token", token),
		client.Var("pinCode", pinCode),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func ConfirmWithdrawTest(c *client.Client, options client.Option, idToken string, transaction_id string, _type model.PaymentType, pinCode string, token string) (struct{ ConfirmTransactionAgent bool }, error) {
	var _response struct{ ConfirmTransactionAgent bool }
	err := c.Post(
		`
        mutation ConfirmTransactionAgent($transaction_id: String!, $type: PaymentType!, $token: String!, $pinCode: String!) {
			confirmTransactionAgent(transaction_id: $transaction_id, type: $type, token: $token, pinCode: $pinCode)
			}
	`,
		&_response,
		client.Var("transaction_id", transaction_id),
		client.Var("type", _type),
		client.Var("token", token),
		client.Var("pinCode", pinCode),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}
