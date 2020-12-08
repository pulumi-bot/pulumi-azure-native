// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.StreamAnalytics.V20170401Preview.Inputs
{

    /// <summary>
    /// Describes a Power BI output data source.
    /// </summary>
    public sealed class PowerBIOutputDataSourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Authentication Mode.
        /// </summary>
        [Input("authenticationMode")]
        public InputUnion<string, Pulumi.AzureNextGen.StreamAnalytics.V20170401Preview.AuthenticationMode>? AuthenticationMode { get; set; }

        /// <summary>
        /// The name of the Power BI dataset. Required on PUT (CreateOrReplace) requests.
        /// </summary>
        [Input("dataset")]
        public Input<string>? Dataset { get; set; }

        /// <summary>
        /// The ID of the Power BI group.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The name of the Power BI group. Use this property to help remember which specific Power BI group id was used.
        /// </summary>
        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        /// <summary>
        /// A refresh token that can be used to obtain a valid access token that can then be used to authenticate with the data source. A valid refresh token is currently only obtainable via the Azure Portal. It is recommended to put a dummy string value here when creating the data source and then going to the Azure Portal to authenticate the data source which will update this property with a valid refresh token. Required on PUT (CreateOrReplace) requests.
        /// </summary>
        [Input("refreshToken")]
        public Input<string>? RefreshToken { get; set; }

        /// <summary>
        /// The name of the Power BI table under the specified dataset. Required on PUT (CreateOrReplace) requests.
        /// </summary>
        [Input("table")]
        public Input<string>? Table { get; set; }

        /// <summary>
        /// The user display name of the user that was used to obtain the refresh token. Use this property to help remember which user was used to obtain the refresh token.
        /// </summary>
        [Input("tokenUserDisplayName")]
        public Input<string>? TokenUserDisplayName { get; set; }

        /// <summary>
        /// The user principal name (UPN) of the user that was used to obtain the refresh token. Use this property to help remember which user was used to obtain the refresh token.
        /// </summary>
        [Input("tokenUserPrincipalName")]
        public Input<string>? TokenUserPrincipalName { get; set; }

        /// <summary>
        /// Indicates the type of data source output will be written to. Required on PUT (CreateOrReplace) requests.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public PowerBIOutputDataSourceArgs()
        {
        }
    }
}
