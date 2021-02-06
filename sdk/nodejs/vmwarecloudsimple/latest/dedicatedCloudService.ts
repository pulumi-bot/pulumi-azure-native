// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Dedicated cloud service model
 * Latest API Version: 2019-04-01.
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
        return new DedicatedCloudService(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:vmwarecloudsimple/latest:DedicatedCloudService';

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
     * gateway Subnet for the account. It will collect the subnet address and always treat it as /28
     */
    public readonly gatewaySubnet!: pulumi.Output<string>;
    /**
     * indicates whether account onboarded or not in a given region
     */
    public /*out*/ readonly isAccountOnboarded!: pulumi.Output<string>;
    /**
     * Azure region
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * {dedicatedCloudServiceName}
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * total nodes purchased
     */
    public /*out*/ readonly nodes!: pulumi.Output<number>;
    /**
     * link to a service management web portal
     */
    public /*out*/ readonly serviceURL!: pulumi.Output<string>;
    /**
     * The list of tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
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
    constructor(name: string, args: DedicatedCloudServiceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.dedicatedCloudServiceName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'dedicatedCloudServiceName'");
            }
            if ((!args || args.gatewaySubnet === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'gatewaySubnet'");
            }
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["dedicatedCloudServiceName"] = args ? args.dedicatedCloudServiceName : undefined;
            inputs["gatewaySubnet"] = args ? args.gatewaySubnet : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["isAccountOnboarded"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["nodes"] = undefined /*out*/;
            inputs["serviceURL"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["gatewaySubnet"] = undefined /*out*/;
            inputs["isAccountOnboarded"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["nodes"] = undefined /*out*/;
            inputs["serviceURL"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:vmwarecloudsimple/v20190401:DedicatedCloudService" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(DedicatedCloudService.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DedicatedCloudService resource.
 */
export interface DedicatedCloudServiceArgs {
    /**
     * dedicated cloud Service name
     */
    readonly dedicatedCloudServiceName: pulumi.Input<string>;
    /**
     * gateway Subnet for the account. It will collect the subnet address and always treat it as /28
     */
    readonly gatewaySubnet: pulumi.Input<string>;
    /**
     * Azure region
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the resource group
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The list of tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
