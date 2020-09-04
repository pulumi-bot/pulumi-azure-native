// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.Latest
{
    /// <summary>
    /// The integration account session.
    /// 
    /// ## Example Usage
    /// ### Create or update an integration account session
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var integrationAccountSession = new AzureRM.Logic.Latest.IntegrationAccountSession("integrationAccountSession", new AzureRM.Logic.Latest.IntegrationAccountSessionArgs
    ///         {
    ///             Content = 
    ///             {
    ///                 { "controlNumber", "1234" },
    ///                 { "controlNumberChangedTime", "2017-02-21T22:30:11.9923759Z" },
    ///             },
    ///             IntegrationAccountName = "testia123",
    ///             ResourceGroupName = "testrg123",
    ///             SessionName = "testsession123-ICN",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class IntegrationAccountSession : Pulumi.CustomResource
    {
        /// <summary>
        /// The changed time.
        /// </summary>
        [Output("changedTime")]
        public Output<string> ChangedTime { get; private set; } = null!;

        /// <summary>
        /// The session content.
        /// </summary>
        [Output("content")]
        public Output<ImmutableDictionary<string, object>?> Content { get; private set; } = null!;

        /// <summary>
        /// The created time.
        /// </summary>
        [Output("createdTime")]
        public Output<string> CreatedTime { get; private set; } = null!;

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
        /// Create a IntegrationAccountSession resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IntegrationAccountSession(string name, IntegrationAccountSessionArgs args, CustomResourceOptions? options = null)
            : base("azurerm:logic/latest:IntegrationAccountSession", name, args ?? new IntegrationAccountSessionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IntegrationAccountSession(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:logic/latest:IntegrationAccountSession", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:logic/v20160601:IntegrationAccountSession"},
                    new Pulumi.Alias { Type = "azurerm:logic/v20180701preview:IntegrationAccountSession"},
                    new Pulumi.Alias { Type = "azurerm:logic/v20190501:IntegrationAccountSession"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IntegrationAccountSession resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IntegrationAccountSession Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new IntegrationAccountSession(name, id, options);
        }
    }

    public sealed class IntegrationAccountSessionArgs : Pulumi.ResourceArgs
    {
        [Input("content")]
        private InputMap<object>? _content;

        /// <summary>
        /// The session content.
        /// </summary>
        public InputMap<object> Content
        {
            get => _content ?? (_content = new InputMap<object>());
            set => _content = value;
        }

        /// <summary>
        /// The integration account name.
        /// </summary>
        [Input("integrationAccountName", required: true)]
        public Input<string> IntegrationAccountName { get; set; } = null!;

        /// <summary>
        /// The resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The integration account session name.
        /// </summary>
        [Input("sessionName", required: true)]
        public Input<string> SessionName { get; set; } = null!;

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

        public IntegrationAccountSessionArgs()
        {
        }
    }
}
