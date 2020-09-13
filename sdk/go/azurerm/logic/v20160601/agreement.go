// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The integration account agreement.
//
// ## Example Usage
// ### Create or update an agreement
//
// ```go
// package main
//
// import (
// 	logic "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/logic/v20160601"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := logic.NewAgreement(ctx, "agreement", &logic.AgreementArgs{
// 			AgreementName: pulumi.String("testAgreement"),
// 			AgreementType: pulumi.String("AS2"),
// 			Content: &logic.AgreementContentArgs{
// 				AS2: &logic.AS2AgreementContentArgs{
// 					ReceiveAgreement: &logic.AS2OneWayAgreementArgs{
// 						ProtocolSettings: &logic.AS2ProtocolSettingsArgs{
// 							AcknowledgementConnectionSettings: &logic.AS2AcknowledgementConnectionSettingsArgs{
// 								IgnoreCertificateNameMismatch: pulumi.Bool(true),
// 								KeepHttpConnectionAlive:       pulumi.Bool(true),
// 								SupportHttpStatusCodeContinue: pulumi.Bool(true),
// 								UnfoldHttpHeaders:             pulumi.Bool(true),
// 							},
// 							EnvelopeSettings: &logic.AS2EnvelopeSettingsArgs{
// 								AutogenerateFileName:                    pulumi.Bool(true),
// 								FileNameTemplate:                        pulumi.String("Test"),
// 								MessageContentType:                      pulumi.String("text/plain"),
// 								SuspendMessageOnFileNameGenerationError: pulumi.Bool(true),
// 								TransmitFileNameInMimeHeader:            pulumi.Bool(true),
// 							},
// 							ErrorSettings: &logic.AS2ErrorSettingsArgs{
// 								ResendIfMdnNotReceived:  pulumi.Bool(true),
// 								SuspendDuplicateMessage: pulumi.Bool(true),
// 							},
// 							MdnSettings: &logic.AS2MdnSettingsArgs{
// 								DispositionNotificationTo:  pulumi.String("http://tempuri.org"),
// 								MdnText:                    pulumi.String("Sample"),
// 								MicHashingAlgorithm:        pulumi.String("SHA1"),
// 								NeedMdn:                    pulumi.Bool(true),
// 								ReceiptDeliveryUrl:         pulumi.String("http://tempuri.org"),
// 								SendInboundMdnToMessageBox: pulumi.Bool(true),
// 								SendMdnAsynchronously:      pulumi.Bool(true),
// 								SignMdn:                    pulumi.Bool(true),
// 								SignOutboundMdnIfOptional:  pulumi.Bool(true),
// 							},
// 							MessageConnectionSettings: &logic.AS2MessageConnectionSettingsArgs{
// 								IgnoreCertificateNameMismatch: pulumi.Bool(true),
// 								KeepHttpConnectionAlive:       pulumi.Bool(true),
// 								SupportHttpStatusCodeContinue: pulumi.Bool(true),
// 								UnfoldHttpHeaders:             pulumi.Bool(true),
// 							},
// 							SecuritySettings: &logic.AS2SecuritySettingsArgs{
// 								EnableNrrForInboundDecodedMessages:  pulumi.Bool(true),
// 								EnableNrrForInboundEncodedMessages:  pulumi.Bool(true),
// 								EnableNrrForInboundMdn:              pulumi.Bool(true),
// 								EnableNrrForOutboundDecodedMessages: pulumi.Bool(true),
// 								EnableNrrForOutboundEncodedMessages: pulumi.Bool(true),
// 								EnableNrrForOutboundMdn:             pulumi.Bool(true),
// 								OverrideGroupSigningCertificate:     pulumi.Bool(false),
// 							},
// 							ValidationSettings: &logic.AS2ValidationSettingsArgs{
// 								CheckCertificateRevocationListOnReceive: pulumi.Bool(true),
// 								CheckCertificateRevocationListOnSend:    pulumi.Bool(true),
// 								CheckDuplicateMessage:                   pulumi.Bool(true),
// 								CompressMessage:                         pulumi.Bool(true),
// 								EncryptMessage:                          pulumi.Bool(false),
// 								EncryptionAlgorithm:                     pulumi.String("AES128"),
// 								InterchangeDuplicatesValidityDays:       pulumi.Int(100),
// 								OverrideMessageProperties:               pulumi.Bool(true),
// 								SignMessage:                             pulumi.Bool(false),
// 							},
// 						},
// 						ReceiverBusinessIdentity: &logic.BusinessIdentityArgs{
// 							Qualifier: pulumi.String("ZZ"),
// 							Value:     pulumi.String("ZZ"),
// 						},
// 						SenderBusinessIdentity: &logic.BusinessIdentityArgs{
// 							Qualifier: pulumi.String("AA"),
// 							Value:     pulumi.String("AA"),
// 						},
// 					},
// 					SendAgreement: &logic.AS2OneWayAgreementArgs{
// 						ProtocolSettings: &logic.AS2ProtocolSettingsArgs{
// 							AcknowledgementConnectionSettings: &logic.AS2AcknowledgementConnectionSettingsArgs{
// 								IgnoreCertificateNameMismatch: pulumi.Bool(true),
// 								KeepHttpConnectionAlive:       pulumi.Bool(true),
// 								SupportHttpStatusCodeContinue: pulumi.Bool(true),
// 								UnfoldHttpHeaders:             pulumi.Bool(true),
// 							},
// 							EnvelopeSettings: &logic.AS2EnvelopeSettingsArgs{
// 								AutogenerateFileName:                    pulumi.Bool(true),
// 								FileNameTemplate:                        pulumi.String("Test"),
// 								MessageContentType:                      pulumi.String("text/plain"),
// 								SuspendMessageOnFileNameGenerationError: pulumi.Bool(true),
// 								TransmitFileNameInMimeHeader:            pulumi.Bool(true),
// 							},
// 							ErrorSettings: &logic.AS2ErrorSettingsArgs{
// 								ResendIfMdnNotReceived:  pulumi.Bool(true),
// 								SuspendDuplicateMessage: pulumi.Bool(true),
// 							},
// 							MdnSettings: &logic.AS2MdnSettingsArgs{
// 								DispositionNotificationTo:  pulumi.String("http://tempuri.org"),
// 								MdnText:                    pulumi.String("Sample"),
// 								MicHashingAlgorithm:        pulumi.String("SHA1"),
// 								NeedMdn:                    pulumi.Bool(true),
// 								ReceiptDeliveryUrl:         pulumi.String("http://tempuri.org"),
// 								SendInboundMdnToMessageBox: pulumi.Bool(true),
// 								SendMdnAsynchronously:      pulumi.Bool(true),
// 								SignMdn:                    pulumi.Bool(true),
// 								SignOutboundMdnIfOptional:  pulumi.Bool(true),
// 							},
// 							MessageConnectionSettings: &logic.AS2MessageConnectionSettingsArgs{
// 								IgnoreCertificateNameMismatch: pulumi.Bool(true),
// 								KeepHttpConnectionAlive:       pulumi.Bool(true),
// 								SupportHttpStatusCodeContinue: pulumi.Bool(true),
// 								UnfoldHttpHeaders:             pulumi.Bool(true),
// 							},
// 							SecuritySettings: &logic.AS2SecuritySettingsArgs{
// 								EnableNrrForInboundDecodedMessages:  pulumi.Bool(true),
// 								EnableNrrForInboundEncodedMessages:  pulumi.Bool(true),
// 								EnableNrrForInboundMdn:              pulumi.Bool(true),
// 								EnableNrrForOutboundDecodedMessages: pulumi.Bool(true),
// 								EnableNrrForOutboundEncodedMessages: pulumi.Bool(true),
// 								EnableNrrForOutboundMdn:             pulumi.Bool(true),
// 								OverrideGroupSigningCertificate:     pulumi.Bool(false),
// 							},
// 							ValidationSettings: &logic.AS2ValidationSettingsArgs{
// 								CheckCertificateRevocationListOnReceive: pulumi.Bool(true),
// 								CheckCertificateRevocationListOnSend:    pulumi.Bool(true),
// 								CheckDuplicateMessage:                   pulumi.Bool(true),
// 								CompressMessage:                         pulumi.Bool(true),
// 								EncryptMessage:                          pulumi.Bool(false),
// 								EncryptionAlgorithm:                     pulumi.String("AES128"),
// 								InterchangeDuplicatesValidityDays:       pulumi.Int(100),
// 								OverrideMessageProperties:               pulumi.Bool(true),
// 								SignMessage:                             pulumi.Bool(false),
// 							},
// 						},
// 						ReceiverBusinessIdentity: &logic.BusinessIdentityArgs{
// 							Qualifier: pulumi.String("AA"),
// 							Value:     pulumi.String("AA"),
// 						},
// 						SenderBusinessIdentity: &logic.BusinessIdentityArgs{
// 							Qualifier: pulumi.String("ZZ"),
// 							Value:     pulumi.String("ZZ"),
// 						},
// 					},
// 				},
// 			},
// 			GuestIdentity: &logic.BusinessIdentityArgs{
// 				Qualifier: pulumi.String("AA"),
// 				Value:     pulumi.String("AA"),
// 			},
// 			GuestPartner: pulumi.String("GuestPartner"),
// 			HostIdentity: &logic.BusinessIdentityArgs{
// 				Qualifier: pulumi.String("ZZ"),
// 				Value:     pulumi.String("ZZ"),
// 			},
// 			HostPartner:            pulumi.String("HostPartner"),
// 			IntegrationAccountName: pulumi.String("testIntegrationAccount"),
// 			Location:               pulumi.String("westus"),
// 			Metadata:               nil,
// 			ResourceGroupName:      pulumi.String("testResourceGroup"),
// 			Tags: pulumi.StringMap{
// 				"IntegrationAccountAgreement": pulumi.String("<IntegrationAccountAgreementName>"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Agreement struct {
	pulumi.CustomResourceState

	// The agreement type.
	AgreementType pulumi.StringOutput `pulumi:"agreementType"`
	// The changed time.
	ChangedTime pulumi.StringOutput `pulumi:"changedTime"`
	// The agreement content.
	Content AgreementContentResponseOutput `pulumi:"content"`
	// The created time.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// The business identity of the guest partner.
	GuestIdentity BusinessIdentityResponseOutput `pulumi:"guestIdentity"`
	// The integration account partner that is set as guest partner for this agreement.
	GuestPartner pulumi.StringOutput `pulumi:"guestPartner"`
	// The business identity of the host partner.
	HostIdentity BusinessIdentityResponseOutput `pulumi:"hostIdentity"`
	// The integration account partner that is set as host partner for this agreement.
	HostPartner pulumi.StringOutput `pulumi:"hostPartner"`
	// The resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The metadata.
	Metadata pulumi.MapOutput `pulumi:"metadata"`
	// Gets the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Gets the resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAgreement registers a new resource with the given unique name, arguments, and options.
