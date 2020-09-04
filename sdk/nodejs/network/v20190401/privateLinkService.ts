// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Private link service resource.
 */
export class PrivateLinkService extends pulumi.CustomResource {
    /**
     * Get an existing PrivateLinkService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PrivateLinkService {
        return new PrivateLinkService(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20190401:PrivateLinkService';

    /**
     * Returns true if the given object is an instance of PrivateLinkService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrivateLinkService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrivateLinkService.__pulumiType;
    }

    /**
     * The alias of the private link service.
     */
    public /*out*/ readonly alias!: pulumi.Output<string>;
    /**
     * The auto-approval list of the private link service.
     */
    public readonly autoApproval!: pulumi.Output<outputs.network.v20190401.PrivateLinkServicePropertiesResponseAutoApproval | undefined>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * The list of Fqdn.
     */
    public readonly fqdns!: pulumi.Output<string[] | undefined>;
    /**
     * An array of references to the private link service IP configuration.
     */
    public readonly ipConfigurations!: pulumi.Output<outputs.network.v20190401.PrivateLinkServiceIpConfigurationResponse[] | undefined>;
    /**
     * An array of references to the load balancer IP configurations.
     */
    public readonly loadBalancerFrontendIpConfigurations!: pulumi.Output<outputs.network.v20190401.FrontendIPConfigurationResponse[] | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Gets an array of references to the network interfaces created for this private link service.
     */
    public /*out*/ readonly networkInterfaces!: pulumi.Output<outputs.network.v20190401.NetworkInterfaceResponse[]>;
    /**
     * An array of list about connections to the private endpoint.
     */
    public readonly privateEndpointConnections!: pulumi.Output<outputs.network.v20190401.PrivateEndpointConnectionResponse[] | undefined>;
    /**
     * The provisioning state of the private link service.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The visibility list of the private link service.
     */
    public readonly visibility!: pulumi.Output<outputs.network.v20190401.PrivateLinkServicePropertiesResponseVisibility | undefined>;

    /**
     * Create a PrivateLinkService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrivateLinkServiceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serviceName === undefined) {
                throw new Error("Missing required property 'serviceName'");
            }
            inputs["autoApproval"] = args ? args.autoApproval : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["fqdns"] = args ? args.fqdns : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["ipConfigurations"] = args ? args.ipConfigurations : undefined;
            inputs["loadBalancerFrontendIpConfigurations"] = args ? args.loadBalancerFrontendIpConfigurations : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["privateEndpointConnections"] = args ? args.privateEndpointConnections : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["visibility"] = args ? args.visibility : undefined;
            inputs["alias"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["networkInterfaces"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["alias"] = undefined /*out*/;
            inputs["autoApproval"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["fqdns"] = undefined /*out*/;
            inputs["ipConfigurations"] = undefined /*out*/;
            inputs["loadBalancerFrontendIpConfigurations"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["networkInterfaces"] = undefined /*out*/;
            inputs["privateEndpointConnections"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["visibility"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:PrivateLinkService" }, { type: "azurerm:network/v20190601:PrivateLinkService" }, { type: "azurerm:network/v20190701:PrivateLinkService" }, { type: "azurerm:network/v20190801:PrivateLinkService" }, { type: "azurerm:network/v20190901:PrivateLinkService" }, { type: "azurerm:network/v20191101:PrivateLinkService" }, { type: "azurerm:network/v20191201:PrivateLinkService" }, { type: "azurerm:network/v20200301:PrivateLinkService" }, { type: "azurerm:network/v20200401:PrivateLinkService" }, { type: "azurerm:network/v20200501:PrivateLinkService" }, { type: "azurerm:network/v20200601:PrivateLinkService" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(PrivateLinkService.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a PrivateLinkService resource.
 */
export interface PrivateLinkServiceArgs {
    /**
     * The auto-approval list of the private link service.
     */
    readonly autoApproval?: pulumi.Input<inputs.network.v20190401.PrivateLinkServicePropertiesAutoApproval>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The list of Fqdn.
     */
    readonly fqdns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * An array of references to the private link service IP configuration.
     */
    readonly ipConfigurations?: pulumi.Input<pulumi.Input<inputs.network.v20190401.PrivateLinkServiceIpConfiguration>[]>;
    /**
     * An array of references to the load balancer IP configurations.
     */
    readonly loadBalancerFrontendIpConfigurations?: pulumi.Input<pulumi.Input<inputs.network.v20190401.FrontendIPConfiguration>[]>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * An array of list about connections to the private endpoint.
     */
    readonly privateEndpointConnections?: pulumi.Input<pulumi.Input<inputs.network.v20190401.PrivateEndpointConnection>[]>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the private link service.
     */
    readonly serviceName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The visibility list of the private link service.
     */
    readonly visibility?: pulumi.Input<inputs.network.v20190401.PrivateLinkServicePropertiesVisibility>;
}
