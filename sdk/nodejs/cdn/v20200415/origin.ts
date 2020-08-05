// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * CDN origin is the source of the content being delivered via CDN. When the edge nodes represented by an endpoint do not have the requested content cached, they attempt to fetch it from one or more of the configured origins.
 */
export class Origin extends pulumi.CustomResource {
    /**
     * Get an existing Origin resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Origin {
        return new Origin(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:cdn/v20200415:Origin';

    /**
     * Returns true if the given object is an instance of Origin.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Origin {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Origin.__pulumiType;
    }

    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The JSON object that contains the properties of the origin.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.cdn.v20200415.OriginPropertiesResponse>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Origin resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OriginArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OriginArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as OriginArgs | undefined;
            if (!args || args.endpointName === undefined) {
                throw new Error("Missing required property 'endpointName'");
            }
            if (!args || args.hostName === undefined) {
                throw new Error("Missing required property 'hostName'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.profileName === undefined) {
                throw new Error("Missing required property 'profileName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["endpointName"] = args ? args.endpointName : undefined;
            inputs["hostName"] = args ? args.hostName : undefined;
            inputs["httpPort"] = args ? args.httpPort : undefined;
            inputs["httpsPort"] = args ? args.httpsPort : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["originHostHeader"] = args ? args.originHostHeader : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["privateLinkAlias"] = args ? args.privateLinkAlias : undefined;
            inputs["privateLinkApprovalMessage"] = args ? args.privateLinkApprovalMessage : undefined;
            inputs["privateLinkLocation"] = args ? args.privateLinkLocation : undefined;
            inputs["privateLinkResourceId"] = args ? args.privateLinkResourceId : undefined;
            inputs["profileName"] = args ? args.profileName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["weight"] = args ? args.weight : undefined;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Origin.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Origin resource.
 */
export interface OriginArgs {
    /**
     * Origin is enabled for load balancing or not
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Name of the endpoint under the profile which is unique globally.
     */
    readonly endpointName: pulumi.Input<string>;
    /**
     * The address of the origin. Domain names, IPv4 addresses, and IPv6 addresses are supported.This should be unique across all origins in an endpoint.
     */
    readonly hostName: pulumi.Input<string>;
    /**
     * The value of the HTTP port. Must be between 1 and 65535.
     */
    readonly httpPort?: pulumi.Input<number>;
    /**
     * The value of the HTTPS port. Must be between 1 and 65535.
     */
    readonly httpsPort?: pulumi.Input<number>;
    /**
     * Resource location.
     */
    readonly location: pulumi.Input<string>;
    /**
     * Name of the origin that is unique within the endpoint.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default. This overrides the host header defined at Endpoint
     */
    readonly originHostHeader?: pulumi.Input<string>;
    /**
     * Priority of origin in given origin group for load balancing. Higher priorities will not be used for load balancing if any lower priority origin is healthy.Must be between 1 and 5
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * The Alias of the Private Link resource. Populating this optional field indicates that this origin is 'Private'
     */
    readonly privateLinkAlias?: pulumi.Input<string>;
    /**
     * A custom message to be included in the approval request to connect to the Private Link.
     */
    readonly privateLinkApprovalMessage?: pulumi.Input<string>;
    /**
     * The location of the Private Link resource. Required only if 'privateLinkResourceId' is populated
     */
    readonly privateLinkLocation?: pulumi.Input<string>;
    /**
     * The Resource Id of the Private Link resource. Populating this optional field indicates that this backend is 'Private'
     */
    readonly privateLinkResourceId?: pulumi.Input<string>;
    /**
     * Name of the CDN profile which is unique within the resource group.
     */
    readonly profileName: pulumi.Input<string>;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Weight of the origin in given origin group for load balancing. Must be between 1 and 1000
     */
    readonly weight?: pulumi.Input<number>;
}
