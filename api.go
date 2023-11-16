package aircall

import (
	"encoding/json"
	"strconv"
)

// Ping API request
func (client *Client) Ping() (*PingResponse, error) {
	data, err := client.Get("/ping", map[string]string{})
	if err != nil {
		return nil, err
	}

	response := &PingResponse{}
	json.Unmarshal(data, &response)

	return response, nil
}

// Company API request
//
// res, err := client.Company()
func (client *Client) Company() (CompanyResponse, error) {
	data, err := client.Get("/company", map[string]string{})
	response := CompanyResponse{}

	if err != nil {
		return response, err
	}

	json.Unmarshal(data, &response)

	return response, nil
}

// Users API request
func (client *Client) Users(paginate Paginate) (UsersResponse, error) {
	params := buildPaginateParamsMap(paginate)

	data, err := client.Get("/users", params)
	response := UsersResponse{}

	if err != nil {
		return response, err
	}

	json.Unmarshal(data, &response)

	return response, nil
}

// User API request
func (client *Client) User(ID int) (UserResponse, error) {
	data, err := client.Get("/users/"+strconv.Itoa(ID), map[string]string{})
	response := UserResponse{}

	if err != nil {
		return response, err
	}

	json.Unmarshal(data, &response)

	return response, nil
}

// Numbers API request
func (client *Client) Numbers(paginate Paginate) (NumbersResponse, error) {
	params := buildPaginateParamsMap(paginate)
	data, err := client.Get("/numbers", params)
	response := NumbersResponse{}

	if err != nil {
		return response, err
	}

	json.Unmarshal(data, &response)

	return response, nil
}

// Number API request
func (client *Client) Number(ID int) (NumberResponse, error) {
	data, err := client.Get("/numbers/"+strconv.Itoa(ID), map[string]string{})
	response := NumberResponse{}

	if err != nil {
		return response, err
	}

	json.Unmarshal(data, &response)

	return response, nil
}

// Calls API request
func (client *Client) Calls(paginate Paginate) (CallsResponse, error) {
	params := buildPaginateParamsMap(paginate)
	data, err := client.Get("/calls", params)
	response := CallsResponse{}

	if err != nil {
		return response, err
	}

	json.Unmarshal(data, &response)

	return response, nil
}

// Call API request
func (client *Client) Call(ID int) (CallResponse, error) {
	data, err := client.Get("/calls/"+strconv.Itoa(ID), map[string]string{})
	response := CallResponse{}

	if err != nil {
		return response, err
	}

	json.Unmarshal(data, &response)

	return response, nil
}

// SearchCalls API request
func (client *Client) SearchCalls(paginate Paginate, search Search) (CallsResponse, error) {
	params := buildPaginateParamsMap(paginate)
	searchParams := buildSearchParamsMap(search)
	for k, v := range searchParams {
		params[k] = v
	}

	data, err := client.Get("/calls", params)
	response := CallsResponse{}

	if err != nil {
		return response, err
	}

	json.Unmarshal(data, &response)

	return response, nil
}

// LinkCall API Request
func (client *Client) LinkCall(ID int, link string) (CallResponse, error) {
	data, err := client.Post("/calls/"+strconv.Itoa(ID)+"/link", LinkCallRequest{
		Link: link,
	})
	response := CallResponse{}

	if err != nil {
		return response, err
	}

	json.Unmarshal(data, &response)

	return response, nil
}

// TransferCall API request
func (client *Client) TransferCall(ID int, userID int) (CallResponse, error) {
	data, err := client.Post("/calls/"+strconv.Itoa(ID)+"/transfers", TransferCallRequest{
		UserID: userID,
	})
	response := CallResponse{}

	if err != nil {
		return response, err
	}

	json.Unmarshal(data, &response)

	return response, nil
}

// CreateInsighCards API request
func (client *Client) CreateInsighCards(ID int, insightCards CallInsightCardRequest) error {
	_, err := client.Post("/calls/"+strconv.Itoa(ID)+"/insight_cards", insightCards)

	if err != nil {
		return err
	}

	return nil
}

// DeleteRecording API request
func (client *Client) DeleteRecording(callID int) (Response, error) {
	data, err := client.Delete("/calls/"+strconv.Itoa(callID)+"/recording", map[string]string{})
	response := Response{}

	if err != nil {
		return response, err
	}

	json.Unmarshal(data, &response)

	return response, nil
}

