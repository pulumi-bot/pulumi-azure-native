// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Billing.V20200301Preview
{
    /// <summary>
    /// The cost allocation rule model definition
    /// </summary>
    public partial class CostAllocationRule : Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the rule. This is a read only value.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Cost allocation rule properties
        /// </summary>
        [Output("properties")]
        public Output<Outputs.CostAllocationRulePropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource type of the rule. This is a read only value of Microsoft.CostManagement/CostAllocationRule.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a CostAllocationRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CostAllocationRule(string name, CostAllocationRuleArgs args, CustomResourceOptions? options = null)
            : base("azurerm:billing/v20200301preview:CostAllocationRule", name, args ?? new CostAllocationRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CostAllocationRule(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:billing/v20200301preview:CostAllocationRule", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing CostAllocationRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CostAllocationRule Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CostAllocationRule(name, id, options);
        }
    }

    public sealed class CostAllocationRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// BillingAccount ID
        /// </summary>
        [Input("billingAccountId", required: true)]
        public Input<string> BillingAccountId { get; set; } = null!;

        /// <summary>
        /// Cost allocation rule properties
        /// </summary>
        [Input("properties")]
        public Input<Inputs.CostAllocationRulePropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// Cost allocation rule name. The name cannot include spaces or any non alphanumeric characters other than '_' and '-'. The max length is 260 characters.
        /// </summary>
        [Input("ruleName", required: true)]
        public Input<string> RuleName { get; set; } = null!;

        public CostAllocationRuleArgs()
        {
        }
    }
}
