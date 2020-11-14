// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// App certificate
type Certificate struct {
	pulumi.CustomResourceState

	// Raw bytes of .cer file
	CerBlob pulumi.StringPtrOutput `pulumi:"cerBlob"`
	// Certificate expiration date
	ExpirationDate pulumi.StringPtrOutput `pulumi:"expirationDate"`
	// Friendly name of the certificate
	FriendlyName pulumi.StringPtrOutput `pulumi:"friendlyName"`
	// Host names the certificate applies to
	HostNames pulumi.StringArrayOutput `pulumi:"hostNames"`
	// Specification for the hosting environment (App Service Environment) to use for the certificate
	HostingEnvironmentProfile HostingEnvironmentProfileResponsePtrOutput `pulumi:"hostingEnvironmentProfile"`
	// Certificate issue Date
	IssueDate pulumi.StringPtrOutput `pulumi:"issueDate"`
	// Certificate issuer
	Issuer pulumi.StringPtrOutput `pulumi:"issuer"`
	// Kind of resource
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource Name
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Certificate password
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Pfx blob
	PfxBlob pulumi.StringPtrOutput `pulumi:"pfxBlob"`
	// Public key hash
	PublicKeyHash pulumi.StringPtrOutput `pulumi:"publicKeyHash"`
	// Self link
	SelfLink pulumi.StringPtrOutput `pulumi:"selfLink"`
	// App name
	SiteName pulumi.StringPtrOutput `pulumi:"siteName"`
	// Subject name of the certificate
	SubjectName pulumi.StringPtrOutput `pulumi:"subjectName"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Certificate thumbprint
	Thumbprint pulumi.StringPtrOutput `pulumi:"thumbprint"`
	// Resource type
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// Is the certificate valid?
	Valid pulumi.BoolPtrOutput `pulumi:"valid"`
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &CertificateArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:web/latest:Certificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20160301:Certificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20180201:Certificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20181101:Certificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20190801:Certificate"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200601:Certificate"),
		},
	})
	opts = append(opts, aliases)
	var resource Certificate
	err := ctx.RegisterResource("azure-nextgen:web/v20150801:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificate gets an existing Certificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	var resource Certificate
	err := ctx.ReadResource("azure-nextgen:web/v20150801:Certificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Certificate resources.
type certificateState struct {
	// Raw bytes of .cer file
	CerBlob *string `pulumi:"cerBlob"`
	// Certificate expiration date
	ExpirationDate *string `pulumi:"expirationDate"`
	// Friendly name of the certificate
	FriendlyName *string `pulumi:"friendlyName"`
	// Host names the certificate applies to
	HostNames []string `pulumi:"hostNames"`
	// Specification for the hosting environment (App Service Environment) to use for the certificate
	HostingEnvironmentProfile *HostingEnvironmentProfileResponse `pulumi:"hostingEnvironmentProfile"`
	// Certificate issue Date
	IssueDate *string `pulumi:"issueDate"`
	// Certificate issuer
	Issuer *string `pulumi:"issuer"`
	// Kind of resource
	Kind *string `pulumi:"kind"`
	// Resource Location
	Location *string `pulumi:"location"`
	// Resource Name
	Name *string `pulumi:"name"`
	// Certificate password
	Password *string `pulumi:"password"`
	// Pfx blob
	PfxBlob *string `pulumi:"pfxBlob"`
	// Public key hash
	PublicKeyHash *string `pulumi:"publicKeyHash"`
	// Self link
	SelfLink *string `pulumi:"selfLink"`
	// App name
	SiteName *string `pulumi:"siteName"`
	// Subject name of the certificate
	SubjectName *string `pulumi:"subjectName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Certificate thumbprint
	Thumbprint *string `pulumi:"thumbprint"`
	// Resource type
	Type *string `pulumi:"type"`
	// Is the certificate valid?
	Valid *bool `pulumi:"valid"`
}

type CertificateState struct {
	// Raw bytes of .cer file
	CerBlob pulumi.StringPtrInput
	// Certificate expiration date
	ExpirationDate pulumi.StringPtrInput
	// Friendly name of the certificate
	FriendlyName pulumi.StringPtrInput
	// Host names the certificate applies to
	HostNames pulumi.StringArrayInput
	// Specification for the hosting environment (App Service Environment) to use for the certificate
	HostingEnvironmentProfile HostingEnvironmentProfileResponsePtrInput
	// Certificate issue Date
	IssueDate pulumi.StringPtrInput
	// Certificate issuer
	Issuer pulumi.StringPtrInput
	// Kind of resource
	Kind pulumi.StringPtrInput
	// Resource Location
	Location pulumi.StringPtrInput
	// Resource Name
	Name pulumi.StringPtrInput
	// Certificate password
	Password pulumi.StringPtrInput
	// Pfx blob
	PfxBlob pulumi.StringPtrInput
	// Public key hash
	PublicKeyHash pulumi.StringPtrInput
	// Self link
	SelfLink pulumi.StringPtrInput
	// App name
	SiteName pulumi.StringPtrInput
	// Subject name of the certificate
	SubjectName pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Certificate thumbprint
	Thumbprint pulumi.StringPtrInput
	// Resource type
	Type pulumi.StringPtrInput
	// Is the certificate valid?
	Valid pulumi.BoolPtrInput
}

