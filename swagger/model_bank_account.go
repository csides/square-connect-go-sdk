/*
 * Square
 *
 * Use Square APIs to manage and run business including payment, customer, product, inventory, and employee management.
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Represents a bank account. For more information about  linking a bank account to a Square account, see  [Bank Accounts API](https://developer.squareup.com/docs/bank-accounts-api).
type BankAccount struct {
	// The unique, Square-issued identifier for the bank account.
	Id string `json:"id"`
	// The last few digits of the account number.
	AccountNumberSuffix string           `json:"account_number_suffix"`
	Country             *Country         `json:"country"`
	Currency            *Currency        `json:"currency"`
	AccountType         *BankAccountType `json:"account_type"`
	// Name of the account holder. This name must match the name  on the targeted bank account record.
	HolderName string `json:"holder_name"`
	// Primary identifier for the bank. For more information, see  [Bank Accounts API](https://developer.squareup.com/docs/bank-accounts-api).
	PrimaryBankIdentificationNumber string `json:"primary_bank_identification_number"`
	// Secondary identifier for the bank. For more information, see  [Bank Accounts API](https://developer.squareup.com/docs/bank-accounts-api).
	SecondaryBankIdentificationNumber string `json:"secondary_bank_identification_number,omitempty"`
	// Reference identifier that will be displayed to UK bank account owners when collecting direct debit authorization. Only required for UK bank accounts.
	DebitMandateReferenceId string `json:"debit_mandate_reference_id,omitempty"`
	// Client-provided identifier for linking the banking account to an entity in a third-party system (for example, a bank account number or a user identifier).
	ReferenceId string `json:"reference_id,omitempty"`
	// The location to which the bank account belongs.
	LocationId string             `json:"location_id,omitempty"`
	Status     *BankAccountStatus `json:"status"`
	// Indicates whether it is possible for Square to send money to this bank account.
	Creditable bool `json:"creditable"`
	// Indicates whether it is possible for Square to take money from this  bank account.
	Debitable bool `json:"debitable"`
	// A Square-assigned, unique identifier for the bank account based on the account information. The account fingerprint can be used to compare account entries and determine if the they represent the same real-world bank account.
	Fingerprint string `json:"fingerprint,omitempty"`
	// The current version of the `BankAccount`.
	Version int32 `json:"version,omitempty"`
	// Read only. Name of actual financial institution.  For example \"Bank of America\".
	BankName string `json:"bank_name,omitempty"`
}
