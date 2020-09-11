// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Address information for domain registration.
type Address struct {
	// First line of an Address.
	Address1 string `pulumi:"address1"`
	// The second line of the Address. Optional.
	Address2 *string `pulumi:"address2"`
	// The city for the address.
	City string `pulumi:"city"`
	// The country for the address.
	Country string `pulumi:"country"`
	// The postal code for the address.
	PostalCode string `pulumi:"postalCode"`
	// The state or province for the address.
	State string `pulumi:"state"`
}

// AddressInput is an input type that accepts AddressArgs and AddressOutput values.
// You can construct a concrete instance of `AddressInput` via:
//
//          AddressArgs{...}
type AddressInput interface {
	pulumi.Input

	ToAddressOutput() AddressOutput
	ToAddressOutputWithContext(context.Context) AddressOutput
}

// Address information for domain registration.
type AddressArgs struct {
	// First line of an Address.
	Address1 pulumi.StringInput `pulumi:"address1"`
	// The second line of the Address. Optional.
	Address2 pulumi.StringPtrInput `pulumi:"address2"`
	// The city for the address.
	City pulumi.StringInput `pulumi:"city"`
	// The country for the address.
	Country pulumi.StringInput `pulumi:"country"`
	// The postal code for the address.
	PostalCode pulumi.StringInput `pulumi:"postalCode"`
	// The state or province for the address.
	State pulumi.StringInput `pulumi:"state"`
}

func (AddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Address)(nil)).Elem()
}

func (i AddressArgs) ToAddressOutput() AddressOutput {
	return i.ToAddressOutputWithContext(context.Background())
}

func (i AddressArgs) ToAddressOutputWithContext(ctx context.Context) AddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressOutput)
}

func (i AddressArgs) ToAddressPtrOutput() AddressPtrOutput {
	return i.ToAddressPtrOutputWithContext(context.Background())
}

func (i AddressArgs) ToAddressPtrOutputWithContext(ctx context.Context) AddressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressOutput).ToAddressPtrOutputWithContext(ctx)
}

// AddressPtrInput is an input type that accepts AddressArgs, AddressPtr and AddressPtrOutput values.
// You can construct a concrete instance of `AddressPtrInput` via:
//
//          AddressArgs{...}
//
//  or:
//
//          nil
type AddressPtrInput interface {
	pulumi.Input

	ToAddressPtrOutput() AddressPtrOutput
	ToAddressPtrOutputWithContext(context.Context) AddressPtrOutput
}

type addressPtrType AddressArgs

func AddressPtr(v *AddressArgs) AddressPtrInput {
	return (*addressPtrType)(v)
}

func (*addressPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Address)(nil)).Elem()
}

func (i *addressPtrType) ToAddressPtrOutput() AddressPtrOutput {
	return i.ToAddressPtrOutputWithContext(context.Background())
}

func (i *addressPtrType) ToAddressPtrOutputWithContext(ctx context.Context) AddressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressPtrOutput)
}

// Address information for domain registration.
type AddressResponse struct {
	// First line of an Address.
	Address1 string `pulumi:"address1"`
	// The second line of the Address. Optional.
	Address2 *string `pulumi:"address2"`
	// The city for the address.
	City string `pulumi:"city"`
	// The country for the address.
	Country string `pulumi:"country"`
	// The postal code for the address.
	PostalCode string `pulumi:"postalCode"`
	// The state or province for the address.
	State string `pulumi:"state"`
}

// Address information for domain registration.
type AddressResponseOutput struct{ *pulumi.OutputState }

func (AddressResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressResponse)(nil)).Elem()
}

func (o AddressResponseOutput) ToAddressResponseOutput() AddressResponseOutput {
	return o
}

func (o AddressResponseOutput) ToAddressResponseOutputWithContext(ctx context.Context) AddressResponseOutput {
	return o
}

func (o AddressResponseOutput) ToAddressResponsePtrOutput() AddressResponsePtrOutput {
	return o.ToAddressResponsePtrOutputWithContext(context.Background())
}

func (o AddressResponseOutput) ToAddressResponsePtrOutputWithContext(ctx context.Context) AddressResponsePtrOutput {
	return o.ApplyT(func(v AddressResponse) *AddressResponse {
		return &v
	}).(AddressResponsePtrOutput)
}

