// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * ## Example Usage
 * ### Create JIT network access policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const jitNetworkAccessPolicy = new azurerm.security.v20200101.JitNetworkAccessPolicy("jitNetworkAccessPolicy", {
 *     ascLocation: "westeurope",
 *     jitNetworkAccessPolicyName: "default",
 *     kind: "Basic",
 *     requests: [{
 *         requestor: "barbara@contoso.com",
 *         startTimeUtc: "2018-05-17T08:06:45.5691611Z",
 *         virtualMachines: [{
 *             id: "/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Compute/virtualMachines/vm1",
 *             ports: [{
 *                 allowedSourceAddressPrefix: "192.127.0.2",
 *                 endTimeUtc: "2018-05-17T09:06:45.5691611Z",
 *                 number: 3389,
 *                 status: "Initiated",
 *                 statusReason: "UserRequested",
 *             }],
 *         }],
 *     }],
 *     resourceGroupName: "myRg1",
 *     virtualMachines: [{
 *         id: "/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg1/providers/Microsoft.Compute/virtualMachines/vm1",
 *         ports: [
 *             {
 *                 allowedSourceAddressPrefix: "*",
 *                 maxRequestAccessDuration: "PT3H",
 *                 number: 22,
 *                 protocol: "*",
 *             },
 *             {
 *                 allowedSourceAddressPrefix: "*",
 *                 maxRequestAccessDuration: "PT3H",
 *                 number: 3389,
 *                 protocol: "*",
 *             },
 *         ],
 *     }],
 * });
 *
 * ```
 */
export class JitNetworkAccessPolicy extends pulumi.CustomResource {
    /**
     * Get an existing JitNetworkAccessPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): JitNetworkAccessPolicy {
        return new JitNetworkAccessPolicy(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:security/v20200101:JitNetworkAccessPolicy';

    /**
     * Returns true if the given object is an instance of JitNetworkAccessPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is JitNetworkAccessPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === JitNetworkAccessPolicy.__pulumiType;
    }

    /**
     * Kind of the resource
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * Location where the resource is stored
     */
    public /*out*/ readonly location!: pulumi.Output<string>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Gets the provisioning state of the Just-in-Time policy.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    public readonly requests!: pulumi.Output<outputs.security.v20200101.JitNetworkAccessRequestResponse[] | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Configurations for Microsoft.Compute/virtualMachines resource type.
     */
    public readonly virtualMachines!: pulumi.Output<outputs.security.v20200101.JitNetworkAccessPolicyVirtualMachineResponse[]>;

    /**
     * Create a JitNetworkAccessPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: JitNetworkAccessPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: JitNetworkAccessPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as JitNetworkAccessPolicyArgs | undefined;
            if (!args || args.ascLocation === undefined) {
                throw new Error("Missing required property 'ascLocation'");
            }
            if (!args || args.jitNetworkAccessPolicyName === undefined) {
                throw new Error("Missing required property 'jitNetworkAccessPolicyName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.virtualMachines === undefined) {
                throw new Error("Missing required property 'virtualMachines'");
            }
            inputs["ascLocation"] = args ? args.ascLocation : undefined;
            inputs["jitNetworkAccessPolicyName"] = args ? args.jitNetworkAccessPolicyName : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["requests"] = args ? args.requests : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["virtualMachines"] = args ? args.virtualMachines : undefined;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:security/latest:JitNetworkAccessPolicy" }, { type: "azurerm:security/v20150601preview:JitNetworkAccessPolicy" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(JitNetworkAccessPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a JitNetworkAccessPolicy resource.
 */
export interface JitNetworkAccessPolicyArgs {
    /**
     * The location where ASC stores the data of the subscription. can be retrieved from Get locations
     */
    readonly ascLocation: pulumi.Input<string>;
    /**
     * Name of a Just-in-Time access configuration policy.
     */
    readonly jitNetworkAccessPolicyName: pulumi.Input<string>;
    /**
     * Kind of the resource
     */
    readonly kind?: pulumi.Input<string>;
    readonly requests?: pulumi.Input<pulumi.Input<inputs.security.v20200101.JitNetworkAccessRequest>[]>;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Configurations for Microsoft.Compute/virtualMachines resource type.
     */
    readonly virtualMachines: pulumi.Input<pulumi.Input<inputs.security.v20200101.JitNetworkAccessPolicyVirtualMachine>[]>;
}