func NewAgreement(ctx *pulumi.Context,
	name string, args *AgreementArgs, opts ...pulumi.ResourceOption) (*Agreement, error) {
	if args == nil || args.AgreementName == nil {
		return nil, errors.New("missing required argument 'AgreementName'")
	}
	if args == nil || args.AgreementType == nil {
		return nil, errors.New("missing required argument 'AgreementType'")
	}
	if args == nil || args.Content == nil {
		return nil, errors.New("missing required argument 'Content'")
	}
	if args == nil || args.GuestIdentity == nil {
		return nil, errors.New("missing required argument 'GuestIdentity'")
	}
	if args == nil || args.GuestPartner == nil {
		return nil, errors.New("missing required argument 'GuestPartner'")
	}
	if args == nil || args.HostIdentity == nil {
		return nil, errors.New("missing required argument 'HostIdentity'")
	}
	if args == nil || args.HostPartner == nil {
		return nil, errors.New("missing required argument 'HostPartner'")
	}
	if args == nil || args.IntegrationAccountName == nil {
		return nil, errors.New("missing required argument 'IntegrationAccountName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &AgreementArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:logic/latest:Agreement"),
		},
		{
			Type: pulumi.String("azurerm:logic/v20150801preview:Agreement"),
		},
		{
			Type: pulumi.String("azurerm:logic/v20180701preview:Agreement"),
		},
		{
			Type: pulumi.String("azurerm:logic/v20190501:Agreement"),
		},
	})
	opts = append(opts, aliases)
	var resource Agreement
	err := ctx.RegisterResource("azurerm:logic/v20160601:Agreement", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAgreement gets an existing Agreement resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAgreement(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AgreementState, opts ...pulumi.ResourceOption) (*Agreement, error) {
	var resource Agreement
	err := ctx.ReadResource("azurerm:logic/v20160601:Agreement", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Agreement resources.
type agreementState struct {
	// The agreement type.
	AgreementType *string `pulumi:"agreementType"`
	// The changed time.
	ChangedTime *string `pulumi:"changedTime"`
	// The agreement content.
	Content *AgreementContentResponse `pulumi:"content"`
	// The created time.
	CreatedTime *string `pulumi:"createdTime"`
	// The business identity of the guest partner.
	GuestIdentity *BusinessIdentityResponse `pulumi:"guestIdentity"`
	// The integration account partner that is set as guest partner for this agreement.
	GuestPartner *string `pulumi:"guestPartner"`
	// The business identity of the host partner.
	HostIdentity *BusinessIdentityResponse `pulumi:"hostIdentity"`
	// The integration account partner that is set as host partner for this agreement.
	HostPartner *string `pulumi:"hostPartner"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The metadata.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// Gets the resource name.
	Name *string `pulumi:"name"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Gets the resource type.
	Type *string `pulumi:"type"`
}

type AgreementState struct {
	// The agreement type.
	AgreementType pulumi.StringPtrInput
	// The changed time.
	ChangedTime pulumi.StringPtrInput
	// The agreement content.
	Content AgreementContentResponsePtrInput
	// The created time.
	CreatedTime pulumi.StringPtrInput
	// The business identity of the guest partner.
	GuestIdentity BusinessIdentityResponsePtrInput
	// The integration account partner that is set as guest partner for this agreement.
	GuestPartner pulumi.StringPtrInput
	// The business identity of the host partner.
	HostIdentity BusinessIdentityResponsePtrInput
	// The integration account partner that is set as host partner for this agreement.
	HostPartner pulumi.StringPtrInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The metadata.
	Metadata pulumi.MapInput
	// Gets the resource name.
	Name pulumi.StringPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// Gets the resource type.
	Type pulumi.StringPtrInput
}

func (AgreementState) ElementType() reflect.Type {
	return reflect.TypeOf((*agreementState)(nil)).Elem()
}

type agreementArgs struct {
	// The integration account agreement name.
	AgreementName string `pulumi:"agreementName"`
	// The agreement type.
	AgreementType string `pulumi:"agreementType"`
	// The agreement content.
	Content AgreementContent `pulumi:"content"`
	// The business identity of the guest partner.
	GuestIdentity BusinessIdentity `pulumi:"guestIdentity"`
	// The integration account partner that is set as guest partner for this agreement.
	GuestPartner string `pulumi:"guestPartner"`
	// The business identity of the host partner.
	HostIdentity BusinessIdentity `pulumi:"hostIdentity"`
	// The integration account partner that is set as host partner for this agreement.
	HostPartner string `pulumi:"hostPartner"`
	// The integration account name.
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The metadata.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Agreement resource.
type AgreementArgs struct {
	// The integration account agreement name.
	AgreementName pulumi.StringInput
	// The agreement type.
	AgreementType pulumi.StringInput
	// The agreement content.
	Content AgreementContentInput
	// The business identity of the guest partner.
	GuestIdentity BusinessIdentityInput
	// The integration account partner that is set as guest partner for this agreement.
	GuestPartner pulumi.StringInput
	// The business identity of the host partner.
	HostIdentity BusinessIdentityInput
	// The integration account partner that is set as host partner for this agreement.
	HostPartner pulumi.StringInput
	// The integration account name.
	IntegrationAccountName pulumi.StringInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The metadata.
	Metadata pulumi.MapInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (AgreementArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*agreementArgs)(nil)).Elem()
}
