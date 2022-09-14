package transaction

import (
	"server/graph/model"

	"github.com/99designs/gqlgen/client"
)

func InitiaTopupTest(c *client.Client, options client.Option, idToken string, pincode string, topup model.TopUpInput) (struct{ AddTopUp string }, error) {
	var _response struct {
		AddTopUp string
	}

	err := c.Post(
		`
        mutation TOPUP($topup: TopUpInput!, $pinCode: String!){
            addTopUp(topup: $topup, pinCode: $pinCode)
        }
	`,
		&_response,
		client.Var("topup", topup),
		client.Var("pinCode", pincode),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func InitiateWithdrawTest(c *client.Client, options client.Option, idToken string, pincode string, withdraw model.WithdrawInput) (struct{ AddWithDraw string }, error) {
	var _response struct{ AddWithDraw string }
	err := c.Post(
		`
		mutation WITHDRAW($withdraw: WithdrawInput!, $pinCode: String!){
			addWithDraw(withdraw: $withdraw, pinCode: $pinCode)
		}
	`,
		&_response,
		client.Var("withdraw", withdraw),
		client.Var("pinCode", pincode),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func CancelTopUpUserTest(c *client.Client, options client.Option, idToken string, transaction_id string, _type model.PaymentType, pinCode string) (struct{ CancelTransactionUser bool }, error) {
	var _response struct{ CancelTransactionUser bool }
	err := c.Post(
		`
        mutation CancelTransactionUser($transaction_id: String, $type: PaymentType, $pinCode: String!) {
            cancelTransactionUser(transaction_id: $transaction_id, type: $type, pinCode: $pinCode)
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

func CancelWithdrawUserTest(c *client.Client, options client.Option, idToken string, transaction_id string, _type model.PaymentType, pinCode string) (struct{ CancelTransactionUser bool }, error) {
	var _response struct{ CancelTransactionUser bool }
	err := c.Post(
		`
        mutation CancelTransactionUser($transaction_id: String, $type: PaymentType, $pinCode: String!) {
            cancelTransactionUser(transaction_id: $transaction_id, type: $type, pinCode: $pinCode)
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

func SendMoneyToPayeeTest(c *client.Client, options client.Option, idToken string, address string, token string, amount float32, pinCode string, destinationUser string) (struct{ CreateTransfer bool }, error) {
	var _response struct{ CreateTransfer bool }
	err := c.Post(
		`
		mutation createTransfer($address: String, $token: String!, $amount: Float!, $pinCode: String!, $destinationUser: String!) {
			createTransfer(address: $address, token: $token, amount: $amount, pinCode: $pinCode, destinationUser: $destinationUser)
		}
	`,
		&_response,
		client.Var("address", address),
		client.Var("token", token),
		client.Var("amount", amount),
		client.Var("pinCode", pinCode),
		client.Var("destinationUser", destinationUser),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func LoadAllUserTransactionTest(c *client.Client, options client.Option, idToken string) (struct{ GetActivity []interface{} }, error) {
	var _response struct{ GetActivity []interface{} }
	err := c.Post(
		`
        query GetActivity{
            getActivity {
            _id
            status
            type
            creator{
              _id
              first_name
              last_name
              photoUrl
            }
            amount
            agency {
              _id
              title
              address
              status
            }
            destinationUser {
              _id
              first_name
              last_name
              photoUrl
            }
            destination
            transactionId
            createdAt
            updatedAt
            shortId
            }
        }
	`,
		&_response,
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func AuthenticateForPaymentTest(c *client.Client, options client.Option, idToken string, amount float32, ref string) (struct {
	AuthenticateForPayment model.Paiement
}, error) {
	var _response struct {
		AuthenticateForPayment model.Paiement
	}
	err := c.Post(
		`
		mutation AuthenticateForPayment($amount: Float!, $ref: String!) {
			authenticateForPayment(amount: $amount, ref: $ref) {
				... on Paiement {
					_id
					status
				}
				... on UserSmall {
					_id
					first_name
					
				}
				... on EnterpriseSmall {
					_id
					walletPublicKey
					
				}
			}
		}
	`,
		&_response,
		client.Var("amount", amount),
		client.Var("ref", ref),
		client.AddHeader("Authorization", "Basic "+idToken),
		options,
	)

	return _response, err
}

func GetTransactionByIdTest(c *client.Client, options client.Option, idToken string, id string) (struct{ GetTransactionById model.Paiement }, error) {
	var _response struct{ GetTransactionById model.Paiement }
	err := c.Post(
		`
        query GetTransactionById($id: String!) {
            getTransactionById(id: $id,) {
                    _id
                    status
                    cancellor {
                        _id
                    }
            }
        }
	`,
		&_response,
		client.Var("id", id),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}
