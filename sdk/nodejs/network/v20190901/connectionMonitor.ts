// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Information about the connection monitor.
 */
export class ConnectionMonitor extends pulumi.CustomResource {
    /**
     * Get an existing ConnectionMonitor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ConnectionMonitor {
        return new ConnectionMonitor(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20190901:ConnectionMonitor';

    /**
     * Returns true if the given object is an instance of ConnectionMonitor.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConnectionMonitor {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConnectionMonitor.__pulumiType;
    }

    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string | undefined>;
    /**
     * Connection monitor location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Name of the connection monitor.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Properties of the connection monitor result.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.network.v20190901.ConnectionMonitorResultPropertiesResponse>;
    /**
     * Connection monitor tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Connection monitor type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ConnectionMonitor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectionMonitorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConnectionMonitorArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ConnectionMonitorArgs | undefined;
            if (!args || args.destination === undefined) {
                throw new Error("Missing required property 'destination'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.networkWatcherName === undefined) {
                throw new Error("Missing required property 'networkWatcherName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.source === undefined) {
                throw new Error("Missing required property 'source'");
            }
            inputs["autoStart"] = args ? args.autoStart : undefined;
            inputs["destination"] = args ? args.destination : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["monitoringIntervalInSeconds"] = args ? args.monitoringIntervalInSeconds : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkWatcherName"] = args ? args.networkWatcherName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["source"] = args ? args.source : undefined;
            inputs["tags"] = args ? args.tags : undefined;
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
        super(ConnectionMonitor.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ConnectionMonitor resource.
 */
export interface ConnectionMonitorArgs {
    /**
     * Determines if the connection monitor will start automatically once created.
     */
    readonly autoStart?: pulumi.Input<boolean>;
    /**
     * Describes the destination of connection monitor.
     */
    readonly destination: pulumi.Input<inputs.network.v20190901.ConnectionMonitorDestination>;
    /**
     * Connection monitor location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Monitoring interval in seconds.
     */
    readonly monitoringIntervalInSeconds?: pulumi.Input<number>;
    /**
     * The name of the connection monitor.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The name of the Network Watcher resource.
     */
    readonly networkWatcherName: pulumi.Input<string>;
    /**
     * The name of the resource group containing Network Watcher.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Describes the source of connection monitor.
     */
    readonly source: pulumi.Input<inputs.network.v20190901.ConnectionMonitorSource>;
    /**
     * Connection monitor tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
