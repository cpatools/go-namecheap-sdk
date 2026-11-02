package namecheap

import (
	"context"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

// addressGetInfoDefaultResponse is the shape the API returns for the default
// address: the id is 0 and the flag element is named Default_YN.
const addressGetInfoDefaultResponse = `<?xml version="1.0" encoding="utf-8"?>
<ApiResponse Status="OK" xmlns="http://api.namecheap.com/xml.response">
	<Errors />
	<CommandResponse Type="namecheap.users.address.getInfo">
		<GetAddressInfoResult>
			<AddressId>0</AddressId>
			<AddressName>Main Address</AddressName>
			<Default_YN>true</Default_YN>
			<Country>US</Country>
		</GetAddressInfoResult>
	</CommandResponse>
</ApiResponse>`

func TestUsersAddressService_GetInfo(t *testing.T) {
	t.Parallel()

	t.Run("success parse and request", func(t *testing.T) {
		t.Parallel()
		var sent url.Values
		client := usersMockClient(t, addressGetInfoResponse, &sent)

		resp, err := client.UsersAddress.GetInfoWithContext(context.Background(), NewAddressID(777))
		mustNoError(t, err)

		assert.Equal(t, "namecheap.users.address.getInfo", sent.Get("Command"))
		assert.Equal(t, "777", sent.Get("AddressId"))

		got := resp.GetAddressInfoResult
		mustNotNil(t, got)
		assert.Equal(t, "Home", *got.AddressName)
		assert.Equal(t, "85284", *got.Zip)
		assert.Equal(t, "NameCheap.com", *got.Organization)
	})

	t.Run("zero id is a valid id", func(t *testing.T) {
		t.Parallel()
		var sent url.Values
		client := usersMockClient(t, addressGetInfoResponse, &sent)

		_, err := client.UsersAddress.GetInfoWithContext(context.Background(), NewAddressID(0))
		mustNoError(t, err)

		assert.Equal(t, "0", sent.Get("AddressId"))
	})

	t.Run("unset id no http", func(t *testing.T) {
		t.Parallel()
		client := setupClient(nil)
		_, err := client.UsersAddress.GetInfoWithContext(context.Background(), AddressID{})
		assertInvalidArguments(t, err, "AddressId")
	})

	t.Run("negative id no http", func(t *testing.T) {
		t.Parallel()
		client := setupClient(nil)
		_, err := client.UsersAddress.GetInfoWithContext(context.Background(), NewAddressID(-1))
		assertInvalidArguments(t, err, "AddressId")
	})

	t.Run("default flag is parsed from Default_YN", func(t *testing.T) {
		t.Parallel()
		client := usersMockClient(t, addressGetInfoDefaultResponse, nil)

		resp, err := client.UsersAddress.GetInfoWithContext(context.Background(), NewAddressID(0))
		mustNoError(t, err)

		got := resp.GetAddressInfoResult
		mustNotNil(t, got)
		mustNotNil(t, got.DefaultYN)
		assert.True(t, *got.DefaultYN)
	})

	t.Run("api error mapped to APIError", func(t *testing.T) {
		t.Parallel()
		client := usersMockClient(t, apiErrorXML(4022336, "Failed to retrieve user's address"), nil)
		_, err := client.UsersAddress.GetInfoWithContext(context.Background(), NewAddressID(777))
		assertAPIError(t, err, 4022336)
	})
}
