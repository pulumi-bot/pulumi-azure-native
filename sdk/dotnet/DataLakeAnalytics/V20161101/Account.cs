// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataLakeAnalytics.V20161101
{
    /// <summary>
    /// A Data Lake Analytics account object, containing all information associated with the named Data Lake Analytics account.
    /// </summary>
    public partial class Account : Pulumi.CustomResource
    {
        /// <summary>
        /// The resource location.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The properties defined by Data Lake Analytics all properties are specific to each resource provider.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.DataLakeAnalyticsAccountPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// The resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Account resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Account(string name, AccountArgs args, CustomResourceOptions? options = null)
            : base("azurerm:datalakeanalytics/v20161101:Account", name, args ?? new AccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Account(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:datalakeanalytics/v20161101:Account", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Account resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Account Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Account(name, id, options);
        }
    }

    public sealed class AccountArgs : Pulumi.ResourceArgs
    {
        [Input("computePolicies")]
        private InputList<Inputs.CreateComputePolicyWithAccountParametersArgs>? _computePolicies;

        /// <summary>
        /// The list of compute policies associated with this account.
        /// </summary>
        public InputList<Inputs.CreateComputePolicyWithAccountParametersArgs> ComputePolicies
        {
            get => _computePolicies ?? (_computePolicies = new InputList<Inputs.CreateComputePolicyWithAccountParametersArgs>());
            set => _computePolicies = value;
        }

        [Input("dataLakeStoreAccounts", required: true)]
        private InputList<Inputs.AddDataLakeStoreWithAccountParametersArgs>? _dataLakeStoreAccounts;

        /// <summary>
        /// The list of Data Lake Store accounts associated with this account.
        /// </summary>
        public InputList<Inputs.AddDataLakeStoreWithAccountParametersArgs> DataLakeStoreAccounts
        {
            get => _dataLakeStoreAccounts ?? (_dataLakeStoreAccounts = new InputList<Inputs.AddDataLakeStoreWithAccountParametersArgs>());
            set => _dataLakeStoreAccounts = value;
        }

        /// <summary>
        /// The default Data Lake Store account associated with this account.
        /// </summary>
        [Input("defaultDataLakeStoreAccount", required: true)]
        public Input<string> DefaultDataLakeStoreAccount { get; set; } = null!;

        /// <summary>
        /// The current state of allowing or disallowing IPs originating within Azure through the firewall. If the firewall is disabled, this is not enforced.
        /// </summary>
        [Input("firewallAllowAzureIps")]
        public Input<string>? FirewallAllowAzureIps { get; set; }

        [Input("firewallRules")]
        private InputList<Inputs.CreateFirewallRuleWithAccountParametersArgs>? _firewallRules;

        /// <summary>
        /// The list of firewall rules associated with this account.
        /// </summary>
        public InputList<Inputs.CreateFirewallRuleWithAccountParametersArgs> FirewallRules
        {
            get => _firewallRules ?? (_firewallRules = new InputList<Inputs.CreateFirewallRuleWithAccountParametersArgs>());
            set => _firewallRules = value;
        }

        /// <summary>
        /// The current state of the IP address firewall for this account.
        /// </summary>
        [Input("firewallState")]
        public Input<string>? FirewallState { get; set; }

        /// <summary>
        /// The resource location.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The maximum supported degree of parallelism for this account.
        /// </summary>
        [Input("maxDegreeOfParallelism")]
        public Input<int>? MaxDegreeOfParallelism { get; set; }

        /// <summary>
        /// The maximum supported degree of parallelism per job for this account.
        /// </summary>
        [Input("maxDegreeOfParallelismPerJob")]
        public Input<int>? MaxDegreeOfParallelismPerJob { get; set; }

        /// <summary>
        /// The maximum supported jobs running under the account at the same time.
        /// </summary>
        [Input("maxJobCount")]
        public Input<int>? MaxJobCount { get; set; }

        /// <summary>
        /// The minimum supported priority per job for this account.
        /// </summary>
        [Input("minPriorityPerJob")]
        public Input<int>? MinPriorityPerJob { get; set; }

        /// <summary>
        /// The name of the Data Lake Analytics account.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The commitment tier for the next month.
        /// </summary>
        [Input("newTier")]
        public Input<string>? NewTier { get; set; }

        /// <summary>
        /// The number of days that job metadata is retained.
        /// </summary>
        [Input("queryStoreRetention")]
        public Input<int>? QueryStoreRetention { get; set; }

        /// <summary>
        /// The name of the Azure resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("storageAccounts")]
        private InputList<Inputs.AddStorageAccountWithAccountParametersArgs>? _storageAccounts;

        /// <summary>
        /// The list of Azure Blob Storage accounts associated with this account.
        /// </summary>
        public InputList<Inputs.AddStorageAccountWithAccountParametersArgs> StorageAccounts
        {
            get => _storageAccounts ?? (_storageAccounts = new InputList<Inputs.AddStorageAccountWithAccountParametersArgs>());
            set => _storageAccounts = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public AccountArgs()
        {
        }
    }
}
