// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Solutions.V20180601
{
    /// <summary>
    /// Information about managed application definition.
    /// </summary>
    public partial class ApplicationDefinition : Pulumi.CustomResource
    {
        /// <summary>
        /// The collection of managed application artifacts. The portal will use the files specified as artifacts to construct the user experience of creating a managed application from a managed application definition.
        /// </summary>
        [Output("artifacts")]
        public Output<ImmutableArray<Outputs.ApplicationArtifactResponseResult>> Artifacts { get; private set; } = null!;

        /// <summary>
        /// The managed application provider authorizations.
        /// </summary>
        [Output("authorizations")]
        public Output<ImmutableArray<Outputs.ApplicationProviderAuthorizationResponseResult>> Authorizations { get; private set; } = null!;

        /// <summary>
        /// The createUiDefinition json for the backing template with Microsoft.Solutions/applications resource. It can be a JObject or well-formed JSON string.
        /// </summary>
        [Output("createUiDefinition")]
        public Output<ImmutableDictionary<string, object>?> CreateUiDefinition { get; private set; } = null!;

        /// <summary>
        /// The managed application definition description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The managed application definition display name.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The identity of the resource.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.IdentityResponseResult?> Identity { get; private set; } = null!;

        /// <summary>
        /// A value indicating whether the package is enabled or not.
        /// </summary>
        [Output("isEnabled")]
        public Output<string?> IsEnabled { get; private set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The managed application lock level.
        /// </summary>
        [Output("lockLevel")]
        public Output<string> LockLevel { get; private set; } = null!;

        /// <summary>
        /// The inline main template json which has resources to be provisioned. It can be a JObject or well-formed JSON string.
        /// </summary>
        [Output("mainTemplate")]
        public Output<ImmutableDictionary<string, object>?> MainTemplate { get; private set; } = null!;

        /// <summary>
        /// ID of the resource that manages this resource.
        /// </summary>
        [Output("managedBy")]
        public Output<string?> ManagedBy { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The managed application definition package file Uri. Use this element
        /// </summary>
        [Output("packageFileUri")]
        public Output<string?> PackageFileUri { get; private set; } = null!;

        /// <summary>
        /// The SKU of the resource.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.SkuResponseResult?> Sku { get; private set; } = null!;

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
        /// Create a ApplicationDefinition resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationDefinition(string name, ApplicationDefinitionArgs args, CustomResourceOptions? options = null)
            : base("azurerm:solutions/v20180601:ApplicationDefinition", name, args ?? new ApplicationDefinitionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationDefinition(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:solutions/v20180601:ApplicationDefinition", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:solutions/latest:ApplicationDefinition"},
                    new Pulumi.Alias { Type = "azurerm:solutions/v20170901:ApplicationDefinition"},
                    new Pulumi.Alias { Type = "azurerm:solutions/v20190701:ApplicationDefinition"},
                    new Pulumi.Alias { Type = "azurerm:solutions/v20200821preview:ApplicationDefinition"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ApplicationDefinition resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationDefinition Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ApplicationDefinition(name, id, options);
        }
    }

    public sealed class ApplicationDefinitionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the managed application definition.
        /// </summary>
        [Input("applicationDefinitionName", required: true)]
        public Input<string> ApplicationDefinitionName { get; set; } = null!;

        [Input("artifacts")]
        private InputList<Inputs.ApplicationArtifactArgs>? _artifacts;

        /// <summary>
        /// The collection of managed application artifacts. The portal will use the files specified as artifacts to construct the user experience of creating a managed application from a managed application definition.
        /// </summary>
        public InputList<Inputs.ApplicationArtifactArgs> Artifacts
        {
            get => _artifacts ?? (_artifacts = new InputList<Inputs.ApplicationArtifactArgs>());
            set => _artifacts = value;
        }

        [Input("authorizations", required: true)]
        private InputList<Inputs.ApplicationProviderAuthorizationArgs>? _authorizations;

        /// <summary>
        /// The managed application provider authorizations.
        /// </summary>
        public InputList<Inputs.ApplicationProviderAuthorizationArgs> Authorizations
        {
            get => _authorizations ?? (_authorizations = new InputList<Inputs.ApplicationProviderAuthorizationArgs>());
            set => _authorizations = value;
        }

        [Input("createUiDefinition")]
        private InputMap<object>? _createUiDefinition;

        /// <summary>
        /// The createUiDefinition json for the backing template with Microsoft.Solutions/applications resource. It can be a JObject or well-formed JSON string.
        /// </summary>
        public InputMap<object> CreateUiDefinition
        {
            get => _createUiDefinition ?? (_createUiDefinition = new InputMap<object>());
            set => _createUiDefinition = value;
        }

        /// <summary>
        /// The managed application definition description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The managed application definition display name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The identity of the resource.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.IdentityArgs>? Identity { get; set; }

        /// <summary>
        /// A value indicating whether the package is enabled or not.
        /// </summary>
        [Input("isEnabled")]
        public Input<string>? IsEnabled { get; set; }

        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The managed application lock level.
        /// </summary>
        [Input("lockLevel", required: true)]
        public Input<string> LockLevel { get; set; } = null!;

        [Input("mainTemplate")]
        private InputMap<object>? _mainTemplate;

        /// <summary>
        /// The inline main template json which has resources to be provisioned. It can be a JObject or well-formed JSON string.
        /// </summary>
        public InputMap<object> MainTemplate
        {
            get => _mainTemplate ?? (_mainTemplate = new InputMap<object>());
            set => _mainTemplate = value;
        }

        /// <summary>
        /// ID of the resource that manages this resource.
        /// </summary>
        [Input("managedBy")]
        public Input<string>? ManagedBy { get; set; }

        /// <summary>
        /// The managed application definition package file Uri. Use this element
        /// </summary>
        [Input("packageFileUri")]
        public Input<string>? PackageFileUri { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The SKU of the resource.
        /// </summary>
        [Input("sku")]
        public Input<Inputs.SkuArgs>? Sku { get; set; }

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

        public ApplicationDefinitionArgs()
        {
        }
    }
}
