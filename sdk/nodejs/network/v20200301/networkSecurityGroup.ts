// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * NetworkSecurityGroup resource.
 */
export class NetworkSecurityGroup extends pulumi.CustomResource {
    /**
     * Get an existing NetworkSecurityGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): NetworkSecurityGroup {
        return new NetworkSecurityGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:network/v20200301:NetworkSecurityGroup';

    /**
     * Returns true if the given object is an instance of NetworkSecurityGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkSecurityGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkSecurityGroup.__pulumiType;
    }

    /**
     * The default security rules of network security group.
     */
    public /*out*/ readonly defaultSecurityRules!: pulumi.Output<outputs.network.v20200301.SecurityRuleResponse[]>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * A collection of references to flow log resources.
     */
    public /*out*/ readonly flowLogs!: pulumi.Output<outputs.network.v20200301.FlowLogResponse[]>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * A collection of references to network interfaces.
     */
    public /*out*/ readonly networkInterfaces!: pulumi.Output<outputs.network.v20200301.NetworkInterfaceResponse[]>;
    /**
     * The provisioning state of the network security group resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The resource GUID property of the network security group resource.
     */
    public /*out*/ readonly resourceGuid!: pulumi.Output<string>;
    /**
     * A collection of security rules of the network security group.
     */
    public readonly securityRules!: pulumi.Output<outputs.network.v20200301.SecurityRuleResponse[] | undefined>;
    /**
     * A collection of references to subnets.
     */
    public /*out*/ readonly subnets!: pulumi.Output<outputs.network.v20200301.SubnetResponse[]>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a NetworkSecurityGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkSecurityGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.networkSecurityGroupName === undefined) {
                throw new Error("Missing required property 'networkSecurityGroupName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["id"] = args ? args.id : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["networkSecurityGroupName"] = args ? args.networkSecurityGroupName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["securityRules"] = args ? args.securityRules : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["defaultSecurityRules"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["flowLogs"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["networkInterfaces"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resourceGuid"] = undefined /*out*/;
            inputs["subnets"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["defaultSecurityRules"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["flowLogs"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["networkInterfaces"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resourceGuid"] = undefined /*out*/;
            inputs["securityRules"] = undefined /*out*/;
            inputs["subnets"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:network/latest:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20150501preview:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20150615:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20160330:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20160601:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20160901:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20161201:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20170301:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20170601:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20170801:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20170901:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20171001:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20171101:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20180101:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20180201:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20180401:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20180601:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20180701:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20180801:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20181001:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20181101:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20181201:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20190201:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20190401:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20190601:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20190701:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20190801:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20190901:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20191101:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20191201:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20200401:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20200501:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20200601:NetworkSecurityGroup" }, { type: "azure-nextgen:network/v20200701:NetworkSecurityGroup" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(NetworkSecurityGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a NetworkSecurityGroup resource.
 */
export interface NetworkSecurityGroupArgs {
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the network security group.
     */
    readonly networkSecurityGroupName: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A collection of security rules of the network security group.
     */
    readonly securityRules?: pulumi.Input<pulumi.Input<inputs.network.v20200301.SecurityRule>[]>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
