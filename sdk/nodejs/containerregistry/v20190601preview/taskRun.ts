// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The task run that has the ARM resource and properties.
 * The task run will have the information of request and result of a run.
 */
export class TaskRun extends pulumi.CustomResource {
    /**
     * Get an existing TaskRun resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TaskRun {
        return new TaskRun(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:containerregistry/v20190601preview:TaskRun';

    /**
     * Returns true if the given object is an instance of TaskRun.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TaskRun {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TaskRun.__pulumiType;
    }

    /**
     * How the run should be forced to rerun even if the run request configuration has not changed
     */
    public readonly forceUpdateTag!: pulumi.Output<string | undefined>;
    /**
     * Identity for the resource.
     */
    public readonly identity!: pulumi.Output<outputs.containerregistry.v20190601preview.IdentityPropertiesResponse | undefined>;
    /**
     * The location of the resource
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The provisioning state of this task run
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The request (parameters) for the run
     */
    public readonly runRequest!: pulumi.Output<outputs.containerregistry.v20190601preview.RunRequestResponse | undefined>;
    /**
     * The result of this task run
     */
    public /*out*/ readonly runResult!: pulumi.Output<outputs.containerregistry.v20190601preview.RunResponse>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a TaskRun resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TaskRunArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TaskRunArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as TaskRunArgs | undefined;
            if (!args || args.registryName === undefined) {
                throw new Error("Missing required property 'registryName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.taskRunName === undefined) {
                throw new Error("Missing required property 'taskRunName'");
            }
            inputs["forceUpdateTag"] = args ? args.forceUpdateTag : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["registryName"] = args ? args.registryName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["runRequest"] = args ? args.runRequest : undefined;
            inputs["taskRunName"] = args ? args.taskRunName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["runResult"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(TaskRun.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a TaskRun resource.
 */
export interface TaskRunArgs {
    /**
     * How the run should be forced to rerun even if the run request configuration has not changed
     */
    readonly forceUpdateTag?: pulumi.Input<string>;
    /**
     * Identity for the resource.
     */
    readonly identity?: pulumi.Input<inputs.containerregistry.v20190601preview.IdentityProperties>;
    /**
     * The location of the resource
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the container registry.
     */
    readonly registryName: pulumi.Input<string>;
    /**
     * The name of the resource group to which the container registry belongs.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The request (parameters) for the run
     */
    readonly runRequest?: pulumi.Input<inputs.containerregistry.v20190601preview.RunRequest>;
    /**
     * The name of the task run.
     */
    readonly taskRunName: pulumi.Input<string>;
}
