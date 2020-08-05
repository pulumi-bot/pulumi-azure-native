// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Hybrid Connection contract. This is used to configure a Hybrid Connection.
 */
export class WebAppHybridConnection extends pulumi.CustomResource {
    /**
     * Get an existing WebAppHybridConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WebAppHybridConnection {
        return new WebAppHybridConnection(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:web/v20180201:WebAppHybridConnection';

    /**
     * Returns true if the given object is an instance of WebAppHybridConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebAppHybridConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebAppHybridConnection.__pulumiType;
    }

    /**
     * Kind of resource.
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * Resource Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * HybridConnection resource specific properties
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.web.v20180201.HybridConnectionResponseProperties>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a WebAppHybridConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebAppHybridConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebAppHybridConnectionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as WebAppHybridConnectionArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.namespaceName === undefined) {
                throw new Error("Missing required property 'namespaceName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["hostname"] = args ? args.hostname : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namespaceName"] = args ? args.namespaceName : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["relayArmUri"] = args ? args.relayArmUri : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sendKeyName"] = args ? args.sendKeyName : undefined;
            inputs["sendKeyValue"] = args ? args.sendKeyValue : undefined;
            inputs["serviceBusNamespace"] = args ? args.serviceBusNamespace : undefined;
            inputs["serviceBusSuffix"] = args ? args.serviceBusSuffix : undefined;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(WebAppHybridConnection.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a WebAppHybridConnection resource.
 */
export interface WebAppHybridConnectionArgs {
    /**
     * The hostname of the endpoint.
     */
    readonly hostname?: pulumi.Input<string>;
    /**
     * Kind of resource.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * The name of the Service Bus relay.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The namespace for this hybrid connection.
     */
    readonly namespaceName: pulumi.Input<string>;
    /**
     * The port of the endpoint.
     */
    readonly port?: pulumi.Input<number>;
    /**
     * The ARM URI to the Service Bus relay.
     */
    readonly relayArmUri?: pulumi.Input<string>;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the Service Bus key which has Send permissions. This is used to authenticate to Service Bus.
     */
    readonly sendKeyName?: pulumi.Input<string>;
    /**
     * The value of the Service Bus key. This is used to authenticate to Service Bus. In ARM this key will not be returned
     * normally, use the POST /listKeys API instead.
     */
    readonly sendKeyValue?: pulumi.Input<string>;
    /**
     * The name of the Service Bus namespace.
     */
    readonly serviceBusNamespace?: pulumi.Input<string>;
    /**
     * The suffix for the service bus endpoint. By default this is .servicebus.windows.net
     */
    readonly serviceBusSuffix?: pulumi.Input<string>;
}
