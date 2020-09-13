// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181001

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Response to get user settings
//
// ## Example Usage
// ### PutUserSettings
//
// ```go
// package main
//
// import (
// 	portal "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/portal/v20181001"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := portal.NewUserSettings(ctx, "userSettings", &portal.UserSettingsArgs{
// 			UserSettingsName: pulumi.String("cloudconsole"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type UserSettings struct {
	pulumi.CustomResourceState

	// The cloud shell user settings properties.
	Properties UserPropertiesResponseOutput `pulumi:"properties"`
}

// NewUserSettings registers a new resource with the given unique name, arguments, and options.
func NewUserSettings(ctx *pulumi.Context,
	name string, args *UserSettingsArgs, opts ...pulumi.ResourceOption) (*UserSettings, error) {
	if args == nil || args.Properties == nil {
		return nil, errors.New("missing required argument 'Properties'")
	}
	if args == nil || args.UserSettingsName == nil {
		return nil, errors.New("missing required argument 'UserSettingsName'")
	}
	if args == nil {
		args = &UserSettingsArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:portal/latest:UserSettings"),
		},
	})
	opts = append(opts, aliases)
	var resource UserSettings
	err := ctx.RegisterResource("azurerm:portal/v20181001:UserSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserSettings gets an existing UserSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserSettingsState, opts ...pulumi.ResourceOption) (*UserSettings, error) {
	var resource UserSettings
	err := ctx.ReadResource("azurerm:portal/v20181001:UserSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserSettings resources.
type userSettingsState struct {
	// The cloud shell user settings properties.
	Properties *UserPropertiesResponse `pulumi:"properties"`
}

type UserSettingsState struct {
	// The cloud shell user settings properties.
	Properties UserPropertiesResponsePtrInput
}

func (UserSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*userSettingsState)(nil)).Elem()
}

type userSettingsArgs struct {
	// The cloud shell user settings properties.
	Properties UserProperties `pulumi:"properties"`
	// The name of the user settings
	UserSettingsName string `pulumi:"userSettingsName"`
}

// The set of arguments for constructing a UserSettings resource.
type UserSettingsArgs struct {
	// The cloud shell user settings properties.
	Properties UserPropertiesInput
	// The name of the user settings
	UserSettingsName pulumi.StringInput
}

func (UserSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userSettingsArgs)(nil)).Elem()
}
