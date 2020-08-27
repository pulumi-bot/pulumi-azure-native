// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getHostingEnvironment(args: GetHostingEnvironmentArgs, opts?: pulumi.InvokeOptions): Promise<GetHostingEnvironmentResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:web/v20150801:getHostingEnvironment", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetHostingEnvironmentArgs {
    /**
     * Name of hostingEnvironment (App Service Environment)
     */
    readonly name: string;
    /**
     * Name of resource group
     */
    readonly resourceGroupName: string;
}

/**
 * Description of an hostingEnvironment (App Service Environment)
 */
export interface GetHostingEnvironmentResult {
    /**
     * List of comma separated strings describing which VM sizes are allowed for front-ends
     */
    readonly allowedMultiSizes?: string;
    /**
     * List of comma separated strings describing which VM sizes are allowed for workers
     */
    readonly allowedWorkerSizes?: string;
    /**
     * Api Management Account associated with this Hosting Environment
     */
    readonly apiManagementAccountId?: string;
    /**
     * Custom settings for changing the behavior of the hosting environment
     */
    readonly clusterSettings?: outputs.web.v20150801.NameValuePairResponse[];
    /**
     * Edition of the metadata database for the hostingEnvironment (App Service Environment) e.g. "Standard"
     */
    readonly databaseEdition?: string;
    /**
     * Service objective of the metadata database for the hostingEnvironment (App Service Environment) e.g. "S0"
     */
    readonly databaseServiceObjective?: string;
    /**
     * DNS suffix of the hostingEnvironment (App Service Environment)
     */
    readonly dnsSuffix?: string;
    /**
     * Current total, used, and available worker capacities
     */
    readonly environmentCapacities?: outputs.web.v20150801.StampCapacityResponse[];
    /**
     * True/false indicating whether the hostingEnvironment (App Service Environment) is healthy
     */
    readonly environmentIsHealthy?: boolean;
    /**
     * Detailed message about with results of the last check of the hostingEnvironment (App Service Environment)
     */
    readonly environmentStatus?: string;
    /**
     * Specifies which endpoints to serve internally in the hostingEnvironment's (App Service Environment) VNET
     */
    readonly internalLoadBalancingMode?: InternalLoadBalancingMode;
    /**
     * Number of IP SSL addresses reserved for this hostingEnvironment (App Service Environment)
     */
    readonly ipsslAddressCount?: number;
    /**
     * Kind of resource
     */
    readonly kind?: string;
    /**
     * Last deployment action on this hostingEnvironment (App Service Environment)
     */
    readonly lastAction?: string;
    /**
     * Result of the last deployment action on this hostingEnvironment (App Service Environment)
     */
    readonly lastActionResult?: string;
    /**
     * Resource Location
     */
    readonly location: string;
    /**
     * Maximum number of VMs in this hostingEnvironment (App Service Environment)
     */
    readonly maximumNumberOfMachines?: number;
    /**
     * Number of front-end instances
     */
    readonly multiRoleCount?: number;
    /**
     * Front-end VM size, e.g. "Medium", "Large"
     */
    readonly multiSize?: string;
    /**
     * Resource Name
     */
    readonly name?: string;
    /**
     * Access control list for controlling traffic to the hostingEnvironment (App Service Environment)
     */
    readonly networkAccessControlList?: outputs.web.v20150801.NetworkAccessControlEntryResponse[];
    /**
     * Provisioning state of the hostingEnvironment (App Service Environment)
     */
    readonly provisioningState?: ProvisioningState;
    /**
     * Resource group of the hostingEnvironment (App Service Environment)
     */
    readonly resourceGroup?: string;
    /**
     * Current status of the hostingEnvironment (App Service Environment)
     */
    readonly status: HostingEnvironmentStatus;
    /**
     * Subscription of the hostingEnvironment (App Service Environment)
     */
    readonly subscriptionId?: string;
    /**
     * True/false indicating whether the hostingEnvironment is suspended. The environment can be suspended e.g. when the management endpoint is no longer available
     *             (most likely because NSG blocked the incoming traffic)
     */
    readonly suspended?: boolean;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type?: string;
    /**
     * Number of upgrade domains of this hostingEnvironment (App Service Environment)
     */
    readonly upgradeDomains?: number;
    /**
     * Description of IP SSL mapping for this hostingEnvironment (App Service Environment)
     */
    readonly vipMappings?: outputs.web.v20150801.VirtualIPMappingResponse[];
    /**
     * Description of the hostingEnvironment's (App Service Environment) virtual network
     */
    readonly virtualNetwork?: outputs.web.v20150801.VirtualNetworkProfileResponse;
    /**
     * Name of the hostingEnvironment's (App Service Environment) virtual network
     */
    readonly vnetName?: string;
    /**
     * Resource group of the hostingEnvironment's (App Service Environment) virtual network
     */
    readonly vnetResourceGroupName?: string;
    /**
     * Subnet of the hostingEnvironment's (App Service Environment) virtual network
     */
    readonly vnetSubnetName?: string;
    /**
     * Description of worker pools with worker size ids, VM sizes, and number of workers in each pool
     */
    readonly workerPools?: outputs.web.v20150801.WorkerPoolResponse[];
}
