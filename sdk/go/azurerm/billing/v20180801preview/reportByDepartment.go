// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180801preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A report resource.
type ReportByDepartment struct {
	pulumi.CustomResourceState

	// Has definition for the report.
	Definition ReportDefinitionResponseOutput `pulumi:"definition"`
	// Has delivery information for the report.
	DeliveryInfo ReportDeliveryInfoResponseOutput `pulumi:"deliveryInfo"`
	// The format of the report being delivered.
	Format pulumi.StringPtrOutput `pulumi:"format"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Has schedule information for the report.
	Schedule ReportScheduleResponsePtrOutput `pulumi:"schedule"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewReportByDepartment registers a new resource with the given unique name, arguments, and options.
func NewReportByDepartment(ctx *pulumi.Context,
	name string, args *ReportByDepartmentArgs, opts ...pulumi.ResourceOption) (*ReportByDepartment, error) {
	if args == nil || args.Definition == nil {
		return nil, errors.New("missing required argument 'Definition'")
	}
	if args == nil || args.DeliveryInfo == nil {
		return nil, errors.New("missing required argument 'DeliveryInfo'")
	}
	if args == nil || args.DepartmentId == nil {
		return nil, errors.New("missing required argument 'DepartmentId'")
	}
	if args == nil || args.ReportName == nil {
		return nil, errors.New("missing required argument 'ReportName'")
	}
	if args == nil {
		args = &ReportByDepartmentArgs{}
	}
	var resource ReportByDepartment
	err := ctx.RegisterResource("azurerm:billing/v20180801preview:ReportByDepartment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReportByDepartment gets an existing ReportByDepartment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReportByDepartment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReportByDepartmentState, opts ...pulumi.ResourceOption) (*ReportByDepartment, error) {
	var resource ReportByDepartment
	err := ctx.ReadResource("azurerm:billing/v20180801preview:ReportByDepartment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReportByDepartment resources.
type reportByDepartmentState struct {
	// Has definition for the report.
	Definition *ReportDefinitionResponse `pulumi:"definition"`
	// Has delivery information for the report.
	DeliveryInfo *ReportDeliveryInfoResponse `pulumi:"deliveryInfo"`
	// The format of the report being delivered.
	Format *string `pulumi:"format"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Has schedule information for the report.
	Schedule *ReportScheduleResponse `pulumi:"schedule"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type ReportByDepartmentState struct {
	// Has definition for the report.
	Definition ReportDefinitionResponsePtrInput
	// Has delivery information for the report.
	DeliveryInfo ReportDeliveryInfoResponsePtrInput
	// The format of the report being delivered.
	Format pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Has schedule information for the report.
	Schedule ReportScheduleResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (ReportByDepartmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*reportByDepartmentState)(nil)).Elem()
}

type reportByDepartmentArgs struct {
	// Has definition for the report.
	Definition ReportDefinition `pulumi:"definition"`
	// Has delivery information for the report.
	DeliveryInfo ReportDeliveryInfo `pulumi:"deliveryInfo"`
	// Department ID
	DepartmentId string `pulumi:"departmentId"`
	// The format of the report being delivered.
	Format *string `pulumi:"format"`
	// Report Name.
	ReportName string `pulumi:"reportName"`
	// Has schedule information for the report.
	Schedule *ReportSchedule `pulumi:"schedule"`
}

// The set of arguments for constructing a ReportByDepartment resource.
type ReportByDepartmentArgs struct {
	// Has definition for the report.
	Definition ReportDefinitionInput
	// Has delivery information for the report.
	DeliveryInfo ReportDeliveryInfoInput
	// Department ID
	DepartmentId pulumi.StringInput
	// The format of the report being delivered.
	Format pulumi.StringPtrInput
	// Report Name.
	ReportName pulumi.StringInput
	// Has schedule information for the report.
	Schedule ReportSchedulePtrInput
}

func (ReportByDepartmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*reportByDepartmentArgs)(nil)).Elem()
}
