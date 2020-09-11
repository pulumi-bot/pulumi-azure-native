// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents the filter actions which are allowed for the transformation of a message that have been matched by a filter expression.
type Action struct {
	// This property is reserved for future use. An integer value showing the compatibility level, currently hard-coded to 20.
	CompatibilityLevel *int `pulumi:"compatibilityLevel"`
	// Value that indicates whether the rule action requires preprocessing.
	RequiresPreprocessing *bool `pulumi:"requiresPreprocessing"`
	// SQL expression. e.g. MyProperty='ABC'
	SqlExpression *string `pulumi:"sqlExpression"`
}

// ActionInput is an input type that accepts ActionArgs and ActionOutput values.
// You can construct a concrete instance of `ActionInput` via:
//
//          ActionArgs{...}
type ActionInput interface {
	pulumi.Input

	ToActionOutput() ActionOutput
	ToActionOutputWithContext(context.Context) ActionOutput
}

// Represents the filter actions which are allowed for the transformation of a message that have been matched by a filter expression.
type ActionArgs struct {
	// This property is reserved for future use. An integer value showing the compatibility level, currently hard-coded to 20.
	CompatibilityLevel pulumi.IntPtrInput `pulumi:"compatibilityLevel"`
	// Value that indicates whether the rule action requires preprocessing.
	RequiresPreprocessing pulumi.BoolPtrInput `pulumi:"requiresPreprocessing"`
	// SQL expression. e.g. MyProperty='ABC'
	SqlExpression pulumi.StringPtrInput `pulumi:"sqlExpression"`
}

func (ActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Action)(nil)).Elem()
}

func (i ActionArgs) ToActionOutput() ActionOutput {
	return i.ToActionOutputWithContext(context.Background())
}

func (i ActionArgs) ToActionOutputWithContext(ctx context.Context) ActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionOutput)
}

func (i ActionArgs) ToActionPtrOutput() ActionPtrOutput {
	return i.ToActionPtrOutputWithContext(context.Background())
}

func (i ActionArgs) ToActionPtrOutputWithContext(ctx context.Context) ActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionOutput).ToActionPtrOutputWithContext(ctx)
}

// ActionPtrInput is an input type that accepts ActionArgs, ActionPtr and ActionPtrOutput values.
// You can construct a concrete instance of `ActionPtrInput` via:
//
//          ActionArgs{...}
//
//  or:
//
//          nil
type ActionPtrInput interface {
	pulumi.Input

	ToActionPtrOutput() ActionPtrOutput
	ToActionPtrOutputWithContext(context.Context) ActionPtrOutput
}

type actionPtrType ActionArgs

func ActionPtr(v *ActionArgs) ActionPtrInput {
	return (*actionPtrType)(v)
}

func (*actionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Action)(nil)).Elem()
}

func (i *actionPtrType) ToActionPtrOutput() ActionPtrOutput {
	return i.ToActionPtrOutputWithContext(context.Background())
}

func (i *actionPtrType) ToActionPtrOutputWithContext(ctx context.Context) ActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionPtrOutput)
}

// Represents the filter actions which are allowed for the transformation of a message that have been matched by a filter expression.
type ActionResponse struct {
	// This property is reserved for future use. An integer value showing the compatibility level, currently hard-coded to 20.
	CompatibilityLevel *int `pulumi:"compatibilityLevel"`
	// Value that indicates whether the rule action requires preprocessing.
	RequiresPreprocessing *bool `pulumi:"requiresPreprocessing"`
	// SQL expression. e.g. MyProperty='ABC'
	SqlExpression *string `pulumi:"sqlExpression"`
}

// Represents the filter actions which are allowed for the transformation of a message that have been matched by a filter expression.
type ActionResponseOutput struct{ *pulumi.OutputState }

func (ActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionResponse)(nil)).Elem()
}

func (o ActionResponseOutput) ToActionResponseOutput() ActionResponseOutput {
	return o
}

func (o ActionResponseOutput) ToActionResponseOutputWithContext(ctx context.Context) ActionResponseOutput {
	return o
}

func (o ActionResponseOutput) ToActionResponsePtrOutput() ActionResponsePtrOutput {
	return o.ToActionResponsePtrOutputWithContext(context.Background())
}

func (o ActionResponseOutput) ToActionResponsePtrOutputWithContext(ctx context.Context) ActionResponsePtrOutput {
	return o.ApplyT(func(v ActionResponse) *ActionResponse {
		return &v
	}).(ActionResponsePtrOutput)
}