func (CertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateState)(nil)).Elem()
}

type certificateArgs struct {
	// Raw bytes of .cer file
	CerBlob *string `pulumi:"cerBlob"`
	// Certificate expiration date
	ExpirationDate *string `pulumi:"expirationDate"`
	// Friendly name of the certificate
	FriendlyName *string `pulumi:"friendlyName"`
	// Host names the certificate applies to
	HostNames []string `pulumi:"hostNames"`
	// Specification for the hosting environment (App Service Environment) to use for the certificate
	HostingEnvironmentProfile *HostingEnvironmentProfile `pulumi:"hostingEnvironmentProfile"`
	// Resource Id
	Id *string `pulumi:"id"`
	// Certificate issue Date
	IssueDate *string `pulumi:"issueDate"`
	// Certificate issuer
	Issuer *string `pulumi:"issuer"`
	// Kind of resource
	Kind *string `pulumi:"kind"`
	// Resource Location
	Location string `pulumi:"location"`
	// Resource Name
	Name string `pulumi:"name"`
	// Certificate password
	Password *string `pulumi:"password"`
	// Pfx blob
	PfxBlob *string `pulumi:"pfxBlob"`
	// Public key hash
	PublicKeyHash *string `pulumi:"publicKeyHash"`
	// Name of the resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Self link
	SelfLink *string `pulumi:"selfLink"`
	// App name
	SiteName *string `pulumi:"siteName"`
	// Subject name of the certificate
	SubjectName *string `pulumi:"subjectName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Certificate thumbprint
	Thumbprint *string `pulumi:"thumbprint"`
	// Resource type
	Type *string `pulumi:"type"`
	// Is the certificate valid?
	Valid *bool `pulumi:"valid"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// Raw bytes of .cer file
	CerBlob pulumi.StringPtrInput
	// Certificate expiration date
	ExpirationDate pulumi.StringPtrInput
	// Friendly name of the certificate
	FriendlyName pulumi.StringPtrInput
	// Host names the certificate applies to
	HostNames pulumi.StringArrayInput
	// Specification for the hosting environment (App Service Environment) to use for the certificate
	HostingEnvironmentProfile HostingEnvironmentProfilePtrInput
	// Resource Id
	Id pulumi.StringPtrInput
	// Certificate issue Date
	IssueDate pulumi.StringPtrInput
	// Certificate issuer
	Issuer pulumi.StringPtrInput
	// Kind of resource
	Kind pulumi.StringPtrInput
	// Resource Location
	Location pulumi.StringInput
	// Resource Name
	Name pulumi.StringInput
	// Certificate password
	Password pulumi.StringPtrInput
	// Pfx blob
	PfxBlob pulumi.StringPtrInput
	// Public key hash
	PublicKeyHash pulumi.StringPtrInput
	// Name of the resource group
	ResourceGroupName pulumi.StringInput
	// Self link
	SelfLink pulumi.StringPtrInput
	// App name
	SiteName pulumi.StringPtrInput
	// Subject name of the certificate
	SubjectName pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Certificate thumbprint
	Thumbprint pulumi.StringPtrInput
	// Resource type
	Type pulumi.StringPtrInput
	// Is the certificate valid?
	Valid pulumi.BoolPtrInput
}

func (CertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateArgs)(nil)).Elem()
}

type CertificateInput interface {
	pulumi.Input

	ToCertificateOutput() CertificateOutput
	ToCertificateOutputWithContext(ctx context.Context) CertificateOutput
}

func (Certificate) ElementType() reflect.Type {
	return reflect.TypeOf((*Certificate)(nil)).Elem()
}

func (i Certificate) ToCertificateOutput() CertificateOutput {
	return i.ToCertificateOutputWithContext(context.Background())
}

func (i Certificate) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateOutput)
}

type CertificateOutput struct {
	*pulumi.OutputState
}

func (CertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateOutput)(nil)).Elem()
}

func (o CertificateOutput) ToCertificateOutput() CertificateOutput {
	return o
}

func (o CertificateOutput) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CertificateOutput{})
}
