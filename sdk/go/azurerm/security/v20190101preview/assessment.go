// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190101preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Security assessment on a resource
//
// ## Example Usage
// ### Create security recommendation task on a resource
//
// ```go
//
// ```
type Assessment struct {
	pulumi.CustomResourceState

	// Additional data regarding the assessment
	AdditionalData pulumi.StringMapOutput `pulumi:"additionalData"`
	// User friendly display name of the assessment
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Links relevant to the assessment
	Links AssessmentLinksResponsePtrOutput `pulumi:"links"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Details of the resource that was assessed
	ResourceDetails pulumi.AnyOutput `pulumi:"resourceDetails"`
	// The result of the assessment
	Status AssessmentStatusResponseOutput `pulumi:"status"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAssessment registers a new resource with the given unique name, arguments, and options.
func NewAssessment(ctx *pulumi.Context,
	name string, args *AssessmentArgs, opts ...pulumi.ResourceOption) (*Assessment, error) {
	if args == nil || args.AssessmentName == nil {
		return nil, errors.New("missing required argument 'AssessmentName'")
	}
	if args == nil || args.ResourceDetails == nil {
		return nil, errors.New("missing required argument 'ResourceDetails'")
	}
	if args == nil || args.ResourceId == nil {
		return nil, errors.New("missing required argument 'ResourceId'")
	}
	if args == nil || args.Status == nil {
		return nil, errors.New("missing required argument 'Status'")
	}
	if args == nil {
		args = &AssessmentArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:security/latest:Assessment"),
		},
		{
			Type: pulumi.String("azurerm:security/v20200101:Assessment"),
		},
	})
	opts = append(opts, aliases)
	var resource Assessment
	err := ctx.RegisterResource("azurerm:security/v20190101preview:Assessment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAssessment gets an existing Assessment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAssessment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssessmentState, opts ...pulumi.ResourceOption) (*Assessment, error) {
	var resource Assessment
	err := ctx.ReadResource("azurerm:security/v20190101preview:Assessment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Assessment resources.
type assessmentState struct {
	// Additional data regarding the assessment
	AdditionalData map[string]string `pulumi:"additionalData"`
	// User friendly display name of the assessment
	DisplayName *string `pulumi:"displayName"`
	// Links relevant to the assessment
	Links *AssessmentLinksResponse `pulumi:"links"`
	// Resource name
	Name *string `pulumi:"name"`
	// Details of the resource that was assessed
	ResourceDetails interface{} `pulumi:"resourceDetails"`
	// The result of the assessment
	Status *AssessmentStatusResponse `pulumi:"status"`
	// Resource type
	Type *string `pulumi:"type"`
}

type AssessmentState struct {
	// Additional data regarding the assessment
	AdditionalData pulumi.StringMapInput
	// User friendly display name of the assessment
	DisplayName pulumi.StringPtrInput
	// Links relevant to the assessment
	Links AssessmentLinksResponsePtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Details of the resource that was assessed
	ResourceDetails pulumi.Input
	// The result of the assessment
	Status AssessmentStatusResponsePtrInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (AssessmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*assessmentState)(nil)).Elem()
}

type assessmentArgs struct {
	// Additional data regarding the assessment
	AdditionalData map[string]string `pulumi:"additionalData"`
	// The Assessment Key - Unique key for the assessment type
	AssessmentName string `pulumi:"assessmentName"`
	// Details of the resource that was assessed
	ResourceDetails interface{} `pulumi:"resourceDetails"`
	// The identifier of the resource.
	ResourceId string `pulumi:"resourceId"`
	// The result of the assessment
	Status AssessmentStatus `pulumi:"status"`
}

// The set of arguments for constructing a Assessment resource.
type AssessmentArgs struct {
	// Additional data regarding the assessment
	AdditionalData pulumi.StringMapInput
	// The Assessment Key - Unique key for the assessment type
	AssessmentName pulumi.StringInput
	// Details of the resource that was assessed
	ResourceDetails pulumi.Input
	// The identifier of the resource.
	ResourceId pulumi.StringInput
	// The result of the assessment
	Status AssessmentStatusInput
}

func (AssessmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assessmentArgs)(nil)).Elem()
}