// This property is reserved for future use. An integer value showing the compatibility level, currently hard-coded to 20.
func (o ActionResponseOutput) CompatibilityLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ActionResponse) *int { return v.CompatibilityLevel }).(pulumi.IntPtrOutput)
}

// Value that indicates whether the rule action requires preprocessing.
func (o ActionResponseOutput) RequiresPreprocessing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ActionResponse) *bool { return v.RequiresPreprocessing }).(pulumi.BoolPtrOutput)
}

// SQL expression. e.g. MyProperty='ABC'
func (o ActionResponseOutput) SqlExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActionResponse) *string { return v.SqlExpression }).(pulumi.StringPtrOutput)
}

type ActionResponsePtrOutput struct{ *pulumi.OutputState }

func (ActionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionResponse)(nil)).Elem()
}

func (o ActionResponsePtrOutput) ToActionResponsePtrOutput() ActionResponsePtrOutput {
	return o
}

func (o ActionResponsePtrOutput) ToActionResponsePtrOutputWithContext(ctx context.Context) ActionResponsePtrOutput {
	return o
}

func (o ActionResponsePtrOutput) Elem() ActionResponseOutput {
	return o.ApplyT(func(v *ActionResponse) ActionResponse { return *v }).(ActionResponseOutput)
}

// This property is reserved for future use. An integer value showing the compatibility level, currently hard-coded to 20.
func (o ActionResponsePtrOutput) CompatibilityLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ActionResponse) *int {
		if v == nil {
			return nil
		}
		return v.CompatibilityLevel
	}).(pulumi.IntPtrOutput)
}

// Value that indicates whether the rule action requires preprocessing.
func (o ActionResponsePtrOutput) RequiresPreprocessing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ActionResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RequiresPreprocessing
	}).(pulumi.BoolPtrOutput)
}

// SQL expression. e.g. MyProperty='ABC'
func (o ActionResponsePtrOutput) SqlExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActionResponse) *string {
		if v == nil {
			return nil
		}
		return v.SqlExpression
	}).(pulumi.StringPtrOutput)
}

// Represents the correlation filter expression.
type CorrelationFilter struct {
	// Content type of the message.
	ContentType *string `pulumi:"contentType"`
	// Identifier of the correlation.
	CorrelationId *string `pulumi:"correlationId"`
	// Application specific label.
	Label *string `pulumi:"label"`
	// Identifier of the message.
	MessageId *string `pulumi:"messageId"`
	// dictionary object for custom filters
	Properties map[string]string `pulumi:"properties"`
	// Address of the queue to reply to.
	ReplyTo *string `pulumi:"replyTo"`
	// Session identifier to reply to.
	ReplyToSessionId *string `pulumi:"replyToSessionId"`
	// Value that indicates whether the rule action requires preprocessing.
	RequiresPreprocessing *bool `pulumi:"requiresPreprocessing"`
	// Session identifier.
	SessionId *string `pulumi:"sessionId"`
	// Address to send to.
	To *string `pulumi:"to"`
}

// CorrelationFilterInput is an input type that accepts CorrelationFilterArgs and CorrelationFilterOutput values.
// You can construct a concrete instance of `CorrelationFilterInput` via:
//
//          CorrelationFilterArgs{...}
type CorrelationFilterInput interface {
	pulumi.Input

	ToCorrelationFilterOutput() CorrelationFilterOutput
	ToCorrelationFilterOutputWithContext(context.Context) CorrelationFilterOutput
}

// Represents the correlation filter expression.
type CorrelationFilterArgs struct {
	// Content type of the message.
	ContentType pulumi.StringPtrInput `pulumi:"contentType"`
	// Identifier of the correlation.
	CorrelationId pulumi.StringPtrInput `pulumi:"correlationId"`
	// Application specific label.
	Label pulumi.StringPtrInput `pulumi:"label"`
	// Identifier of the message.
	MessageId pulumi.StringPtrInput `pulumi:"messageId"`
	// dictionary object for custom filters
	Properties pulumi.StringMapInput `pulumi:"properties"`
	// Address of the queue to reply to.
	ReplyTo pulumi.StringPtrInput `pulumi:"replyTo"`
	// Session identifier to reply to.
	ReplyToSessionId pulumi.StringPtrInput `pulumi:"replyToSessionId"`
	// Value that indicates whether the rule action requires preprocessing.
	RequiresPreprocessing pulumi.BoolPtrInput `pulumi:"requiresPreprocessing"`
	// Session identifier.
	SessionId pulumi.StringPtrInput `pulumi:"sessionId"`
	// Address to send to.
	To pulumi.StringPtrInput `pulumi:"to"`
}

