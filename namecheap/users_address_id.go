package namecheap

// AddressID identifies an address-book entry as returned by address.getList.
// Its zero value is "not set" and is rejected by the methods that take it, so a
// forgotten id is never sent as 0.
type AddressID struct {
	id  int
	set bool
}

// NewAddressID wraps an address id. Valid ids are not negative; 0 is valid.
func NewAddressID(id int) AddressID { return AddressID{id: id, set: true} }

// IsSet reports whether the id was built by NewAddressID.
func (a AddressID) IsSet() bool { return a.set }

// Int returns the numeric id.
func (a AddressID) Int() int { return a.id }