// First line of an Address.
func (o AddressResponseOutput) Address1() pulumi.StringOutput {
	return o.ApplyT(func(v AddressResponse) string { return v.Address1 }).(pulumi.StringOutput)
}

// The second line of the Address. Optional.
func (o AddressResponseOutput) Address2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AddressResponse) *string { return v.Address2 }).(pulumi.StringPtrOutput)
}

// The city for the address.
func (o AddressResponseOutput) City() pulumi.StringOutput {
	return o.ApplyT(func(v AddressResponse) string { return v.City }).(pulumi.StringOutput)
}

// The country for the address.
func (o AddressResponseOutput) Country() pulumi.StringOutput {
	return o.ApplyT(func(v AddressResponse) string { return v.Country }).(pulumi.StringOutput)
}

// The postal code for the address.
func (o AddressResponseOutput) PostalCode() pulumi.StringOutput {
	return o.ApplyT(func(v AddressResponse) string { return v.PostalCode }).(pulumi.StringOutput)
}

// The state or province for the address.
func (o AddressResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v AddressResponse) string { return v.State }).(pulumi.StringOutput)
}

type AddressResponsePtrOutput struct{ *pulumi.OutputState }

func (AddressResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressResponse)(nil)).Elem()
}

func (o AddressResponsePtrOutput) ToAddressResponsePtrOutput() AddressResponsePtrOutput {
	return o
}

func (o AddressResponsePtrOutput) ToAddressResponsePtrOutputWithContext(ctx context.Context) AddressResponsePtrOutput {
	return o
}

func (o AddressResponsePtrOutput) Elem() AddressResponseOutput {
	return o.ApplyT(func(v *AddressResponse) AddressResponse { return *v }).(AddressResponseOutput)
}

// First line of an Address.
func (o AddressResponsePtrOutput) Address1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AddressResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Address1
	}).(pulumi.StringPtrOutput)
}

// The second line of the Address. Optional.
func (o AddressResponsePtrOutput) Address2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.Address2
	}).(pulumi.StringPtrOutput)
}

// The city for the address.
func (o AddressResponsePtrOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AddressResponse) *string {
		if v == nil {
			return nil
		}
		return &v.City
	}).(pulumi.StringPtrOutput)
}

// The country for the address.
func (o AddressResponsePtrOutput) Country() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AddressResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Country
	}).(pulumi.StringPtrOutput)
}

// The postal code for the address.
func (o AddressResponsePtrOutput) PostalCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AddressResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PostalCode
	}).(pulumi.StringPtrOutput)
}

// The state or province for the address.
func (o AddressResponsePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AddressResponse) *string {
		if v == nil {
			return nil
		}
		return &v.State
	}).(pulumi.StringPtrOutput)
}

// Contact information for domain registration. If 'Domain Privacy' option is not selected then the contact information is made publicly available through the Whois
// directories as per ICANN requirements.
type Contact struct {
	// Mailing address.
	AddressMailing *Address `pulumi:"addressMailing"`
	// Email address.
	Email string `pulumi:"email"`
	// Fax number.
	Fax *string `pulumi:"fax"`
	// Job title.
	JobTitle *string `pulumi:"jobTitle"`
	// First name.
	NameFirst string `pulumi:"nameFirst"`
	// Last name.
	NameLast string `pulumi:"nameLast"`
	// Middle name.
	NameMiddle *string `pulumi:"nameMiddle"`
	// Organization contact belongs to.
	Organization *string `pulumi:"organization"`
	// Phone number.
	Phone string `pulumi:"phone"`
}

// ContactInput is an input type that accepts ContactArgs and ContactOutput values.
// You can construct a concrete instance of `ContactInput` via:
//
//          ContactArgs{...}
type ContactInput interface {
	pulumi.Input

	ToContactOutput() ContactOutput
	ToContactOutputWithContext(context.Context) ContactOutput
}

