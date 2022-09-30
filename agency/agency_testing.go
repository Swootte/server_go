package agency

import (
	"server/graph/model"

	"github.com/99designs/gqlgen/client"
)

func CreateAgencyTest(c *client.Client, options client.Option, idToken string, pincode string, agency model.AgencyInpyt) (struct{ AddAGency string }, error) {
	var _response struct{ AddAGency string }
	err := c.Post(
		`
        mutation AddAGency($agency: AgencyInpyt!, $pinCode: String!) {
            addAGency(agency: $agency, pinCode: $pinCode)
        }
		`,
		&_response,
		client.Var("agency", agency),
		client.Var("pinCode", pincode),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)
	return _response, err
}

func LoadAlAgenciesTest(c *client.Client, options client.Option, idToken string) (struct{ RetrieveAllAgnecies []*model.Agency }, error) {
	var _response struct {
		RetrieveAllAgnecies []*model.Agency
	}
	err := c.Post(
		`
			query RetrieveAllAgnecies{
				retrieveAllAgnecies {
					_id
					title
				}
			}
		`,
		&_response,
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}
