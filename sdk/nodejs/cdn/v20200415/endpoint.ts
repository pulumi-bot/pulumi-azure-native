// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * CDN endpoint is the entity within a CDN profile containing configuration information such as origin, protocol, content caching and delivery behavior. The CDN endpoint uses the URL format <endpointname>.azureedge.net.
 */
export class Endpoint extends pulumi.CustomResource {
    /**
     * Get an existing Endpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Endpoint {
        return new Endpoint(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:cdn/v20200415:Endpoint';

    /**
     * Returns true if the given object is an instance of Endpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Endpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Endpoint.__pulumiType;
    }

    /**
     * List of content types on which compression applies. The value should be a valid MIME type.
     */
    public readonly contentTypesToCompress!: pulumi.Output<string[] | undefined>;
    /**
     * A reference to the origin group.
     */
    public readonly defaultOriginGroup!: pulumi.Output<outputs.cdn.v20200415.ResourceReferenceResponse | undefined>;
    /**
     * A policy that specifies the delivery rules to be used for an endpoint.
     */
    public readonly deliveryPolicy!: pulumi.Output<outputs.cdn.v20200415.EndpointPropertiesUpdateParametersResponseDeliveryPolicy | undefined>;
    /**
     * List of rules defining the user's geo access within a CDN endpoint. Each geo filter defines an access rule to a specified path or content, e.g. block APAC for path /pictures/
     */
    public readonly geoFilters!: pulumi.Output<outputs.cdn.v20200415.GeoFilterResponse[] | undefined>;
    /**
     * The host name of the endpoint structured as {endpointName}.{DNSZone}, e.g. contoso.azureedge.net
     */
    public /*out*/ readonly hostName!: pulumi.Output<string>;
    /**
     * Indicates whether content compression is enabled on CDN. Default value is false. If compression is enabled, content will be served as compressed if user requests for a compressed version. Content won't be compressed on CDN when requested content is smaller than 1 byte or larger than 1 MB.
     */
    public readonly isCompressionEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates whether HTTP traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
     */
    public readonly isHttpAllowed!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates whether HTTPS traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
     */
    public readonly isHttpsAllowed!: pulumi.Output<boolean | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Specifies what scenario the customer wants this CDN endpoint to optimize for, e.g. Download, Media services. With this information, CDN can apply scenario driven optimization.
     */
    public readonly optimizationType!: pulumi.Output<string | undefined>;
    /**
     * The origin groups comprising of origins that are used for load balancing the traffic based on availability.
     */
    public readonly originGroups!: pulumi.Output<outputs.cdn.v20200415.DeepCreatedOriginGroupResponse[] | undefined>;
    /**
     * The host header value sent to the origin with each request. This property at Endpoint is only allowed when endpoint uses single origin and can be overridden by the same property specified at origin.If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default.
     */
    public readonly originHostHeader!: pulumi.Output<string | undefined>;
    /**
     * A directory path on the origin that CDN can use to retrieve content from, e.g. contoso.cloudapp.net/originpath.
     */
    public readonly originPath!: pulumi.Output<string | undefined>;
    /**
     * The source of the content being delivered via CDN.
     */
    public readonly origins!: pulumi.Output<outputs.cdn.v20200415.DeepCreatedOriginResponse[]>;
    /**
     * Path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the origin path. This property is only relevant when using a single origin.
     */
    public readonly probePath!: pulumi.Output<string | undefined>;
    /**
     * Provisioning status of the endpoint.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Defines how CDN caches requests that include query strings. You can ignore any query strings when caching, bypass caching to prevent requests that contain query strings from being cached, or cache every request with a unique URL.
     */
    public readonly queryStringCachingBehavior!: pulumi.Output<string | undefined>;
    /**
     * Resource status of the endpoint.
     */
    public /*out*/ readonly resourceState!: pulumi.Output<string>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * List of keys used to validate the signed URL hashes.
     */
    public readonly urlSigningKeys!: pulumi.Output<outputs.cdn.v20200415.UrlSigningKeyResponse[] | undefined>;
    /**
     * Defines the Web Application Firewall policy for the endpoint (if applicable)
     */
    public readonly webApplicationFirewallPolicyLink!: pulumi.Output<outputs.cdn.v20200415.EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLink | undefined>;