// Contact information for domain registration. If 'Domain Privacy' option is not selected then the contact information is made publicly available through the Whois
// directories as per ICANN requirements.
type ContactArgs struct {
	// Mailing address.
	AddressMailing AddressPtrInput `pulumi:"addressMailing"`
	// Email address.
	Email pulumi.StringInput `pulumi:"email"`
	// Fax number.
	Fax pulumi.StringPtrInput `pulumi:"fax"`
	// Job title.
	JobTitle pulumi.StringPtrInput `pulumi:"jobTitle"`
	// First name.
	NameFirst pulumi.StringInput `pulumi:"nameFirst"`
	// Last name.
	NameLast pulumi.StringInput `pulumi:"nameLast"`
	// Middle name.
	NameMiddle pulumi.StringPtrInput `pulumi:"nameMiddle"`
	// Organization contact belongs to.
	Organization pulumi.StringPtrInput `pulumi:"organization"`
	// Phone number.
	Phone pulumi.StringInput `pulumi:"phone"`
}

func (ContactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Contact)(nil)).Elem()
}

func (i ContactArgs) ToContactOutput() ContactOutput {
	return i.ToContactOutputWithContext(context.Background())
}

func (i ContactArgs) ToContactOutputWithContext(ctx context.Context) ContactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactOutput)
}

func (i ContactArgs) ToContactPtrOutput() ContactPtrOutput {
	return i.ToContactPtrOutputWithContext(context.Background())
}

func (i ContactArgs) ToContactPtrOutputWithContext(ctx context.Context) ContactPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactOutput).ToContactPtrOutputWithContext(ctx)
}

// ContactPtrInput is an input type that accepts ContactArgs, ContactPtr and ContactPtrOutput values.
// You can construct a concrete instance of `ContactPtrInput` via:
//
//          ContactArgs{...}
//
//  or:
//
//          nil
type ContactPtrInput interface {
	pulumi.Input

	ToContactPtrOutput() ContactPtrOutput
	ToContactPtrOutputWithContext(context.Context) ContactPtrOutput
}

type contactPtrType ContactArgs

func ContactPtr(v *ContactArgs) ContactPtrInput {
	return (*contactPtrType)(v)
}

func (*contactPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Contact)(nil)).Elem()
}

func (i *contactPtrType) ToContactPtrOutput() ContactPtrOutput {
	return i.ToContactPtrOutputWithContext(context.Background())
}

func (i *contactPtrType) ToContactPtrOutputWithContext(ctx context.Context) ContactPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactPtrOutput)
}

// Contact information for domain registration. If 'Domain Privacy' option is not selected then the contact information is made publicly available through the Whois
// directories as per ICANN requirements.
type ContactResponse struct {
	// Mailing address.
	AddressMailing *AddressResponse `pulumi:"addressMailing"`
	// Email address.
	Email string `pulumi:"email"`
	// Fax number.
	Fax *string `pulumi:"fax"`
	// Job title.
	JobTitle *string `pulumi:"jobTitle"`
	// First name.
	NameFirst string `pulumi:"nameFirst"`
	// Last name.
	NameLast string `pulumi:"nameLast"`
	// Middle name.
	NameMiddle *string `pulumi:"nameMiddle"`
	// Organization contact belongs to.
	Organization *string `pulumi:"organization"`
	// Phone number.
	Phone string `pulumi:"phone"`
}

// Contact information for domain registration. If 'Domain Privacy' option is not selected then the contact information is made publicly available through the Whois
// directories as per ICANN requirements.
type ContactResponseOutput struct{ *pulumi.OutputState }

func (ContactResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContactResponse)(nil)).Elem()
}

func (o ContactResponseOutput) ToContactResponseOutput() ContactResponseOutput {
	return o
}

func (o ContactResponseOutput) ToContactResponseOutputWithContext(ctx context.Context) ContactResponseOutput {
	return o
}

func (o ContactResponseOutput) ToContactResponsePtrOutput() ContactResponsePtrOutput {
	return o.ToContactResponsePtrOutputWithContext(context.Background())
}

func (o ContactResponseOutput) ToContactResponsePtrOutputWithContext(ctx context.Context) ContactResponsePtrOutput {
	return o.ApplyT(func(v ContactResponse) *ContactResponse {
		return &v
	}).(ContactResponsePtrOutput)
}

// Mailing address.
func (o ContactResponseOutput) AddressMailing() AddressResponsePtrOutput {
	return o.ApplyT(func(v ContactResponse) *AddressResponse { return v.AddressMailing }).(AddressResponsePtrOutput)
}

// Email address.
func (o ContactResponseOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v ContactResponse) string { return v.Email }).(pulumi.StringOutput)
}

