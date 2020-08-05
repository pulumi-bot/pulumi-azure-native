// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * NetworkVirtualAppliance Resource.
 */
export class NetworkVirtualAppliance extends pulumi.CustomResource {
    /**
     * Get an existing NetworkVirtualAppliance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): NetworkVirtualAppliance {
        return new NetworkVirtualAppliance(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20200501:NetworkVirtualAppliance';

    /**
     * Returns true if the given object is an instance of NetworkVirtualAppliance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkVirtualAppliance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkVirtualAppliance.__pulumiType;
    }

    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The service principal that has read access to cloud-init and config blob.
     */
    public readonly identity!: pulumi.Output<outputs.network.v20200501.ManagedServiceIdentityResponse | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Properties of the Network Virtual Appliance.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.network.v20200501.NetworkVirtualAppliancePropertiesFormatResponse>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a NetworkVirtualAppliance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkVirtualApplianceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkVirtualApplianceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as NetworkVirtualApplianceArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["bootStrapConfigurationBlobs"] = args ? args.bootStrapConfigurationBlobs : undefined;
            inputs["cloudInitConfiguration"] = args ? args.cloudInitConfiguration : undefined;
            inputs["cloudInitConfigurationBlobs"] = args ? args.cloudInitConfigurationBlobs : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nvaSku"] = args ? args.nvaSku : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["virtualApplianceAsn"] = args ? args.virtualApplianceAsn : undefined;
            inputs["virtualHub"] = args ? args.virtualHub : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(NetworkVirtualAppliance.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a NetworkVirtualAppliance resource.
 */
export interface NetworkVirtualApplianceArgs {
    /**
     * BootStrapConfigurationBlobs storage URLs.
     */
    readonly bootStrapConfigurationBlobs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * CloudInitConfiguration string in plain text.
     */
    readonly cloudInitConfiguration?: pulumi.Input<string>;
    /**
     * CloudInitConfigurationBlob storage URLs.
     */
    readonly cloudInitConfigurationBlobs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * The service principal that has read access to cloud-init and config blob.
     */
    readonly identity?: pulumi.Input<inputs.network.v20200501.ManagedServiceIdentity>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of Network Virtual Appliance.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Network Virtual Appliance SKU.
     */
    readonly nvaSku?: pulumi.Input<inputs.network.v20200501.VirtualApplianceSkuProperties>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * VirtualAppliance ASN.
     */
    readonly virtualApplianceAsn?: pulumi.Input<number>;
    /**
     * The Virtual Hub where Network Virtual Appliance is being deployed.
     */
    readonly virtualHub?: pulumi.Input<inputs.network.v20200501.SubResource>;
}
