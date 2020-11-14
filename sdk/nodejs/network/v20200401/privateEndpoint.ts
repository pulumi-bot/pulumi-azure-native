// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Private endpoint resource.
 */
export class PrivateEndpoint extends pulumi.CustomResource {
    /**
     * Get an existing PrivateEndpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PrivateEndpoint {
        return new PrivateEndpoint(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:network/v20200401:PrivateEndpoint';

    /**
     * Returns true if the given object is an instance of PrivateEndpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrivateEndpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrivateEndpoint.__pulumiType;
    }

    /**
     * An array of custom dns configurations.
     */
    public readonly customDnsConfigs!: pulumi.Output<outputs.network.v20200401.CustomDnsConfigPropertiesFormatResponse[] | undefined>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * A grouping of information about the connection to the remote resource. Used when the network admin does not have access to approve connections to the remote resource.
     */
    public readonly manualPrivateLinkServiceConnections!: pulumi.Output<outputs.network.v20200401.PrivateLinkServiceConnectionResponse[] | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * An array of references to the network interfaces created for this private endpoint.
     */
    public /*out*/ readonly networkInterfaces!: pulumi.Output<outputs.network.v20200401.NetworkInterfaceResponse[]>;
    /**
     * A grouping of information about the connection to the remote resource.
     */
    public readonly privateLinkServiceConnections!: pulumi.Output<outputs.network.v20200401.PrivateLinkServiceConnectionResponse[] | undefined>;
    /**
     * The provisioning state of the private endpoint resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The ID of the subnet from which the private IP will be allocated.
     */
    public readonly subnet!: pulumi.Output<outputs.network.v20200401.SubnetResponse | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a PrivateEndpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrivateEndpointArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.privateEndpointName === undefined) {
                throw new Error("Missing required property 'privateEndpointName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["customDnsConfigs"] = args ? args.customDnsConfigs : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["manualPrivateLinkServiceConnections"] = args ? args.manualPrivateLinkServiceConnections : undefined;
            inputs["privateEndpointName"] = args ? args.privateEndpointName : undefined;
            inputs["privateLinkServiceConnections"] = args ? args.privateLinkServiceConnections : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["subnet"] = args ? args.subnet : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["networkInterfaces"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["customDnsConfigs"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["manualPrivateLinkServiceConnections"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["networkInterfaces"] = undefined /*out*/;
            inputs["privateLinkServiceConnections"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["subnet"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:network/latest:PrivateEndpoint" }, { type: "azure-nextgen:network/v20190401:PrivateEndpoint" }, { type: "azure-nextgen:network/v20190601:PrivateEndpoint" }, { type: "azure-nextgen:network/v20190701:PrivateEndpoint" }, { type: "azure-nextgen:network/v20190801:PrivateEndpoint" }, { type: "azure-nextgen:network/v20190901:PrivateEndpoint" }, { type: "azure-nextgen:network/v20191101:PrivateEndpoint" }, { type: "azure-nextgen:network/v20191201:PrivateEndpoint" }, { type: "azure-nextgen:network/v20200301:PrivateEndpoint" }, { type: "azure-nextgen:network/v20200501:PrivateEndpoint" }, { type: "azure-nextgen:network/v20200601:PrivateEndpoint" }, { type: "azure-nextgen:network/v20200701:PrivateEndpoint" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(PrivateEndpoint.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a PrivateEndpoint resource.
 */
export interface PrivateEndpointArgs {
    /**
     * An array of custom dns configurations.
     */
    readonly customDnsConfigs?: pulumi.Input<pulumi.Input<inputs.network.v20200401.CustomDnsConfigPropertiesFormat>[]>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * A grouping of information about the connection to the remote resource. Used when the network admin does not have access to approve connections to the remote resource.
     */
    readonly manualPrivateLinkServiceConnections?: pulumi.Input<pulumi.Input<inputs.network.v20200401.PrivateLinkServiceConnection>[]>;
    /**
     * The name of the private endpoint.
     */
    readonly privateEndpointName: pulumi.Input<string>;
    /**
     * A grouping of information about the connection to the remote resource.
     */
    readonly privateLinkServiceConnections?: pulumi.Input<pulumi.Input<inputs.network.v20200401.PrivateLinkServiceConnection>[]>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The ID of the subnet from which the private IP will be allocated.
     */
    readonly subnet?: pulumi.Input<inputs.network.v20200401.Subnet>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
