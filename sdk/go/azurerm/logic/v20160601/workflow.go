// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The workflow type.
type Workflow struct {
	pulumi.CustomResourceState

	// Gets the access endpoint.
	AccessEndpoint pulumi.StringOutput `pulumi:"accessEndpoint"`
	// Gets the changed time.
	ChangedTime pulumi.StringOutput `pulumi:"changedTime"`
	// Gets the created time.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// The definition. See [Schema reference for Workflow Definition Language in Azure Logic Apps](https://aka.ms/logic-apps-workflow-definition-language).
	Definition pulumi.MapOutput `pulumi:"definition"`
	// The integration account.
	IntegrationAccount ResourceReferenceResponsePtrOutput `pulumi:"integrationAccount"`
	// The resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Gets the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The parameters.
	Parameters WorkflowParameterResponseMapOutput `pulumi:"parameters"`
	// Gets the provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The sku.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// The state.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Gets the resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Gets the version.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewWorkflow registers a new resource with the given unique name, arguments, and options.
func NewWorkflow(ctx *pulumi.Context,
	name string, args *WorkflowArgs, opts ...pulumi.ResourceOption) (*Workflow, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.WorkflowName == nil {
		return nil, errors.New("missing required argument 'WorkflowName'")
	}
	if args == nil {
		args = &WorkflowArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:logic/latest:Workflow"),
		},
		{
			Type: pulumi.String("azurerm:logic/v20150201preview:Workflow"),
		},
		{
			Type: pulumi.String("azurerm:logic/v20180701preview:Workflow"),
		},
		{
			Type: pulumi.String("azurerm:logic/v20190501:Workflow"),
		},
	})
	opts = append(opts, aliases)
	var resource Workflow
	err := ctx.RegisterResource("azurerm:logic/v20160601:Workflow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkflow gets an existing Workflow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkflow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkflowState, opts ...pulumi.ResourceOption) (*Workflow, error) {
	var resource Workflow
	err := ctx.ReadResource("azurerm:logic/v20160601:Workflow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workflow resources.
type workflowState struct {
	// Gets the access endpoint.
	AccessEndpoint *string `pulumi:"accessEndpoint"`
	// Gets the changed time.
	ChangedTime *string `pulumi:"changedTime"`
	// Gets the created time.
	CreatedTime *string `pulumi:"createdTime"`
	// The definition. See [Schema reference for Workflow Definition Language in Azure Logic Apps](https://aka.ms/logic-apps-workflow-definition-language).
	Definition map[string]interface{} `pulumi:"definition"`
	// The integration account.
	IntegrationAccount *ResourceReferenceResponse `pulumi:"integrationAccount"`
	// The resource location.
	Location *string `pulumi:"location"`
	// Gets the resource name.
	Name *string `pulumi:"name"`
	// The parameters.
	Parameters map[string]WorkflowParameterResponse `pulumi:"parameters"`
	// Gets the provisioning state.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The sku.
	Sku *SkuResponse `pulumi:"sku"`
	// The state.
	State *string `pulumi:"state"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Gets the resource type.
	Type *string `pulumi:"type"`
	// Gets the version.
	Version *string `pulumi:"version"`
}

type WorkflowState struct {
	// Gets the access endpoint.
	AccessEndpoint pulumi.StringPtrInput
	// Gets the changed time.
	ChangedTime pulumi.StringPtrInput
	// Gets the created time.
	CreatedTime pulumi.StringPtrInput
	// The definition. See [Schema reference for Workflow Definition Language in Azure Logic Apps](https://aka.ms/logic-apps-workflow-definition-language).
	Definition pulumi.MapInput
	// The integration account.
	IntegrationAccount ResourceReferenceResponsePtrInput
	// The resource location.
	Location pulumi.StringPtrInput
	// Gets the resource name.
	Name pulumi.StringPtrInput
	// The parameters.
	Parameters WorkflowParameterResponseMapInput
	// Gets the provisioning state.
	ProvisioningState pulumi.StringPtrInput
	// The sku.
	Sku SkuResponsePtrInput
	// The state.
	State pulumi.StringPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// Gets the resource type.
	Type pulumi.StringPtrInput
	// Gets the version.
	Version pulumi.StringPtrInput
}

func (WorkflowState) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowState)(nil)).Elem()
}

type workflowArgs struct {
	// The definition. See [Schema reference for Workflow Definition Language in Azure Logic Apps](https://aka.ms/logic-apps-workflow-definition-language).
	Definition map[string]interface{} `pulumi:"definition"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The parameters.
	Parameters map[string]WorkflowParameter `pulumi:"parameters"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The sku.
	Sku *Sku `pulumi:"sku"`
	// The state.
	State *string `pulumi:"state"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The workflow name.
	WorkflowName string `pulumi:"workflowName"`
}

// The set of arguments for constructing a Workflow resource.
type WorkflowArgs struct {
	// The definition. See [Schema reference for Workflow Definition Language in Azure Logic Apps](https://aka.ms/logic-apps-workflow-definition-language).
	Definition pulumi.MapInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The parameters.
	Parameters WorkflowParameterMapInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The sku.
	Sku SkuPtrInput
	// The state.
	State pulumi.StringPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// The workflow name.
	WorkflowName pulumi.StringInput
}

func (WorkflowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowArgs)(nil)).Elem()
}
