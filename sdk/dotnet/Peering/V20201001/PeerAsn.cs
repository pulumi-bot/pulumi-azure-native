// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Peering.V20201001
{
    /// <summary>
    /// The essential information related to the peer's ASN.
    /// </summary>
    public partial class PeerAsn : Pulumi.CustomResource
    {
        /// <summary>
        /// The error message for the validation state
        /// </summary>
        [Output("errorMessage")]
        public Output<string> ErrorMessage { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Autonomous System Number (ASN) of the peer.
        /// </summary>
        [Output("peerAsn")]
        public Output<int?> PeerAsnValue { get; private set; } = null!;

        /// <summary>
        /// The contact details of the peer.
        /// </summary>
        [Output("peerContactDetail")]
        public Output<ImmutableArray<Outputs.ContactDetailResponse>> PeerContactDetail { get; private set; } = null!;

        /// <summary>
        /// The name of the peer.
        /// </summary>
        [Output("peerName")]
        public Output<string?> PeerName { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The validation state of the ASN associated with the peer.
        /// </summary>
        [Output("validationState")]
        public Output<string?> ValidationState { get; private set; } = null!;


        /// <summary>
        /// Create a PeerAsn resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PeerAsn(string name, PeerAsnArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:peering/v20201001:PeerAsn", name, args ?? new PeerAsnArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PeerAsn(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:peering/v20201001:PeerAsn", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:peering/latest:PeerAsn"},
                    new Pulumi.Alias { Type = "azure-nextgen:peering/v20190801preview:PeerAsn"},
                    new Pulumi.Alias { Type = "azure-nextgen:peering/v20190901preview:PeerAsn"},
                    new Pulumi.Alias { Type = "azure-nextgen:peering/v20200101preview:PeerAsn"},
                    new Pulumi.Alias { Type = "azure-nextgen:peering/v20200401:PeerAsn"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PeerAsn resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PeerAsn Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PeerAsn(name, id, options);
        }
    }

    public sealed class PeerAsnArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Autonomous System Number (ASN) of the peer.
        /// </summary>
        [Input("peerAsn")]
        public Input<int>? PeerAsn { get; set; }

        /// <summary>
        /// The peer ASN name.
        /// </summary>
        [Input("peerAsnName", required: true)]
        public Input<string> PeerAsnName { get; set; } = null!;

        [Input("peerContactDetail")]
        private InputList<Inputs.ContactDetailArgs>? _peerContactDetail;

        /// <summary>
        /// The contact details of the peer.
        /// </summary>
        public InputList<Inputs.ContactDetailArgs> PeerContactDetail
        {
            get => _peerContactDetail ?? (_peerContactDetail = new InputList<Inputs.ContactDetailArgs>());
            set => _peerContactDetail = value;
        }

        /// <summary>
        /// The name of the peer.
        /// </summary>
        [Input("peerName")]
        public Input<string>? PeerName { get; set; }

        /// <summary>
        /// The validation state of the ASN associated with the peer.
        /// </summary>
        [Input("validationState")]
        public InputUnion<string, Pulumi.AzureNextGen.Peering.V20201001.ValidationState>? ValidationState { get; set; }

        public PeerAsnArgs()
        {
        }
    }
}
