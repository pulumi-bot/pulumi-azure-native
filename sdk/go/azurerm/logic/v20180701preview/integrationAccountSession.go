// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180701preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The integration account session.
//
// ## Example Usage
// ### Create or update an integration account session
//
// ```go
// package main
//
// import (
// 	logic "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/logic/v20180701preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := logic.NewIntegrationAccountSession(ctx, "integrationAccountSession", &logic.IntegrationAccountSessionArgs{
// 			Content: pulumi.StringMap{
// 				"controlNumber":            pulumi.String("1234"),
// 				"controlNumberChangedTime": pulumi.String("2017-02-21T22:30:11.9923759Z"),
// 			},
// 			IntegrationAccountName: pulumi.String("testia123"),
// 			ResourceGroupName:      pulumi.String("testrg123"),
// 			SessionName:            pulumi.String("testsession123-ICN"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type IntegrationAccountSession struct {
	pulumi.CustomResourceState

	// The changed time.
	ChangedTime pulumi.StringOutput `pulumi:"changedTime"`
	// The session content.
	Content pulumi.MapOutput `pulumi:"content"`
	// The created time.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// The resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Gets the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Gets the resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIntegrationAccountSession registers a new resource with the given unique name, arguments, and options.
func NewIntegrationAccountSession(ctx *pulumi.Context,
	name string, args *IntegrationAccountSessionArgs, opts ...pulumi.ResourceOption) (*IntegrationAccountSession, error) {
	if args == nil || args.IntegrationAccountName == nil {
		return nil, errors.New("missing required argument 'IntegrationAccountName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.SessionName == nil {
		return nil, errors.New("missing required argument 'SessionName'")
	}
	if args == nil {
		args = &IntegrationAccountSessionArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:logic/latest:IntegrationAccountSession"),
		},
		{
			Type: pulumi.String("azurerm:logic/v20160601:IntegrationAccountSession"),
		},
		{
			Type: pulumi.String("azurerm:logic/v20190501:IntegrationAccountSession"),
		},
	})
	opts = append(opts, aliases)
	var resource IntegrationAccountSession
	err := ctx.RegisterResource("azurerm:logic/v20180701preview:IntegrationAccountSession", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationAccountSession gets an existing IntegrationAccountSession resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationAccountSession(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationAccountSessionState, opts ...pulumi.ResourceOption) (*IntegrationAccountSession, error) {
	var resource IntegrationAccountSession
	err := ctx.ReadResource("azurerm:logic/v20180701preview:IntegrationAccountSession", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationAccountSession resources.
type integrationAccountSessionState struct {
	// The changed time.
	ChangedTime *string `pulumi:"changedTime"`
	// The session content.
	Content map[string]interface{} `pulumi:"content"`
	// The created time.
	CreatedTime *string `pulumi:"createdTime"`
	// The resource location.
	Location *string `pulumi:"location"`
	// Gets the resource name.
	Name *string `pulumi:"name"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Gets the resource type.
	Type *string `pulumi:"type"`
}

type IntegrationAccountSessionState struct {
	// The changed time.
	ChangedTime pulumi.StringPtrInput
	// The session content.
	Content pulumi.MapInput
	// The created time.
	CreatedTime pulumi.StringPtrInput
	// The resource location.
	Location pulumi.StringPtrInput
	// Gets the resource name.
	Name pulumi.StringPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// Gets the resource type.
	Type pulumi.StringPtrInput
}

func (IntegrationAccountSessionState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountSessionState)(nil)).Elem()
}

type integrationAccountSessionArgs struct {
	// The session content.
	Content map[string]interface{} `pulumi:"content"`
	// The integration account name.
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The integration account session name.
	SessionName string `pulumi:"sessionName"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a IntegrationAccountSession resource.
type IntegrationAccountSessionArgs struct {
	// The session content.
	Content pulumi.MapInput
	// The integration account name.
	IntegrationAccountName pulumi.StringInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The integration account session name.
	SessionName pulumi.StringInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (IntegrationAccountSessionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountSessionArgs)(nil)).Elem()
}
