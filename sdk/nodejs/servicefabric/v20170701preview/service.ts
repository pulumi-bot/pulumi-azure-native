// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The service resource.
 */
export class Service extends pulumi.CustomResource {
    /**
     * Get an existing Service resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Service {
        return new Service(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:servicefabric/v20170701preview:Service';

    /**
     * Returns true if the given object is an instance of Service.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Service {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Service.__pulumiType;
    }

    /**
     * A list that describes the correlation of the service with other services.
     */
    public readonly correlationScheme!: pulumi.Output<outputs.servicefabric.v20170701preview.ServiceCorrelationDescriptionResponse[] | undefined>;
    /**
     * Specifies the move cost for the service.
     */
    public readonly defaultMoveCost!: pulumi.Output<string | undefined>;
    /**
     * Azure resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Azure resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Describes how the service is partitioned.
     */
    public readonly partitionDescription!: pulumi.Output<outputs.servicefabric.v20170701preview.PartitionSchemeDescriptionResponse | undefined>;
    /**
     * The placement constraints as a string. Placement constraints are boolean expressions on node properties and allow for restricting a service to particular nodes based on the service requirements. For example, to place a service on nodes where NodeType is blue specify the following: "NodeColor == blue)".
     */
    public readonly placementConstraints!: pulumi.Output<string | undefined>;
    /**
     * The current deployment or provisioning state, which only appears in the response
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The kind of service (Stateless or Stateful).
     */
    public readonly serviceKind!: pulumi.Output<string>;
    /**
     * The service load metrics is given as an array of ServiceLoadMetricDescription objects.
     */
    public readonly serviceLoadMetrics!: pulumi.Output<outputs.servicefabric.v20170701preview.ServiceLoadMetricDescriptionResponse[] | undefined>;
    /**
     * A list that describes the correlation of the service with other services.
     */
    public readonly servicePlacementPolicies!: pulumi.Output<outputs.servicefabric.v20170701preview.ServicePlacementPolicyDescriptionResponse[] | undefined>;
    /**
     * The name of the service type
     */
    public readonly serviceTypeName!: pulumi.Output<string | undefined>;
    /**
     * Azure resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Service resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ServiceArgs | undefined;
            if (!args || args.applicationName === undefined) {
                throw new Error("Missing required property 'applicationName'");
            }
            if (!args || args.clusterName === undefined) {
                throw new Error("Missing required property 'clusterName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serviceKind === undefined) {
                throw new Error("Missing required property 'serviceKind'");
            }
            if (!args || args.serviceName === undefined) {
                throw new Error("Missing required property 'serviceName'");
            }
            inputs["applicationName"] = args ? args.applicationName : undefined;
            inputs["clusterName"] = args ? args.clusterName : undefined;
            inputs["correlationScheme"] = args ? args.correlationScheme : undefined;
            inputs["defaultMoveCost"] = args ? args.defaultMoveCost : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["partitionDescription"] = args ? args.partitionDescription : undefined;
            inputs["placementConstraints"] = args ? args.placementConstraints : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serviceKind"] = args ? args.serviceKind : undefined;
            inputs["serviceLoadMetrics"] = args ? args.serviceLoadMetrics : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["servicePlacementPolicies"] = args ? args.servicePlacementPolicies : undefined;
            inputs["serviceTypeName"] = args ? args.serviceTypeName : undefined;
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
        const aliasOpts = { aliases: [{ type: "azurerm:servicefabric/latest:Service" }, { type: "azurerm:servicefabric/v20190301:Service" }, { type: "azurerm:servicefabric/v20190301preview:Service" }, { type: "azurerm:servicefabric/v20190601preview:Service" }, { type: "azurerm:servicefabric/v20191101preview:Service" }, { type: "azurerm:servicefabric/v20200301:Service" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Service.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Service resource.
 */
export interface ServiceArgs {
    /**
     * The name of the application resource.
     */
    readonly applicationName: pulumi.Input<string>;
    /**
     * The name of the cluster resource.
     */
    readonly clusterName: pulumi.Input<string>;
    /**
     * A list that describes the correlation of the service with other services.
     */
    readonly correlationScheme?: pulumi.Input<pulumi.Input<inputs.servicefabric.v20170701preview.ServiceCorrelationDescription>[]>;
    /**
     * Specifies the move cost for the service.
     */
    readonly defaultMoveCost?: pulumi.Input<string>;
    /**
     * Azure resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Describes how the service is partitioned.
     */
    readonly partitionDescription?: pulumi.Input<inputs.servicefabric.v20170701preview.PartitionSchemeDescription>;
    /**
     * The placement constraints as a string. Placement constraints are boolean expressions on node properties and allow for restricting a service to particular nodes based on the service requirements. For example, to place a service on nodes where NodeType is blue specify the following: "NodeColor == blue)".
     */
    readonly placementConstraints?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The kind of service (Stateless or Stateful).
     */
    readonly serviceKind: pulumi.Input<string>;
    /**
     * The service load metrics is given as an array of ServiceLoadMetricDescription objects.
     */
    readonly serviceLoadMetrics?: pulumi.Input<pulumi.Input<inputs.servicefabric.v20170701preview.ServiceLoadMetricDescription>[]>;
    /**
     * The name of the service resource in the format of {applicationName}~{serviceName}.
     */
    readonly serviceName: pulumi.Input<string>;
    /**
     * A list that describes the correlation of the service with other services.
     */
    readonly servicePlacementPolicies?: pulumi.Input<pulumi.Input<inputs.servicefabric.v20170701preview.ServicePlacementPolicyDescription>[]>;
    /**
     * The name of the service type
     */
    readonly serviceTypeName?: pulumi.Input<string>;
}