func (CorrelationFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CorrelationFilter)(nil)).Elem()
}

func (i CorrelationFilterArgs) ToCorrelationFilterOutput() CorrelationFilterOutput {
	return i.ToCorrelationFilterOutputWithContext(context.Background())
}

func (i CorrelationFilterArgs) ToCorrelationFilterOutputWithContext(ctx context.Context) CorrelationFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorrelationFilterOutput)
}

func (i CorrelationFilterArgs) ToCorrelationFilterPtrOutput() CorrelationFilterPtrOutput {
	return i.ToCorrelationFilterPtrOutputWithContext(context.Background())
}

func (i CorrelationFilterArgs) ToCorrelationFilterPtrOutputWithContext(ctx context.Context) CorrelationFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorrelationFilterOutput).ToCorrelationFilterPtrOutputWithContext(ctx)
}

// CorrelationFilterPtrInput is an input type that accepts CorrelationFilterArgs, CorrelationFilterPtr and CorrelationFilterPtrOutput values.
// You can construct a concrete instance of `CorrelationFilterPtrInput` via:
//
//          CorrelationFilterArgs{...}
//
//  or:
//
//          nil
type CorrelationFilterPtrInput interface {
	pulumi.Input

	ToCorrelationFilterPtrOutput() CorrelationFilterPtrOutput
	ToCorrelationFilterPtrOutputWithContext(context.Context) CorrelationFilterPtrOutput
}

type correlationFilterPtrType CorrelationFilterArgs

func CorrelationFilterPtr(v *CorrelationFilterArgs) CorrelationFilterPtrInput {
	return (*correlationFilterPtrType)(v)
}

func (*correlationFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CorrelationFilter)(nil)).Elem()
}

func (i *correlationFilterPtrType) ToCorrelationFilterPtrOutput() CorrelationFilterPtrOutput {
	return i.ToCorrelationFilterPtrOutputWithContext(context.Background())
}

func (i *correlationFilterPtrType) ToCorrelationFilterPtrOutputWithContext(ctx context.Context) CorrelationFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorrelationFilterPtrOutput)
}

// Represents the correlation filter expression.
type CorrelationFilterResponse struct {
	// Content type of the message.
	ContentType *string `pulumi:"contentType"`
	// Identifier of the correlation.
	CorrelationId *string `pulumi:"correlationId"`
	// Application specific label.
	Label *string `pulumi:"label"`
	// Identifier of the message.
	MessageId *string `pulumi:"messageId"`
	// dictionary object for custom filters
	Properties map[string]string `pulumi:"properties"`
	// Address of the queue to reply to.
	ReplyTo *string `pulumi:"replyTo"`
	// Session identifier to reply to.
	ReplyToSessionId *string `pulumi:"replyToSessionId"`
	// Value that indicates whether the rule action requires preprocessing.
	RequiresPreprocessing *bool `pulumi:"requiresPreprocessing"`
	// Session identifier.
	SessionId *string `pulumi:"sessionId"`
	// Address to send to.
	To *string `pulumi:"to"`
}

// Represents the correlation filter expression.
type CorrelationFilterResponseOutput struct{ *pulumi.OutputState }

func (CorrelationFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CorrelationFilterResponse)(nil)).Elem()
}

func (o CorrelationFilterResponseOutput) ToCorrelationFilterResponseOutput() CorrelationFilterResponseOutput {
	return o
}

func (o CorrelationFilterResponseOutput) ToCorrelationFilterResponseOutputWithContext(ctx context.Context) CorrelationFilterResponseOutput {
	return o
}

func (o CorrelationFilterResponseOutput) ToCorrelationFilterResponsePtrOutput() CorrelationFilterResponsePtrOutput {
	return o.ToCorrelationFilterResponsePtrOutputWithContext(context.Background())
}

func (o CorrelationFilterResponseOutput) ToCorrelationFilterResponsePtrOutputWithContext(ctx context.Context) CorrelationFilterResponsePtrOutput {
	return o.ApplyT(func(v CorrelationFilterResponse) *CorrelationFilterResponse {
		return &v
	}).(CorrelationFilterResponsePtrOutput)
}

// Content type of the message.
func (o CorrelationFilterResponseOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilterResponse) *string { return v.ContentType }).(pulumi.StringPtrOutput)
}

// Identifier of the correlation.
func (o CorrelationFilterResponseOutput) CorrelationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilterResponse) *string { return v.CorrelationId }).(pulumi.StringPtrOutput)
}

