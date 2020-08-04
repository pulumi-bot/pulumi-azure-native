// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Dedicated cloud service model
 */
export class DedicatedCloudService extends pulumi.CustomResource {
    /**
     * Get an existing DedicatedCloudService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DedicatedCloudService {
        return new DedicatedCloudService(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:vmwarecloudsimple/v20190401:DedicatedCloudService';

    /**
     * Returns true if the given object is an instance of DedicatedCloudService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DedicatedCloudService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DedicatedCloudService.__pulumiType;
    }

    /**
     * Azure region
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * {dedicatedCloudServiceName}
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The properties of Dedicated Node Service
     */
    public readonly properties!: pulumi.Output<outputs.vmwarecloudsimple.v20190401.DedicatedCloudServicePropertiesResponse>;
    /**
     * The list of tags
     */
    public readonly tags!: pulumi.Output<outputs.vmwarecloudsimple.v20190401.TagsResponse | undefined>;
    /**
     * {resourceProviderNamespace}/{resourceType}
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a DedicatedCloudService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DedicatedCloudServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DedicatedCloudServiceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as DedicatedCloudServiceArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DedicatedCloudService.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DedicatedCloudService resource.
 */
export interface DedicatedCloudServiceArgs {
    /**
     * Azure region
     */
    readonly location: pulumi.Input<string>;
    /**
     * dedicated cloud Service name
     */
    readonly name: pulumi.Input<string>;
    /**
     * The properties of Dedicated Node Service
     */
    readonly properties?: pulumi.Input<inputs.vmwarecloudsimple.v20190401.DedicatedCloudServiceProperties>;
    /**
     * The name of the resource group
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The list of tags
     */
    readonly tags?: pulumi.Input<inputs.vmwarecloudsimple.v20190401.Tags>;
}
