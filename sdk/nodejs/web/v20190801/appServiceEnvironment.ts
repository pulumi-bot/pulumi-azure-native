// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * App Service Environment ARM resource.
 */
export class AppServiceEnvironment extends pulumi.CustomResource {
    /**
     * Get an existing AppServiceEnvironment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AppServiceEnvironment {
        return new AppServiceEnvironment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:web/v20190801:AppServiceEnvironment';

    /**
     * Returns true if the given object is an instance of AppServiceEnvironment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppServiceEnvironment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppServiceEnvironment.__pulumiType;
    }

    /**
     * List of comma separated strings describing which VM sizes are allowed for front-ends.
     */
    public /*out*/ readonly allowedMultiSizes!: pulumi.Output<string>;
    /**
     * List of comma separated strings describing which VM sizes are allowed for workers.
     */
    public /*out*/ readonly allowedWorkerSizes!: pulumi.Output<string>;
    /**
     * API Management Account associated with the App Service Environment.
     */
    public readonly apiManagementAccountId!: pulumi.Output<string | undefined>;
    /**
     * Custom settings for changing the behavior of the App Service Environment.
     */
    public readonly clusterSettings!: pulumi.Output<outputs.web.v20190801.NameValuePairResponse[] | undefined>;
    /**
     * Edition of the metadata database for the App Service Environment, e.g. "Standard".
     */
    public /*out*/ readonly databaseEdition!: pulumi.Output<string>;
    /**
     * Service objective of the metadata database for the App Service Environment, e.g. "S0".
     */
    public /*out*/ readonly databaseServiceObjective!: pulumi.Output<string>;
    /**
     * Default Scale Factor for FrontEnds.
     */
    public /*out*/ readonly defaultFrontEndScaleFactor!: pulumi.Output<number>;
    /**
     * DNS suffix of the App Service Environment.
     */
    public readonly dnsSuffix!: pulumi.Output<string | undefined>;
    /**
     * True/false indicating whether the App Service Environment is suspended. The environment can be suspended e.g. when the management endpoint is no longer available
     * (most likely because NSG blocked the incoming traffic).
     */
    public readonly dynamicCacheEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Current total, used, and available worker capacities.
     */
    public /*out*/ readonly environmentCapacities!: pulumi.Output<outputs.web.v20190801.StampCapacityResponse[]>;
    /**
     * True/false indicating whether the App Service Environment is healthy.
     */
    public /*out*/ readonly environmentIsHealthy!: pulumi.Output<boolean>;
    /**
     * Detailed message about with results of the last check of the App Service Environment.
     */
    public /*out*/ readonly environmentStatus!: pulumi.Output<string>;
    /**
     * Scale factor for front-ends.
     */
    public readonly frontEndScaleFactor!: pulumi.Output<number | undefined>;
    /**
     * Flag that displays whether an ASE has linux workers or not
     */
    public readonly hasLinuxWorkers!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment.
     */
    public readonly internalLoadBalancingMode!: pulumi.Output<string | undefined>;
    /**
     * Number of IP SSL addresses reserved for the App Service Environment.
     */
    public readonly ipsslAddressCount!: pulumi.Output<number | undefined>;
    /**
     * Kind of resource.
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * Last deployment action on the App Service Environment.
     */
    public /*out*/ readonly lastAction!: pulumi.Output<string>;
    /**
     * Result of the last deployment action on the App Service Environment.
     */
    public /*out*/ readonly lastActionResult!: pulumi.Output<string>;
    /**
     * Resource Location.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Maximum number of VMs in the App Service Environment.
     */
    public /*out*/ readonly maximumNumberOfMachines!: pulumi.Output<number>;
    /**
     * Number of front-end instances.
     */
    public readonly multiRoleCount!: pulumi.Output<number | undefined>;
    /**
     * Front-end VM size, e.g. "Medium", "Large".
     */
    public readonly multiSize!: pulumi.Output<string | undefined>;
    /**
     * Resource Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Access control list for controlling traffic to the App Service Environment.
     */
    public readonly networkAccessControlList!: pulumi.Output<outputs.web.v20190801.NetworkAccessControlEntryResponse[] | undefined>;
    /**
     * Provisioning state of the App Service Environment.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Resource group of the App Service Environment.
     */
    public /*out*/ readonly resourceGroup!: pulumi.Output<string>;
    /**
     * Key Vault ID for ILB App Service Environment default SSL certificate
     */
    public readonly sslCertKeyVaultId!: pulumi.Output<string | undefined>;
    /**
     * Key Vault Secret Name for ILB App Service Environment default SSL certificate
     */
    public readonly sslCertKeyVaultSecretName!: pulumi.Output<string | undefined>;
    /**
     * Current status of the App Service Environment.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Subscription of the App Service Environment.
     */
    public /*out*/ readonly subscriptionId!: pulumi.Output<string>;
    /**
     * <code>true</code> if the App Service Environment is suspended; otherwise, <code>false</code>. The environment can be suspended, e.g. when the management endpoint is no longer available
     *  (most likely because NSG blocked the incoming traffic).
     */
    public readonly suspended!: pulumi.Output<boolean | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Number of upgrade domains of the App Service Environment.
     */
    public /*out*/ readonly upgradeDomains!: pulumi.Output<number>;
    /**
     * User added ip ranges to whitelist on ASE db
     */
    public readonly userWhitelistedIpRanges!: pulumi.Output<string[] | undefined>;
    /**
     * Description of IP SSL mapping for the App Service Environment.
     */
    public /*out*/ readonly vipMappings!: pulumi.Output<outputs.web.v20190801.VirtualIPMappingResponse[]>;
    /**
     * Description of the Virtual Network.
     */
    public readonly virtualNetwork!: pulumi.Output<outputs.web.v20190801.VirtualNetworkProfileResponse>;
    /**
     * Name of the Virtual Network for the App Service Environment.
     */
    public readonly vnetName!: pulumi.Output<string | undefined>;
    /**
     * Resource group of the Virtual Network.
     */
    public readonly vnetResourceGroupName!: pulumi.Output<string | undefined>;
    /**
     * Subnet of the Virtual Network.
     */
    public readonly vnetSubnetName!: pulumi.Output<string | undefined>;
    /**
     * Description of worker pools with worker size IDs, VM sizes, and number of workers in each pool.
     */
    public readonly workerPools!: pulumi.Output<outputs.web.v20190801.WorkerPoolResponse[]>;

