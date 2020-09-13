// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The integration account map.
//
// ## Example Usage
// ### Create or update a map
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	logic "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/logic/v20160601"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := logic.NewMap(ctx, "_map", &logic.MapArgs{
// 			Content: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "<?xml version=\"1.0\" encoding=\"UTF-16\"?>\n<xsl:stylesheet xmlns:xsl=\"http://www.w3.org/1999/XSL/Transform\" xmlns:msxsl=\"urn:schemas-microsoft-com:xslt\" xmlns:var=\"http://schemas.microsoft.com/BizTalk/2003/var\" exclude-result-prefixes=\"msxsl var s0 userCSharp\" version=\"1.0\" xmlns:ns0=\"http://BizTalk_Server_Project4.StringFunctoidsDestinationSchema\" xmlns:s0=\"http://BizTalk_Server_Project4.StringFunctoidsSourceSchema\" xmlns:userCSharp=\"http://schemas.microsoft.com/BizTalk/2003/userCSharp\">\n  <xsl:import href=\"http://btsfunctoids.blob.core.windows.net/functoids/functoids.xslt\" />\n  <xsl:output omit-xml-declaration=\"yes\" method=\"xml\" version=\"1.0\" />\n  <xsl:template match=\"/\">\n    <xsl:apply-templates select=\"/s0:Root\" />\n  </xsl:template>\n  <xsl:template match=\"/s0:Root\">\n    <xsl:variable name=\"var:v1\" select=\"userCSharp:StringFind(string(StringFindSource/text()) , &quot;SearchString&quot;)\" />\n    <xsl:variable name=\"var:v2\" select=\"userCSharp:StringLeft(string(StringLeftSource/text()) , &quot;2&quot;)\" />\n    <xsl:variable name=\"var:v3\" select=\"userCSharp:StringRight(string(StringRightSource/text()) , &quot;2&quot;)\" />\n    <xsl:variable name=\"var:v4\" select=\"userCSharp:StringUpperCase(string(UppercaseSource/text()))\" />\n    <xsl:variable name=\"var:v5\" select=\"userCSharp:StringLowerCase(string(LowercaseSource/text()))\" />\n    <xsl:variable name=\"var:v6\" select=\"userCSharp:StringSize(string(SizeSource/text()))\" />\n    <xsl:variable name=\"var:v7\" select=\"userCSharp:StringSubstring(string(StringExtractSource/text()) , &quot;0&quot; , &quot;2&quot;)\" />\n    <xsl:variable name=\"var:v8\" select=\"userCSharp:StringConcat(string(StringConcatSource/text()))\" />\n    <xsl:variable name=\"var:v9\" select=\"userCSharp:StringTrimLeft(string(StringLeftTrimSource/text()))\" />\n    <xsl:variable name=\"var:v10\" select=\"userCSharp:StringTrimRight(string(StringRightTrimSource/text()))\" />\n    <ns0:Root>\n      <StringFindDestination>\n        <xsl:value-of select=\"", "$", "var:v1\" />\n      </StringFindDestination>\n      <StringLeftDestination>\n        <xsl:value-of select=\"", "$", "var:v2\" />\n      </StringLeftDestination>\n      <StringRightDestination>\n        <xsl:value-of select=\"", "$", "var:v3\" />\n      </StringRightDestination>\n      <UppercaseDestination>\n        <xsl:value-of select=\"", "$", "var:v4\" />\n      </UppercaseDestination>\n      <LowercaseDestination>\n        <xsl:value-of select=\"", "$", "var:v5\" />\n      </LowercaseDestination>\n      <SizeDestination>\n        <xsl:value-of select=\"", "$", "var:v6\" />\n      </SizeDestination>\n      <StringExtractDestination>\n        <xsl:value-of select=\"", "$", "var:v7\" />\n      </StringExtractDestination>\n      <StringConcatDestination>\n        <xsl:value-of select=\"", "$", "var:v8\" />\n      </StringConcatDestination>\n      <StringLeftTrimDestination>\n        <xsl:value-of select=\"", "$", "var:v9\" />\n      </StringLeftTrimDestination>\n      <StringRightTrimDestination>\n        <xsl:value-of select=\"", "$", "var:v10\" />\n      </StringRightTrimDestination>\n    </ns0:Root>\n  </xsl:template>\n</xsl:stylesheet>")),
// 			ContentType:            pulumi.String("application/xml"),
// 			IntegrationAccountName: pulumi.String("testIntegrationAccount"),
// 			Location:               pulumi.String("westus"),
// 			MapName:                pulumi.String("testMap"),
// 			MapType:                pulumi.String("Xslt"),
// 			Metadata:               nil,
// 			ResourceGroupName:      pulumi.String("testResourceGroup"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Map struct {
	pulumi.CustomResourceState

	// The changed time.
	ChangedTime pulumi.StringOutput `pulumi:"changedTime"`
	// The content.
	Content pulumi.StringPtrOutput `pulumi:"content"`
	// The content link.
	ContentLink ContentLinkResponseOutput `pulumi:"contentLink"`
	// The content type.
	ContentType pulumi.StringPtrOutput `pulumi:"contentType"`
	// The created time.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// The resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The map type.
	MapType pulumi.StringOutput `pulumi:"mapType"`
	// The metadata.
	Metadata pulumi.MapOutput `pulumi:"metadata"`
	// Gets the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The parameters schema of integration account map.
	ParametersSchema IntegrationAccountMapPropertiesResponseParametersSchemaPtrOutput `pulumi:"parametersSchema"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Gets the resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMap registers a new resource with the given unique name, arguments, and options.
func NewMap(ctx *pulumi.Context,
	name string, args *MapArgs, opts ...pulumi.ResourceOption) (*Map, error) {
	if args == nil || args.IntegrationAccountName == nil {
		return nil, errors.New("missing required argument 'IntegrationAccountName'")
	}
	if args == nil || args.MapName == nil {
		return nil, errors.New("missing required argument 'MapName'")
	}
	if args == nil || args.MapType == nil {
		return nil, errors.New("missing required argument 'MapType'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &MapArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:logic/latest:Map"),
		},
		{
			Type: pulumi.String("azurerm:logic/v20150801preview:Map"),
		},
		{
			Type: pulumi.String("azurerm:logic/v20180701preview:Map"),
		},
		{
			Type: pulumi.String("azurerm:logic/v20190501:Map"),
		},
	})
	opts = append(opts, aliases)
	var resource Map
	err := ctx.RegisterResource("azurerm:logic/v20160601:Map", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMap gets an existing Map resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MapState, opts ...pulumi.ResourceOption) (*Map, error) {
	var resource Map
	err := ctx.ReadResource("azurerm:logic/v20160601:Map", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Map resources.
type mapState struct {
	// The changed time.
	ChangedTime *string `pulumi:"changedTime"`
	// The content.
	Content *string `pulumi:"content"`
	// The content link.
	ContentLink *ContentLinkResponse `pulumi:"contentLink"`
	// The content type.
	ContentType *string `pulumi:"contentType"`
	// The created time.
	CreatedTime *string `pulumi:"createdTime"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The map type.
	MapType *string `pulumi:"mapType"`
	// The metadata.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// Gets the resource name.
	Name *string `pulumi:"name"`
	// The parameters schema of integration account map.
	ParametersSchema *IntegrationAccountMapPropertiesResponseParametersSchema `pulumi:"parametersSchema"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Gets the resource type.
	Type *string `pulumi:"type"`
}

type MapState struct {
	// The changed time.
	ChangedTime pulumi.StringPtrInput
	// The content.
	Content pulumi.StringPtrInput
	// The content link.
	ContentLink ContentLinkResponsePtrInput
	// The content type.
	ContentType pulumi.StringPtrInput
	// The created time.
	CreatedTime pulumi.StringPtrInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The map type.
	MapType pulumi.StringPtrInput
	// The metadata.
	Metadata pulumi.MapInput
	// Gets the resource name.
	Name pulumi.StringPtrInput
	// The parameters schema of integration account map.
	ParametersSchema IntegrationAccountMapPropertiesResponseParametersSchemaPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// Gets the resource type.
	Type pulumi.StringPtrInput
}

func (MapState) ElementType() reflect.Type {
	return reflect.TypeOf((*mapState)(nil)).Elem()
}

type mapArgs struct {
	// The content.
	Content *string `pulumi:"content"`
	// The content type.
	ContentType *string `pulumi:"contentType"`
	// The integration account name.
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The integration account map name.
	MapName string `pulumi:"mapName"`
	// The map type.
	MapType string `pulumi:"mapType"`
	// The metadata.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// The parameters schema of integration account map.
	ParametersSchema *IntegrationAccountMapPropertiesParametersSchema `pulumi:"parametersSchema"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Map resource.
type MapArgs struct {
	// The content.
	Content pulumi.StringPtrInput
	// The content type.
	ContentType pulumi.StringPtrInput
	// The integration account name.
	IntegrationAccountName pulumi.StringInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The integration account map name.
	MapName pulumi.StringInput
	// The map type.
	MapType pulumi.StringInput
	// The metadata.
	Metadata pulumi.MapInput
	// The parameters schema of integration account map.
	ParametersSchema IntegrationAccountMapPropertiesParametersSchemaPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (MapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mapArgs)(nil)).Elem()
}
