// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getService(args: GetServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:search/v20200313:getService", {
        "resourceGroupName": args.resourceGroupName,
        "searchServiceName": args.searchServiceName,
    }, opts);
}

export interface GetServiceArgs {
    /**
     * The name of the resource group within the current subscription. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the Azure Cognitive Search service associated with the specified resource group.
     */
    readonly searchServiceName: string;
}

/**
 * Describes an Azure Cognitive Search service and its current state.
 */
export interface GetServiceResult {
    /**
     * Applicable only for the standard3 SKU. You can set this property to enable up to 3 high density partitions that allow up to 1000 indexes, which is much higher than the maximum indexes allowed for any other SKU. For the standard3 SKU, the value is either 'default' or 'highDensity'. For all other SKUs, this value must be 'default'.
     */
    readonly hostingMode?: HostingMode;
    /**
     * The identity of the resource.
     */
    readonly identity?: outputs.search.v20200313.IdentityResponse;
    /**
     * The geographic location of the resource. This must be one of the supported and registered Azure Geo Regions (for example, West US, East US, Southeast Asia, and so forth). This property is required when creating a new resource.
     */
    readonly location?: string;
    /**
     * The name of the resource.
     */
    readonly name: string;
    /**
     * Network specific rules that determine how the Azure Cognitive Search service may be reached.
     */
    readonly networkRuleSet?: outputs.search.v20200313.NetworkRuleSetResponse;
    /**
     * The number of partitions in the Search service; if specified, it can be 1, 2, 3, 4, 6, or 12. Values greater than 1 are only valid for standard SKUs. For 'standard3' services with hostingMode set to 'highDensity', the allowed values are between 1 and 3.
     */
    readonly partitionCount?: number;
    /**
     * The list of private endpoint connections to the Azure Cognitive Search service.
     */
    readonly privateEndpointConnections: outputs.search.v20200313.PrivateEndpointConnectionResponse[];
    /**
     * The state of the last provisioning operation performed on the Search service. Provisioning is an intermediate state that occurs while service capacity is being established. After capacity is set up, provisioningState changes to either 'succeeded' or 'failed'. Client applications can poll provisioning status (the recommended polling interval is from 30 seconds to one minute) by using the Get Search Service operation to see when an operation is completed. If you are using the free service, this value tends to come back as 'succeeded' directly in the call to Create Search service. This is because the free service uses capacity that is already set up.
     */
    readonly provisioningState: ProvisioningState;
    /**
     * This value can be set to 'enabled' to avoid breaking changes on existing customer resources and templates. If set to 'disabled', traffic over public interface is not allowed, and private endpoint connections would be the exclusive access method.
     */
    readonly publicNetworkAccess?: PublicNetworkAccess;
    /**
     * The number of replicas in the Search service. If specified, it must be a value between 1 and 12 inclusive for standard SKUs or between 1 and 3 inclusive for basic SKU.
     */
    readonly replicaCount?: number;
    /**
     * The list of shared private link resources managed by the Azure Cognitive Search service.
     */
    readonly sharedPrivateLinkResources: outputs.search.v20200313.SharedPrivateLinkResourceResponse[];
    /**
     * The SKU of the Search Service, which determines price tier and capacity limits. This property is required when creating a new Search Service.
     */
    readonly sku?: outputs.search.v20200313.SkuResponse;
    /**
     * The status of the Search service. Possible values include: 'running': The Search service is running and no provisioning operations are underway. 'provisioning': The Search service is being provisioned or scaled up or down. 'deleting': The Search service is being deleted. 'degraded': The Search service is degraded. This can occur when the underlying search units are not healthy. The Search service is most likely operational, but performance might be slow and some requests might be dropped. 'disabled': The Search service is disabled. In this state, the service will reject all API requests. 'error': The Search service is in an error state. If your service is in the degraded, disabled, or error states, it means the Azure Cognitive Search team is actively investigating the underlying issue. Dedicated services in these states are still chargeable based on the number of search units provisioned.
     */
    readonly status: SearchServiceStatus;
    /**
     * The details of the Search service status.
     */
    readonly statusDetails: string;
    /**
     * Tags to help categorize the resource in the Azure portal.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The resource type.
     */
    readonly type: string;
}
