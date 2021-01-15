// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Automation.V20200113Preview
{
    /// <summary>
    /// Definition of the variable.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:automation/v20200113preview:Variable")]
    public partial class Variable : Pulumi.CustomResource
    {
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
        /// Gets or sets the encrypted flag of the variable.
        /// </summary>
        [Output("isEncrypted")]
        public Output<bool?> IsEncrypted { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the last modified time.
        /// </summary>
        [Output("lastModifiedTime")]
        public Output<string?> LastModifiedTime { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the value of the variable.
        /// </summary>
        [Output("value")]
        public Output<string?> Value { get; private set; } = null!;


        /// <summary>
        /// Create a Variable resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Variable(string name, VariableArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:automation/v20200113preview:Variable", name, args ?? new VariableArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Variable(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:automation/v20200113preview:Variable", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:automation/latest:Variable"},
                    new Pulumi.Alias { Type = "azure-nextgen:automation/v20151031:Variable"},
                    new Pulumi.Alias { Type = "azure-nextgen:automation/v20190601:Variable"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Variable resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Variable Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Variable(name, id, options);
        }
    }

    public sealed class VariableArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the automation account.
        /// </summary>
        [Input("automationAccountName", required: true)]
        public Input<string> AutomationAccountName { get; set; } = null!;

        /// <summary>
        /// Gets or sets the description of the variable.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Gets or sets the encrypted flag of the variable.
        /// </summary>
        [Input("isEncrypted")]
        public Input<bool>? IsEncrypted { get; set; }

        /// <summary>
        /// Gets or sets the name of the variable.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Name of an Azure Resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Gets or sets the value of the variable.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        /// <summary>
        /// The variable name.
        /// </summary>
        [Input("variableName", required: true)]
        public Input<string> VariableName { get; set; } = null!;

        public VariableArgs()
        {
        }
    }
}
