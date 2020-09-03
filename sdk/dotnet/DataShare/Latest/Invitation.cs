// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataShare.Latest
{
    /// <summary>
    /// A Invitation data transfer object.
    /// 
    /// ## Example Usage
    /// ### Invitations_Create
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var invitation = new AzureRM.DataShare.Latest.Invitation("invitation", new AzureRM.DataShare.Latest.InvitationArgs
    ///         {
    ///             AccountName = "Account1",
    ///             InvitationName = "Invitation1",
    ///             ResourceGroupName = "SampleResourceGroup",
    ///             ShareName = "Share1",
    ///             TargetEmail = "receiver@microsoft.com",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class Invitation : Pulumi.CustomResource
    {
        /// <summary>
        /// unique invitation id
        /// </summary>
        [Output("invitationId")]
        public Output<string> InvitationId { get; private set; } = null!;

        /// <summary>
        /// The status of the invitation.
        /// </summary>
        [Output("invitationStatus")]
        public Output<string> InvitationStatus { get; private set; } = null!;

        /// <summary>
        /// Name of the azure resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The time the recipient responded to the invitation.
        /// </summary>
        [Output("respondedAt")]
        public Output<string> RespondedAt { get; private set; } = null!;

        /// <summary>
        /// Gets the time at which the invitation was sent.
        /// </summary>
        [Output("sentAt")]
        public Output<string> SentAt { get; private set; } = null!;

        /// <summary>
        /// The target Azure AD Id. Can't be combined with email.
        /// </summary>
        [Output("targetActiveDirectoryId")]
        public Output<string?> TargetActiveDirectoryId { get; private set; } = null!;

        /// <summary>
        /// The email the invitation is directed to.
        /// </summary>
        [Output("targetEmail")]
        public Output<string?> TargetEmail { get; private set; } = null!;

        /// <summary>
        /// The target user or application Id that invitation is being sent to.
        /// Must be specified along TargetActiveDirectoryId. This enables sending
        /// invitations to specific users or applications in an AD tenant.
        /// </summary>
        [Output("targetObjectId")]
        public Output<string?> TargetObjectId { get; private set; } = null!;

        /// <summary>
        /// Type of the azure resource
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Email of the user who created the resource
        /// </summary>
        [Output("userEmail")]
        public Output<string> UserEmail { get; private set; } = null!;

        /// <summary>
        /// Name of the user who created the resource
        /// </summary>
        [Output("userName")]
        public Output<string> UserName { get; private set; } = null!;


        /// <summary>
        /// Create a Invitation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Invitation(string name, InvitationArgs args, CustomResourceOptions? options = null)
            : base("azurerm:datashare/latest:Invitation", name, args ?? new InvitationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Invitation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:datashare/latest:Invitation", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:datashare/v20181101preview:Invitation"},
                    new Pulumi.Alias { Type = "azurerm:datashare/v20191101:Invitation"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Invitation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Invitation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Invitation(name, id, options);
        }
    }

    public sealed class InvitationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the share account.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the invitation.
        /// </summary>
        [Input("invitationName", required: true)]
        public Input<string> InvitationName { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the share to send the invitation for.
        /// </summary>
        [Input("shareName", required: true)]
        public Input<string> ShareName { get; set; } = null!;

        /// <summary>
        /// The target Azure AD Id. Can't be combined with email.
        /// </summary>
        [Input("targetActiveDirectoryId")]
        public Input<string>? TargetActiveDirectoryId { get; set; }

        /// <summary>
        /// The email the invitation is directed to.
        /// </summary>
        [Input("targetEmail")]
        public Input<string>? TargetEmail { get; set; }

        /// <summary>
        /// The target user or application Id that invitation is being sent to.
        /// Must be specified along TargetActiveDirectoryId. This enables sending
        /// invitations to specific users or applications in an AD tenant.
        /// </summary>
        [Input("targetObjectId")]
        public Input<string>? TargetObjectId { get; set; }

        public InvitationArgs()
        {
        }
    }
}
