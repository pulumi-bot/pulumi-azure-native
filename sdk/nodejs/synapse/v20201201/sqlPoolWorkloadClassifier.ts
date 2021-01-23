// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Workload classifier operations for a data warehouse
 */
export class SqlPoolWorkloadClassifier extends pulumi.CustomResource {
    /**
     * Get an existing SqlPoolWorkloadClassifier resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SqlPoolWorkloadClassifier {
        return new SqlPoolWorkloadClassifier(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:synapse/v20201201:SqlPoolWorkloadClassifier';

    /**
     * Returns true if the given object is an instance of SqlPoolWorkloadClassifier.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SqlPoolWorkloadClassifier {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SqlPoolWorkloadClassifier.__pulumiType;
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
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The workload classifier start time for classification.
     */
    public readonly startTime!: pulumi.Output<string | undefined>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a SqlPoolWorkloadClassifier resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SqlPoolWorkloadClassifierArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.memberName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'memberName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.sqlPoolName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'sqlPoolName'");
            }
            if ((!args || args.workloadClassifierName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'workloadClassifierName'");
            }
            if ((!args || args.workloadGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'workloadGroupName'");
            }
            if ((!args || args.workspaceName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'workspaceName'");
            }
            inputs["context"] = args ? args.context : undefined;
            inputs["endTime"] = args ? args.endTime : undefined;
            inputs["importance"] = args ? args.importance : undefined;
            inputs["label"] = args ? args.label : undefined;
            inputs["memberName"] = args ? args.memberName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sqlPoolName"] = args ? args.sqlPoolName : undefined;
            inputs["startTime"] = args ? args.startTime : undefined;
            inputs["workloadClassifierName"] = args ? args.workloadClassifierName : undefined;
            inputs["workloadGroupName"] = args ? args.workloadGroupName : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
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
        const aliasOpts = { aliases: [{ type: "azure-nextgen:synapse/latest:SqlPoolWorkloadClassifier" }, { type: "azure-nextgen:synapse/v20190601preview:SqlPoolWorkloadClassifier" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(SqlPoolWorkloadClassifier.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a SqlPoolWorkloadClassifier resource.
 */
export interface SqlPoolWorkloadClassifierArgs {
    /**
     * The workload classifier context.
     */
    readonly context?: pulumi.Input<string>;
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
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * SQL pool name
     */
    readonly sqlPoolName: pulumi.Input<string>;
    /**
     * The workload classifier start time for classification.
     */
    readonly startTime?: pulumi.Input<string>;
    /**
     * The name of the workload classifier.
     */
    readonly workloadClassifierName: pulumi.Input<string>;
    /**
     * The name of the workload group.
     */
    readonly workloadGroupName: pulumi.Input<string>;
    /**
     * The name of the workspace
     */
    readonly workspaceName: pulumi.Input<string>;
}