// Fax number.
func (o ContactResponseOutput) Fax() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactResponse) *string { return v.Fax }).(pulumi.StringPtrOutput)
}

// Job title.
func (o ContactResponseOutput) JobTitle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactResponse) *string { return v.JobTitle }).(pulumi.StringPtrOutput)
}

// First name.
func (o ContactResponseOutput) NameFirst() pulumi.StringOutput {
	return o.ApplyT(func(v ContactResponse) string { return v.NameFirst }).(pulumi.StringOutput)
}

// Last name.
func (o ContactResponseOutput) NameLast() pulumi.StringOutput {
	return o.ApplyT(func(v ContactResponse) string { return v.NameLast }).(pulumi.StringOutput)
}

// Middle name.
func (o ContactResponseOutput) NameMiddle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactResponse) *string { return v.NameMiddle }).(pulumi.StringPtrOutput)
}

// Organization contact belongs to.
func (o ContactResponseOutput) Organization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContactResponse) *string { return v.Organization }).(pulumi.StringPtrOutput)
}

// Phone number.
func (o ContactResponseOutput) Phone() pulumi.StringOutput {
	return o.ApplyT(func(v ContactResponse) string { return v.Phone }).(pulumi.StringOutput)
}

type ContactResponsePtrOutput struct{ *pulumi.OutputState }

func (ContactResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContactResponse)(nil)).Elem()
}

func (o ContactResponsePtrOutput) ToContactResponsePtrOutput() ContactResponsePtrOutput {
	return o
}

func (o ContactResponsePtrOutput) ToContactResponsePtrOutputWithContext(ctx context.Context) ContactResponsePtrOutput {
	return o
}

func (o ContactResponsePtrOutput) Elem() ContactResponseOutput {
	return o.ApplyT(func(v *ContactResponse) ContactResponse { return *v }).(ContactResponseOutput)
}

// Mailing address.
func (o ContactResponsePtrOutput) AddressMailing() AddressResponsePtrOutput {
	return o.ApplyT(func(v *ContactResponse) *AddressResponse {
		if v == nil {
			return nil
		}
		return v.AddressMailing
	}).(AddressResponsePtrOutput)
}

// Email address.
func (o ContactResponsePtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContactResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Email
	}).(pulumi.StringPtrOutput)
}

// Fax number.
func (o ContactResponsePtrOutput) Fax() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContactResponse) *string {
		if v == nil {
			return nil
		}
		return v.Fax
	}).(pulumi.StringPtrOutput)
}

// Job title.
func (o ContactResponsePtrOutput) JobTitle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContactResponse) *string {
		if v == nil {
			return nil
		}
		return v.JobTitle
	}).(pulumi.StringPtrOutput)
}

// First name.
func (o ContactResponsePtrOutput) NameFirst() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContactResponse) *string {
		if v == nil {
			return nil
		}
		return &v.NameFirst
	}).(pulumi.StringPtrOutput)
}

// Last name.
func (o ContactResponsePtrOutput) NameLast() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContactResponse) *string {
		if v == nil {
			return nil
		}
		return &v.NameLast
	}).(pulumi.StringPtrOutput)
}

// Middle name.
func (o ContactResponsePtrOutput) NameMiddle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContactResponse) *string {
		if v == nil {
			return nil
		}
		return v.NameMiddle
	}).(pulumi.StringPtrOutput)
}

// Organization contact belongs to.
func (o ContactResponsePtrOutput) Organization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContactResponse) *string {
		if v == nil {
			return nil
		}
		return v.Organization
	}).(pulumi.StringPtrOutput)
}

// Phone number.
func (o ContactResponsePtrOutput) Phone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContactResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Phone
	}).(pulumi.StringPtrOutput)
}

// Domain purchase consent object, representing acceptance of applicable legal agreements.
type DomainPurchaseConsent struct {
	// Timestamp when the agreements were accepted.
	AgreedAt *string `pulumi:"agreedAt"`
	// Client IP address.
	AgreedBy *string `pulumi:"agreedBy"`
	// List of applicable legal agreement keys. This list can be retrieved using ListLegalAgreements API under <code>TopLevelDomain</code> resource.
	AgreementKeys []string `pulumi:"agreementKeys"`
}

