// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Security.V20170801Preview
{
    /// <summary>
    /// Security Solution
    /// </summary>
    public partial class IotSecuritySolution : Pulumi.CustomResource
    {
        /// <summary>
        /// List of resources that were automatically discovered as relevant to the security solution.
        /// </summary>
        [Output("autoDiscoveredResources")]
        public Output<ImmutableArray<string>> AutoDiscoveredResources { get; private set; } = null!;

        /// <summary>
        /// Disabled data sources. Disabling these data sources compromises the system.
        /// </summary>
        [Output("disabledDataSources")]
        public Output<ImmutableArray<string>> DisabledDataSources { get; private set; } = null!;

        /// <summary>
        /// Resource display name.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// List of additional export to workspace data options
        /// </summary>
        [Output("export")]
        public Output<ImmutableArray<string>> Export { get; private set; } = null!;

        /// <summary>
        /// IoT Hub resource IDs
        /// </summary>
        [Output("iotHubs")]
        public Output<ImmutableArray<string>> IotHubs { get; private set; } = null!;

        /// <summary>
        /// The resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// List of recommendation configuration
        /// </summary>
        [Output("recommendationsConfiguration")]
        public Output<ImmutableArray<Outputs.RecommendationConfigurationPropertiesResponse>> RecommendationsConfiguration { get; private set; } = null!;

        /// <summary>
        /// Security solution status
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Properties of the solution's user defined resources.
        /// </summary>
        [Output("userDefinedResources")]
        public Output<Outputs.UserDefinedResourcesPropertiesResponse?> UserDefinedResources { get; private set; } = null!;

        /// <summary>
        /// Workspace resource ID
        /// </summary>
        [Output("workspace")]
        public Output<string> Workspace { get; private set; } = null!;


        /// <summary>
        /// Create a IotSecuritySolution resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IotSecuritySolution(string name, IotSecuritySolutionArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:security/v20170801preview:IotSecuritySolution", name, args ?? new IotSecuritySolutionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IotSecuritySolution(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:security/v20170801preview:IotSecuritySolution", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:security/latest:IotSecuritySolution"},
                    new Pulumi.Alias { Type = "azure-nextgen:security/v20190801:IotSecuritySolution"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IotSecuritySolution resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IotSecuritySolution Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new IotSecuritySolution(name, id, options);
        }
    }

    public sealed class IotSecuritySolutionArgs : Pulumi.ResourceArgs
    {
        [Input("disabledDataSources")]
        private InputList<Union<string, Pulumi.AzureNextGen.Security.V20170801Preview.DataSource>>? _disabledDataSources;

        /// <summary>
        /// Disabled data sources. Disabling these data sources compromises the system.
        /// </summary>
        public InputList<Union<string, Pulumi.AzureNextGen.Security.V20170801Preview.DataSource>> DisabledDataSources
        {
            get => _disabledDataSources ?? (_disabledDataSources = new InputList<Union<string, Pulumi.AzureNextGen.Security.V20170801Preview.DataSource>>());
            set => _disabledDataSources = value;
        }

        /// <summary>
        /// Resource display name.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        [Input("export")]
        private InputList<Union<string, Pulumi.AzureNextGen.Security.V20170801Preview.ExportData>>? _export;

        /// <summary>
        /// List of additional export to workspace data options
        /// </summary>
        public InputList<Union<string, Pulumi.AzureNextGen.Security.V20170801Preview.ExportData>> Export
        {
            get => _export ?? (_export = new InputList<Union<string, Pulumi.AzureNextGen.Security.V20170801Preview.ExportData>>());
            set => _export = value;
        }

        [Input("iotHubs", required: true)]
        private InputList<string>? _iotHubs;

        /// <summary>
        /// IoT Hub resource IDs
        /// </summary>
        public InputList<string> IotHubs
        {
            get => _iotHubs ?? (_iotHubs = new InputList<string>());
            set => _iotHubs = value;
        }

        /// <summary>
        /// The resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("recommendationsConfiguration")]
        private InputList<Inputs.RecommendationConfigurationPropertiesArgs>? _recommendationsConfiguration;

        /// <summary>
        /// List of recommendation configuration
        /// </summary>
        public InputList<Inputs.RecommendationConfigurationPropertiesArgs> RecommendationsConfiguration
        {
            get => _recommendationsConfiguration ?? (_recommendationsConfiguration = new InputList<Inputs.RecommendationConfigurationPropertiesArgs>());
            set => _recommendationsConfiguration = value;
        }

        /// <summary>
        /// The name of the resource group within the user's subscription. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The solution manager name
        /// </summary>
        [Input("solutionName", required: true)]
        public Input<string> SolutionName { get; set; } = null!;

        /// <summary>
        /// Security solution status
        /// </summary>
        [Input("status")]
        public InputUnion<string, Pulumi.AzureNextGen.Security.V20170801Preview.SecuritySolutionStatus>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Properties of the solution's user defined resources.
        /// </summary>
        [Input("userDefinedResources")]
        public Input<Inputs.UserDefinedResourcesPropertiesArgs>? UserDefinedResources { get; set; }

        /// <summary>
        /// Workspace resource ID
        /// </summary>
        [Input("workspace", required: true)]
        public Input<string> Workspace { get; set; } = null!;

        public IotSecuritySolutionArgs()
        {
        }
    }
}
