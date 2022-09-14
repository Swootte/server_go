package contact

import (
	"server/graph/model"

	"github.com/99designs/gqlgen/client"
)

func AddContactTest(c *client.Client, options client.Option, idToken string, contact string) (struct{ CreateContact bool }, error) {
	var _response struct{ CreateContact bool }
	err := c.Post(
		`
		mutation CreateContact($contact: String!) {
			createContact(contact: $contact)
		}
	`,
		&_response,
		client.Var("contact", contact),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)
	return _response, err
}

func LoadAllContactsTest(c *client.Client, options client.Option, idToken string, searchText string) (struct{ GetAllUserContact []model.UserSmall }, error) {
	var _response struct{ GetAllUserContact []model.UserSmall }
	err := c.Post(
		`
		query GetAllUserContact($searchText: String!){
			getAllUserContact(searchText: $searchText) {
				_id
				first_name
				last_name
			}
		}
	`,
		&_response,
		client.Var("searchText", searchText),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func RemoveContactTest(c *client.Client, options client.Option, idToken string, contact_id string) (struct{ RemoveContact bool }, error) {
	var _response struct{ RemoveContact bool }
	err := c.Post(
		`
		mutation RemoveContact($contact_id: String!) {
			removeContact(contact_id: $contact_id)
		}
	`,
		&_response,
		client.Var("contact_id", contact_id),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}

func LoadNonAddedContactTest(c *client.Client, options client.Option, idToken string, searchText string) (struct{ GetAllContactNotAdded []model.UserSmall }, error) {
	var _response struct{ GetAllContactNotAdded []model.UserSmall }
	err := c.Post(
		`
		query GetAllContactNotAdded($searchText: String!){
			getAllContactNotAdded(searchText: $searchText) {
				_id
				first_name
				last_name
				address

			}
		}
	`,
		&_response,
		client.Var("searchText", searchText),
		client.AddHeader("Authorization", "Bearer "+idToken),
		options,
	)

	return _response, err
}