// DomainPurchaseConsentInput is an input type that accepts DomainPurchaseConsentArgs and DomainPurchaseConsentOutput values.
// You can construct a concrete instance of `DomainPurchaseConsentInput` via:
//
//          DomainPurchaseConsentArgs{...}
type DomainPurchaseConsentInput interface {
	pulumi.Input

	ToDomainPurchaseConsentOutput() DomainPurchaseConsentOutput
	ToDomainPurchaseConsentOutputWithContext(context.Context) DomainPurchaseConsentOutput
}

// Domain purchase consent object, representing acceptance of applicable legal agreements.
type DomainPurchaseConsentArgs struct {
	// Timestamp when the agreements were accepted.
	AgreedAt pulumi.StringPtrInput `pulumi:"agreedAt"`
	// Client IP address.
	AgreedBy pulumi.StringPtrInput `pulumi:"agreedBy"`
	// List of applicable legal agreement keys. This list can be retrieved using ListLegalAgreements API under <code>TopLevelDomain</code> resource.
	AgreementKeys pulumi.StringArrayInput `pulumi:"agreementKeys"`
}

func (DomainPurchaseConsentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainPurchaseConsent)(nil)).Elem()
}

func (i DomainPurchaseConsentArgs) ToDomainPurchaseConsentOutput() DomainPurchaseConsentOutput {
	return i.ToDomainPurchaseConsentOutputWithContext(context.Background())
}

func (i DomainPurchaseConsentArgs) ToDomainPurchaseConsentOutputWithContext(ctx context.Context) DomainPurchaseConsentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainPurchaseConsentOutput)
}

func (i DomainPurchaseConsentArgs) ToDomainPurchaseConsentPtrOutput() DomainPurchaseConsentPtrOutput {
	return i.ToDomainPurchaseConsentPtrOutputWithContext(context.Background())
}

func (i DomainPurchaseConsentArgs) ToDomainPurchaseConsentPtrOutputWithContext(ctx context.Context) DomainPurchaseConsentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainPurchaseConsentOutput).ToDomainPurchaseConsentPtrOutputWithContext(ctx)
}

// DomainPurchaseConsentPtrInput is an input type that accepts DomainPurchaseConsentArgs, DomainPurchaseConsentPtr and DomainPurchaseConsentPtrOutput values.
// You can construct a concrete instance of `DomainPurchaseConsentPtrInput` via:
//
//          DomainPurchaseConsentArgs{...}
//
//  or:
//
//          nil
type DomainPurchaseConsentPtrInput interface {
	pulumi.Input

	ToDomainPurchaseConsentPtrOutput() DomainPurchaseConsentPtrOutput
	ToDomainPurchaseConsentPtrOutputWithContext(context.Context) DomainPurchaseConsentPtrOutput
}

type domainPurchaseConsentPtrType DomainPurchaseConsentArgs

func DomainPurchaseConsentPtr(v *DomainPurchaseConsentArgs) DomainPurchaseConsentPtrInput {
	return (*domainPurchaseConsentPtrType)(v)
}

func (*domainPurchaseConsentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainPurchaseConsent)(nil)).Elem()
}

func (i *domainPurchaseConsentPtrType) ToDomainPurchaseConsentPtrOutput() DomainPurchaseConsentPtrOutput {
	return i.ToDomainPurchaseConsentPtrOutputWithContext(context.Background())
}

func (i *domainPurchaseConsentPtrType) ToDomainPurchaseConsentPtrOutputWithContext(ctx context.Context) DomainPurchaseConsentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainPurchaseConsentPtrOutput)
}

// Domain purchase consent object, representing acceptance of applicable legal agreements.
type DomainPurchaseConsentResponse struct {
	// Timestamp when the agreements were accepted.
	AgreedAt *string `pulumi:"agreedAt"`
	// Client IP address.
	AgreedBy *string `pulumi:"agreedBy"`
	// List of applicable legal agreement keys. This list can be retrieved using ListLegalAgreements API under <code>TopLevelDomain</code> resource.
	AgreementKeys []string `pulumi:"agreementKeys"`
}

// Domain purchase consent object, representing acceptance of applicable legal agreements.
type DomainPurchaseConsentResponseOutput struct{ *pulumi.OutputState }

func (DomainPurchaseConsentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainPurchaseConsentResponse)(nil)).Elem()
}

