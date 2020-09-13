// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Definition of the certificate.
//
// ## Example Usage
// ### Create or update a certificate
//
// ```go
// package main
//
// import (
// 	automation "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/automation/latest"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := automation.NewCertificate(ctx, "certificate", &automation.CertificateArgs{
// 			AutomationAccountName: pulumi.String("myAutomationAccount18"),
// 			Base64Value:           pulumi.String("base 64 value of cert"),
// 			CertificateName:       pulumi.String("testCert"),
// 			Description:           pulumi.String("Sample Cert"),
// 			IsExportable:          pulumi.Bool(false),
// 			Name:                  pulumi.String("testCert"),
// 			ResourceGroupName:     pulumi.String("rg"),
// 			Thumbprint:            pulumi.String("thumbprint of cert"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Certificate struct {
	pulumi.CustomResourceState

	// Gets the creation time.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// Gets or sets the description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Gets the expiry time of the certificate.
	ExpiryTime pulumi.StringOutput `pulumi:"expiryTime"`
	// Gets the is exportable flag of the certificate.
	IsExportable pulumi.BoolOutput `pulumi:"isExportable"`
	// Gets the last modified time.
	LastModifiedTime pulumi.StringOutput `pulumi:"lastModifiedTime"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets the thumbprint of the certificate.
	Thumbprint pulumi.StringOutput `pulumi:"thumbprint"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil || args.AutomationAccountName == nil {
		return nil, errors.New("missing required argument 'AutomationAccountName'")
	}
	if args == nil || args.Base64Value == nil {
		return nil, errors.New("missing required argument 'Base64Value'")
	}
	if args == nil || args.CertificateName == nil {
		return nil, errors.New("missing required argument 'CertificateName'")
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
			Type: pulumi.String("azurerm:automation/v20151031:Certificate"),
		},
	})
	opts = append(opts, aliases)
	var resource Certificate
	err := ctx.RegisterResource("azurerm:automation/latest:Certificate", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:automation/latest:Certificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Certificate resources.
type certificateState struct {
	// Gets the creation time.
	CreationTime *string `pulumi:"creationTime"`
	// Gets or sets the description.
	Description *string `pulumi:"description"`
	// Gets the expiry time of the certificate.
	ExpiryTime *string `pulumi:"expiryTime"`
	// Gets the is exportable flag of the certificate.
	IsExportable *bool `pulumi:"isExportable"`
	// Gets the last modified time.
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// Gets the thumbprint of the certificate.
	Thumbprint *string `pulumi:"thumbprint"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type CertificateState struct {
	// Gets the creation time.
	CreationTime pulumi.StringPtrInput
	// Gets or sets the description.
	Description pulumi.StringPtrInput
	// Gets the expiry time of the certificate.
	ExpiryTime pulumi.StringPtrInput
	// Gets the is exportable flag of the certificate.
	IsExportable pulumi.BoolPtrInput
	// Gets the last modified time.
	LastModifiedTime pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// Gets the thumbprint of the certificate.
	Thumbprint pulumi.StringPtrInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (CertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateState)(nil)).Elem()
}

type certificateArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// Gets or sets the base64 encoded value of the certificate.
	Base64Value string `pulumi:"base64Value"`
	// The parameters supplied to the create or update certificate operation.
	CertificateName string `pulumi:"certificateName"`
	// Gets or sets the description of the certificate.
	Description *string `pulumi:"description"`
	// Gets or sets the is exportable flag of the certificate.
	IsExportable *bool `pulumi:"isExportable"`
	// Gets or sets the name of the certificate.
	Name string `pulumi:"name"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets the thumbprint of the certificate.
	Thumbprint *string `pulumi:"thumbprint"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput
	// Gets or sets the base64 encoded value of the certificate.
	Base64Value pulumi.StringInput
	// The parameters supplied to the create or update certificate operation.
	CertificateName pulumi.StringInput
	// Gets or sets the description of the certificate.
	Description pulumi.StringPtrInput
	// Gets or sets the is exportable flag of the certificate.
	IsExportable pulumi.BoolPtrInput
	// Gets or sets the name of the certificate.
	Name pulumi.StringInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
	// Gets or sets the thumbprint of the certificate.
	Thumbprint pulumi.StringPtrInput
}

func (CertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateArgs)(nil)).Elem()
}