// Application specific label.
func (o CorrelationFilterResponseOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilterResponse) *string { return v.Label }).(pulumi.StringPtrOutput)
}

// Identifier of the message.
func (o CorrelationFilterResponseOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilterResponse) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

// dictionary object for custom filters
func (o CorrelationFilterResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v CorrelationFilterResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

// Address of the queue to reply to.
func (o CorrelationFilterResponseOutput) ReplyTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilterResponse) *string { return v.ReplyTo }).(pulumi.StringPtrOutput)
}

// Session identifier to reply to.
func (o CorrelationFilterResponseOutput) ReplyToSessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilterResponse) *string { return v.ReplyToSessionId }).(pulumi.StringPtrOutput)
}

// Value that indicates whether the rule action requires preprocessing.
func (o CorrelationFilterResponseOutput) RequiresPreprocessing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CorrelationFilterResponse) *bool { return v.RequiresPreprocessing }).(pulumi.BoolPtrOutput)
}

// Session identifier.
func (o CorrelationFilterResponseOutput) SessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilterResponse) *string { return v.SessionId }).(pulumi.StringPtrOutput)
}

// Address to send to.
func (o CorrelationFilterResponseOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilterResponse) *string { return v.To }).(pulumi.StringPtrOutput)
}

type CorrelationFilterResponsePtrOutput struct{ *pulumi.OutputState }

func (CorrelationFilterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CorrelationFilterResponse)(nil)).Elem()
}

func (o CorrelationFilterResponsePtrOutput) ToCorrelationFilterResponsePtrOutput() CorrelationFilterResponsePtrOutput {
	return o
}

func (o CorrelationFilterResponsePtrOutput) ToCorrelationFilterResponsePtrOutputWithContext(ctx context.Context) CorrelationFilterResponsePtrOutput {
	return o
}

func (o CorrelationFilterResponsePtrOutput) Elem() CorrelationFilterResponseOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) CorrelationFilterResponse { return *v }).(CorrelationFilterResponseOutput)
}

// Content type of the message.
func (o CorrelationFilterResponsePtrOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.ContentType
	}).(pulumi.StringPtrOutput)
}

// Identifier of the correlation.
func (o CorrelationFilterResponsePtrOutput) CorrelationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.CorrelationId
	}).(pulumi.StringPtrOutput)
}

// Application specific label.
func (o CorrelationFilterResponsePtrOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.Label
	}).(pulumi.StringPtrOutput)
}

// Identifier of the message.
func (o CorrelationFilterResponsePtrOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.MessageId
	}).(pulumi.StringPtrOutput)
}

// dictionary object for custom filters
func (o CorrelationFilterResponsePtrOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(pulumi.StringMapOutput)
}

// Address of the queue to reply to.
func (o CorrelationFilterResponsePtrOutput) ReplyTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.ReplyTo
	}).(pulumi.StringPtrOutput)
}

// Session identifier to reply to.
func (o CorrelationFilterResponsePtrOutput) ReplyToSessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.ReplyToSessionId
	}).(pulumi.StringPtrOutput)
}

// Value that indicates whether the rule action requires preprocessing.
func (o CorrelationFilterResponsePtrOutput) RequiresPreprocessing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RequiresPreprocessing
	}).(pulumi.BoolPtrOutput)
}

// Session identifier.
func (o CorrelationFilterResponsePtrOutput) SessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.SessionId
	}).(pulumi.StringPtrOutput)
}

// Address to send to.
func (o CorrelationFilterResponsePtrOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.To
	}).(pulumi.StringPtrOutput)
}

// Message Count Details.
type MessageCountDetailsResponse struct {
	// Number of active messages in the queue, topic, or subscription.
	ActiveMessageCount int `pulumi:"activeMessageCount"`
	// Number of messages that are dead lettered.
	DeadLetterMessageCount int `pulumi:"deadLetterMessageCount"`
	// Number of scheduled messages.
	ScheduledMessageCount int `pulumi:"scheduledMessageCount"`
	// Number of messages transferred into dead letters.
	TransferDeadLetterMessageCount int `pulumi:"transferDeadLetterMessageCount"`
	// Number of messages transferred to another queue, topic, or subscription.
	TransferMessageCount int `pulumi:"transferMessageCount"`
}

// Message Count Details.
type MessageCountDetailsResponseOutput struct{ *pulumi.OutputState }

func (MessageCountDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MessageCountDetailsResponse)(nil)).Elem()
}

