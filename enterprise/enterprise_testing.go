package enterprise

import (
	"server/graph/model"

	"github.com/99designs/gqlgen/client"
)

func CreateEnterpriseTest(c *client.Client, options client.Option, idToken string, enterprise model.EnterpriseInput) (struct {
	CreateEnterprise struct {
		ID                   string       `json:"_id" bson:"_id"`
		Name                 string       `json:"name" bson:"name"`
		Type                 string       `json:"type" bson:"type"`
		LogoUrl              string       `json:"logoUrl" bson:"logoUrl"`
		PublishableKey       string       `json:"publishableKey" bson:"publishableKey"`
		Private_key          string       `json:"private_key" bson:"private_key"`
		WalletPublicKey      string       `json:"walletPublicKey" bson:"walletPublicKey"`
		DefaultEnterprise    bool         `json:"default_enterprise" bson:"default_enterprise"`
		Country              string       `json:"country" bson:"country"`
		Description          string       `json:"description" bson:"description"`
		SellingPhysicalGoods bool         `json:"sellingPhysicalGoods" bson:"sellingPhysicalGoods"`
		SelfShippingProduct  string       `json:"selfShippingProduct" bson:"selfShippingProduct"`
		ShippingDelay        string       `json:"shippingDelay" bson:"shippingDelay"`
		TransactionLibele    string       `json:"transactionLibele" bson:"transactionLibele"`
		AbregedLibele        string       `json:"abregedLibele" bson:"abregedLibele"`
		Phone                model.Phone  `json:"phone" bson:"phone"`
		Email                string       `json:"email" bson:"email"`
		Sector               string       `json:"sector" bson:"sector"`
		RCCM                 string       `json:"rccm" bson:"rccm"`
		Website              string       `json:"website" bson:"website"`
		Person               model.Person `json:"person" bson:"person"`
	}
}, error) {
	var _response struct {
		CreateEnterprise struct {
			ID                   string       `json:"_id" bson:"_id"`
			Name                 string       `json:"name" bson:"name"`
			Type                 string       `json:"type" bson:"type"`
			LogoUrl              string       `json:"logoUrl" bson:"logoUrl"`
			PublishableKey       string       `json:"publishableKey" bson:"publishableKey"`
			Private_key          string       `json:"private_key" bson:"private_key"`
			WalletPublicKey      string       `json:"walletPublicKey" bson:"walletPublicKey"`
			DefaultEnterprise    bool         `json:"default_enterprise" bson:"default_enterprise"`
			Country              string       `json:"country" bson:"country"`
			Description          string       `json:"description" bson:"description"`
			SellingPhysicalGoods bool         `json:"sellingPhysicalGoods" bson:"sellingPhysicalGoods"`
			SelfShippingProduct  string       `json:"selfShippingProduct" bson:"selfShippingProduct"`
			ShippingDelay        string       `json:"shippingDelay" bson:"shippingDelay"`
			TransactionLibele    string       `json:"transactionLibele" bson:"transactionLibele"`
			AbregedLibele        string       `json:"abregedLibele" bson:"abregedLibele"`
			Phone                model.Phone  `json:"phone" bson:"phone"`
			Email                string       `json:"email" bson:"email"`
			Sector               string       `json:"sector" bson:"sector"`
			RCCM                 string       `json:"rccm" bson:"rccm"`
			Website              string       `json:"website" bson:"website"`
			Person               model.Person `json:"person" bson:"person"`
		}
	}
	err := c.Post(
		`
		mutation CreateEnterprise($enterprise: EnterpriseInput!) {
			createEnterprise(enterprise: $enterprise) {
				_id
				name
				type
				logoUrl
				publishableKey
				private_key
				walletPublicKey
				default_enterprise
				country
				description
				sellingPhysicalGoods
				selfShippingProduct
				shippingDelay
				transactionLibele
				abregedLibele
				phone {
					phone
					dialcode
				}
				email
				sector
				rccm
				website
				person {
					first_name
					last_name
					email
					address
					state
					city
					zip
				}
			}
		}
	`,
		&_response,
		client.Var("enterprise", enterprise),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func RemoveEnterpriseTest(c *client.Client, options client.Option, idToken string, pincode string, enterpriseId string) (struct{ RemoveEnterprise []model.Enterprise }, error) {
	var _response struct{ RemoveEnterprise []model.Enterprise }
	err := c.Post(
		`
		mutation RemoveEnterprise($enterpriseId: String!, $pinCode: String!) {
			removeEnterprise(enterpriseId: $enterpriseId, pinCode: $pinCode){
				_id
				default_enterprise
			}
		}
	`,
		&_response,
		client.Var("enterpriseId", enterpriseId),
		client.Var("pinCode", pincode),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func RecreateTestPublishableKeyTest(c *client.Client, options client.Option, idToken string, pincode string, enterpriseId string) (struct {
	RecreateEnterprisePublishableKey []struct {
		ID                string `json:"_id" bson:"_id"`
		Type              string `json:"type" bson:"type"`
		Name              string `json:"name" bson:"name"`
		Website           string `json:"website" bson:"website"`
		LogoUrl           string `json:"logoUrl" bson:"logoUrl"`
		Creator           string `json:"creator" bson:"creator"`
		CreatedAt         string `json:"createdAt" bson:"createdAt"`
		UpdatedAt         string `json:"updatedAt" bson:"updatedAt"`
		PublishableKey    string `json:"publishableKey" bson:"publishableKey"`
		Private_key       string `json:"private_key" bson:"private_key"`
		WalletPublicKey   string `json:"walletPublicKey" bson:"walletPublicKey"`
		DefaultEnterprise bool   `json:"default_enterprise" bson:"default_enterprise"`
	}
}, error) {
	var _response struct {
		RecreateEnterprisePublishableKey []struct {
			ID                string `json:"_id" bson:"_id"`
			Type              string `json:"type" bson:"type"`
			Name              string `json:"name" bson:"name"`
			Website           string `json:"website" bson:"website"`
			LogoUrl           string `json:"logoUrl" bson:"logoUrl"`
			Creator           string `json:"creator" bson:"creator"`
			CreatedAt         string `json:"createdAt" bson:"createdAt"`
			UpdatedAt         string `json:"updatedAt" bson:"updatedAt"`
			PublishableKey    string `json:"publishableKey" bson:"publishableKey"`
			Private_key       string `json:"private_key" bson:"private_key"`
			WalletPublicKey   string `json:"walletPublicKey" bson:"walletPublicKey"`
			DefaultEnterprise bool   `json:"default_enterprise" bson:"default_enterprise"`
		}
	}
	err := c.Post(
		`
        mutation RecreatePublishableKey($enterpriseId: String!, $pinCode: String!) {
            recreateEnterprisePublishableKey(enterpriseId: $enterpriseId, pinCode: $pinCode) {
                    _id
                    type
                    name
                    website
                    logoUrl
                    creator
                    createdAt
                    updatedAt
                    publishableKey
                    private_key
                    walletPublicKey
                    default_enterprise
            }
        }
	`,
		&_response,
		client.Var("enterpriseId", enterpriseId),
		client.Var("pinCode", pincode),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func RecreateTestPrivate_keyTest(c *client.Client, options client.Option, idToken string, pincode string, enterpriseId string) (struct {
	RecreateEnterprisePrivateKey []struct {
		ID                string `json:"_id" bson:"_id"`
		Type              string `json:"type" bson:"type"`
		Name              string `json:"name" bson:"name"`
		Website           string `json:"website" bson:"website"`
		LogoUrl           string `json:"logoUrl" bson:"logoUrl"`
		Creator           string `json:"creator" bson:"creator"`
		CreatedAt         string `json:"createdAt" bson:"createdAt"`
		UpdatedAt         string `json:"updatedAt" bson:"updatedAt"`
		PublishableKey    string `json:"publishableKey" bson:"publishableKey"`
		Private_key       string `json:"private_key" bson:"private_key"`
		WalletPublicKey   string `json:"walletPublicKey" bson:"walletPublicKey"`
		DefaultEnterprise bool   `json:"default_enterprise" bson:"default_enterprise"`
	}
}, error) {
	var _response struct {
		RecreateEnterprisePrivateKey []struct {
			ID                string `json:"_id" bson:"_id"`
			Type              string `json:"type" bson:"type"`
			Name              string `json:"name" bson:"name"`
			Website           string `json:"website" bson:"website"`
			LogoUrl           string `json:"logoUrl" bson:"logoUrl"`
			Creator           string `json:"creator" bson:"creator"`
			CreatedAt         string `json:"createdAt" bson:"createdAt"`
			UpdatedAt         string `json:"updatedAt" bson:"updatedAt"`
			PublishableKey    string `json:"publishableKey" bson:"publishableKey"`
			Private_key       string `json:"private_key" bson:"private_key"`
			WalletPublicKey   string `json:"walletPublicKey" bson:"walletPublicKey"`
			DefaultEnterprise bool   `json:"default_enterprise" bson:"default_enterprise"`
		}
	}
	err := c.Post(
		`
        mutation RecreatePrivate_key($enterpriseId: String!, $pinCode: String!) {
            recreateEnterprisePrivateKey(enterpriseId: $enterpriseId, pinCode: $pinCode) {
                    _id
                    type
                    name
                    website
                    logoUrl
                    creator
                    createdAt
                    updatedAt
                    publishableKey
                    private_key
                    walletPublicKey
                    default_enterprise
            }
        }
	`,
		&_response,
		client.Var("enterpriseId", enterpriseId),
		client.Var("pinCode", pincode),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)
	return _response, err
}

func UpdateEnterpriseTypeTest(c *client.Client, options client.Option, idToken string, enterpriseId string, _type string, _country string) (struct{ UpdateEnterpriseType []model.Enterprise }, error) {
	var _response struct{ UpdateEnterpriseType []model.Enterprise }
	err := c.Post(
		`
        mutation UpdateEnterpriseType($enterpriseId: String!, $type: String!, $country: String!) {
            updateEnterpriseType(enterpriseId: $enterpriseId, type: $type, country: $country) {
                    _id
                    name
                    type
                    logoUrl
                    publishableKey
                    private_key
                    walletPublicKey
                    default_enterprise
                    country
                    description
                    sellingPhysicalGoods
                    selfShippingProduct
                    shippingDelay
                    transactionLibele
                    abregedLibele
					phone {
						phone
						dialcode
					}
                    email
                    sector
                    rccm
                    website
                    person {
                        first_name
                        last_name
                        email
                        address
                        state
                        city
                        zip
                    }
            }
        }
	`,
		&_response,
		client.Var("enterpriseId", enterpriseId),
		client.Var("type", _type),
		client.Var("country", _country),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func UpdatePersonnalInformationTest(c *client.Client, options client.Option, idToken string, enterpriseId string, first_name string, last_name string, email string, address string, city string, state string, zip string) (struct{ UpdatePersonnalInformation []model.Enterprise }, error) {
	var _response struct{ UpdatePersonnalInformation []model.Enterprise }
	err := c.Post(
		`
        mutation UpdatePersonnalInformation($enterpriseId: String!, $first_name: String!, $last_name: String!, $email: String!, $address: String!, $city: String!, $state: String!, $zip: String!) {
            updatePersonnalInformation(enterpriseId: $enterpriseId, first_name: $first_name, last_name: $last_name, email: $email, address: $address, city: $city, state: $state, zip: $zip) {
                	_id
                    name
                    type
                    logoUrl
                    publishableKey
                    private_key
                    walletPublicKey
                    default_enterprise
                    country
                    description
                    sellingPhysicalGoods
                    selfShippingProduct
                    shippingDelay
                    transactionLibele
                    abregedLibele
					phone {
						phone
						dialcode
					}
                    email
                    sector
                    rccm
                    website
                    person {
                        first_name
                        last_name
                        email
                        address
                        state
                        city
                        zip
                    }
            }
        }
	`,
		&_response,
		client.Var("enterpriseId", enterpriseId),
		client.Var("first_name", first_name),
		client.Var("last_name", last_name),
		client.Var("email", email),
		client.Var("address", address),
		client.Var("city", city),
		client.Var("state", state),
		client.Var("zip", zip),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func UpdateEnterpriseInformationTest(c *client.Client, options client.Option, idToken string, enterpriseId string, rccm string, sector string, website string, description string) (struct{ UpdateEnterpriseInformation []model.Enterprise }, error) {
	var _response struct{ UpdateEnterpriseInformation []model.Enterprise }
	err := c.Post(
		`
        mutation UpdateEnterpriseInformation($enterpriseId: String!, $rccm: String!, $sector: String!, $website: String, $description: String) {
            updateEnterpriseInformation(enterpriseId: $enterpriseId, rccm: $rccm, sector: $sector, website: $website, description: $description) {
                    _id
                    name
                    type
                    logoUrl
                    publishableKey
                    private_key
                    walletPublicKey
                    default_enterprise
                    country
                    description
                    sellingPhysicalGoods
                    selfShippingProduct
                    shippingDelay
                    transactionLibele
                    abregedLibele
					phone {
						phone
						dialcode
					}
                    email
                    sector
                    rccm
                    website
                    person {
                        first_name
                        last_name
                        email
                        address
                        state
                        city
                        zip
                    }
            }
        }
	`,
		&_response,
		client.Var("enterpriseId", enterpriseId),
		client.Var("rccm", rccm),
		client.Var("sector", sector),
		client.Var("website", website),
		client.Var("description", description),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func UpdateExecutionInformationTest(c *client.Client, options client.Option, idToken string, enterpriseId string, sellingPyshicalGoods bool, selfShipping bool, shippingDelay string) (struct{ UpdateExecutionInformation []model.Enterprise }, error) {
	var _response struct{ UpdateExecutionInformation []model.Enterprise }
	err := c.Post(
		`
        mutation UUpdateExecutionInformation($enterpriseId: String!, $sellingPyshicalGoods: Boolean, $selfShipping: Boolean, $shippingDelay: String) {
            updateExecutionInformation(enterpriseId: $enterpriseId, sellingPyshicalGoods: $sellingPyshicalGoods, selfShipping: $selfShipping, shippingDelay: $shippingDelay) {
                _id
                    name
                    type
                    logoUrl
                    walletPublicKey
                    default_enterprise
                    country
                    description
                    sellingPhysicalGoods
                    selfShippingProduct
                    shippingDelay
                    transactionLibele
                    abregedLibele
					phone {
						phone
						dialcode
					}
                    email
                    sector
                    rccm
                    website
                    person {
                        first_name
                        last_name
                        email
                        address
                        state
                        city
                        zip
                    }
            }
        }
	`,
		&_response,
		client.Var("enterpriseId", enterpriseId),
		client.Var("sellingPyshicalGoods", sellingPyshicalGoods),
		client.Var("selfShipping", selfShipping),
		client.Var("shippingDelay", shippingDelay),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func UpdatePublicInformationTest(c *client.Client, options client.Option, idToken string, enterpriseId string, name string, libelle string, libelleAbreged string, email string, phone model.PhoneInput) (struct{ UpdatePublicInformation []model.Enterprise }, error) {
	var _response struct{ UpdatePublicInformation []model.Enterprise }
	err := c.Post(
		`
        mutation UpdatePublicInformation($enterpriseId: String!, $name: String!, $libelle: String!, $libelleAbreged: String!, $email: String, $phone: PhoneInput!) {
            updatePublicInformation(enterpriseId: $enterpriseId, name: $name, libelle: $libelle, libelleAbreged: $libelleAbreged, email: $email, phone: $phone) {
                    _id
                    name
                    type
                    logoUrl
                    walletPublicKey
                    default_enterprise
                    country
                    description
                    sellingPhysicalGoods
                    selfShippingProduct
                    shippingDelay
                    transactionLibele
                    abregedLibele
					phone {
						phone
						dialcode
					}
                    email
                    sector
                    rccm
                    website
                    person {
                        first_name
                        last_name
                        email
                        address
                        state
                        city
                        zip
                    }
            }
        }
	`,
		&_response,
		client.Var("enterpriseId", enterpriseId),
		client.Var("name", name),
		client.Var("libelle", libelle),
		client.Var("libelleAbreged", libelleAbreged),
		client.Var("email", email),
		client.Var("phone", phone),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func GetProfilNetChartDataTest(c *client.Client, options client.Option, idToken string, enterpriseId string, from string, to string) (struct{ GetProfilNetChartData model.ChartData }, error) {
	var _response struct{ GetProfilNetChartData model.ChartData }
	err := c.Post(
		`
		query GetProfilNetChartData($enterpriseId: String!, $from: String!, $to: String!) {
			getProfilNetChartData(enterpriseId: $enterpriseId, from: $from, to: $to) {
				currentTotal
				formerTotal
				pourcentageDifference
				isPositive
				chart

			}
		}
	`,
		&_response,
		client.Var("enterpriseId", enterpriseId),
		client.Var("from", from),
		client.Var("to", to),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)
	return _response, err
}

func GetProfilBrutChartDataTest(c *client.Client, options client.Option, idToken string, enterpriseId string, from string, to string) (struct{ GetProfilBrutChartData model.ChartData }, error) {
	var _response struct{ GetProfilBrutChartData model.ChartData }
	err := c.Post(
		`
		query GetProfilBrutChartData($enterpriseId: String!, $from: String!, $to: String!) {
			getProfilBrutChartData(enterpriseId: $enterpriseId, from: $from, to: $to) {
				currentTotal
				formerTotal
				pourcentageDifference
				isPositive
				chart

			}
		}
	`,
		&_response,
		client.Var("enterpriseId", enterpriseId),
		client.Var("from", from),
		client.Var("to", to),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func GetProfilNonCarpturedChartDataTest(c *client.Client, options client.Option, idToken string, enterpriseId string, from string, to string) (struct{ GetProfilNonCarpturedChartData model.ChartData }, error) {
	var _response struct{ GetProfilNonCarpturedChartData model.ChartData }
	err := c.Post(
		`
		query GetProfilNonCarpturedChartData($enterpriseId: String!, $from: String!, $to: String!) {
			getProfilNonCarpturedChartData(enterpriseId: $enterpriseId, from: $from, to: $to) {
				currentTotal
				formerTotal
				pourcentageDifference
				isPositive
				chart
			}
		}
	`,
		&_response,
		client.Var("enterpriseId", enterpriseId),
		client.Var("from", from),
		client.Var("to", to),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func GetEnterpriseBalanceTest(c *client.Client, options client.Option, idToken string, enterpriseId string) (struct{ GetEnterpriseBalance float64 }, error) {
	var _response struct{ GetEnterpriseBalance float64 }
	err := c.Post(
		`
		query GetEnterpriseBalance($enterpriseId: String!) {
			getEnterpriseBalance(enterpriseId: $enterpriseId)
		}
	`,
		&_response,
		client.Var("enterpriseId", enterpriseId),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)
	return _response, err
}

func GetPdfTest(c *client.Client, options client.Option, idToken string, enterpriseId string) (struct{ GetPdf *string }, error) {
	var _response struct{ GetPdf *string }
	err := c.Post(
		`
		query  GetPdf($enterpriseId: String!) {
			getPdf(enterpriseId: $enterpriseId)
		}
	`,
		&_response,
		client.Var("enterpriseId", enterpriseId),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func GetAllTransactionByEnterpriseIdTest(c *client.Client, options client.Option, idToken string, enterpriseId string, from string, to string, limit int, skip int) (struct{ GetAllTransactionByEnterpriseId model.TransactionWithPageInfo }, error) {
	var _response struct{ GetAllTransactionByEnterpriseId model.TransactionWithPageInfo }
	err := c.Post(
		`
		query  GetAllTransactionByEnterpriseId($enterpriseId: String!, $from: String!, $to: String!, $limit: Float!, $skip: Float!) {
			getAllTransactionByEnterpriseId(enterpriseId: $enterpriseId, from: $from, to: $to, limit: $limit, skip: $skip) {
				transactions {
					_id
					status
					type
					creator{
						_id
					}
					token
					amount
					cancellor {
						_id
					}
					
					destination
					validator{
						_id
					}
					transactionId
					createdAt
					description
					updatedAt
					shortId
					destinationUser{
						_id
					}
					enterprise {
						_id
					}
				}
				pageTotal
			}
		}
	`,
		&_response,
		client.Var("enterpriseId", enterpriseId),
		client.Var("from", from),
		client.Var("to", to),
		client.Var("limit", limit),
		client.Var("skip", skip),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func GetSuccessFullTransactionByEnterpriseIdTest(c *client.Client, options client.Option, idToken string, enterpriseId string, from string, to string, limit int, skip int) (struct{ GetSuccessFullTransactionByEnterpriseId model.TransactionWithPageInfo }, error) {
	var _response struct{ GetSuccessFullTransactionByEnterpriseId model.TransactionWithPageInfo }
	err := c.Post(
		`
		query  GetSuccessFullTransactionByEnterpriseId($enterpriseId: String!, $from: String!, $to: String!, $limit: Float!, $skip: Float!) {
			getSuccessFullTransactionByEnterpriseId(enterpriseId: $enterpriseId, from: $from, to: $to, limit: $limit, skip: $skip) {
				transactions {
					_id
					status
					type
					creator{
						_id
					}
					token
					amount
					cancellor {
						_id
					}
					
					destination
					validator{
						_id
					}
					transactionId
					createdAt
					description
					updatedAt
					shortId
					destinationUser{
						_id
					}
					enterprise {
						_id
					}
				}
				pageTotal
			}
		}
	`,
		&_response,
		client.Var("enterpriseId", enterpriseId),
		client.Var("from", from),
		client.Var("to", to),
		client.Var("limit", limit),
		client.Var("skip", skip),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func GetRefundedTransactionByEnterpriseIdTest(c *client.Client, options client.Option, idToken string, enterpriseId string, from string, to string, limit int, skip int) (struct {
	GetRefundedTransactionByEnterpriseId model.TransactionWithPageInfo
}, error) {
	var _response struct {
		GetRefundedTransactionByEnterpriseId model.TransactionWithPageInfo
	}
	err := c.Post(
		`
		query  GetRefundedTransactionByEnterpriseId($enterpriseId: String!, $from: String!, $to: String!, $limit: Float!, $skip: Float!) {
			getRefundedTransactionByEnterpriseId(enterpriseId: $enterpriseId, from: $from, to: $to, limit: $limit, skip: $skip) {
				transactions {
					_id
					status
					type
					creator{
						_id
					}
					token
					amount
					cancellor {
						_id
					}
					
					destination
					validator{
						_id
					}
					transactionId
					createdAt
					description
					updatedAt
					shortId
					destinationUser{
						_id
					}
					enterprise {
						_id
					}
				}
				pageTotal
			}
		}
	`,
		&_response,
		client.Var("enterpriseId", enterpriseId),
		client.Var("from", from),
		client.Var("to", to),
		client.Var("limit", limit),
		client.Var("skip", skip),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func GetNonCapturedTransactionByEnterpriseIdTest(c *client.Client, options client.Option, idToken string, enterpriseId string, from string, to string, limit int, skip int) (struct {
	GetNonCapturedTransactionByEnterpriseId model.TransactionWithPageInfo
}, error) {
	var _response struct {
		GetNonCapturedTransactionByEnterpriseId model.TransactionWithPageInfo
	}
	err := c.Post(
		`
		query  GetNonCapturedTransactionByEnterpriseId($enterpriseId: String!, $from: String!, $to: String!, $limit: Float!, $skip: Float!) {
			getNonCapturedTransactionByEnterpriseId(enterpriseId: $enterpriseId, from: $from, to: $to, limit: $limit, skip: $skip) {
				transactions {
					_id
					status
					type
					creator{
						_id
					}
					token
					amount
					cancellor {
						_id
					}
					
					destination
					validator{
						_id
					}
					transactionId
					createdAt
					description
					updatedAt
					shortId
					destinationUser{
						_id
					}
					enterprise {
						_id
					}
				}
				pageTotal
			}
		}
	`,
		&_response,
		client.Var("enterpriseId", enterpriseId),
		client.Var("from", from),
		client.Var("to", to),
		client.Var("limit", limit),
		client.Var("skip", skip),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func GetFailedTransactionByEnterpriseIdTest(c *client.Client, options client.Option, idToken string, enterpriseId string, from string, to string, limit int, skip int) (struct{ GetFailedTransactionByEnterpriseId model.TransactionWithPageInfo }, error) {
	var _response struct{ GetFailedTransactionByEnterpriseId model.TransactionWithPageInfo }
	err := c.Post(
		`
		query  GetFailedTransactionByEnterpriseId($enterpriseId: String!, $from: String!, $to: String!, $limit: Float!, $skip: Float!) {
			getFailedTransactionByEnterpriseId(enterpriseId: $enterpriseId, from: $from, to: $to, limit: $limit, skip: $skip) {
				transactions {
					_id
					status
					type
					creator{
						_id
					}
					token
					amount
					cancellor {
						_id
					}
					
					destination
					validator{
						_id
					}
					transactionId
					createdAt
					description
					updatedAt
					shortId
					destinationUser{
						_id
					}
					enterprise {
						_id
					}
				}
				pageTotal
			}
		}
	`,
		&_response,
		client.Var("enterpriseId", enterpriseId),
		client.Var("from", from),
		client.Var("to", to),
		client.Var("limit", limit),
		client.Var("skip", skip),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func GetEnterpriseByIdTest(c *client.Client, options client.Option, idToken string, enterpriseId string) (struct{ GetEntrepriseById model.Enterprise }, error) {
	var _response struct{ GetEntrepriseById model.Enterprise }
	err := c.Post(
		`
        query  GetEntrepriseById($enterpriseId: String!) {
            getEntrepriseById(enterpriseId: $enterpriseId) {
                    _id
                    name
                    type
                    logoUrl
                    publishableKey
                    private_key
            }
        }
	`,
		&_response,
		client.Var("enterpriseId", enterpriseId),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func GetAllEnterpriseTest(c *client.Client, options client.Option, idToken string) (struct{ GetAllUserEnterprise []*model.Enterprise }, error) {
	var _response struct{ GetAllUserEnterprise []*model.Enterprise }
	err := c.Post(
		`
        query  GetAllUserEnterprise{
            getAllUserEnterprise {
					_id
					name
					type
					logoUrl
					publishableKey
					private_key
					walletPublicKey
					default_enterprise

            }
        }
	`,
		&_response,
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)
	return _response, err
}

func ChangeDefaultEnterpriseTest(c *client.Client, options client.Option, idToken string, enterpriseId string) (struct{ ChangeDefaultEnterprise []*model.Enterprise }, error) {
	var _response struct{ ChangeDefaultEnterprise []*model.Enterprise }
	err := c.Post(
		`
        mutation  ChangeDefaultEnterprise($enterpriseId: String!){
            changeDefaultEnterprise(enterpriseId: $enterpriseId) {
                    _id
                    name
                    type
                    logoUrl
                    publishableKey
                    private_key
                    walletPublicKey
                    default_enterprise

            }
        }
	`,
		&_response,
		client.Var("enterpriseId", enterpriseId),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func PayUnConfirmedTransactionTest(c *client.Client, options client.Option, idToken string, enterpriseId string, pinCode string, transactionId string) (struct{ PayUnConfirmedTransaction model.Paiement }, error) {
	var _response struct{ PayUnConfirmedTransaction model.Paiement }
	err := c.Post(
		`
        mutation PayUnConfirmedTransaction($enterpriseId: String!, $pinCode: String!, $transactionId: String!) {
            payUnConfirmedTransaction(enterpriseId: $enterpriseId, pinCode: $pinCode, transactionId: $transactionId) {
                _id
                status
                type
            }
        }
	`,
		&_response,
		client.Var("enterpriseId", enterpriseId),
		client.Var("pinCode", pinCode),
		client.Var("transactionId", transactionId),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func RefundTransactionTest(c *client.Client, options client.Option, idToken string, enterpriseId string, pinCode string, transactionId string) (struct{ RefundTransaction bool }, error) {
	var _response struct{ RefundTransaction bool }
	err := c.Post(
		`
		mutation RefundTransaction($enterpriseId: String!, $pinCode: String!, $transactionId: String!) {
			refundTransaction(enterpriseId: $enterpriseId, pinCode: $pinCode, transactionId: $transactionId)
		}
	`,
		&_response,
		client.Var("enterpriseId", enterpriseId),
		client.Var("pinCode", pinCode),
		client.Var("transactionId", transactionId),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func TransferMoneyEnterpriseTest(c *client.Client, options client.Option, idToken string, enterpriseId string, pinCode string, publicKey string, amount float32) (struct{ TransferMoneyEnterprise bool }, error) {
	var _response struct{ TransferMoneyEnterprise bool }
	err := c.Post(
		`
        mutation TransferMoneyEnterprise($enterpriseId: String!, $pinCode: String!, $publicKey: String!, $amount: Float!) {
            transferMoneyEnterprise(enterpriseId: $enterpriseId, pinCode: $pinCode, publicKey: $publicKey, amount: $amount)
        }
	`,
		&_response,
		client.Var("enterpriseId", enterpriseId),
		client.Var("pinCode", pinCode),
		client.Var("publicKey", publicKey),
		client.Var("amount", amount),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}