    /**
     * Create a Endpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EndpointArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.endpointName === undefined) {
                throw new Error("Missing required property 'endpointName'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.origins === undefined) {
                throw new Error("Missing required property 'origins'");
            }
            if (!args || args.profileName === undefined) {
                throw new Error("Missing required property 'profileName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["contentTypesToCompress"] = args ? args.contentTypesToCompress : undefined;
            inputs["defaultOriginGroup"] = args ? args.defaultOriginGroup : undefined;
            inputs["deliveryPolicy"] = args ? args.deliveryPolicy : undefined;
            inputs["endpointName"] = args ? args.endpointName : undefined;
            inputs["geoFilters"] = args ? args.geoFilters : undefined;
            inputs["isCompressionEnabled"] = args ? args.isCompressionEnabled : undefined;
            inputs["isHttpAllowed"] = args ? args.isHttpAllowed : undefined;
            inputs["isHttpsAllowed"] = args ? args.isHttpsAllowed : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["optimizationType"] = args ? args.optimizationType : undefined;
            inputs["originGroups"] = args ? args.originGroups : undefined;
            inputs["originHostHeader"] = args ? args.originHostHeader : undefined;
            inputs["originPath"] = args ? args.originPath : undefined;
            inputs["origins"] = args ? args.origins : undefined;
            inputs["probePath"] = args ? args.probePath : undefined;
            inputs["profileName"] = args ? args.profileName : undefined;
            inputs["queryStringCachingBehavior"] = args ? args.queryStringCachingBehavior : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["urlSigningKeys"] = args ? args.urlSigningKeys : undefined;
            inputs["webApplicationFirewallPolicyLink"] = args ? args.webApplicationFirewallPolicyLink : undefined;
            inputs["hostName"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resourceState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["contentTypesToCompress"] = undefined /*out*/;
            inputs["defaultOriginGroup"] = undefined /*out*/;
            inputs["deliveryPolicy"] = undefined /*out*/;
            inputs["geoFilters"] = undefined /*out*/;
            inputs["hostName"] = undefined /*out*/;
            inputs["isCompressionEnabled"] = undefined /*out*/;
            inputs["isHttpAllowed"] = undefined /*out*/;
            inputs["isHttpsAllowed"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["optimizationType"] = undefined /*out*/;
            inputs["originGroups"] = undefined /*out*/;
            inputs["originHostHeader"] = undefined /*out*/;
            inputs["originPath"] = undefined /*out*/;
            inputs["origins"] = undefined /*out*/;
            inputs["probePath"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["queryStringCachingBehavior"] = undefined /*out*/;
            inputs["resourceState"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["urlSigningKeys"] = undefined /*out*/;
            inputs["webApplicationFirewallPolicyLink"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:cdn/latest:Endpoint" }, { type: "azurerm:cdn/v20150601:Endpoint" }, { type: "azurerm:cdn/v20160402:Endpoint" }, { type: "azurerm:cdn/v20161002:Endpoint" }, { type: "azurerm:cdn/v20170402:Endpoint" }, { type: "azurerm:cdn/v20171012:Endpoint" }, { type: "azurerm:cdn/v20190415:Endpoint" }, { type: "azurerm:cdn/v20190615:Endpoint" }, { type: "azurerm:cdn/v20190615preview:Endpoint" }, { type: "azurerm:cdn/v20191231:Endpoint" }, { type: "azurerm:cdn/v20200331:Endpoint" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Endpoint.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Endpoint resource.
 */
export interface EndpointArgs {
    /**
     * List of content types on which compression applies. The value should be a valid MIME type.
     */
    readonly contentTypesToCompress?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A reference to the origin group.
     */
    readonly defaultOriginGroup?: pulumi.Input<inputs.cdn.v20200415.ResourceReference>;
    /**
     * A policy that specifies the delivery rules to be used for an endpoint.
     */
    readonly deliveryPolicy?: pulumi.Input<inputs.cdn.v20200415.EndpointPropertiesUpdateParametersDeliveryPolicy>;
    /**
     * Name of the endpoint under the profile which is unique globally.
     */
    readonly endpointName: pulumi.Input<string>;
    /**
     * List of rules defining the user's geo access within a CDN endpoint. Each geo filter defines an access rule to a specified path or content, e.g. block APAC for path /pictures/
     */
    readonly geoFilters?: pulumi.Input<pulumi.Input<inputs.cdn.v20200415.GeoFilter>[]>;
    /**
     * Indicates whether content compression is enabled on CDN. Default value is false. If compression is enabled, content will be served as compressed if user requests for a compressed version. Content won't be compressed on CDN when requested content is smaller than 1 byte or larger than 1 MB.
     */
    readonly isCompressionEnabled?: pulumi.Input<boolean>;
    /**
     * Indicates whether HTTP traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
     */
    readonly isHttpAllowed?: pulumi.Input<boolean>;
    /**
     * Indicates whether HTTPS traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
     */
    readonly isHttpsAllowed?: pulumi.Input<boolean>;
    /**
     * Resource location.
     */
    readonly location: pulumi.Input<string>;
    /**
     * Specifies what scenario the customer wants this CDN endpoint to optimize for, e.g. Download, Media services. With this information, CDN can apply scenario driven optimization.
     */
    readonly optimizationType?: pulumi.Input<string>;
    /**
     * The origin groups comprising of origins that are used for load balancing the traffic based on availability.
     */
    readonly originGroups?: pulumi.Input<pulumi.Input<inputs.cdn.v20200415.DeepCreatedOriginGroup>[]>;
    /**
     * The host header value sent to the origin with each request. This property at Endpoint is only allowed when endpoint uses single origin and can be overridden by the same property specified at origin.If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default.
     */
    readonly originHostHeader?: pulumi.Input<string>;
    /**
     * A directory path on the origin that CDN can use to retrieve content from, e.g. contoso.cloudapp.net/originpath.
     */
    readonly originPath?: pulumi.Input<string>;
    /**
     * The source of the content being delivered via CDN.
     */
    readonly origins: pulumi.Input<pulumi.Input<inputs.cdn.v20200415.DeepCreatedOrigin>[]>;
    /**
     * Path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the origin path. This property is only relevant when using a single origin.
     */
    readonly probePath?: pulumi.Input<string>;
    /**
     * Name of the CDN profile which is unique within the resource group.
     */
    readonly profileName: pulumi.Input<string>;
    /**
     * Defines how CDN caches requests that include query strings. You can ignore any query strings when caching, bypass caching to prevent requests that contain query strings from being cached, or cache every request with a unique URL.
     */
    readonly queryStringCachingBehavior?: pulumi.Input<string>;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List of keys used to validate the signed URL hashes.
     */
    readonly urlSigningKeys?: pulumi.Input<pulumi.Input<inputs.cdn.v20200415.UrlSigningKey>[]>;
    /**
     * Defines the Web Application Firewall policy for the endpoint (if applicable)
     */
    readonly webApplicationFirewallPolicyLink?: pulumi.Input<inputs.cdn.v20200415.EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLink>;
}
