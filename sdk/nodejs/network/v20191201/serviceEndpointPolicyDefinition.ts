// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Service Endpoint policy definitions.
 *
 * ## Example Usage
 * ### Create service endpoint policy definition
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const serviceEndpointPolicyDefinition = new azurerm.network.v20191201.ServiceEndpointPolicyDefinition("serviceEndpointPolicyDefinition", {
 *     description: "Storage Service EndpointPolicy Definition",
 *     resourceGroupName: "rg1",
 *     service: "Microsoft.Storage",
 *     serviceEndpointPolicyDefinitionName: "testDefinition",
 *     serviceEndpointPolicyName: "testPolicy",
 *     serviceResources: [
 *         "/subscriptions/subid1",
 *         "/subscriptions/subid1/resourceGroups/storageRg",
 *         "/subscriptions/subid1/resourceGroups/storageRg/providers/Microsoft.Storage/storageAccounts/stAccount",
 *     ],
 * });
 *
 * ```
 */
export class ServiceEndpointPolicyDefinition extends pulumi.CustomResource {
    /**
     * Get an existing ServiceEndpointPolicyDefinition resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ServiceEndpointPolicyDefinition {
        return new ServiceEndpointPolicyDefinition(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20191201:ServiceEndpointPolicyDefinition';

    /**
     * Returns true if the given object is an instance of ServiceEndpointPolicyDefinition.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceEndpointPolicyDefinition {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceEndpointPolicyDefinition.__pulumiType;
    }

    /**
     * A description for this rule. Restricted to 140 chars.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The provisioning state of the service endpoint policy definition resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Service endpoint name.
     */
    public readonly service!: pulumi.Output<string | undefined>;
    /**
     * A list of service resources.
     */
    public readonly serviceResources!: pulumi.Output<string[] | undefined>;

    /**
     * Create a ServiceEndpointPolicyDefinition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceEndpointPolicyDefinitionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serviceEndpointPolicyDefinitionName === undefined) {
                throw new Error("Missing required property 'serviceEndpointPolicyDefinitionName'");
            }
            if (!args || args.serviceEndpointPolicyName === undefined) {
                throw new Error("Missing required property 'serviceEndpointPolicyName'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["service"] = args ? args.service : undefined;
            inputs["serviceEndpointPolicyDefinitionName"] = args ? args.serviceEndpointPolicyDefinitionName : undefined;
            inputs["serviceEndpointPolicyName"] = args ? args.serviceEndpointPolicyName : undefined;
            inputs["serviceResources"] = args ? args.serviceResources : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
        } else {
            inputs["description"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["service"] = undefined /*out*/;
            inputs["serviceResources"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:ServiceEndpointPolicyDefinition" }, { type: "azurerm:network/v20180701:ServiceEndpointPolicyDefinition" }, { type: "azurerm:network/v20180801:ServiceEndpointPolicyDefinition" }, { type: "azurerm:network/v20181001:ServiceEndpointPolicyDefinition" }, { type: "azurerm:network/v20181101:ServiceEndpointPolicyDefinition" }, { type: "azurerm:network/v20181201:ServiceEndpointPolicyDefinition" }, { type: "azurerm:network/v20190201:ServiceEndpointPolicyDefinition" }, { type: "azurerm:network/v20190401:ServiceEndpointPolicyDefinition" }, { type: "azurerm:network/v20190601:ServiceEndpointPolicyDefinition" }, { type: "azurerm:network/v20190701:ServiceEndpointPolicyDefinition" }, { type: "azurerm:network/v20190801:ServiceEndpointPolicyDefinition" }, { type: "azurerm:network/v20190901:ServiceEndpointPolicyDefinition" }, { type: "azurerm:network/v20191101:ServiceEndpointPolicyDefinition" }, { type: "azurerm:network/v20200301:ServiceEndpointPolicyDefinition" }, { type: "azurerm:network/v20200401:ServiceEndpointPolicyDefinition" }, { type: "azurerm:network/v20200501:ServiceEndpointPolicyDefinition" }, { type: "azurerm:network/v20200601:ServiceEndpointPolicyDefinition" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ServiceEndpointPolicyDefinition.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ServiceEndpointPolicyDefinition resource.
 */
export interface ServiceEndpointPolicyDefinitionArgs {
    /**
     * A description for this rule. Restricted to 140 chars.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * The name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Service endpoint name.
     */
    readonly service?: pulumi.Input<string>;
    /**
     * The name of the service endpoint policy definition name.
     */
    readonly serviceEndpointPolicyDefinitionName: pulumi.Input<string>;
    /**
     * The name of the service endpoint policy.
     */
    readonly serviceEndpointPolicyName: pulumi.Input<string>;
    /**
     * A list of service resources.
     */
    readonly serviceResources?: pulumi.Input<pulumi.Input<string>[]>;
}