// DeleteVoicemail API request
func (client *Client) DeleteVoicemail(callID int) (Response, error) {
	data, err := client.Delete("/calls/"+strconv.Itoa(callID)+"/voicemail", map[string]string{})
	response := Response{}

	if err != nil {
		return response, err
	}

	json.Unmarshal(data, &response)

	return response, nil
}

// Contacts API request
func (client *Client) Contacts(paginate Paginate) (ContactsResponse, error) {
	params := buildPaginateParamsMap(paginate)
	data, err := client.Get("/contacts", params)
	response := ContactsResponse{}

	if err != nil {
		return response, err
	}

	json.Unmarshal(data, &response)

	return response, nil
}

// Contact API request
func (client *Client) Contact(ID int) (ContactResponse, error) {
	data, err := client.Get("/contacts/"+strconv.Itoa(ID), map[string]string{})
	response := ContactResponse{}

	if err != nil {
		return response, err
	}

	json.Unmarshal(data, &response)

	return response, nil
}

// CreateContact API request
func (client *Client) CreateContact(contact ContactRequest) (ContactResponse, error) {
	data, err := client.Post("/contacts", contact)
	response := ContactResponse{}

	if err != nil {
		return response, err
	}

	json.Unmarshal(data, &response)

	return response, nil
}

// UpdateContact API request
func (client *Client) UpdateContact(ID int, contact ContactRequest) (ContactResponse, error) {
	data, err := client.Post("/contacts/"+strconv.Itoa(ID), contact)
	response := ContactResponse{}

	if err != nil {
		return response, err
	}

	json.Unmarshal(data, &response)

	return response, nil
}

// SearchContacts API request
func (client *Client) SearchContacts(paginate Paginate, search Search) (ContactsResponse, error) {
	params := buildPaginateParamsMap(paginate)
	searchParams := buildSearchParamsMap(search)
	for k, v := range searchParams {
		params[k] = v
	}

	data, err := client.Get("/contacts/search", params)
	response := ContactsResponse{}

	if err != nil {
		return response, err
	}

	json.Unmarshal(data, &response)

	return response, nil
}

// DeleteContact API request
func (client *Client) DeleteContact(ID int) (Response, error) {
	data, err := client.Delete("/contacts/"+strconv.Itoa(ID), map[string]string{})
	response := Response{}

	if err != nil {
		return response, err
	}

	json.Unmarshal(data, &response)

	return response, nil
}

// Webhooks API request
func (client *Client) Webhooks(paginate Paginate) (WebhooksResponse, error) {
	params := buildPaginateParamsMap(paginate)
	data, err := client.Get("/webhooks", params)

	response := WebhooksResponse{}

	if err != nil {
		return response, err
	}

	json.Unmarshal(data, &response)

	return response, nil
}

// CreateWebhook API request
func (client *Client) CreateWebhook(webhook WebhookRequest) (WebhookResponse, error) {
	data, err := client.Post("/webhooks", webhook)
	response := WebhookResponse{}

	if err != nil {
		return response, err
	}

	json.Unmarshal(data, &response)

	return response, nil
}

// DeleteWebhook API request
func (client *Client) DeleteWebhook(ID string) (Response, error) {
	data, err := client.Delete("/webhooks/" + ID, map[string]string{})
	response := Response{}

	if err != nil {
		return response, err
	}

	json.Unmarshal(data, &response)

	return response, nil
}

func buildSearchParamsMap(search Search) map[string]string {
	params := make(map[string]string)

	if search.PhoneNumber != "" {
		params["phone_number"] = search.PhoneNumber
	}
	if search.Email != "" {
		params["email"] = search.Email
	}

	return params
}

func buildPaginateParamsMap(paginate Paginate) map[string]string {
	params := make(map[string]string)

	if paginate.Page > 0 {
		params["page"] = strconv.Itoa(paginate.Page)
	}
	if paginate.PerPage > 0 {
		params["per_page"] = strconv.Itoa(paginate.PerPage)
	}
	if paginate.Order != "" {
		params["order"] = paginate.Order
	}
	if paginate.From > 0 {
		params["from"] = strconv.Itoa(paginate.From)
	}
	if paginate.To > 0 {
		params["to"] = strconv.Itoa(paginate.To)
	}

	return params
}
