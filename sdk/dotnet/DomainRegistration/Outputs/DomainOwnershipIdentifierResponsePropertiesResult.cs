// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DomainRegistration.Outputs
{

    [OutputType]
    public sealed class DomainOwnershipIdentifierResponsePropertiesResult
    {
        /// <summary>
        /// Ownership Id.
        /// </summary>
        public readonly string? OwnershipId;

        [OutputConstructor]
        private DomainOwnershipIdentifierResponsePropertiesResult(string? ownershipId)
        {
            OwnershipId = ownershipId;
        }
    }
}