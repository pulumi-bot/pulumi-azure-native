// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.V20190501
{
    /// <summary>
    /// The managed api definition.
    /// 
    /// ## Example Usage
    /// ### Gets the integration service environment managed Apis
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var integrationServiceEnvironmentManagedApi = new AzureRM.Logic.V20190501.IntegrationServiceEnvironmentManagedApi("integrationServiceEnvironmentManagedApi", new AzureRM.Logic.V20190501.IntegrationServiceEnvironmentManagedApiArgs
    ///         {
    ///             ApiName = "servicebus",
    ///             IntegrationServiceEnvironmentName = "testIntegrationServiceEnvironment",
    ///             ResourceGroup = "testResourceGroup",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class IntegrationServiceEnvironmentManagedApi : Pulumi.CustomResource
    {
        /// <summary>
        /// The resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Gets the resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The api resource properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ApiResourcePropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// The resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Gets the resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a IntegrationServiceEnvironmentManagedApi resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IntegrationServiceEnvironmentManagedApi(string name, IntegrationServiceEnvironmentManagedApiArgs args, CustomResourceOptions? options = null)
            : base("azurerm:logic/v20190501:IntegrationServiceEnvironmentManagedApi", name, args ?? new IntegrationServiceEnvironmentManagedApiArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IntegrationServiceEnvironmentManagedApi(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:logic/v20190501:IntegrationServiceEnvironmentManagedApi", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:logic/latest:IntegrationServiceEnvironmentManagedApi"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IntegrationServiceEnvironmentManagedApi resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IntegrationServiceEnvironmentManagedApi Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new IntegrationServiceEnvironmentManagedApi(name, id, options);
        }
    }

    public sealed class IntegrationServiceEnvironmentManagedApiArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The api name.
        /// </summary>
        [Input("apiName", required: true)]
        public Input<string> ApiName { get; set; } = null!;

        /// <summary>
        /// The integration service environment name.
        /// </summary>
        [Input("integrationServiceEnvironmentName", required: true)]
        public Input<string> IntegrationServiceEnvironmentName { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroup", required: true)]
        public Input<string> ResourceGroup { get; set; } = null!;

        public IntegrationServiceEnvironmentManagedApiArgs()
        {
        }
    }
}
