// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Registration information.
//
// ## Example Usage
// ### Create or update an Azure Stack registration.
//
// ```go
// package main
//
// import (
// 	azurestack "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/azurestack/latest"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := azurestack.NewRegistration(ctx, "registration", &azurestack.RegistrationArgs{
// 			Location:          pulumi.String("global"),
// 			RegistrationName:  pulumi.String("testregistration"),
// 			RegistrationToken: pulumi.String("EyjIAWXSAw5nTw9KZWWiOiJeZxZlbg9wBwvUdCiSIM9iaMVjdeLkijoinwIzyJa2Ytgtowm2yy00OdG4lTlLyJmtztHjZGfJZTC0NZK1iIWiY2XvdWRJzCi6iJy5nDy0oDk1LTNHmWeTnDUwyS05oDI0LTrINzYwoGq5mjAzziIsim1HCmtldHBsYwnLu3LuZGljYXrpB25FBmfIbgVkIJp0CNvLLCJOYXJkd2FYzuLUZM8iOlt7IM51bunvcMVZiJoYlCjcaw9ZiJPBIjNkzDJHmda3yte5ndqZMdq4YmZkZmi5oDM3OTY3ZwNMIL0SIM5PyYI6WyJLZTy0ztJJMwZKy2m0OWNLODDLMwm2zTm0ymzKyjmWySisiJA3njlHmtdlY2q4NjRjnwFIZtC1YZi5ZGyZodM3Y2vjIl0siMnwDsi6wyi2oDUZoTbiY2RhNDa0ymrKoWe4YtK5otblzWrJzGyzNCISIjmYnzC4M2vmnZdIoDRKM2i5ytfkmJlhnDc1zdhLzWm1il0sim5HBwuiOijIqzF1MTvhmDIXmIIsimrpc2SiolsioWNlZjVhnZM1otQ0nDu3NmjlN2M3zmfjzmyZMTJhZtiiLcjLZjLmmZJhmWVhytG0NTu0OTqZNWu1Mda0MZbIYtfjyijdLCj1DWlKijoinwM5Mwu3NjytMju5Os00oTIwlWi0OdmTnGzHotiWm2RjyTCxIIwiBWvTb3J5ijPbijAYZDA3M2fjNzu0YTRMZTfhodkxzDnkogY5ZtAWzdyXIiwINZcWzThLnDQ4otrJndAzZGI5MGzlYtY1ZJA5ZdfiNMQIXX1DlcJpC3n1zxiiOijZb21lB25LIIWIdmVyC2LVbiI6IJeuMcJ9"),
// 			ResourceGroup:     pulumi.String("azurestack"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Registration struct {
	pulumi.CustomResourceState

	// Specifies the billing mode for the Azure Stack registration.
	BillingModel pulumi.StringPtrOutput `pulumi:"billingModel"`
	// The identifier of the registered Azure Stack.
	CloudId pulumi.StringPtrOutput `pulumi:"cloudId"`
	// The entity tag used for optimistic concurrency when modifying the resource.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Location of the resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// Name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The object identifier associated with the Azure Stack connecting to Azure.
	ObjectId pulumi.StringPtrOutput `pulumi:"objectId"`
	// Custom tags for the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Type of Resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRegistration registers a new resource with the given unique name, arguments, and options.
func NewRegistration(ctx *pulumi.Context,
	name string, args *RegistrationArgs, opts ...pulumi.ResourceOption) (*Registration, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.RegistrationName == nil {
		return nil, errors.New("missing required argument 'RegistrationName'")
	}
	if args == nil || args.RegistrationToken == nil {
		return nil, errors.New("missing required argument 'RegistrationToken'")
	}
	if args == nil || args.ResourceGroup == nil {
		return nil, errors.New("missing required argument 'ResourceGroup'")
	}
	if args == nil {
		args = &RegistrationArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:azurestack/v20170601:Registration"),
		},
	})
	opts = append(opts, aliases)
	var resource Registration
	err := ctx.RegisterResource("azurerm:azurestack/latest:Registration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistration gets an existing Registration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistrationState, opts ...pulumi.ResourceOption) (*Registration, error) {
	var resource Registration
	err := ctx.ReadResource("azurerm:azurestack/latest:Registration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Registration resources.
type registrationState struct {
	// Specifies the billing mode for the Azure Stack registration.
	BillingModel *string `pulumi:"billingModel"`
	// The identifier of the registered Azure Stack.
	CloudId *string `pulumi:"cloudId"`
	// The entity tag used for optimistic concurrency when modifying the resource.
	Etag *string `pulumi:"etag"`
	// Location of the resource.
	Location *string `pulumi:"location"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// The object identifier associated with the Azure Stack connecting to Azure.
	ObjectId *string `pulumi:"objectId"`
	// Custom tags for the resource.
	Tags map[string]string `pulumi:"tags"`
	// Type of Resource.
	Type *string `pulumi:"type"`
}

type RegistrationState struct {
	// Specifies the billing mode for the Azure Stack registration.
	BillingModel pulumi.StringPtrInput
	// The identifier of the registered Azure Stack.
	CloudId pulumi.StringPtrInput
	// The entity tag used for optimistic concurrency when modifying the resource.
	Etag pulumi.StringPtrInput
	// Location of the resource.
	Location pulumi.StringPtrInput
	// Name of the resource.
	Name pulumi.StringPtrInput
	// The object identifier associated with the Azure Stack connecting to Azure.
	ObjectId pulumi.StringPtrInput
	// Custom tags for the resource.
	Tags pulumi.StringMapInput
	// Type of Resource.
	Type pulumi.StringPtrInput
}

func (RegistrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*registrationState)(nil)).Elem()
}

type registrationArgs struct {
	// Location of the resource.
	Location string `pulumi:"location"`
	// Name of the Azure Stack registration.
	RegistrationName string `pulumi:"registrationName"`
	// The token identifying registered Azure Stack
	RegistrationToken string `pulumi:"registrationToken"`
	// Name of the resource group.
	ResourceGroup string `pulumi:"resourceGroup"`
}

// The set of arguments for constructing a Registration resource.
type RegistrationArgs struct {
	// Location of the resource.
	Location pulumi.StringInput
	// Name of the Azure Stack registration.
	RegistrationName pulumi.StringInput
	// The token identifying registered Azure Stack
	RegistrationToken pulumi.StringInput
	// Name of the resource group.
	ResourceGroup pulumi.StringInput
}

func (RegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registrationArgs)(nil)).Elem()
}
