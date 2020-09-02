// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearningCompute.V20170601Preview.Outputs
{

    [OutputType]
    public sealed class SslConfigurationResponseResult
    {
        /// <summary>
        /// The SSL cert data in PEM format encoded as base64 string
        /// </summary>
        public readonly string? Cert;
        /// <summary>
        /// The SSL key data in PEM format encoded as base64 string. This is not returned in response of GET/PUT on the resource.. To see this please call listKeys API.
        /// </summary>
        public readonly string? Key;
        /// <summary>
        /// SSL status. Allowed values are Enabled and Disabled.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private SslConfigurationResponseResult(
            string? cert,

            string? key,

            string? status)
        {
            Cert = cert;
            Key = key;
            Status = status;
        }
    }
}
