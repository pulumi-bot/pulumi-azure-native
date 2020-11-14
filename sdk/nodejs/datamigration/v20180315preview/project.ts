// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * A project resource
 */
export class Project extends pulumi.CustomResource {
    /**
     * Get an existing Project resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Project {
        return new Project(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:datamigration/v20180315preview:Project';

    /**
     * Returns true if the given object is an instance of Project.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Project {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Project.__pulumiType;
    }

    /**
     * UTC Date and time when project was created
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * List of DatabaseInfo
     */
    public readonly databasesInfo!: pulumi.Output<outputs.datamigration.v20180315preview.DatabaseInfoResponse[] | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The project's provisioning state
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Information for connecting to source
     */
    public readonly sourceConnectionInfo!: pulumi.Output<outputs.datamigration.v20180315preview.SqlConnectionInfoResponse | undefined>;
    /**
     * Source platform for the project
     */
    public readonly sourcePlatform!: pulumi.Output<string>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Information for connecting to target
     */
    public readonly targetConnectionInfo!: pulumi.Output<outputs.datamigration.v20180315preview.SqlConnectionInfoResponse | undefined>;
    /**
     * Target platform for the project
     */
    public readonly targetPlatform!: pulumi.Output<string>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Project resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.groupName === undefined) {
                throw new Error("Missing required property 'groupName'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.projectName === undefined) {
                throw new Error("Missing required property 'projectName'");
            }
            if (!args || args.serviceName === undefined) {
                throw new Error("Missing required property 'serviceName'");
            }
            if (!args || args.sourcePlatform === undefined) {
                throw new Error("Missing required property 'sourcePlatform'");
            }
            if (!args || args.targetPlatform === undefined) {
                throw new Error("Missing required property 'targetPlatform'");
            }
            inputs["databasesInfo"] = args ? args.databasesInfo : undefined;
            inputs["groupName"] = args ? args.groupName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["projectName"] = args ? args.projectName : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["sourceConnectionInfo"] = args ? args.sourceConnectionInfo : undefined;
            inputs["sourcePlatform"] = args ? args.sourcePlatform : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["targetConnectionInfo"] = args ? args.targetConnectionInfo : undefined;
            inputs["targetPlatform"] = args ? args.targetPlatform : undefined;
            inputs["creationTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["creationTime"] = undefined /*out*/;
            inputs["databasesInfo"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["sourceConnectionInfo"] = undefined /*out*/;
            inputs["sourcePlatform"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["targetConnectionInfo"] = undefined /*out*/;
            inputs["targetPlatform"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:datamigration/latest:Project" }, { type: "azure-nextgen:datamigration/v20171115preview:Project" }, { type: "azure-nextgen:datamigration/v20180331preview:Project" }, { type: "azure-nextgen:datamigration/v20180419:Project" }, { type: "azure-nextgen:datamigration/v20180715preview:Project" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Project.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Project resource.
 */
export interface ProjectArgs {
    /**
     * List of DatabaseInfo
     */
    readonly databasesInfo?: pulumi.Input<pulumi.Input<inputs.datamigration.v20180315preview.DatabaseInfo>[]>;
    /**
     * Name of the resource group
     */
    readonly groupName: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location: pulumi.Input<string>;
    /**
     * Name of the project
     */
    readonly projectName: pulumi.Input<string>;
    /**
     * Name of the service
     */
    readonly serviceName: pulumi.Input<string>;
    /**
     * Information for connecting to source
     */
    readonly sourceConnectionInfo?: pulumi.Input<inputs.datamigration.v20180315preview.SqlConnectionInfo>;
    /**
     * Source platform for the project
     */
    readonly sourcePlatform: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Information for connecting to target
     */
    readonly targetConnectionInfo?: pulumi.Input<inputs.datamigration.v20180315preview.SqlConnectionInfo>;
    /**
     * Target platform for the project
     */
    readonly targetPlatform: pulumi.Input<string>;
}
