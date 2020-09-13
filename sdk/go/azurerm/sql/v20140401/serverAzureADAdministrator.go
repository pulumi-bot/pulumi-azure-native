// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20140401

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An server Active Directory Administrator.
//
// ## Example Usage
// ### Create/Update a server administrator
//
// ```go
// package main
//
// import (
// 	sql "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/sql/v20140401"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sql.NewServerAzureADAdministrator(ctx, "serverAzureADAdministrator", &sql.ServerAzureADAdministratorArgs{
// 			AdministratorName: pulumi.String("activeDirectory"),
// 			AdministratorType: pulumi.String("ActiveDirectory"),
// 			Login:             pulumi.String("bob@contoso.com"),
// 			ResourceGroupName: pulumi.String("sqlcrudtest-4799"),
// 			ServerName:        pulumi.String("sqlcrudtest-6440"),
// 			Sid:               pulumi.String("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c"),
// 			TenantId:          pulumi.String("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type ServerAzureADAdministrator struct {
	pulumi.CustomResourceState

	// The type of administrator.
	AdministratorType pulumi.StringOutput `pulumi:"administratorType"`
	// The server administrator login value.
	Login pulumi.StringOutput `pulumi:"login"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The server administrator Sid (Secure ID).
	Sid pulumi.StringOutput `pulumi:"sid"`
	// The server Active Directory Administrator tenant id.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewServerAzureADAdministrator registers a new resource with the given unique name, arguments, and options.
func NewServerAzureADAdministrator(ctx *pulumi.Context,
	name string, args *ServerAzureADAdministratorArgs, opts ...pulumi.ResourceOption) (*ServerAzureADAdministrator, error) {
	if args == nil || args.AdministratorName == nil {
		return nil, errors.New("missing required argument 'AdministratorName'")
	}
	if args == nil || args.AdministratorType == nil {
		return nil, errors.New("missing required argument 'AdministratorType'")
	}
	if args == nil || args.Login == nil {
		return nil, errors.New("missing required argument 'Login'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServerName == nil {
		return nil, errors.New("missing required argument 'ServerName'")
	}
	if args == nil || args.Sid == nil {
		return nil, errors.New("missing required argument 'Sid'")
	}
	if args == nil || args.TenantId == nil {
		return nil, errors.New("missing required argument 'TenantId'")
	}
	if args == nil {
		args = &ServerAzureADAdministratorArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:sql/latest:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azurerm:sql/v20180601preview:ServerAzureADAdministrator"),
		},
		{
			Type: pulumi.String("azurerm:sql/v20190601preview:ServerAzureADAdministrator"),
		},
	})
	opts = append(opts, aliases)
	var resource ServerAzureADAdministrator
	err := ctx.RegisterResource("azurerm:sql/v20140401:ServerAzureADAdministrator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerAzureADAdministrator gets an existing ServerAzureADAdministrator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerAzureADAdministrator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerAzureADAdministratorState, opts ...pulumi.ResourceOption) (*ServerAzureADAdministrator, error) {
	var resource ServerAzureADAdministrator
	err := ctx.ReadResource("azurerm:sql/v20140401:ServerAzureADAdministrator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerAzureADAdministrator resources.
type serverAzureADAdministratorState struct {
	// The type of administrator.
	AdministratorType *string `pulumi:"administratorType"`
	// The server administrator login value.
	Login *string `pulumi:"login"`
	// Resource name.
	Name *string `pulumi:"name"`
	// The server administrator Sid (Secure ID).
	Sid *string `pulumi:"sid"`
	// The server Active Directory Administrator tenant id.
	TenantId *string `pulumi:"tenantId"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type ServerAzureADAdministratorState struct {
	// The type of administrator.
	AdministratorType pulumi.StringPtrInput
	// The server administrator login value.
	Login pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// The server administrator Sid (Secure ID).
	Sid pulumi.StringPtrInput
	// The server Active Directory Administrator tenant id.
	TenantId pulumi.StringPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (ServerAzureADAdministratorState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverAzureADAdministratorState)(nil)).Elem()
}

type serverAzureADAdministratorArgs struct {
	// Name of the server administrator resource.
	AdministratorName string `pulumi:"administratorName"`
	// The type of administrator.
	AdministratorType string `pulumi:"administratorType"`
	// The server administrator login value.
	Login string `pulumi:"login"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The server administrator Sid (Secure ID).
	Sid string `pulumi:"sid"`
	// The server Active Directory Administrator tenant id.
	TenantId string `pulumi:"tenantId"`
}

// The set of arguments for constructing a ServerAzureADAdministrator resource.
type ServerAzureADAdministratorArgs struct {
	// Name of the server administrator resource.
	AdministratorName pulumi.StringInput
	// The type of administrator.
	AdministratorType pulumi.StringInput
	// The server administrator login value.
	Login pulumi.StringInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
	// The server administrator Sid (Secure ID).
	Sid pulumi.StringInput
	// The server Active Directory Administrator tenant id.
	TenantId pulumi.StringInput
}

func (ServerAzureADAdministratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverAzureADAdministratorArgs)(nil)).Elem()
}
