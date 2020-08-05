// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataBoxEdge.V20190301
{
    /// <summary>
    /// Represents a user who has access to one or more shares on the Data Box Edge/Gateway device.
    /// </summary>
    public partial class User : Pulumi.CustomResource
    {
        /// <summary>
        /// The object name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The storage account credential properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.UserPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// The hierarchical type of the object.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a User resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public User(string name, UserArgs args, CustomResourceOptions? options = null)
            : base("azurerm:databoxedge/v20190301:User", name, args ?? new UserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private User(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:databoxedge/v20190301:User", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing User resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static User Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new User(name, id, options);
        }
    }

    public sealed class UserArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The device name.
        /// </summary>
        [Input("deviceName", required: true)]
        public Input<string> DeviceName { get; set; } = null!;

        /// <summary>
        /// The password details.
        /// </summary>
        [Input("encryptedPassword")]
        public Input<Inputs.AsymmetricEncryptedSecretArgs>? EncryptedPassword { get; set; }

        /// <summary>
        /// The user name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("shareAccessRights")]
        private InputList<Inputs.ShareAccessRightArgs>? _shareAccessRights;

        /// <summary>
        /// List of shares that the user has rights on. This field should not be specified during user creation.
        /// </summary>
        public InputList<Inputs.ShareAccessRightArgs> ShareAccessRights
        {
            get => _shareAccessRights ?? (_shareAccessRights = new InputList<Inputs.ShareAccessRightArgs>());
            set => _shareAccessRights = value;
        }

        public UserArgs()
        {
        }
    }
}
