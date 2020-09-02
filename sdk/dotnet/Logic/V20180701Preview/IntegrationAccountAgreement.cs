// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.V20180701Preview
{
    /// <summary>
    /// The integration account agreement.
    /// </summary>
    public partial class IntegrationAccountAgreement : Pulumi.CustomResource
    {
        /// <summary>
        /// The agreement type.
        /// </summary>
        [Output("agreementType")]
        public Output<string> AgreementType { get; private set; } = null!;

        /// <summary>
        /// The changed time.
        /// </summary>
        [Output("changedTime")]
        public Output<string> ChangedTime { get; private set; } = null!;

        /// <summary>
        /// The agreement content.
        /// </summary>
        [Output("content")]
        public Output<Outputs.AgreementContentResponseResult> Content { get; private set; } = null!;

        /// <summary>
        /// The created time.
        /// </summary>
        [Output("createdTime")]
        public Output<string> CreatedTime { get; private set; } = null!;

        /// <summary>
        /// The business identity of the guest partner.
        /// </summary>
        [Output("guestIdentity")]
        public Output<Outputs.BusinessIdentityResponseResult> GuestIdentity { get; private set; } = null!;

        /// <summary>
        /// The integration account partner that is set as guest partner for this agreement.
        /// </summary>
        [Output("guestPartner")]
        public Output<string> GuestPartner { get; private set; } = null!;

        /// <summary>
        /// The business identity of the host partner.
        /// </summary>
        [Output("hostIdentity")]
        public Output<Outputs.BusinessIdentityResponseResult> HostIdentity { get; private set; } = null!;

        /// <summary>
        /// The integration account partner that is set as host partner for this agreement.
        /// </summary>
        [Output("hostPartner")]
        public Output<string> HostPartner { get; private set; } = null!;

        /// <summary>
        /// The resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The metadata.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, object>?> Metadata { get; private set; } = null!;

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
        /// Create a IntegrationAccountAgreement resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IntegrationAccountAgreement(string name, IntegrationAccountAgreementArgs args, CustomResourceOptions? options = null)
            : base("azurerm:logic/v20180701preview:IntegrationAccountAgreement", name, args ?? new IntegrationAccountAgreementArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IntegrationAccountAgreement(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:logic/v20180701preview:IntegrationAccountAgreement", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:logic/latest:IntegrationAccountAgreement"},
                    new Pulumi.Alias { Type = "azurerm:logic/v20150801preview:IntegrationAccountAgreement"},
                    new Pulumi.Alias { Type = "azurerm:logic/v20160601:IntegrationAccountAgreement"},
                    new Pulumi.Alias { Type = "azurerm:logic/v20190501:IntegrationAccountAgreement"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IntegrationAccountAgreement resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IntegrationAccountAgreement Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new IntegrationAccountAgreement(name, id, options);
        }
    }

    public sealed class IntegrationAccountAgreementArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The integration account agreement name.
        /// </summary>
        [Input("agreementName", required: true)]
        public Input<string> AgreementName { get; set; } = null!;

        /// <summary>
        /// The agreement type.
        /// </summary>
        [Input("agreementType", required: true)]
        public Input<string> AgreementType { get; set; } = null!;

        /// <summary>
        /// The agreement content.
        /// </summary>
        [Input("content", required: true)]
        public Input<Inputs.AgreementContentArgs> Content { get; set; } = null!;

        /// <summary>
        /// The business identity of the guest partner.
        /// </summary>
        [Input("guestIdentity", required: true)]
        public Input<Inputs.BusinessIdentityArgs> GuestIdentity { get; set; } = null!;

        /// <summary>
        /// The integration account partner that is set as guest partner for this agreement.
        /// </summary>
        [Input("guestPartner", required: true)]
        public Input<string> GuestPartner { get; set; } = null!;

        /// <summary>
        /// The business identity of the host partner.
        /// </summary>
        [Input("hostIdentity", required: true)]
        public Input<Inputs.BusinessIdentityArgs> HostIdentity { get; set; } = null!;

        /// <summary>
        /// The integration account partner that is set as host partner for this agreement.
        /// </summary>
        [Input("hostPartner", required: true)]
        public Input<string> HostPartner { get; set; } = null!;

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

        [Input("metadata")]
        private InputMap<object>? _metadata;

        /// <summary>
        /// The metadata.
        /// </summary>
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

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

        public IntegrationAccountAgreementArgs()
        {
        }
    }
}
