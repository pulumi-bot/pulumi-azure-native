// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Network watcher in a resource group.
 */
export class NetworkWatcher extends pulumi.CustomResource {
    /**
     * Get an existing NetworkWatcher resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): NetworkWatcher {
        return new NetworkWatcher(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20190201:NetworkWatcher';

    /**
     * Returns true if the given object is an instance of NetworkWatcher.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkWatcher {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkWatcher.__pulumiType;
    }

    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The provisioning state of the resource.
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a NetworkWatcher resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkWatcherArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.networkWatcherName === undefined) {
                throw new Error("Missing required property 'networkWatcherName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["etag"] = args ? args.etag : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["networkWatcherName"] = args ? args.networkWatcherName : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["etag"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:NetworkWatcher" }, { type: "azurerm:network/v20160901:NetworkWatcher" }, { type: "azurerm:network/v20161201:NetworkWatcher" }, { type: "azurerm:network/v20170301:NetworkWatcher" }, { type: "azurerm:network/v20170601:NetworkWatcher" }, { type: "azurerm:network/v20170801:NetworkWatcher" }, { type: "azurerm:network/v20170901:NetworkWatcher" }, { type: "azurerm:network/v20171001:NetworkWatcher" }, { type: "azurerm:network/v20171101:NetworkWatcher" }, { type: "azurerm:network/v20180101:NetworkWatcher" }, { type: "azurerm:network/v20180201:NetworkWatcher" }, { type: "azurerm:network/v20180401:NetworkWatcher" }, { type: "azurerm:network/v20180601:NetworkWatcher" }, { type: "azurerm:network/v20180701:NetworkWatcher" }, { type: "azurerm:network/v20180801:NetworkWatcher" }, { type: "azurerm:network/v20181001:NetworkWatcher" }, { type: "azurerm:network/v20181101:NetworkWatcher" }, { type: "azurerm:network/v20181201:NetworkWatcher" }, { type: "azurerm:network/v20190401:NetworkWatcher" }, { type: "azurerm:network/v20190601:NetworkWatcher" }, { type: "azurerm:network/v20190701:NetworkWatcher" }, { type: "azurerm:network/v20190801:NetworkWatcher" }, { type: "azurerm:network/v20190901:NetworkWatcher" }, { type: "azurerm:network/v20191101:NetworkWatcher" }, { type: "azurerm:network/v20191201:NetworkWatcher" }, { type: "azurerm:network/v20200301:NetworkWatcher" }, { type: "azurerm:network/v20200401:NetworkWatcher" }, { type: "azurerm:network/v20200501:NetworkWatcher" }, { type: "azurerm:network/v20200601:NetworkWatcher" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(NetworkWatcher.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a NetworkWatcher resource.
 */
export interface NetworkWatcherArgs {
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the network watcher.
     */
    readonly networkWatcherName: pulumi.Input<string>;
    /**
     * The provisioning state of the resource.
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
