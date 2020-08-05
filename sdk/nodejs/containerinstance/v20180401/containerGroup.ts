// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A container group.
 */
export class ContainerGroup extends pulumi.CustomResource {
    /**
     * Get an existing ContainerGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ContainerGroup {
        return new ContainerGroup(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:containerinstance/v20180401:ContainerGroup';

    /**
     * Returns true if the given object is an instance of ContainerGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ContainerGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ContainerGroup.__pulumiType;
    }

    /**
     * The resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The resource name.
     */
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly properties!: pulumi.Output<outputs.containerinstance.v20180401.ContainerGroupResponseProperties>;
    /**
     * The resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ContainerGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContainerGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ContainerGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ContainerGroupArgs | undefined;
            if (!args || args.containers === undefined) {
                throw new Error("Missing required property 'containers'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.osType === undefined) {
                throw new Error("Missing required property 'osType'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["containers"] = args ? args.containers : undefined;
            inputs["imageRegistryCredentials"] = args ? args.imageRegistryCredentials : undefined;
            inputs["ipAddress"] = args ? args.ipAddress : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["osType"] = args ? args.osType : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["restartPolicy"] = args ? args.restartPolicy : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["volumes"] = args ? args.volumes : undefined;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ContainerGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ContainerGroup resource.
 */
export interface ContainerGroupArgs {
    /**
     * The containers within the container group.
     */
    readonly containers: pulumi.Input<pulumi.Input<inputs.containerinstance.v20180401.Container>[]>;
    /**
     * The image registry credentials by which the container group is created from.
     */
    readonly imageRegistryCredentials?: pulumi.Input<pulumi.Input<inputs.containerinstance.v20180401.ImageRegistryCredential>[]>;
    /**
     * The IP address type of the container group.
     */
    readonly ipAddress?: pulumi.Input<inputs.containerinstance.v20180401.IpAddress>;
    /**
     * The resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the container group.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The operating system type required by the containers in the container group.
     */
    readonly osType: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Restart policy for all containers within the container group. 
     * - `Always` Always restart
     * - `OnFailure` Restart on failure
     * - `Never` Never restart
     */
    readonly restartPolicy?: pulumi.Input<string>;
    /**
     * The resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The list of volumes that can be mounted by containers in this container group.
     */
    readonly volumes?: pulumi.Input<pulumi.Input<inputs.containerinstance.v20180401.Volume>[]>;
}
