// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Subnet in a virtual network resource.
 */
export class Subnet extends pulumi.CustomResource {
    /**
     * Get an existing Subnet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Subnet {
        return new Subnet(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20190701:Subnet';

    /**
     * Returns true if the given object is an instance of Subnet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Subnet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Subnet.__pulumiType;
    }

    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * Properties of the subnet.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.network.v20190701.SubnetPropertiesFormatResponse>;

    /**
     * Create a Subnet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubnetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SubnetArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as SubnetArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.virtualNetworkName === undefined) {
                throw new Error("Missing required property 'virtualNetworkName'");
            }
            inputs["addressPrefix"] = args ? args.addressPrefix : undefined;
            inputs["addressPrefixes"] = args ? args.addressPrefixes : undefined;
            inputs["delegations"] = args ? args.delegations : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["natGateway"] = args ? args.natGateway : undefined;
            inputs["networkSecurityGroup"] = args ? args.networkSecurityGroup : undefined;
            inputs["privateEndpointNetworkPolicies"] = args ? args.privateEndpointNetworkPolicies : undefined;
            inputs["privateLinkServiceNetworkPolicies"] = args ? args.privateLinkServiceNetworkPolicies : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceNavigationLinks"] = args ? args.resourceNavigationLinks : undefined;
            inputs["routeTable"] = args ? args.routeTable : undefined;
            inputs["serviceAssociationLinks"] = args ? args.serviceAssociationLinks : undefined;
            inputs["serviceEndpointPolicies"] = args ? args.serviceEndpointPolicies : undefined;
            inputs["serviceEndpoints"] = args ? args.serviceEndpoints : undefined;
            inputs["virtualNetworkName"] = args ? args.virtualNetworkName : undefined;
            inputs["properties"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Subnet.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Subnet resource.
 */
export interface SubnetArgs {
    /**
     * The address prefix for the subnet.
     */
    readonly addressPrefix?: pulumi.Input<string>;
    /**
     * List of address prefixes for the subnet.
     */
    readonly addressPrefixes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An array of references to the delegations on the subnet.
     */
    readonly delegations?: pulumi.Input<pulumi.Input<inputs.network.v20190701.Delegation>[]>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * The name of the subnet.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Nat gateway associated with this subnet.
     */
    readonly natGateway?: pulumi.Input<inputs.network.v20190701.SubResource>;
    /**
     * The reference of the NetworkSecurityGroup resource.
     */
    readonly networkSecurityGroup?: pulumi.Input<inputs.network.v20190701.NetworkSecurityGroup>;
    /**
     * Enable or Disable apply network policies on private end point in the subnet.
     */
    readonly privateEndpointNetworkPolicies?: pulumi.Input<string>;
    /**
     * Enable or Disable apply network policies on private link service in the subnet.
     */
    readonly privateLinkServiceNetworkPolicies?: pulumi.Input<string>;
    /**
     * The provisioning state of the subnet resource.
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * An array of references to the external resources using subnet.
     */
    readonly resourceNavigationLinks?: pulumi.Input<pulumi.Input<inputs.network.v20190701.ResourceNavigationLink>[]>;
    /**
     * The reference of the RouteTable resource.
     */
    readonly routeTable?: pulumi.Input<inputs.network.v20190701.RouteTable>;
    /**
     * An array of references to services injecting into this subnet.
     */
    readonly serviceAssociationLinks?: pulumi.Input<pulumi.Input<inputs.network.v20190701.ServiceAssociationLink>[]>;
    /**
     * An array of service endpoint policies.
     */
    readonly serviceEndpointPolicies?: pulumi.Input<pulumi.Input<inputs.network.v20190701.ServiceEndpointPolicy>[]>;
    /**
     * An array of service endpoints.
     */
    readonly serviceEndpoints?: pulumi.Input<pulumi.Input<inputs.network.v20190701.ServiceEndpointPropertiesFormat>[]>;
    /**
     * The name of the virtual network.
     */
    readonly virtualNetworkName: pulumi.Input<string>;
}
