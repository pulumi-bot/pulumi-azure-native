// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.Inputs
{

    /// <summary>
    /// Parameters supplied to the Create Group operation.
    /// </summary>
    public sealed class GroupCreateParametersPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Group description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Group name.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Identifier of the external groups, this property contains the id of the group from the external identity provider, e.g. for Azure Active Directory `aad://&lt;tenant&gt;.onmicrosoft.com/groups/&lt;group object id&gt;`; otherwise the value is null.
        /// </summary>
        [Input("externalId")]
        public Input<string>? ExternalId { get; set; }

        /// <summary>
        /// Group type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public GroupCreateParametersPropertiesArgs()
        {
        }
    }
}