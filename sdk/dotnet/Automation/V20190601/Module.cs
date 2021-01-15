// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Automation.V20190601
{
    /// <summary>
    /// Definition of the module type.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:automation/v20190601:Module")]
    public partial class Module : Pulumi.CustomResource
    {
        /// <summary>
        /// Gets or sets the activity count of the module.
        /// </summary>
        [Output("activityCount")]
        public Output<int?> ActivityCount { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the contentLink of the module.
        /// </summary>
        [Output("contentLink")]
        public Output<Outputs.ContentLinkResponse?> ContentLink { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the creation time.
        /// </summary>
        [Output("creationTime")]
        public Output<string?> CreationTime { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the error info of the module.
        /// </summary>
        [Output("error")]
        public Output<Outputs.ModuleErrorInfoResponse?> Error { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the etag of the resource.
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// Gets or sets type of module, if its composite or not.
        /// </summary>
        [Output("isComposite")]
        public Output<bool?> IsComposite { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the isGlobal flag of the module.
        /// </summary>
        [Output("isGlobal")]
        public Output<bool?> IsGlobal { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the last modified time.
        /// </summary>
        [Output("lastModifiedTime")]
        public Output<string?> LastModifiedTime { get; private set; } = null!;

        /// <summary>
        /// The Azure Region where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the provisioning state of the module.
        /// </summary>
        [Output("provisioningState")]
        public Output<string?> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the size in bytes of the module.
        /// </summary>
        [Output("sizeInBytes")]
        public Output<double?> SizeInBytes { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the version of the module.
        /// </summary>
        [Output("version")]
        public Output<string?> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Module resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Module(string name, ModuleArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:automation/v20190601:Module", name, args ?? new ModuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Module(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:automation/v20190601:Module", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:automation/latest:Module"},
                    new Pulumi.Alias { Type = "azure-nextgen:automation/v20151031:Module"},
                    new Pulumi.Alias { Type = "azure-nextgen:automation/v20200113preview:Module"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Module resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Module Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Module(name, id, options);
        }
    }

    public sealed class ModuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the automation account.
        /// </summary>
        [Input("automationAccountName", required: true)]
        public Input<string> AutomationAccountName { get; set; } = null!;

        /// <summary>
        /// Gets or sets the module content link.
        /// </summary>
        [Input("contentLink", required: true)]
        public Input<Inputs.ContentLinkArgs> ContentLink { get; set; } = null!;

        /// <summary>
        /// Gets or sets the location of the resource.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of module.
        /// </summary>
        [Input("moduleName", required: true)]
        public Input<string> ModuleName { get; set; } = null!;

        /// <summary>
        /// Gets or sets name of the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Name of an Azure Resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Gets or sets the tags attached to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ModuleArgs()
        {
        }
    }
}
