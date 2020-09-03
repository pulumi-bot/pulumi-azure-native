// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Network security rule.
 *
 * ## Example Usage
 * ### Create security rule
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const securityRule = new azurerm.network.v20200401.SecurityRule("securityRule", {
 *     access: "Deny",
 *     destinationAddressPrefix: "11.0.0.0/8",
 *     destinationPortRange: "8080",
 *     direction: "Outbound",
 *     networkSecurityGroupName: "testnsg",
 *     priority: 100,
 *     protocol: "*",
 *     resourceGroupName: "rg1",
 *     securityRuleName: "rule1",
 *     sourceAddressPrefix: "10.0.0.0/8",
 *     sourcePortRange: "*",
 * });
 *
 * ```
 */
export class SecurityRule extends pulumi.CustomResource {
    /**
     * Get an existing SecurityRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SecurityRule {
        return new SecurityRule(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20200401:SecurityRule';

    /**
     * Returns true if the given object is an instance of SecurityRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityRule.__pulumiType;
    }

    /**
     * The network traffic is allowed or denied.
     */
    public readonly access!: pulumi.Output<string>;
    /**
     * A description for this rule. Restricted to 140 chars.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The destination address prefix. CIDR or destination IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
     */
    public readonly destinationAddressPrefix!: pulumi.Output<string | undefined>;
    /**
     * The destination address prefixes. CIDR or destination IP ranges.
     */
    public readonly destinationAddressPrefixes!: pulumi.Output<string[] | undefined>;
    /**
     * The application security group specified as destination.
     */
    public readonly destinationApplicationSecurityGroups!: pulumi.Output<outputs.network.v20200401.ApplicationSecurityGroupResponse[] | undefined>;
    /**
     * The destination port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
     */
    public readonly destinationPortRange!: pulumi.Output<string | undefined>;
    /**
     * The destination port ranges.
     */
    public readonly destinationPortRanges!: pulumi.Output<string[] | undefined>;
    /**
     * The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic.
     */
    public readonly direction!: pulumi.Output<string>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
     */
    public readonly priority!: pulumi.Output<number | undefined>;
    /**
     * Network protocol this rule applies to.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * The provisioning state of the security rule resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from.
     */
    public readonly sourceAddressPrefix!: pulumi.Output<string | undefined>;
    /**
     * The CIDR or source IP ranges.
     */
    public readonly sourceAddressPrefixes!: pulumi.Output<string[] | undefined>;
    /**
     * The application security group specified as source.
     */
    public readonly sourceApplicationSecurityGroups!: pulumi.Output<outputs.network.v20200401.ApplicationSecurityGroupResponse[] | undefined>;
    /**
     * The source port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
     */
    public readonly sourcePortRange!: pulumi.Output<string | undefined>;
    /**
     * The source port ranges.
     */
    public readonly sourcePortRanges!: pulumi.Output<string[] | undefined>;

    /**
     * Create a SecurityRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecurityRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityRuleArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as SecurityRuleArgs | undefined;
            if (!args || args.access === undefined) {
                throw new Error("Missing required property 'access'");
            }
            if (!args || args.direction === undefined) {
                throw new Error("Missing required property 'direction'");
            }
            if (!args || args.networkSecurityGroupName === undefined) {
                throw new Error("Missing required property 'networkSecurityGroupName'");
            }
            if (!args || args.protocol === undefined) {
                throw new Error("Missing required property 'protocol'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.securityRuleName === undefined) {
                throw new Error("Missing required property 'securityRuleName'");
            }
            inputs["access"] = args ? args.access : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["destinationAddressPrefix"] = args ? args.destinationAddressPrefix : undefined;
            inputs["destinationAddressPrefixes"] = args ? args.destinationAddressPrefixes : undefined;
            inputs["destinationApplicationSecurityGroups"] = args ? args.destinationApplicationSecurityGroups : undefined;
            inputs["destinationPortRange"] = args ? args.destinationPortRange : undefined;
            inputs["destinationPortRanges"] = args ? args.destinationPortRanges : undefined;
            inputs["direction"] = args ? args.direction : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkSecurityGroupName"] = args ? args.networkSecurityGroupName : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["securityRuleName"] = args ? args.securityRuleName : undefined;
            inputs["sourceAddressPrefix"] = args ? args.sourceAddressPrefix : undefined;
            inputs["sourceAddressPrefixes"] = args ? args.sourceAddressPrefixes : undefined;
            inputs["sourceApplicationSecurityGroups"] = args ? args.sourceApplicationSecurityGroups : undefined;
            inputs["sourcePortRange"] = args ? args.sourcePortRange : undefined;
            inputs["sourcePortRanges"] = args ? args.sourcePortRanges : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:SecurityRule" }, { type: "azurerm:network/v20150501preview:SecurityRule" }, { type: "azurerm:network/v20150615:SecurityRule" }, { type: "azurerm:network/v20160330:SecurityRule" }, { type: "azurerm:network/v20160601:SecurityRule" }, { type: "azurerm:network/v20160901:SecurityRule" }, { type: "azurerm:network/v20161201:SecurityRule" }, { type: "azurerm:network/v20170301:SecurityRule" }, { type: "azurerm:network/v20170601:SecurityRule" }, { type: "azurerm:network/v20170801:SecurityRule" }, { type: "azurerm:network/v20170901:SecurityRule" }, { type: "azurerm:network/v20171001:SecurityRule" }, { type: "azurerm:network/v20171101:SecurityRule" }, { type: "azurerm:network/v20180101:SecurityRule" }, { type: "azurerm:network/v20180201:SecurityRule" }, { type: "azurerm:network/v20180401:SecurityRule" }, { type: "azurerm:network/v20180601:SecurityRule" }, { type: "azurerm:network/v20180701:SecurityRule" }, { type: "azurerm:network/v20180801:SecurityRule" }, { type: "azurerm:network/v20181001:SecurityRule" }, { type: "azurerm:network/v20181101:SecurityRule" }, { type: "azurerm:network/v20181201:SecurityRule" }, { type: "azurerm:network/v20190201:SecurityRule" }, { type: "azurerm:network/v20190401:SecurityRule" }, { type: "azurerm:network/v20190601:SecurityRule" }, { type: "azurerm:network/v20190701:SecurityRule" }, { type: "azurerm:network/v20190801:SecurityRule" }, { type: "azurerm:network/v20190901:SecurityRule" }, { type: "azurerm:network/v20191101:SecurityRule" }, { type: "azurerm:network/v20191201:SecurityRule" }, { type: "azurerm:network/v20200301:SecurityRule" }, { type: "azurerm:network/v20200501:SecurityRule" }, { type: "azurerm:network/v20200601:SecurityRule" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(SecurityRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a SecurityRule resource.
 */
export interface SecurityRuleArgs {
    /**
     * The network traffic is allowed or denied.
     */
    readonly access: pulumi.Input<string>;
    /**
     * A description for this rule. Restricted to 140 chars.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The destination address prefix. CIDR or destination IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
     */
    readonly destinationAddressPrefix?: pulumi.Input<string>;
    /**
     * The destination address prefixes. CIDR or destination IP ranges.
     */
    readonly destinationAddressPrefixes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The application security group specified as destination.
     */
    readonly destinationApplicationSecurityGroups?: pulumi.Input<pulumi.Input<inputs.network.v20200401.ApplicationSecurityGroup>[]>;
    /**
     * The destination port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
     */
    readonly destinationPortRange?: pulumi.Input<string>;
    /**
     * The destination port ranges.
     */
    readonly destinationPortRanges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic.
     */
    readonly direction: pulumi.Input<string>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * The name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the network security group.
     */
    readonly networkSecurityGroupName: pulumi.Input<string>;
    /**
     * The priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * Network protocol this rule applies to.
     */
    readonly protocol: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the security rule.
     */
    readonly securityRuleName: pulumi.Input<string>;
    /**
     * The CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from.
     */
    readonly sourceAddressPrefix?: pulumi.Input<string>;
    /**
     * The CIDR or source IP ranges.
     */
    readonly sourceAddressPrefixes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The application security group specified as source.
     */
    readonly sourceApplicationSecurityGroups?: pulumi.Input<pulumi.Input<inputs.network.v20200401.ApplicationSecurityGroup>[]>;
    /**
     * The source port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
     */
    readonly sourcePortRange?: pulumi.Input<string>;
    /**
     * The source port ranges.
     */
    readonly sourcePortRanges?: pulumi.Input<pulumi.Input<string>[]>;
}
