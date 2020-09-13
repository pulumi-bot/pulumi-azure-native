// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200201preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Registration assignment.
//
// ## Example Usage
// ### Put Registration Assignment
//
// ```go
// package main
//
// import (
// 	managedservices "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/managedservices/v20200201preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := managedservices.NewRegistrationAssignment(ctx, "registrationAssignment", &managedservices.RegistrationAssignmentArgs{
// 			RegistrationAssignmentId: pulumi.String("26c128c2-fefa-4340-9bb1-6e081c90ada2"),
// 			Scope:                    pulumi.String("subscription/0afefe50-734e-4610-8a82-a144ahf49dea"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type RegistrationAssignment struct {
	pulumi.CustomResourceState

	// Name of the registration assignment.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of a registration assignment.
	Properties RegistrationAssignmentPropertiesResponseOutput `pulumi:"properties"`
	// Type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRegistrationAssignment registers a new resource with the given unique name, arguments, and options.
func NewRegistrationAssignment(ctx *pulumi.Context,
	name string, args *RegistrationAssignmentArgs, opts ...pulumi.ResourceOption) (*RegistrationAssignment, error) {
	if args == nil || args.RegistrationAssignmentId == nil {
		return nil, errors.New("missing required argument 'RegistrationAssignmentId'")
	}
	if args == nil || args.Scope == nil {
		return nil, errors.New("missing required argument 'Scope'")
	}
	if args == nil {
		args = &RegistrationAssignmentArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:managedservices/latest:RegistrationAssignment"),
		},
		{
			Type: pulumi.String("azurerm:managedservices/v20180601preview:RegistrationAssignment"),
		},
		{
			Type: pulumi.String("azurerm:managedservices/v20190401preview:RegistrationAssignment"),
		},
		{
			Type: pulumi.String("azurerm:managedservices/v20190601:RegistrationAssignment"),
		},
		{
			Type: pulumi.String("azurerm:managedservices/v20190901:RegistrationAssignment"),
		},
	})
	opts = append(opts, aliases)
	var resource RegistrationAssignment
	err := ctx.RegisterResource("azurerm:managedservices/v20200201preview:RegistrationAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistrationAssignment gets an existing RegistrationAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistrationAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistrationAssignmentState, opts ...pulumi.ResourceOption) (*RegistrationAssignment, error) {
	var resource RegistrationAssignment
	err := ctx.ReadResource("azurerm:managedservices/v20200201preview:RegistrationAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegistrationAssignment resources.
type registrationAssignmentState struct {
	// Name of the registration assignment.
	Name *string `pulumi:"name"`
	// Properties of a registration assignment.
	Properties *RegistrationAssignmentPropertiesResponse `pulumi:"properties"`
	// Type of the resource.
	Type *string `pulumi:"type"`
}

type RegistrationAssignmentState struct {
	// Name of the registration assignment.
	Name pulumi.StringPtrInput
	// Properties of a registration assignment.
	Properties RegistrationAssignmentPropertiesResponsePtrInput
	// Type of the resource.
	Type pulumi.StringPtrInput
}

func (RegistrationAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*registrationAssignmentState)(nil)).Elem()
}

type registrationAssignmentArgs struct {
	// Properties of a registration assignment.
	Properties *RegistrationAssignmentProperties `pulumi:"properties"`
	// Guid of the registration assignment.
	RegistrationAssignmentId string `pulumi:"registrationAssignmentId"`
	// Scope of the resource.
	Scope string `pulumi:"scope"`
}

// The set of arguments for constructing a RegistrationAssignment resource.
type RegistrationAssignmentArgs struct {
	// Properties of a registration assignment.
	Properties RegistrationAssignmentPropertiesPtrInput
	// Guid of the registration assignment.
	RegistrationAssignmentId pulumi.StringInput
	// Scope of the resource.
	Scope pulumi.StringInput
}

func (RegistrationAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registrationAssignmentArgs)(nil)).Elem()
}
