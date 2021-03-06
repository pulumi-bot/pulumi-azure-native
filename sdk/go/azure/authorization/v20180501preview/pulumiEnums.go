// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The recurrence type : weekly, monthly, etc.
type AccessReviewRecurrencePatternType pulumi.String

const (
	AccessReviewRecurrencePatternTypeWeekly          = AccessReviewRecurrencePatternType("weekly")
	AccessReviewRecurrencePatternTypeAbsoluteMonthly = AccessReviewRecurrencePatternType("absoluteMonthly")
)

func (AccessReviewRecurrencePatternType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e AccessReviewRecurrencePatternType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessReviewRecurrencePatternType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessReviewRecurrencePatternType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AccessReviewRecurrencePatternType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The recurrence range type. The possible values are: endDate, noEnd, numbered.
type AccessReviewRecurrenceRangeType pulumi.String

const (
	AccessReviewRecurrenceRangeTypeEndDate  = AccessReviewRecurrenceRangeType("endDate")
	AccessReviewRecurrenceRangeTypeNoEnd    = AccessReviewRecurrenceRangeType("noEnd")
	AccessReviewRecurrenceRangeTypeNumbered = AccessReviewRecurrenceRangeType("numbered")
)

func (AccessReviewRecurrenceRangeType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e AccessReviewRecurrenceRangeType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessReviewRecurrenceRangeType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessReviewRecurrenceRangeType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AccessReviewRecurrenceRangeType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// This specifies the behavior for the autoReview feature when an access review completes.
type DefaultDecisionType pulumi.String

const (
	DefaultDecisionTypeApprove        = DefaultDecisionType("Approve")
	DefaultDecisionTypeDeny           = DefaultDecisionType("Deny")
	DefaultDecisionTypeRecommendation = DefaultDecisionType("Recommendation")
)

func (DefaultDecisionType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e DefaultDecisionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DefaultDecisionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DefaultDecisionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DefaultDecisionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}
