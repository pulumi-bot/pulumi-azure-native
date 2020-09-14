// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Workload classifier operations for a data warehouse
 */
export class WorkloadClassifier extends pulumi.CustomResource {
    /**
     * Get an existing WorkloadClassifier resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WorkloadClassifier {
        return new WorkloadClassifier(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:sql/v20190601preview:WorkloadClassifier';

    /**
     * Returns true if the given object is an instance of WorkloadClassifier.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WorkloadClassifier {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WorkloadClassifier.__pulumiType;
    }

    /**
     * The workload classifier context.
     */
    public readonly context!: pulumi.Output<string | undefined>;
    /**
     * The workload classifier end time for classification.
     */
    public readonly endTime!: pulumi.Output<string | undefined>;
    /**
     * The workload classifier importance.
     */
    public readonly importance!: pulumi.Output<string | undefined>;
    /**
     * The workload classifier label.
     */
    public readonly label!: pulumi.Output<string | undefined>;
    /**
     * The workload classifier member name.
     */
    public readonly memberName!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The workload classifier start time for classification.
     */
    public readonly startTime!: pulumi.Output<string | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a WorkloadClassifier resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkloadClassifierArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.databaseName === undefined) {
                throw new Error("Missing required property 'databaseName'");
            }
            if (!args || args.memberName === undefined) {
                throw new Error("Missing required property 'memberName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serverName === undefined) {
                throw new Error("Missing required property 'serverName'");
            }
            if (!args || args.workloadClassifierName === undefined) {
                throw new Error("Missing required property 'workloadClassifierName'");
            }
            if (!args || args.workloadGroupName === undefined) {
                throw new Error("Missing required property 'workloadGroupName'");
            }
            inputs["context"] = args ? args.context : undefined;
            inputs["databaseName"] = args ? args.databaseName : undefined;
            inputs["endTime"] = args ? args.endTime : undefined;
            inputs["importance"] = args ? args.importance : undefined;
            inputs["label"] = args ? args.label : undefined;
            inputs["memberName"] = args ? args.memberName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serverName"] = args ? args.serverName : undefined;
            inputs["startTime"] = args ? args.startTime : undefined;
            inputs["workloadClassifierName"] = args ? args.workloadClassifierName : undefined;
            inputs["workloadGroupName"] = args ? args.workloadGroupName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["context"] = undefined /*out*/;
            inputs["endTime"] = undefined /*out*/;
            inputs["importance"] = undefined /*out*/;
            inputs["label"] = undefined /*out*/;
            inputs["memberName"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["startTime"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:sql/latest:WorkloadClassifier" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(WorkloadClassifier.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a WorkloadClassifier resource.
 */
export interface WorkloadClassifierArgs {
    /**
     * The workload classifier context.
     */
    readonly context?: pulumi.Input<string>;
    /**
     * The name of the database.
     */
    readonly databaseName: pulumi.Input<string>;
    /**
     * The workload classifier end time for classification.
     */
    readonly endTime?: pulumi.Input<string>;
    /**
     * The workload classifier importance.
     */
    readonly importance?: pulumi.Input<string>;
    /**
     * The workload classifier label.
     */
    readonly label?: pulumi.Input<string>;
    /**
     * The workload classifier member name.
     */
    readonly memberName: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the server.
     */
    readonly serverName: pulumi.Input<string>;
    /**
     * The workload classifier start time for classification.
     */
    readonly startTime?: pulumi.Input<string>;
    /**
     * The name of the workload classifier to create/update.
     */
    readonly workloadClassifierName: pulumi.Input<string>;
    /**
     * The name of the workload group from which to receive the classifier from.
     */
    readonly workloadGroupName: pulumi.Input<string>;
}
