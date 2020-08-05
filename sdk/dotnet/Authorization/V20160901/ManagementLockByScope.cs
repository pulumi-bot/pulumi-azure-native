// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Authorization.V20160901
{
    /// <summary>
    /// The lock information.
    /// </summary>
    public partial class ManagementLockByScope : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the lock.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The properties of the lock.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ManagementLockPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// The resource type of the lock - Microsoft.Authorization/locks.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ManagementLockByScope resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagementLockByScope(string name, ManagementLockByScopeArgs args, CustomResourceOptions? options = null)
            : base("azurerm:authorization/v20160901:ManagementLockByScope", name, args ?? new ManagementLockByScopeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagementLockByScope(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:authorization/v20160901:ManagementLockByScope", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ManagementLockByScope resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ManagementLockByScope Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ManagementLockByScope(name, id, options);
        }
    }

    public sealed class ManagementLockByScopeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it.
        /// </summary>
        [Input("level", required: true)]
        public Input<string> Level { get; set; } = null!;

        /// <summary>
        /// The name of lock.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Notes about the lock. Maximum of 512 characters.
        /// </summary>
        [Input("notes")]
        public Input<string>? Notes { get; set; }

        [Input("owners")]
        private InputList<Inputs.ManagementLockOwnerArgs>? _owners;

        /// <summary>
        /// The owners of the lock.
        /// </summary>
        public InputList<Inputs.ManagementLockOwnerArgs> Owners
        {
            get => _owners ?? (_owners = new InputList<Inputs.ManagementLockOwnerArgs>());
            set => _owners = value;
        }

        /// <summary>
        /// The scope for the lock. When providing a scope for the assignment, use '/subscriptions/{subscriptionId}' for subscriptions, '/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}' for resource groups, and '/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePathIfPresent}/{resourceType}/{resourceName}' for resources.
        /// </summary>
        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        public ManagementLockByScopeArgs()
        {
        }
    }
}