func (o MessageCountDetailsResponseOutput) ToMessageCountDetailsResponseOutput() MessageCountDetailsResponseOutput {
	return o
}

func (o MessageCountDetailsResponseOutput) ToMessageCountDetailsResponseOutputWithContext(ctx context.Context) MessageCountDetailsResponseOutput {
	return o
}

func (o MessageCountDetailsResponseOutput) ToMessageCountDetailsResponsePtrOutput() MessageCountDetailsResponsePtrOutput {
	return o.ToMessageCountDetailsResponsePtrOutputWithContext(context.Background())
}

func (o MessageCountDetailsResponseOutput) ToMessageCountDetailsResponsePtrOutputWithContext(ctx context.Context) MessageCountDetailsResponsePtrOutput {
	return o.ApplyT(func(v MessageCountDetailsResponse) *MessageCountDetailsResponse {
		return &v
	}).(MessageCountDetailsResponsePtrOutput)
}

// Number of active messages in the queue, topic, or subscription.
func (o MessageCountDetailsResponseOutput) ActiveMessageCount() pulumi.IntOutput {
	return o.ApplyT(func(v MessageCountDetailsResponse) int { return v.ActiveMessageCount }).(pulumi.IntOutput)
}

// Number of messages that are dead lettered.
func (o MessageCountDetailsResponseOutput) DeadLetterMessageCount() pulumi.IntOutput {
	return o.ApplyT(func(v MessageCountDetailsResponse) int { return v.DeadLetterMessageCount }).(pulumi.IntOutput)
}

// Number of scheduled messages.
func (o MessageCountDetailsResponseOutput) ScheduledMessageCount() pulumi.IntOutput {
	return o.ApplyT(func(v MessageCountDetailsResponse) int { return v.ScheduledMessageCount }).(pulumi.IntOutput)
}

// Number of messages transferred into dead letters.
func (o MessageCountDetailsResponseOutput) TransferDeadLetterMessageCount() pulumi.IntOutput {
	return o.ApplyT(func(v MessageCountDetailsResponse) int { return v.TransferDeadLetterMessageCount }).(pulumi.IntOutput)
}

// Number of messages transferred to another queue, topic, or subscription.
func (o MessageCountDetailsResponseOutput) TransferMessageCount() pulumi.IntOutput {
	return o.ApplyT(func(v MessageCountDetailsResponse) int { return v.TransferMessageCount }).(pulumi.IntOutput)
}

type MessageCountDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (MessageCountDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MessageCountDetailsResponse)(nil)).Elem()
}

func (o MessageCountDetailsResponsePtrOutput) ToMessageCountDetailsResponsePtrOutput() MessageCountDetailsResponsePtrOutput {
	return o
}

func (o MessageCountDetailsResponsePtrOutput) ToMessageCountDetailsResponsePtrOutputWithContext(ctx context.Context) MessageCountDetailsResponsePtrOutput {
	return o
}

func (o MessageCountDetailsResponsePtrOutput) Elem() MessageCountDetailsResponseOutput {
	return o.ApplyT(func(v *MessageCountDetailsResponse) MessageCountDetailsResponse { return *v }).(MessageCountDetailsResponseOutput)
}

// Number of active messages in the queue, topic, or subscription.
func (o MessageCountDetailsResponsePtrOutput) ActiveMessageCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MessageCountDetailsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.ActiveMessageCount
	}).(pulumi.IntPtrOutput)
}

// Number of messages that are dead lettered.
func (o MessageCountDetailsResponsePtrOutput) DeadLetterMessageCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MessageCountDetailsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.DeadLetterMessageCount
	}).(pulumi.IntPtrOutput)
}

// Number of scheduled messages.
func (o MessageCountDetailsResponsePtrOutput) ScheduledMessageCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MessageCountDetailsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.ScheduledMessageCount
	}).(pulumi.IntPtrOutput)
}

// Number of messages transferred into dead letters.
func (o MessageCountDetailsResponsePtrOutput) TransferDeadLetterMessageCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MessageCountDetailsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.TransferDeadLetterMessageCount
	}).(pulumi.IntPtrOutput)
}

// Number of messages transferred to another queue, topic, or subscription.
func (o MessageCountDetailsResponsePtrOutput) TransferMessageCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MessageCountDetailsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.TransferMessageCount
	}).(pulumi.IntPtrOutput)
}

