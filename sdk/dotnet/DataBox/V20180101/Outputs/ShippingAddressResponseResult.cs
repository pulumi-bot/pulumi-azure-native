// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataBox.V20180101.Outputs
{

    [OutputType]
    public sealed class ShippingAddressResponseResult
    {
        /// <summary>
        /// Type of address.
        /// </summary>
        public readonly string? AddressType;
        /// <summary>
        /// Name of the City.
        /// </summary>
        public readonly string? City;
        /// <summary>
        /// Name of the company.
        /// </summary>
        public readonly string? CompanyName;
        /// <summary>
        /// Name of the Country.
        /// </summary>
        public readonly string Country;
        /// <summary>
        /// Postal code.
        /// </summary>
        public readonly string PostalCode;
        /// <summary>
        /// Name of the State or Province.
        /// </summary>
        public readonly string? StateOrProvince;
        /// <summary>
        /// Street Address line 1.
        /// </summary>
        public readonly string StreetAddress1;
        /// <summary>
        /// Street Address line 2.
        /// </summary>
        public readonly string? StreetAddress2;
        /// <summary>
        /// Street Address line 3.
        /// </summary>
        public readonly string? StreetAddress3;
        /// <summary>
        /// Extended Zip Code.
        /// </summary>
        public readonly string? ZipExtendedCode;

        [OutputConstructor]
        private ShippingAddressResponseResult(
            string? addressType,

            string? city,

            string? companyName,

            string country,

            string postalCode,

            string? stateOrProvince,

            string streetAddress1,

            string? streetAddress2,

            string? streetAddress3,

            string? zipExtendedCode)
        {
            AddressType = addressType;
            City = city;
            CompanyName = companyName;
            Country = country;
            PostalCode = postalCode;
            StateOrProvince = stateOrProvince;
            StreetAddress1 = streetAddress1;
            StreetAddress2 = streetAddress2;
            StreetAddress3 = streetAddress3;
            ZipExtendedCode = zipExtendedCode;
        }
    }
}