func (o DomainPurchaseConsentResponseOutput) ToDomainPurchaseConsentResponseOutput() DomainPurchaseConsentResponseOutput {
	return o
}

func (o DomainPurchaseConsentResponseOutput) ToDomainPurchaseConsentResponseOutputWithContext(ctx context.Context) DomainPurchaseConsentResponseOutput {
	return o
}

func (o DomainPurchaseConsentResponseOutput) ToDomainPurchaseConsentResponsePtrOutput() DomainPurchaseConsentResponsePtrOutput {
	return o.ToDomainPurchaseConsentResponsePtrOutputWithContext(context.Background())
}

func (o DomainPurchaseConsentResponseOutput) ToDomainPurchaseConsentResponsePtrOutputWithContext(ctx context.Context) DomainPurchaseConsentResponsePtrOutput {
	return o.ApplyT(func(v DomainPurchaseConsentResponse) *DomainPurchaseConsentResponse {
		return &v
	}).(DomainPurchaseConsentResponsePtrOutput)
}

// Timestamp when the agreements were accepted.
func (o DomainPurchaseConsentResponseOutput) AgreedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainPurchaseConsentResponse) *string { return v.AgreedAt }).(pulumi.StringPtrOutput)
}

// Client IP address.
func (o DomainPurchaseConsentResponseOutput) AgreedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainPurchaseConsentResponse) *string { return v.AgreedBy }).(pulumi.StringPtrOutput)
}

// List of applicable legal agreement keys. This list can be retrieved using ListLegalAgreements API under <code>TopLevelDomain</code> resource.
func (o DomainPurchaseConsentResponseOutput) AgreementKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DomainPurchaseConsentResponse) []string { return v.AgreementKeys }).(pulumi.StringArrayOutput)
}

type DomainPurchaseConsentResponsePtrOutput struct{ *pulumi.OutputState }

func (DomainPurchaseConsentResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainPurchaseConsentResponse)(nil)).Elem()
}

func (o DomainPurchaseConsentResponsePtrOutput) ToDomainPurchaseConsentResponsePtrOutput() DomainPurchaseConsentResponsePtrOutput {
	return o
}

func (o DomainPurchaseConsentResponsePtrOutput) ToDomainPurchaseConsentResponsePtrOutputWithContext(ctx context.Context) DomainPurchaseConsentResponsePtrOutput {
	return o
}

func (o DomainPurchaseConsentResponsePtrOutput) Elem() DomainPurchaseConsentResponseOutput {
	return o.ApplyT(func(v *DomainPurchaseConsentResponse) DomainPurchaseConsentResponse { return *v }).(DomainPurchaseConsentResponseOutput)
}

// Timestamp when the agreements were accepted.
func (o DomainPurchaseConsentResponsePtrOutput) AgreedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainPurchaseConsentResponse) *string {
		if v == nil {
			return nil
		}
		return v.AgreedAt
	}).(pulumi.StringPtrOutput)
}

// Client IP address.
func (o DomainPurchaseConsentResponsePtrOutput) AgreedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainPurchaseConsentResponse) *string {
		if v == nil {
			return nil
		}
		return v.AgreedBy
	}).(pulumi.StringPtrOutput)
}

// List of applicable legal agreement keys. This list can be retrieved using ListLegalAgreements API under <code>TopLevelDomain</code> resource.
func (o DomainPurchaseConsentResponsePtrOutput) AgreementKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DomainPurchaseConsentResponse) []string {
		if v == nil {
			return nil
		}
		return v.AgreementKeys
	}).(pulumi.StringArrayOutput)
}

// Details of a hostname derived from a domain.
type HostNameResponse struct {
	// Name of the Azure resource the hostname is assigned to. If it is assigned to a Traffic Manager then it will be the Traffic Manager name otherwise it will be the app name.
	AzureResourceName *string `pulumi:"azureResourceName"`
	// Type of the Azure resource the hostname is assigned to.
	AzureResourceType *string `pulumi:"azureResourceType"`
	// Type of the DNS record.
	CustomHostNameDnsRecordType *string `pulumi:"customHostNameDnsRecordType"`
	// Type of the hostname.
	HostNameType *string `pulumi:"hostNameType"`
	// Name of the hostname.
	Name *string `pulumi:"name"`
	// List of apps the hostname is assigned to. This list will have more than one app only if the hostname is pointing to a Traffic Manager.
	SiteNames []string `pulumi:"siteNames"`
}

