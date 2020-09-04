// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.TimeSeriesInsights.V20170228Preview
{
    /// <summary>
    /// An access policy is used to grant users and applications access to the environment. Roles are assigned to service principals in Azure Active Directory. These roles define the actions the principal can perform through the Time Series Insights data plane APIs.
    /// </summary>
    public partial class AccessPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// An description of the access policy.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The objectId of the principal in Azure Active Directory.
        /// </summary>
        [Output("principalObjectId")]
        public Output<string?> PrincipalObjectId { get; private set; } = null!;

        /// <summary>
        /// The list of roles the principal is assigned on the environment.
        /// </summary>
        [Output("roles")]
        public Output<ImmutableArray<string>> Roles { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a AccessPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccessPolicy(string name, AccessPolicyArgs args, CustomResourceOptions? options = null)
            : base("azurerm:timeseriesinsights/v20170228preview:AccessPolicy", name, args ?? new AccessPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccessPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:timeseriesinsights/v20170228preview:AccessPolicy", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:timeseriesinsights/latest:AccessPolicy"},
                    new Pulumi.Alias { Type = "azurerm:timeseriesinsights/v20171115:AccessPolicy"},
                    new Pulumi.Alias { Type = "azurerm:timeseriesinsights/v20180815preview:AccessPolicy"},
                    new Pulumi.Alias { Type = "azurerm:timeseriesinsights/v20200515:AccessPolicy"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AccessPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccessPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AccessPolicy(name, id, options);
        }
    }

    public sealed class AccessPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the access policy.
        /// </summary>
        [Input("accessPolicyName", required: true)]
        public Input<string> AccessPolicyName { get; set; } = null!;

        /// <summary>
        /// An description of the access policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the Time Series Insights environment associated with the specified resource group.
        /// </summary>
        [Input("environmentName", required: true)]
        public Input<string> EnvironmentName { get; set; } = null!;

        /// <summary>
        /// The objectId of the principal in Azure Active Directory.
        /// </summary>
        [Input("principalObjectId")]
        public Input<string>? PrincipalObjectId { get; set; }

        /// <summary>
        /// Name of an Azure Resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("roles")]
        private InputList<string>? _roles;

        /// <summary>
        /// The list of roles the principal is assigned on the environment.
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        public AccessPolicyArgs()
        {
        }
    }
}
