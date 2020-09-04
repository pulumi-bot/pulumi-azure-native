// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Solutions.V20200821Preview.Outputs
{

    [OutputType]
    public sealed class ApplicationPackageContactResponseResult
    {
        /// <summary>
        /// The contact name.
        /// </summary>
        public readonly string? ContactName;
        /// <summary>
        /// The contact email.
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// The contact phone number.
        /// </summary>
        public readonly string Phone;

        [OutputConstructor]
        private ApplicationPackageContactResponseResult(
            string? contactName,

            string email,

            string phone)
        {
            ContactName = contactName;
            Email = email;
            Phone = phone;
        }
    }
}
