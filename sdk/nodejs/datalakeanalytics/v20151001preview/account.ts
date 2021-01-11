// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * A Data Lake Analytics account object, containing all information associated with the named Data Lake Analytics account.
 */
export class Account extends pulumi.CustomResource {
    /**
     * Get an existing Account resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Account {
        return new Account(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:datalakeanalytics/v20151001preview:Account';

    /**
     * Returns true if the given object is an instance of Account.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Account {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Account.__pulumiType;
    }

    /**
     * The unique identifier associated with this Data Lake Analytics account.
     */
    public /*out*/ readonly accountId!: pulumi.Output<string>;
    /**
     * The list of compute policies associated with this account.
     */
    public readonly computePolicies!: pulumi.Output<outputs.datalakeanalytics.v20151001preview.ComputePolicyResponse[]>;
    /**
     * The account creation time.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * The commitment tier in use for the current month.
     */
    public /*out*/ readonly currentTier!: pulumi.Output<string>;
    /**
     * The list of Data Lake Store accounts associated with this account.
     */
    public readonly dataLakeStoreAccounts!: pulumi.Output<outputs.datalakeanalytics.v20151001preview.DataLakeStoreAccountInformationResponse[] | undefined>;
    /**
     * The current state of the DebugDataAccessLevel for this account.
     */
    public /*out*/ readonly debugDataAccessLevel!: pulumi.Output<string>;
    /**
     * The default Data Lake Store account associated with this account.
     */
    public readonly defaultDataLakeStoreAccount!: pulumi.Output<string>;
    /**
     * The full CName endpoint for this account.
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * The current state of allowing or disallowing IPs originating within Azure through the firewall. If the firewall is disabled, this is not enforced.
     */
    public readonly firewallAllowAzureIps!: pulumi.Output<string>;
    /**
     * The list of firewall rules associated with this account.
     */
    public readonly firewallRules!: pulumi.Output<outputs.datalakeanalytics.v20151001preview.FirewallRuleResponse[]>;
    /**
     * The current state of the IP address firewall for this account.
     */
    public readonly firewallState!: pulumi.Output<string>;
    /**
     * The list of hiveMetastores associated with this account.
     */
    public /*out*/ readonly hiveMetastores!: pulumi.Output<outputs.datalakeanalytics.v20151001preview.HiveMetastoreResponse[]>;
    /**
     * The account last modified time.
     */
    public /*out*/ readonly lastModifiedTime!: pulumi.Output<string>;
    /**
     * The resource location.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The maximum supported degree of parallelism for this account.
     */
    public readonly maxDegreeOfParallelism!: pulumi.Output<number | undefined>;
    /**
     * The maximum supported degree of parallelism per job for this account.
     */
    public readonly maxDegreeOfParallelismPerJob!: pulumi.Output<number>;
    /**
     * The maximum supported jobs running under the account at the same time.
     */
    public readonly maxJobCount!: pulumi.Output<number | undefined>;
    /**
     * The minimum supported priority per job for this account.
     */
    public readonly minPriorityPerJob!: pulumi.Output<number>;
    /**
     * The resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The commitment tier for the next month.
     */
    public readonly newTier!: pulumi.Output<string>;
    /**
     * The provisioning status of the Data Lake Analytics account.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The number of days that job metadata is retained.
     */
    public readonly queryStoreRetention!: pulumi.Output<number | undefined>;
    /**
     * The state of the Data Lake Analytics account.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The list of Azure Blob Storage accounts associated with this account.
     */
    public readonly storageAccounts!: pulumi.Output<outputs.datalakeanalytics.v20151001preview.StorageAccountInformationResponse[]>;
    /**
     * The system defined maximum supported degree of parallelism for this account, which restricts the maximum value of parallelism the user can set for the account.
     */
    public /*out*/ readonly systemMaxDegreeOfParallelism!: pulumi.Output<number>;
    /**
     * The system defined maximum supported jobs running under the account at the same time, which restricts the maximum number of running jobs the user can set for the account.
     */
    public /*out*/ readonly systemMaxJobCount!: pulumi.Output<number>;
    /**
     * The resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string}>;
    /**
     * The resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The list of virtualNetwork rules associated with this account.
     */
    public /*out*/ readonly virtualNetworkRules!: pulumi.Output<outputs.datalakeanalytics.v20151001preview.VirtualNetworkRuleResponse[]>;