    /**
     * Create a AppServiceEnvironment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppServiceEnvironmentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.name === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.virtualNetwork === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'virtualNetwork'");
            }
            if ((!args || args.workerPools === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'workerPools'");
            }
            inputs["apiManagementAccountId"] = args ? args.apiManagementAccountId : undefined;
            inputs["clusterSettings"] = args ? args.clusterSettings : undefined;
            inputs["dnsSuffix"] = args ? args.dnsSuffix : undefined;
            inputs["dynamicCacheEnabled"] = args ? args.dynamicCacheEnabled : undefined;
            inputs["frontEndScaleFactor"] = args ? args.frontEndScaleFactor : undefined;
            inputs["hasLinuxWorkers"] = args ? args.hasLinuxWorkers : undefined;
            inputs["internalLoadBalancingMode"] = args ? args.internalLoadBalancingMode : undefined;
            inputs["ipsslAddressCount"] = args ? args.ipsslAddressCount : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["multiRoleCount"] = args ? args.multiRoleCount : undefined;
            inputs["multiSize"] = args ? args.multiSize : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkAccessControlList"] = args ? args.networkAccessControlList : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sslCertKeyVaultId"] = args ? args.sslCertKeyVaultId : undefined;
            inputs["sslCertKeyVaultSecretName"] = args ? args.sslCertKeyVaultSecretName : undefined;
            inputs["suspended"] = args ? args.suspended : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["userWhitelistedIpRanges"] = args ? args.userWhitelistedIpRanges : undefined;
            inputs["virtualNetwork"] = args ? args.virtualNetwork : undefined;
            inputs["vnetName"] = args ? args.vnetName : undefined;
            inputs["vnetResourceGroupName"] = args ? args.vnetResourceGroupName : undefined;
            inputs["vnetSubnetName"] = args ? args.vnetSubnetName : undefined;
            inputs["workerPools"] = args ? args.workerPools : undefined;
            inputs["allowedMultiSizes"] = undefined /*out*/;
            inputs["allowedWorkerSizes"] = undefined /*out*/;
            inputs["databaseEdition"] = undefined /*out*/;
            inputs["databaseServiceObjective"] = undefined /*out*/;
            inputs["defaultFrontEndScaleFactor"] = undefined /*out*/;
            inputs["environmentCapacities"] = undefined /*out*/;
            inputs["environmentIsHealthy"] = undefined /*out*/;
            inputs["environmentStatus"] = undefined /*out*/;
            inputs["lastAction"] = undefined /*out*/;
            inputs["lastActionResult"] = undefined /*out*/;
            inputs["maximumNumberOfMachines"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resourceGroup"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["subscriptionId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["upgradeDomains"] = undefined /*out*/;
            inputs["vipMappings"] = undefined /*out*/;
        } else {
            inputs["allowedMultiSizes"] = undefined /*out*/;
            inputs["allowedWorkerSizes"] = undefined /*out*/;
            inputs["apiManagementAccountId"] = undefined /*out*/;
            inputs["clusterSettings"] = undefined /*out*/;
            inputs["databaseEdition"] = undefined /*out*/;
            inputs["databaseServiceObjective"] = undefined /*out*/;
            inputs["defaultFrontEndScaleFactor"] = undefined /*out*/;
            inputs["dnsSuffix"] = undefined /*out*/;
            inputs["dynamicCacheEnabled"] = undefined /*out*/;
            inputs["environmentCapacities"] = undefined /*out*/;
            inputs["environmentIsHealthy"] = undefined /*out*/;
            inputs["environmentStatus"] = undefined /*out*/;
            inputs["frontEndScaleFactor"] = undefined /*out*/;
            inputs["hasLinuxWorkers"] = undefined /*out*/;
            inputs["internalLoadBalancingMode"] = undefined /*out*/;
            inputs["ipsslAddressCount"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["lastAction"] = undefined /*out*/;
            inputs["lastActionResult"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["maximumNumberOfMachines"] = undefined /*out*/;
            inputs["multiRoleCount"] = undefined /*out*/;
            inputs["multiSize"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["networkAccessControlList"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resourceGroup"] = undefined /*out*/;
            inputs["sslCertKeyVaultId"] = undefined /*out*/;
            inputs["sslCertKeyVaultSecretName"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["subscriptionId"] = undefined /*out*/;
            inputs["suspended"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["upgradeDomains"] = undefined /*out*/;
            inputs["userWhitelistedIpRanges"] = undefined /*out*/;
            inputs["vipMappings"] = undefined /*out*/;
            inputs["virtualNetwork"] = undefined /*out*/;
            inputs["vnetName"] = undefined /*out*/;
            inputs["vnetResourceGroupName"] = undefined /*out*/;
            inputs["vnetSubnetName"] = undefined /*out*/;
            inputs["workerPools"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:web/latest:AppServiceEnvironment" }, { type: "azure-nextgen:web/v20150801:AppServiceEnvironment" }, { type: "azure-nextgen:web/v20160901:AppServiceEnvironment" }, { type: "azure-nextgen:web/v20180201:AppServiceEnvironment" }, { type: "azure-nextgen:web/v20200601:AppServiceEnvironment" }, { type: "azure-nextgen:web/v20200901:AppServiceEnvironment" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(AppServiceEnvironment.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AppServiceEnvironment resource.
 */
export interface AppServiceEnvironmentArgs {
    /**
     * API Management Account associated with the App Service Environment.
     */
    readonly apiManagementAccountId?: pulumi.Input<string>;
    /**
     * Custom settings for changing the behavior of the App Service Environment.
     */
    readonly clusterSettings?: pulumi.Input<pulumi.Input<inputs.web.v20190801.NameValuePair>[]>;
    /**
     * DNS suffix of the App Service Environment.
     */
    readonly dnsSuffix?: pulumi.Input<string>;
    /**
     * True/false indicating whether the App Service Environment is suspended. The environment can be suspended e.g. when the management endpoint is no longer available
     * (most likely because NSG blocked the incoming traffic).
     */
    readonly dynamicCacheEnabled?: pulumi.Input<boolean>;
    /**
     * Scale factor for front-ends.
     */
    readonly frontEndScaleFactor?: pulumi.Input<number>;
    /**
     * Flag that displays whether an ASE has linux workers or not
     */
    readonly hasLinuxWorkers?: pulumi.Input<boolean>;
    /**
     * Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment.
     */
    readonly internalLoadBalancingMode?: pulumi.Input<enums.web.v20190801.InternalLoadBalancingMode>;
    /**
     * Number of IP SSL addresses reserved for the App Service Environment.
     */
    readonly ipsslAddressCount?: pulumi.Input<number>;
    /**
     * Kind of resource.
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Resource Location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Number of front-end instances.
     */
    readonly multiRoleCount?: pulumi.Input<number>;
    /**
     * Front-end VM size, e.g. "Medium", "Large".
     */
    readonly multiSize?: pulumi.Input<string>;
    /**
     * Name of the App Service Environment.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Access control list for controlling traffic to the App Service Environment.
     */
    readonly networkAccessControlList?: pulumi.Input<pulumi.Input<inputs.web.v20190801.NetworkAccessControlEntry>[]>;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Key Vault ID for ILB App Service Environment default SSL certificate
     */
    readonly sslCertKeyVaultId?: pulumi.Input<string>;
    /**
     * Key Vault Secret Name for ILB App Service Environment default SSL certificate
     */
    readonly sslCertKeyVaultSecretName?: pulumi.Input<string>;
    /**
     * <code>true</code> if the App Service Environment is suspended; otherwise, <code>false</code>. The environment can be suspended, e.g. when the management endpoint is no longer available
     *  (most likely because NSG blocked the incoming traffic).
     */
    readonly suspended?: pulumi.Input<boolean>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * User added ip ranges to whitelist on ASE db
     */
    readonly userWhitelistedIpRanges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Description of the Virtual Network.
     */
    readonly virtualNetwork: pulumi.Input<inputs.web.v20190801.VirtualNetworkProfile>;
    /**
     * Name of the Virtual Network for the App Service Environment.
     */
    readonly vnetName?: pulumi.Input<string>;
    /**
     * Resource group of the Virtual Network.
     */
    readonly vnetResourceGroupName?: pulumi.Input<string>;
    /**
     * Subnet of the Virtual Network.
     */
    readonly vnetSubnetName?: pulumi.Input<string>;
    /**
     * Description of worker pools with worker size IDs, VM sizes, and number of workers in each pool.
     */
    readonly workerPools: pulumi.Input<pulumi.Input<inputs.web.v20190801.WorkerPool>[]>;
}
