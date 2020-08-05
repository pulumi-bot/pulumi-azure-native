// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Properties of an artifact source.
 */
export class ArtifactSource extends pulumi.CustomResource {
    /**
     * Get an existing ArtifactSource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ArtifactSource {
        return new ArtifactSource(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:devtestlab/v20180915:ArtifactSource';

    /**
     * Returns true if the given object is an instance of ArtifactSource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ArtifactSource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ArtifactSource.__pulumiType;
    }

    /**
     * The location of the resource.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The properties of the resource.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.devtestlab.v20180915.ArtifactSourcePropertiesResponse>;
    /**
     * The tags of the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ArtifactSource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ArtifactSourceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ArtifactSourceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ArtifactSourceArgs | undefined;
            if (!args || args.labName === undefined) {
                throw new Error("Missing required property 'labName'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["armTemplateFolderPath"] = args ? args.armTemplateFolderPath : undefined;
            inputs["branchRef"] = args ? args.branchRef : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["folderPath"] = args ? args.folderPath : undefined;
            inputs["labName"] = args ? args.labName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["securityToken"] = args ? args.securityToken : undefined;
            inputs["sourceType"] = args ? args.sourceType : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["uri"] = args ? args.uri : undefined;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ArtifactSource.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ArtifactSource resource.
 */
export interface ArtifactSourceArgs {
    /**
     * The folder containing Azure Resource Manager templates.
     */
    readonly armTemplateFolderPath?: pulumi.Input<string>;
    /**
     * The artifact source's branch reference.
     */
    readonly branchRef?: pulumi.Input<string>;
    /**
     * The artifact source's display name.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The folder containing artifacts.
     */
    readonly folderPath?: pulumi.Input<string>;
    /**
     * The name of the lab.
     */
    readonly labName: pulumi.Input<string>;
    /**
     * The location of the resource.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the artifact source.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The security token to authenticate to the artifact source.
     */
    readonly securityToken?: pulumi.Input<string>;
    /**
     * The artifact source's type.
     */
    readonly sourceType?: pulumi.Input<string>;
    /**
     * Indicates if the artifact source is enabled (values: Enabled, Disabled).
     */
    readonly status?: pulumi.Input<string>;
    /**
     * The tags of the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The artifact source's URI.
     */
    readonly uri?: pulumi.Input<string>;
}
