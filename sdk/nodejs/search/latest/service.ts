// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Describes an Azure Cognitive Search service and its current state.
 */
export class Service extends pulumi.CustomResource {
    /**
     * Get an existing Service resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Service {
        return new Service(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:search/latest:Service';

    /**
     * Returns true if the given object is an instance of Service.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Service {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Service.__pulumiType;
    }

    /**
     * Applicable only for the standard3 SKU. You can set this property to enable up to 3 high density partitions that allow up to 1000 indexes, which is much higher than the maximum indexes allowed for any other SKU. For the standard3 SKU, the value is either 'default' or 'highDensity'. For all other SKUs, this value must be 'default'.
     */
    public readonly hostingMode!: pulumi.Output<string | undefined>;
    /**
     * The identity of the resource.
     */
    public readonly identity!: pulumi.Output<outputs.search.latest.IdentityResponse | undefined>;
    /**
     * The geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Network specific rules that determine how the Azure Cognitive Search service may be reached.
     */
    public readonly networkRuleSet!: pulumi.Output<outputs.search.latest.NetworkRuleSetResponse | undefined>;
    /**
     * The number of partitions in the search service; if specified, it can be 1, 2, 3, 4, 6, or 12. Values greater than 1 are only valid for standard SKUs. For 'standard3' services with hostingMode set to 'highDensity', the allowed values are between 1 and 3.
     */
    public readonly partitionCount!: pulumi.Output<number | undefined>;
    /**
     * The list of private endpoint connections to the Azure Cognitive Search service.
     */
    public /*out*/ readonly privateEndpointConnections!: pulumi.Output<outputs.search.latest.PrivateEndpointConnectionResponse[]>;
    /**
     * The state of the last provisioning operation performed on the search service. Provisioning is an intermediate state that occurs while service capacity is being established. After capacity is set up, provisioningState changes to either 'succeeded' or 'failed'. Client applications can poll provisioning status (the recommended polling interval is from 30 seconds to one minute) by using the Get Search Service operation to see when an operation is completed. If you are using the free service, this value tends to come back as 'succeeded' directly in the call to Create search service. This is because the free service uses capacity that is already set up.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * This value can be set to 'enabled' to avoid breaking changes on existing customer resources and templates. If set to 'disabled', traffic over public interface is not allowed, and private endpoint connections would be the exclusive access method.
     */
    public readonly publicNetworkAccess!: pulumi.Output<string | undefined>;
    /**
     * The number of replicas in the search service. If specified, it must be a value between 1 and 12 inclusive for standard SKUs or between 1 and 3 inclusive for basic SKU.
     */
    public readonly replicaCount!: pulumi.Output<number | undefined>;
    /**
     * The list of shared private link resources managed by the Azure Cognitive Search service.
     */
    public /*out*/ readonly sharedPrivateLinkResources!: pulumi.Output<outputs.search.latest.SharedPrivateLinkResourceResponse[]>;
    /**
     * The SKU of the Search Service, which determines price tier and capacity limits. This property is required when creating a new Search Service.
     */
    public readonly sku!: pulumi.Output<outputs.search.latest.SkuResponse | undefined>;
    /**
     * The status of the search service. Possible values include: 'running': The search service is running and no provisioning operations are underway. 'provisioning': The search service is being provisioned or scaled up or down. 'deleting': The search service is being deleted. 'degraded': The search service is degraded. This can occur when the underlying search units are not healthy. The search service is most likely operational, but performance might be slow and some requests might be dropped. 'disabled': The search service is disabled. In this state, the service will reject all API requests. 'error': The search service is in an error state. If your service is in the degraded, disabled, or error states, it means the Azure Cognitive Search team is actively investigating the underlying issue. Dedicated services in these states are still chargeable based on the number of search units provisioned.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The details of the search service status.
     */
    public /*out*/ readonly statusDetails!: pulumi.Output<string>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Service resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.searchServiceName === undefined) {
                throw new Error("Missing required property 'searchServiceName'");
            }
            inputs["hostingMode"] = args ? args.hostingMode : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["networkRuleSet"] = args ? args.networkRuleSet : undefined;
            inputs["partitionCount"] = args ? args.partitionCount : undefined;
            inputs["publicNetworkAccess"] = args ? args.publicNetworkAccess : undefined;
            inputs["replicaCount"] = args ? args.replicaCount : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["searchServiceName"] = args ? args.searchServiceName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["privateEndpointConnections"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["sharedPrivateLinkResources"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["statusDetails"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["hostingMode"] = undefined /*out*/;
            inputs["identity"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["networkRuleSet"] = undefined /*out*/;
            inputs["partitionCount"] = undefined /*out*/;
            inputs["privateEndpointConnections"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["publicNetworkAccess"] = undefined /*out*/;
            inputs["replicaCount"] = undefined /*out*/;
            inputs["sharedPrivateLinkResources"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["statusDetails"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:search/v20150819:Service" }, { type: "azure-nextgen:search/v20191001preview:Service" }, { type: "azure-nextgen:search/v20200313:Service" }, { type: "azure-nextgen:search/v20200801:Service" }, { type: "azure-nextgen:search/v20200801preview:Service" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Service.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Service resource.
 */
export interface ServiceArgs {
    /**
     * Applicable only for the standard3 SKU. You can set this property to enable up to 3 high density partitions that allow up to 1000 indexes, which is much higher than the maximum indexes allowed for any other SKU. For the standard3 SKU, the value is either 'default' or 'highDensity'. For all other SKUs, this value must be 'default'.
     */
    readonly hostingMode?: pulumi.Input<string>;
    /**
     * The identity of the resource.
     */
    readonly identity?: pulumi.Input<inputs.search.latest.Identity>;
    /**
     * The geo-location where the resource lives
     */
    readonly location: pulumi.Input<string>;
    /**
     * Network specific rules that determine how the Azure Cognitive Search service may be reached.
     */
    readonly networkRuleSet?: pulumi.Input<inputs.search.latest.NetworkRuleSet>;
    /**
     * The number of partitions in the search service; if specified, it can be 1, 2, 3, 4, 6, or 12. Values greater than 1 are only valid for standard SKUs. For 'standard3' services with hostingMode set to 'highDensity', the allowed values are between 1 and 3.
     */
    readonly partitionCount?: pulumi.Input<number>;
    /**
     * This value can be set to 'enabled' to avoid breaking changes on existing customer resources and templates. If set to 'disabled', traffic over public interface is not allowed, and private endpoint connections would be the exclusive access method.
     */
    readonly publicNetworkAccess?: pulumi.Input<string>;
    /**
     * The number of replicas in the search service. If specified, it must be a value between 1 and 12 inclusive for standard SKUs or between 1 and 3 inclusive for basic SKU.
     */
    readonly replicaCount?: pulumi.Input<number>;
    /**
     * The name of the resource group within the current subscription. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the Azure Cognitive Search service to create or update. Search service names must only contain lowercase letters, digits or dashes, cannot use dash as the first two or last one characters, cannot contain consecutive dashes, and must be between 2 and 60 characters in length. Search service names must be globally unique since they are part of the service URI (https://<name>.search.windows.net). You cannot change the service name after the service is created.
     */
    readonly searchServiceName: pulumi.Input<string>;
    /**
     * The SKU of the Search Service, which determines price tier and capacity limits. This property is required when creating a new Search Service.
     */
    readonly sku?: pulumi.Input<inputs.search.latest.Sku>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
