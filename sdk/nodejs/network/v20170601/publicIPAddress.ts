// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Public IP address resource.
 */
export class PublicIPAddress extends pulumi.CustomResource {
    /**
     * Get an existing PublicIPAddress resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PublicIPAddress {
        return new PublicIPAddress(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20170601:PublicIPAddress';

    /**
     * Returns true if the given object is an instance of PublicIPAddress.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PublicIPAddress {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PublicIPAddress.__pulumiType;
    }

    /**
     * The FQDN of the DNS record associated with the public IP address.
     */
    public readonly dnsSettings!: pulumi.Output<outputs.network.v20170601.PublicIPAddressDnsSettingsResponse | undefined>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * The idle timeout of the public IP address.
     */
    public readonly idleTimeoutInMinutes!: pulumi.Output<number | undefined>;
    /**
     * The IP address associated with the public IP address resource.
     */
    public readonly ipAddress!: pulumi.Output<string | undefined>;
    /**
     * The IP configuration associated with the public IP address.
     */
    public /*out*/ readonly ipConfiguration!: pulumi.Output<outputs.network.v20170601.IPConfigurationResponse>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * The public IP address version. Possible values are: 'IPv4' and 'IPv6'.
     */
    public readonly publicIPAddressVersion!: pulumi.Output<string | undefined>;
    /**
     * The public IP allocation method. Possible values are: 'Static' and 'Dynamic'.
     */
    public readonly publicIPAllocationMethod!: pulumi.Output<string | undefined>;
    /**
     * The resource GUID property of the public IP resource.
     */
    public readonly resourceGuid!: pulumi.Output<string | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * A list of availability zones denoting the IP allocated for the resource needs to come from.
     */
    public readonly zones!: pulumi.Output<string[] | undefined>;

    /**
     * Create a PublicIPAddress resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PublicIPAddressArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PublicIPAddressArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as PublicIPAddressArgs | undefined;
            if (!args || args.publicIpAddressName === undefined) {
                throw new Error("Missing required property 'publicIpAddressName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["dnsSettings"] = args ? args.dnsSettings : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["idleTimeoutInMinutes"] = args ? args.idleTimeoutInMinutes : undefined;
            inputs["ipAddress"] = args ? args.ipAddress : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["publicIPAddressVersion"] = args ? args.publicIPAddressVersion : undefined;
            inputs["publicIPAllocationMethod"] = args ? args.publicIPAllocationMethod : undefined;
            inputs["publicIpAddressName"] = args ? args.publicIpAddressName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceGuid"] = args ? args.resourceGuid : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["zones"] = args ? args.zones : undefined;
            inputs["ipConfiguration"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:PublicIPAddress" }, { type: "azurerm:network/v20150501preview:PublicIPAddress" }, { type: "azurerm:network/v20150615:PublicIPAddress" }, { type: "azurerm:network/v20160330:PublicIPAddress" }, { type: "azurerm:network/v20160601:PublicIPAddress" }, { type: "azurerm:network/v20160901:PublicIPAddress" }, { type: "azurerm:network/v20161201:PublicIPAddress" }, { type: "azurerm:network/v20170301:PublicIPAddress" }, { type: "azurerm:network/v20170801:PublicIPAddress" }, { type: "azurerm:network/v20170901:PublicIPAddress" }, { type: "azurerm:network/v20171001:PublicIPAddress" }, { type: "azurerm:network/v20171101:PublicIPAddress" }, { type: "azurerm:network/v20180101:PublicIPAddress" }, { type: "azurerm:network/v20180201:PublicIPAddress" }, { type: "azurerm:network/v20180401:PublicIPAddress" }, { type: "azurerm:network/v20180601:PublicIPAddress" }, { type: "azurerm:network/v20180701:PublicIPAddress" }, { type: "azurerm:network/v20180801:PublicIPAddress" }, { type: "azurerm:network/v20181001:PublicIPAddress" }, { type: "azurerm:network/v20181101:PublicIPAddress" }, { type: "azurerm:network/v20181201:PublicIPAddress" }, { type: "azurerm:network/v20190201:PublicIPAddress" }, { type: "azurerm:network/v20190401:PublicIPAddress" }, { type: "azurerm:network/v20190601:PublicIPAddress" }, { type: "azurerm:network/v20190701:PublicIPAddress" }, { type: "azurerm:network/v20190801:PublicIPAddress" }, { type: "azurerm:network/v20190901:PublicIPAddress" }, { type: "azurerm:network/v20191101:PublicIPAddress" }, { type: "azurerm:network/v20191201:PublicIPAddress" }, { type: "azurerm:network/v20200301:PublicIPAddress" }, { type: "azurerm:network/v20200401:PublicIPAddress" }, { type: "azurerm:network/v20200501:PublicIPAddress" }, { type: "azurerm:network/v20200601:PublicIPAddress" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(PublicIPAddress.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a PublicIPAddress resource.
 */
export interface PublicIPAddressArgs {
    /**
     * The FQDN of the DNS record associated with the public IP address.
     */
    readonly dnsSettings?: pulumi.Input<inputs.network.v20170601.PublicIPAddressDnsSettings>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * The idle timeout of the public IP address.
     */
    readonly idleTimeoutInMinutes?: pulumi.Input<number>;
    /**
     * The IP address associated with the public IP address resource.
     */
    readonly ipAddress?: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * The public IP address version. Possible values are: 'IPv4' and 'IPv6'.
     */
    readonly publicIPAddressVersion?: pulumi.Input<string>;
    /**
     * The public IP allocation method. Possible values are: 'Static' and 'Dynamic'.
     */
    readonly publicIPAllocationMethod?: pulumi.Input<string>;
    /**
     * The name of the public IP address.
     */
    readonly publicIpAddressName: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The resource GUID property of the public IP resource.
     */
    readonly resourceGuid?: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A list of availability zones denoting the IP allocated for the resource needs to come from.
     */
    readonly zones?: pulumi.Input<pulumi.Input<string>[]>;
}