// SKU of the namespace.
type SBSku struct {
	// The specified messaging units for the tier. For Premium tier, capacity are 1,2 and 4.
	Capacity *int `pulumi:"capacity"`
	// Name of this SKU.
	Name string `pulumi:"name"`
	// The billing tier of this particular SKU.
	Tier *string `pulumi:"tier"`
}

// SBSkuInput is an input type that accepts SBSkuArgs and SBSkuOutput values.
// You can construct a concrete instance of `SBSkuInput` via:
//
//          SBSkuArgs{...}
type SBSkuInput interface {
	pulumi.Input

	ToSBSkuOutput() SBSkuOutput
	ToSBSkuOutputWithContext(context.Context) SBSkuOutput
}

// SKU of the namespace.
type SBSkuArgs struct {
	// The specified messaging units for the tier. For Premium tier, capacity are 1,2 and 4.
	Capacity pulumi.IntPtrInput `pulumi:"capacity"`
	// Name of this SKU.
	Name pulumi.StringInput `pulumi:"name"`
	// The billing tier of this particular SKU.
	Tier pulumi.StringPtrInput `pulumi:"tier"`
}

func (SBSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SBSku)(nil)).Elem()
}

func (i SBSkuArgs) ToSBSkuOutput() SBSkuOutput {
	return i.ToSBSkuOutputWithContext(context.Background())
}

func (i SBSkuArgs) ToSBSkuOutputWithContext(ctx context.Context) SBSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SBSkuOutput)
}

func (i SBSkuArgs) ToSBSkuPtrOutput() SBSkuPtrOutput {
	return i.ToSBSkuPtrOutputWithContext(context.Background())
}

func (i SBSkuArgs) ToSBSkuPtrOutputWithContext(ctx context.Context) SBSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SBSkuOutput).ToSBSkuPtrOutputWithContext(ctx)
}

// SBSkuPtrInput is an input type that accepts SBSkuArgs, SBSkuPtr and SBSkuPtrOutput values.
// You can construct a concrete instance of `SBSkuPtrInput` via:
//
//          SBSkuArgs{...}
//
//  or:
//
//          nil
type SBSkuPtrInput interface {
	pulumi.Input

	ToSBSkuPtrOutput() SBSkuPtrOutput
	ToSBSkuPtrOutputWithContext(context.Context) SBSkuPtrOutput
}

type sbskuPtrType SBSkuArgs

func SBSkuPtr(v *SBSkuArgs) SBSkuPtrInput {
	return (*sbskuPtrType)(v)
}

func (*sbskuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SBSku)(nil)).Elem()
}

func (i *sbskuPtrType) ToSBSkuPtrOutput() SBSkuPtrOutput {
	return i.ToSBSkuPtrOutputWithContext(context.Background())
}

func (i *sbskuPtrType) ToSBSkuPtrOutputWithContext(ctx context.Context) SBSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SBSkuPtrOutput)
}

// SKU of the namespace.
type SBSkuResponse struct {
	// The specified messaging units for the tier. For Premium tier, capacity are 1,2 and 4.
	Capacity *int `pulumi:"capacity"`
	// Name of this SKU.
	Name string `pulumi:"name"`
	// The billing tier of this particular SKU.
	Tier *string `pulumi:"tier"`
}

// SKU of the namespace.
type SBSkuResponseOutput struct{ *pulumi.OutputState }

func (SBSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SBSkuResponse)(nil)).Elem()
}

func (o SBSkuResponseOutput) ToSBSkuResponseOutput() SBSkuResponseOutput {
	return o
}

func (o SBSkuResponseOutput) ToSBSkuResponseOutputWithContext(ctx context.Context) SBSkuResponseOutput {
	return o
}

func (o SBSkuResponseOutput) ToSBSkuResponsePtrOutput() SBSkuResponsePtrOutput {
	return o.ToSBSkuResponsePtrOutputWithContext(context.Background())
}

func (o SBSkuResponseOutput) ToSBSkuResponsePtrOutputWithContext(ctx context.Context) SBSkuResponsePtrOutput {
	return o.ApplyT(func(v SBSkuResponse) *SBSkuResponse {
		return &v
	}).(SBSkuResponsePtrOutput)
}

// The specified messaging units for the tier. For Premium tier, capacity are 1,2 and 4.
func (o SBSkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SBSkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

// Name of this SKU.
func (o SBSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SBSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

// The billing tier of this particular SKU.
func (o SBSkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SBSkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SBSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SBSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SBSkuResponse)(nil)).Elem()
}

func (o SBSkuResponsePtrOutput) ToSBSkuResponsePtrOutput() SBSkuResponsePtrOutput {
	return o
}

