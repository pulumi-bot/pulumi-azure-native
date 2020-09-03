// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20170301
{
    /// <summary>
    /// Diagnostic details.
    /// 
    /// ## Example Usage
    /// ### ApiManagementCreateApiDiagnostic
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var apiDiagnostic = new AzureRM.ApiManagement.V20170301.ApiDiagnostic("apiDiagnostic", new AzureRM.ApiManagement.V20170301.ApiDiagnosticArgs
    ///         {
    ///             ApiId = "57d1f7558aa04f15146d9d8a",
    ///             DiagnosticId = "default",
    ///             Enabled = true,
    ///             ResourceGroupName = "rg1",
    ///             ServiceName = "apimService1",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class ApiDiagnostic : Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates whether a diagnostic should receive data or not.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Resource type for API Management resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ApiDiagnostic resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApiDiagnostic(string name, ApiDiagnosticArgs args, CustomResourceOptions? options = null)
            : base("azurerm:apimanagement/v20170301:ApiDiagnostic", name, args ?? new ApiDiagnosticArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApiDiagnostic(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:apimanagement/v20170301:ApiDiagnostic", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:apimanagement/latest:ApiDiagnostic"},
                    new Pulumi.Alias { Type = "azurerm:apimanagement/v20180101:ApiDiagnostic"},
                    new Pulumi.Alias { Type = "azurerm:apimanagement/v20180601preview:ApiDiagnostic"},
                    new Pulumi.Alias { Type = "azurerm:apimanagement/v20190101:ApiDiagnostic"},
                    new Pulumi.Alias { Type = "azurerm:apimanagement/v20191201:ApiDiagnostic"},
                    new Pulumi.Alias { Type = "azurerm:apimanagement/v20191201preview:ApiDiagnostic"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ApiDiagnostic resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApiDiagnostic Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ApiDiagnostic(name, id, options);
        }
    }

    public sealed class ApiDiagnosticArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// API identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        /// <summary>
        /// Diagnostic identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("diagnosticId", required: true)]
        public Input<string> DiagnosticId { get; set; } = null!;

        /// <summary>
        /// Indicates whether a diagnostic should receive data or not.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public ApiDiagnosticArgs()
        {
        }
    }
}
