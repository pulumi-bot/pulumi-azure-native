// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Consumption.V20191001
{
    /// <summary>
    /// A budget resource.
    /// </summary>
    public partial class Budget : Pulumi.CustomResource
    {
        /// <summary>
        /// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
        /// </summary>
        [Output("eTag")]
        public Output<string?> ETag { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The properties of the budget.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.BudgetPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Budget resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Budget(string name, BudgetArgs args, CustomResourceOptions? options = null)
            : base("azurerm:consumption/v20191001:Budget", name, args ?? new BudgetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Budget(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:consumption/v20191001:Budget", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Budget resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Budget Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Budget(name, id, options);
        }
    }

    public sealed class BudgetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
        /// </summary>
        [Input("eTag")]
        public Input<string>? ETag { get; set; }

        /// <summary>
        /// Budget Name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The properties of the budget.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.BudgetPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The scope associated with budget operations. This includes '/subscriptions/{subscriptionId}/' for subscription scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, '/providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for invoiceSection scope.
        /// </summary>
        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        public BudgetArgs()
        {
        }
    }
}