func (o SBSkuResponsePtrOutput) ToSBSkuResponsePtrOutputWithContext(ctx context.Context) SBSkuResponsePtrOutput {
	return o
}

func (o SBSkuResponsePtrOutput) Elem() SBSkuResponseOutput {
	return o.ApplyT(func(v *SBSkuResponse) SBSkuResponse { return *v }).(SBSkuResponseOutput)
}

// The specified messaging units for the tier. For Premium tier, capacity are 1,2 and 4.
func (o SBSkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SBSkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

// Name of this SKU.
func (o SBSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SBSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

// The billing tier of this particular SKU.
func (o SBSkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SBSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

// Represents a filter which is a composition of an expression and an action that is executed in the pub/sub pipeline.
type SqlFilter struct {
	// This property is reserved for future use. An integer value showing the compatibility level, currently hard-coded to 20.
	CompatibilityLevel *int `pulumi:"compatibilityLevel"`
	// Value that indicates whether the rule action requires preprocessing.
	RequiresPreprocessing *bool `pulumi:"requiresPreprocessing"`
	// The SQL expression. e.g. MyProperty='ABC'
	SqlExpression *string `pulumi:"sqlExpression"`
}

// SqlFilterInput is an input type that accepts SqlFilterArgs and SqlFilterOutput values.
// You can construct a concrete instance of `SqlFilterInput` via:
//
//          SqlFilterArgs{...}
type SqlFilterInput interface {
	pulumi.Input

	ToSqlFilterOutput() SqlFilterOutput
	ToSqlFilterOutputWithContext(context.Context) SqlFilterOutput
}

// Represents a filter which is a composition of an expression and an action that is executed in the pub/sub pipeline.
type SqlFilterArgs struct {
	// This property is reserved for future use. An integer value showing the compatibility level, currently hard-coded to 20.
	CompatibilityLevel pulumi.IntPtrInput `pulumi:"compatibilityLevel"`
	// Value that indicates whether the rule action requires preprocessing.
	RequiresPreprocessing pulumi.BoolPtrInput `pulumi:"requiresPreprocessing"`
	// The SQL expression. e.g. MyProperty='ABC'
	SqlExpression pulumi.StringPtrInput `pulumi:"sqlExpression"`
}

func (SqlFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlFilter)(nil)).Elem()
}

func (i SqlFilterArgs) ToSqlFilterOutput() SqlFilterOutput {
	return i.ToSqlFilterOutputWithContext(context.Background())
}

func (i SqlFilterArgs) ToSqlFilterOutputWithContext(ctx context.Context) SqlFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlFilterOutput)
}

func (i SqlFilterArgs) ToSqlFilterPtrOutput() SqlFilterPtrOutput {
	return i.ToSqlFilterPtrOutputWithContext(context.Background())
}

func (i SqlFilterArgs) ToSqlFilterPtrOutputWithContext(ctx context.Context) SqlFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlFilterOutput).ToSqlFilterPtrOutputWithContext(ctx)
}

// SqlFilterPtrInput is an input type that accepts SqlFilterArgs, SqlFilterPtr and SqlFilterPtrOutput values.
// You can construct a concrete instance of `SqlFilterPtrInput` via:
//
//          SqlFilterArgs{...}
//
//  or:
//
//          nil
type SqlFilterPtrInput interface {
	pulumi.Input

	ToSqlFilterPtrOutput() SqlFilterPtrOutput
	ToSqlFilterPtrOutputWithContext(context.Context) SqlFilterPtrOutput
}

type sqlFilterPtrType SqlFilterArgs

func SqlFilterPtr(v *SqlFilterArgs) SqlFilterPtrInput {
	return (*sqlFilterPtrType)(v)
}

func (*sqlFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlFilter)(nil)).Elem()
}

func (i *sqlFilterPtrType) ToSqlFilterPtrOutput() SqlFilterPtrOutput {
	return i.ToSqlFilterPtrOutputWithContext(context.Background())
}

func (i *sqlFilterPtrType) ToSqlFilterPtrOutputWithContext(ctx context.Context) SqlFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlFilterPtrOutput)
}

// Represents a filter which is a composition of an expression and an action that is executed in the pub/sub pipeline.
type SqlFilterResponse struct {
	// This property is reserved for future use. An integer value showing the compatibility level, currently hard-coded to 20.
	CompatibilityLevel *int `pulumi:"compatibilityLevel"`
	// Value that indicates whether the rule action requires preprocessing.
	RequiresPreprocessing *bool `pulumi:"requiresPreprocessing"`
	// The SQL expression. e.g. MyProperty='ABC'
	SqlExpression *string `pulumi:"sqlExpression"`
}

