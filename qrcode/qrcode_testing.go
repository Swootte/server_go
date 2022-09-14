package qrcode

import (
	"server/graph/model"

	"github.com/99designs/gqlgen/client"
)

func GetQrOwnerTest(c *client.Client, options client.Option, idToken string, qrcode string) (struct {
	GetQrOwner model.EnterpriseSmall
}, error) {
	var _response struct {
		GetQrOwner model.EnterpriseSmall
	}
	err := c.Post(
		`
        query GetQrOwner($qrcode: String!) {
            getQrOwner(qrcode: $qrcode) {
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
                        name
                        walletPublicKey
                        
                    }
            }
        }
	`,
		&_response,
		client.Var("qrcode", qrcode),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func GetQrOwnerUserTest(c *client.Client, options client.Option, idToken string, qrcode string) (struct {
	GetQrOwner model.UserSmall
}, error) {
	var _response struct {
		GetQrOwner model.UserSmall
	}
	err := c.Post(
		`
        query GetQrOwner($qrcode: String!) {
            getQrOwner(qrcode: $qrcode) {
                    ... on Paiement {
                        _id
                        status
                    }

                    ... on UserSmall {
                        _id
                        first_name
                        last_name
                        address
                        
                    }


                    ... on EnterpriseSmall {
                        _id
                        walletPublicKey
                        
                    }
            }
        }
	`,
		&_response,
		client.Var("qrcode", qrcode),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func GetQrOwnerPaymentTest(c *client.Client, options client.Option, idToken string, qrcode string) (struct {
	GetQrOwner model.Paiement
}, error) {
	var _response struct {
		GetQrOwner model.Paiement
	}
	err := c.Post(
		`
        query GetQrOwner($qrcode: String!) {
            getQrOwner(qrcode: $qrcode) {
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
		client.Var("qrcode", qrcode),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}