// Details of a hostname derived from a domain.
type HostNameResponseOutput struct{ *pulumi.OutputState }

func (HostNameResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostNameResponse)(nil)).Elem()
}

func (o HostNameResponseOutput) ToHostNameResponseOutput() HostNameResponseOutput {
	return o
}

func (o HostNameResponseOutput) ToHostNameResponseOutputWithContext(ctx context.Context) HostNameResponseOutput {
	return o
}

// Name of the Azure resource the hostname is assigned to. If it is assigned to a Traffic Manager then it will be the Traffic Manager name otherwise it will be the app name.
func (o HostNameResponseOutput) AzureResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameResponse) *string { return v.AzureResourceName }).(pulumi.StringPtrOutput)
}

// Type of the Azure resource the hostname is assigned to.
func (o HostNameResponseOutput) AzureResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameResponse) *string { return v.AzureResourceType }).(pulumi.StringPtrOutput)
}

// Type of the DNS record.
func (o HostNameResponseOutput) CustomHostNameDnsRecordType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameResponse) *string { return v.CustomHostNameDnsRecordType }).(pulumi.StringPtrOutput)
}

// Type of the hostname.
func (o HostNameResponseOutput) HostNameType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameResponse) *string { return v.HostNameType }).(pulumi.StringPtrOutput)
}

// Name of the hostname.
func (o HostNameResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// List of apps the hostname is assigned to. This list will have more than one app only if the hostname is pointing to a Traffic Manager.
func (o HostNameResponseOutput) SiteNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v HostNameResponse) []string { return v.SiteNames }).(pulumi.StringArrayOutput)
}

type HostNameResponseArrayOutput struct{ *pulumi.OutputState }

func (HostNameResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HostNameResponse)(nil)).Elem()
}

func (o HostNameResponseArrayOutput) ToHostNameResponseArrayOutput() HostNameResponseArrayOutput {
	return o
}

func (o HostNameResponseArrayOutput) ToHostNameResponseArrayOutputWithContext(ctx context.Context) HostNameResponseArrayOutput {
	return o
}

func (o HostNameResponseArrayOutput) Index(i pulumi.IntInput) HostNameResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HostNameResponse {
		return vs[0].([]HostNameResponse)[vs[1].(int)]
	}).(HostNameResponseOutput)
}

// Identifies an object.
type NameIdentifierResponse struct {
	// Name of the object.
	Name *string `pulumi:"name"`
}

// Legal agreement for a top level domain.
type TldLegalAgreementResponse struct {
	// Unique identifier for the agreement.
	AgreementKey string `pulumi:"agreementKey"`
	// Agreement details.
	Content string `pulumi:"content"`
	// Agreement title.
	Title string `pulumi:"title"`
	// URL where a copy of the agreement details is hosted.
	Url *string `pulumi:"url"`
}

func init() {
	pulumi.RegisterOutputType(AddressOutput{})
	pulumi.RegisterOutputType(AddressPtrOutput{})
	pulumi.RegisterOutputType(AddressResponseOutput{})
	pulumi.RegisterOutputType(AddressResponsePtrOutput{})
	pulumi.RegisterOutputType(ContactOutput{})
	pulumi.RegisterOutputType(ContactPtrOutput{})
	pulumi.RegisterOutputType(ContactResponseOutput{})
	pulumi.RegisterOutputType(ContactResponsePtrOutput{})
	pulumi.RegisterOutputType(DomainPurchaseConsentOutput{})
	pulumi.RegisterOutputType(DomainPurchaseConsentPtrOutput{})
	pulumi.RegisterOutputType(DomainPurchaseConsentResponseOutput{})
	pulumi.RegisterOutputType(DomainPurchaseConsentResponsePtrOutput{})
	pulumi.RegisterOutputType(HostNameResponseOutput{})
	pulumi.RegisterOutputType(HostNameResponseArrayOutput{})
	pulumi.RegisterOutputType(NameIdentifierResponseOutput{})
	pulumi.RegisterOutputType(NameIdentifierResponseArrayOutput{})
	pulumi.RegisterOutputType(TldLegalAgreementResponseOutput{})
	pulumi.RegisterOutputType(TldLegalAgreementResponseArrayOutput{})
}