// Represents a filter which is a composition of an expression and an action that is executed in the pub/sub pipeline.
type SqlFilterResponseOutput struct{ *pulumi.OutputState }

func (SqlFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlFilterResponse)(nil)).Elem()
}

func (o SqlFilterResponseOutput) ToSqlFilterResponseOutput() SqlFilterResponseOutput {
	return o
}

func (o SqlFilterResponseOutput) ToSqlFilterResponseOutputWithContext(ctx context.Context) SqlFilterResponseOutput {
	return o
}

func (o SqlFilterResponseOutput) ToSqlFilterResponsePtrOutput() SqlFilterResponsePtrOutput {
	return o.ToSqlFilterResponsePtrOutputWithContext(context.Background())
}

func (o SqlFilterResponseOutput) ToSqlFilterResponsePtrOutputWithContext(ctx context.Context) SqlFilterResponsePtrOutput {
	return o.ApplyT(func(v SqlFilterResponse) *SqlFilterResponse {
		return &v
	}).(SqlFilterResponsePtrOutput)
}

// This property is reserved for future use. An integer value showing the compatibility level, currently hard-coded to 20.
func (o SqlFilterResponseOutput) CompatibilityLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SqlFilterResponse) *int { return v.CompatibilityLevel }).(pulumi.IntPtrOutput)
}

// Value that indicates whether the rule action requires preprocessing.
func (o SqlFilterResponseOutput) RequiresPreprocessing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SqlFilterResponse) *bool { return v.RequiresPreprocessing }).(pulumi.BoolPtrOutput)
}

// The SQL expression. e.g. MyProperty='ABC'
func (o SqlFilterResponseOutput) SqlExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlFilterResponse) *string { return v.SqlExpression }).(pulumi.StringPtrOutput)
}

type SqlFilterResponsePtrOutput struct{ *pulumi.OutputState }

func (SqlFilterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlFilterResponse)(nil)).Elem()
}

func (o SqlFilterResponsePtrOutput) ToSqlFilterResponsePtrOutput() SqlFilterResponsePtrOutput {
	return o
}

func (o SqlFilterResponsePtrOutput) ToSqlFilterResponsePtrOutputWithContext(ctx context.Context) SqlFilterResponsePtrOutput {
	return o
}

func (o SqlFilterResponsePtrOutput) Elem() SqlFilterResponseOutput {
	return o.ApplyT(func(v *SqlFilterResponse) SqlFilterResponse { return *v }).(SqlFilterResponseOutput)
}

// This property is reserved for future use. An integer value showing the compatibility level, currently hard-coded to 20.
func (o SqlFilterResponsePtrOutput) CompatibilityLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SqlFilterResponse) *int {
		if v == nil {
			return nil
		}
		return v.CompatibilityLevel
	}).(pulumi.IntPtrOutput)
}

// Value that indicates whether the rule action requires preprocessing.
func (o SqlFilterResponsePtrOutput) RequiresPreprocessing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SqlFilterResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RequiresPreprocessing
	}).(pulumi.BoolPtrOutput)
}

// The SQL expression. e.g. MyProperty='ABC'
func (o SqlFilterResponsePtrOutput) SqlExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.SqlExpression
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ActionOutput{})
	pulumi.RegisterOutputType(ActionPtrOutput{})
	pulumi.RegisterOutputType(ActionResponseOutput{})
	pulumi.RegisterOutputType(ActionResponsePtrOutput{})
	pulumi.RegisterOutputType(CorrelationFilterOutput{})
	pulumi.RegisterOutputType(CorrelationFilterPtrOutput{})
	pulumi.RegisterOutputType(CorrelationFilterResponseOutput{})
	pulumi.RegisterOutputType(CorrelationFilterResponsePtrOutput{})
	pulumi.RegisterOutputType(MessageCountDetailsResponseOutput{})
	pulumi.RegisterOutputType(MessageCountDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(SBSkuOutput{})
	pulumi.RegisterOutputType(SBSkuPtrOutput{})
	pulumi.RegisterOutputType(SBSkuResponseOutput{})
	pulumi.RegisterOutputType(SBSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SqlFilterOutput{})
	pulumi.RegisterOutputType(SqlFilterPtrOutput{})
	pulumi.RegisterOutputType(SqlFilterResponseOutput{})
	pulumi.RegisterOutputType(SqlFilterResponsePtrOutput{})
}
