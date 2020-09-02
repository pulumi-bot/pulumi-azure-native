// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The integration account certificate.
type Certificate struct {
	pulumi.CustomResourceState

	// The changed time.
	ChangedTime pulumi.StringOutput `pulumi:"changedTime"`
	// The created time.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// The key details in the key vault.
	Key KeyVaultKeyReferenceResponsePtrOutput `pulumi:"key"`
	// The resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The metadata.
	Metadata pulumi.MapOutput `pulumi:"metadata"`
	// Gets the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The public certificate.
	PublicCertificate pulumi.StringPtrOutput `pulumi:"publicCertificate"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Gets the resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil || args.CertificateName == nil {
		return nil, errors.New("missing required argument 'CertificateName'")
	}
	if args == nil || args.IntegrationAccountName == nil {
		return nil, errors.New("missing required argument 'IntegrationAccountName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &CertificateArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:logic/latest:Certificate"),
		},
		{
			Type: pulumi.String("azurerm:logic/v20150801preview:Certificate"),
		},
		{
			Type: pulumi.String("azurerm:logic/v20180701preview:Certificate"),
		},
		{
			Type: pulumi.String("azurerm:logic/v20190501:Certificate"),
		},
	})
	opts = append(opts, aliases)
	var resource Certificate
	err := ctx.RegisterResource("azurerm:logic/v20160601:Certificate", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:logic/v20160601:Certificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Certificate resources.
type certificateState struct {
	// The changed time.
	ChangedTime *string `pulumi:"changedTime"`
	// The created time.
	CreatedTime *string `pulumi:"createdTime"`
	// The key details in the key vault.
	Key *KeyVaultKeyReferenceResponse `pulumi:"key"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The metadata.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// Gets the resource name.
	Name *string `pulumi:"name"`
	// The public certificate.
	PublicCertificate *string `pulumi:"publicCertificate"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Gets the resource type.
	Type *string `pulumi:"type"`
}

type CertificateState struct {
	// The changed time.
	ChangedTime pulumi.StringPtrInput
	// The created time.
	CreatedTime pulumi.StringPtrInput
	// The key details in the key vault.
	Key KeyVaultKeyReferenceResponsePtrInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The metadata.
	Metadata pulumi.MapInput
	// Gets the resource name.
	Name pulumi.StringPtrInput
	// The public certificate.
	PublicCertificate pulumi.StringPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// Gets the resource type.
	Type pulumi.StringPtrInput
}

func (CertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateState)(nil)).Elem()
}

type certificateArgs struct {
	// The integration account certificate name.
	CertificateName string `pulumi:"certificateName"`
	// The integration account name.
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	// The key details in the key vault.
	Key *KeyVaultKeyReference `pulumi:"key"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The metadata.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// The public certificate.
	PublicCertificate *string `pulumi:"publicCertificate"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// The integration account certificate name.
	CertificateName pulumi.StringInput
	// The integration account name.
	IntegrationAccountName pulumi.StringInput
	// The key details in the key vault.
	Key KeyVaultKeyReferencePtrInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The metadata.
	Metadata pulumi.MapInput
	// The public certificate.
	PublicCertificate pulumi.StringPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (CertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateArgs)(nil)).Elem()
}