    /**
     * Create a Account resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccountArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.accountName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.dataLakeStoreAccounts === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'dataLakeStoreAccounts'");
            }
            if ((!args || args.defaultDataLakeStoreAccount === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'defaultDataLakeStoreAccount'");
            }
            if ((!args || args.location === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["computePolicies"] = args ? args.computePolicies : undefined;
            inputs["dataLakeStoreAccounts"] = args ? args.dataLakeStoreAccounts : undefined;
            inputs["defaultDataLakeStoreAccount"] = args ? args.defaultDataLakeStoreAccount : undefined;
            inputs["firewallAllowAzureIps"] = args ? args.firewallAllowAzureIps : undefined;
            inputs["firewallRules"] = args ? args.firewallRules : undefined;
            inputs["firewallState"] = args ? args.firewallState : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["maxDegreeOfParallelism"] = args ? args.maxDegreeOfParallelism : undefined;
            inputs["maxDegreeOfParallelismPerJob"] = args ? args.maxDegreeOfParallelismPerJob : undefined;
            inputs["maxJobCount"] = args ? args.maxJobCount : undefined;
            inputs["minPriorityPerJob"] = args ? args.minPriorityPerJob : undefined;
            inputs["newTier"] = args ? args.newTier : undefined;
            inputs["queryStoreRetention"] = (args ? args.queryStoreRetention : undefined) || 30;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["storageAccounts"] = args ? args.storageAccounts : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["accountId"] = undefined /*out*/;
            inputs["creationTime"] = undefined /*out*/;
            inputs["currentTier"] = undefined /*out*/;
            inputs["debugDataAccessLevel"] = undefined /*out*/;
            inputs["endpoint"] = undefined /*out*/;
            inputs["hiveMetastores"] = undefined /*out*/;
            inputs["lastModifiedTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["systemMaxDegreeOfParallelism"] = undefined /*out*/;
            inputs["systemMaxJobCount"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["virtualNetworkRules"] = undefined /*out*/;
        } else {
            inputs["accountId"] = undefined /*out*/;
            inputs["computePolicies"] = undefined /*out*/;
            inputs["creationTime"] = undefined /*out*/;
            inputs["currentTier"] = undefined /*out*/;
            inputs["dataLakeStoreAccounts"] = undefined /*out*/;
            inputs["debugDataAccessLevel"] = undefined /*out*/;
            inputs["defaultDataLakeStoreAccount"] = undefined /*out*/;
            inputs["endpoint"] = undefined /*out*/;
            inputs["firewallAllowAzureIps"] = undefined /*out*/;
            inputs["firewallRules"] = undefined /*out*/;
            inputs["firewallState"] = undefined /*out*/;
            inputs["hiveMetastores"] = undefined /*out*/;
            inputs["lastModifiedTime"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["maxDegreeOfParallelism"] = undefined /*out*/;
            inputs["maxDegreeOfParallelismPerJob"] = undefined /*out*/;
            inputs["maxJobCount"] = undefined /*out*/;
            inputs["minPriorityPerJob"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["newTier"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["queryStoreRetention"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["storageAccounts"] = undefined /*out*/;
            inputs["systemMaxDegreeOfParallelism"] = undefined /*out*/;
            inputs["systemMaxJobCount"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["virtualNetworkRules"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:datalakeanalytics/latest:Account" }, { type: "azure-nextgen:datalakeanalytics/v20161101:Account" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Account.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Account resource.
 */
export interface AccountArgs {
    /**
     * The name of the Data Lake Analytics account to retrieve.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * The list of compute policies associated with this account.
     */
    readonly computePolicies?: pulumi.Input<pulumi.Input<inputs.datalakeanalytics.v20151001preview.CreateComputePolicyWithAccountParameters>[]>;
    /**
     * The list of Data Lake Store accounts associated with this account.
     */
    readonly dataLakeStoreAccounts: pulumi.Input<pulumi.Input<inputs.datalakeanalytics.v20151001preview.AddDataLakeStoreWithAccountParameters>[]>;
    /**
     * The default Data Lake Store account associated with this account.
     */
    readonly defaultDataLakeStoreAccount: pulumi.Input<string>;
    /**
     * The current state of allowing or disallowing IPs originating within Azure through the firewall. If the firewall is disabled, this is not enforced.
     */
    readonly firewallAllowAzureIps?: pulumi.Input<enums.datalakeanalytics.v20151001preview.FirewallAllowAzureIpsState>;
    /**
     * The list of firewall rules associated with this account.
     */
    readonly firewallRules?: pulumi.Input<pulumi.Input<inputs.datalakeanalytics.v20151001preview.CreateFirewallRuleWithAccountParameters>[]>;
    /**
     * The current state of the IP address firewall for this account.
     */
    readonly firewallState?: pulumi.Input<enums.datalakeanalytics.v20151001preview.FirewallState>;
    /**
     * The resource location.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The maximum supported degree of parallelism for this account.
     */
    readonly maxDegreeOfParallelism?: pulumi.Input<number>;
    /**
     * The maximum supported degree of parallelism per job for this account.
     */
    readonly maxDegreeOfParallelismPerJob?: pulumi.Input<number>;
    /**
     * The maximum supported jobs running under the account at the same time.
     */
    readonly maxJobCount?: pulumi.Input<number>;
    /**
     * The minimum supported priority per job for this account.
     */
    readonly minPriorityPerJob?: pulumi.Input<number>;
    /**
     * The commitment tier for the next month.
     */
    readonly newTier?: pulumi.Input<enums.datalakeanalytics.v20151001preview.TierType>;
    /**
     * The number of days that job metadata is retained.
     */
    readonly queryStoreRetention?: pulumi.Input<number>;
    /**
     * The name of the Azure resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The list of Azure Blob Storage accounts associated with this account.
     */
    readonly storageAccounts?: pulumi.Input<pulumi.Input<inputs.datalakeanalytics.v20151001preview.AddStorageAccountWithAccountParameters>[]>;
    /**
     * The resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
