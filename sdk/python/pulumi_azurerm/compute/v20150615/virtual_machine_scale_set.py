# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class VirtualMachineScaleSet(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    Resource location
    """
    name: pulumi.Output[str]
    """
    Resource name
    """
    properties: pulumi.Output[dict]
    """
    Describes the properties of a Virtual Machine Scale Set.
      * `over_provision` (`bool`) - Specifies whether the Virtual Machine Scale Set should be overprovisioned.
      * `provisioning_state` (`str`) - The provisioning state, which only appears in the response.
      * `upgrade_policy` (`dict`) - The upgrade policy.
        * `mode` (`str`) - Specifies the mode of an upgrade to virtual machines in the scale set.<br /><br /> Possible values are:<br /><br /> **Manual** - You  control the application of updates to virtual machines in the scale set. You do this by using the manualUpgrade action.<br /><br /> **Automatic** - All virtual machines in the scale set are  automatically updated at the same time.

      * `virtual_machine_profile` (`dict`) - The virtual machine profile.
        * `extension_profile` (`dict`) - The virtual machine scale set extension profile.
          * `extensions` (`list`) - The virtual machine scale set child extension resources.
            * `id` (`str`) - Resource Id
            * `name` (`str`) - The name of the extension.
            * `properties` (`dict`) - Describes the properties of a Virtual Machine Scale Set Extension.
              * `auto_upgrade_minor_version` (`bool`) - Whether the extension handler should be automatically upgraded across minor versions.
              * `protected_settings` (`dict`) - Json formatted protected settings for the extension.
              * `provisioning_state` (`str`) - The provisioning state, which only appears in the response.
              * `publisher` (`str`) - The name of the extension handler publisher.
              * `settings` (`dict`) - Json formatted public settings for the extension.
              * `type` (`str`) - The type of the extension handler.
              * `type_handler_version` (`str`) - The type version of the extension handler.

        * `network_profile` (`dict`) - The virtual machine scale set network profile.
          * `network_interface_configurations` (`list`) - The list of network configurations.
            * `id` (`str`) - Resource Id
            * `name` (`str`) - The network configuration name.
            * `properties` (`dict`) - Describes a virtual machine scale set network profile's IP configuration.
              * `ip_configurations` (`list`) - The virtual machine scale set IP Configuration.
                * `id` (`str`) - Resource Id
                * `name` (`str`) - The IP configuration name.
                * `properties` (`dict`) - Describes a virtual machine scale set network profile's IP configuration properties.
                  * `load_balancer_backend_address_pools` (`list`) - The load balancer backend address pools.
                    * `id` (`str`) - Resource Id

                  * `load_balancer_inbound_nat_pools` (`list`) - The load balancer inbound nat pools.
                  * `subnet` (`dict`) - The subnet.
                    * `id` (`str`) - The ARM resource id in the form of /subscriptions/{SubscriptionId}/resourceGroups/{ResourceGroupName}/...

              * `primary` (`bool`) - Whether this is a primary NIC on a virtual machine.

        * `os_profile` (`dict`) - The virtual machine scale set OS profile.
          * `admin_password` (`str`) - Specifies the password of the administrator account. <br><br> **Minimum-length (Windows):** 8 characters <br><br> **Minimum-length (Linux):** 6 characters <br><br> **Max-length (Windows):** 123 characters <br><br> **Max-length (Linux):** 72 characters <br><br> **Complexity requirements:** 3 out of 4 conditions below need to be fulfilled <br> Has lower characters <br>Has upper characters <br> Has a digit <br> Has a special character (Regex match [\W_]) <br><br> **Disallowed values:** "abc@123", "P@$$w0rd", "P@ssw0rd", "P@ssword123", "Pa$$word", "pass@word1", "Password!", "Password1", "Password22", "iloveyou!" <br><br> For resetting the password, see [How to reset the Remote Desktop service or its login password in a Windows VM](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-reset-rdp?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> For resetting root password, see [Manage users, SSH, and check or repair disks on Azure Linux VMs using the VMAccess Extension](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-using-vmaccess-extension?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json#reset-root-password)
          * `admin_username` (`str`) - Specifies the name of the administrator account. <br><br> **Windows-only restriction:** Cannot end in "." <br><br> **Disallowed values:** "administrator", "admin", "user", "user1", "test", "user2", "test1", "user3", "admin1", "1", "123", "a", "actuser", "adm", "admin2", "aspnet", "backup", "console", "david", "guest", "john", "owner", "root", "server", "sql", "support", "support_388945a0", "sys", "test2", "test3", "user4", "user5". <br><br> **Minimum-length (Linux):** 1  character <br><br> **Max-length (Linux):** 64 characters <br><br> **Max-length (Windows):** 20 characters  <br><br><li> For root access to the Linux VM, see [Using root privileges on Linux virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-use-root-privileges?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json)<br><li> For a list of built-in system users on Linux that should not be used in this field, see [Selecting User Names for Linux on Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-usernames?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json)
          * `computer_name_prefix` (`str`) - Specifies the computer name prefix for all of the virtual machines in the scale set. Computer name prefixes must be 1 to 15 characters long.
          * `custom_data` (`str`) - A base-64 encoded string of custom data.
          * `linux_configuration` (`dict`) - The Linux Configuration of the OS profile.
            * `disable_password_authentication` (`bool`) - Specifies whether password authentication should be disabled.
            * `ssh` (`dict`) - Specifies the ssh key configuration for a Linux OS.
              * `public_keys` (`list`) - The list of SSH public keys used to authenticate with linux based VMs.
                * `key_data` (`str`) - SSH public key certificate used to authenticate with the VM through ssh. The key needs to be at least 2048-bit and in ssh-rsa format. <br><br> For creating ssh keys, see [Create SSH keys on Linux and Mac for Linux VMs in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-mac-create-ssh-keys?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json).
                * `path` (`str`) - Specifies the full path on the created VM where ssh public key is stored. If the file already exists, the specified key is appended to the file. Example: /home/user/.ssh/authorized_keys

          * `secrets` (`list`) - The List of certificates for addition to the VM.
            * `source_vault` (`dict`) - The relative URL of the Key Vault containing all of the certificates in VaultCertificates.
            * `vault_certificates` (`list`) - The list of key vault references in SourceVault which contain certificates.
              * `certificate_store` (`str`) - For Windows VMs, specifies the certificate store on the Virtual Machine to which the certificate should be added. The specified certificate store is implicitly in the LocalMachine account. <br><br>For Linux VMs, the certificate file is placed under the /var/lib/waagent directory, with the file name &lt;UppercaseThumbprint&gt;.crt for the X509 certificate file and &lt;UppercaseThumbprint&gt;.prv for private key. Both of these files are .pem formatted.
              * `certificate_url` (`str`) - This is the URL of a certificate that has been uploaded to Key Vault as a secret. For adding a secret to the Key Vault, see [Add a key or secret to the key vault](https://docs.microsoft.com/azure/key-vault/key-vault-get-started/#add). In this case, your certificate needs to be It is the Base64 encoding of the following JSON Object which is encoded in UTF-8: <br><br> {<br>  "data":"<Base64-encoded-certificate>",<br>  "dataType":"pfx",<br>  "password":"<pfx-file-password>"<br>}

          * `windows_configuration` (`dict`) - The Windows Configuration of the OS profile.
            * `additional_unattend_content` (`list`) - Specifies additional base-64 encoded XML formatted information that can be included in the Unattend.xml file, which is used by Windows Setup.
              * `component_name` (`str`) - The component name. Currently, the only allowable value is Microsoft-Windows-Shell-Setup.
              * `content` (`str`) - Specifies the XML formatted content that is added to the unattend.xml file for the specified path and component. The XML must be less than 4KB and must include the root element for the setting or feature that is being inserted.
              * `pass_name` (`str`) - The pass name. Currently, the only allowable value is OobeSystem.
              * `setting_name` (`str`) - Specifies the name of the setting to which the content applies. Possible values are: FirstLogonCommands and AutoLogon.

            * `enable_automatic_updates` (`bool`) - Indicates whether virtual machine is enabled for automatic updates.
            * `provision_vm_agent` (`bool`) - Indicates whether virtual machine agent should be provisioned on the virtual machine. <br><br> When this property is not specified in the request body, default behavior is to set it to true.  This will ensure that VM Agent is installed on the VM so that extensions can be added to the VM later.
            * `time_zone` (`str`) - Specifies the time zone of the virtual machine. e.g. "Pacific Standard Time"
            * `win_rm` (`dict`) - Specifies the Windows Remote Management listeners. This enables remote Windows PowerShell.
              * `listeners` (`list`) - The list of Windows Remote Management listeners
                * `certificate_url` (`str`) - This is the URL of a certificate that has been uploaded to Key Vault as a secret. For adding a secret to the Key Vault, see [Add a key or secret to the key vault](https://docs.microsoft.com/azure/key-vault/key-vault-get-started/#add). In this case, your certificate needs to be It is the Base64 encoding of the following JSON Object which is encoded in UTF-8: <br><br> {<br>  "data":"<Base64-encoded-certificate>",<br>  "dataType":"pfx",<br>  "password":"<pfx-file-password>"<br>}
                * `protocol` (`str`) - Specifies the protocol of listener. <br><br> Possible values are: <br>**http** <br><br> **https**

        * `storage_profile` (`dict`) - The virtual machine scale set storage profile.
          * `image_reference` (`dict`) - The image reference.
            * `offer` (`str`) - Specifies the offer of the platform image or marketplace image used to create the virtual machine.
            * `publisher` (`str`) - The image publisher.
            * `sku` (`str`) - The image SKU.
            * `version` (`str`) - Specifies the version of the platform image or marketplace image used to create the virtual machine. The allowed formats are Major.Minor.Build or 'latest'. Major, Minor, and Build are decimal numbers. Specify 'latest' to use the latest version of an image available at deploy time. Even if you use 'latest', the VM image will not automatically update after deploy time even if a new version becomes available.

          * `os_disk` (`dict`) - The OS disk.
            * `caching` (`str`) - Specifies the caching requirements. <br><br> Possible values are: <br><br> **None** <br><br> **ReadOnly** <br><br> **ReadWrite** <br><br> Default: **None for Standard storage. ReadOnly for Premium storage**
            * `create_option` (`str`) - Specifies how the virtual machines in the scale set should be created.<br><br> The only allowed value is: **FromImage** \u2013 This value is used when you are using an image to create the virtual machine. If you are using a platform image, you also use the imageReference element described above. If you are using a marketplace image, you  also use the plan element previously described.
            * `image` (`dict`) - The Source User Image VirtualHardDisk. This VirtualHardDisk will be copied before using it to attach to the Virtual Machine. If SourceImage is provided, the destination VirtualHardDisk should not exist.
              * `uri` (`str`) - Specifies the virtual hard disk's uri.

            * `name` (`str`) - The disk name.
            * `os_type` (`str`) - This property allows you to specify the type of the OS that is included in the disk if creating a VM from user-image or a specialized VHD. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**
            * `vhd_containers` (`list`) - The list of virtual hard disk container uris.
    """
    sku: pulumi.Output[dict]
    """
    The virtual machine scale set sku.
      * `capacity` (`float`) - Specifies the number of virtual machines in the scale set.
      * `name` (`str`) - The sku name.
      * `tier` (`str`) - Specifies the tier of virtual machines in a scale set.<br /><br /> Possible Values:<br /><br /> **Standard**<br /><br /> **Basic**
    """
    tags: pulumi.Output[dict]
    """
    Resource tags
    """
    type: pulumi.Output[str]
    """
    Resource type
    """
    def __init__(__self__, resource_name, opts=None, location=None, name=None, over_provision=None, provisioning_state=None, resource_group_name=None, sku=None, tags=None, upgrade_policy=None, virtual_machine_profile=None, __props__=None, __name__=None, __opts__=None):
        """
        Describes a Virtual Machine Scale Set.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] name: The name of the VM scale set to create or update.
        :param pulumi.Input[bool] over_provision: Specifies whether the Virtual Machine Scale Set should be overprovisioned.
        :param pulumi.Input[str] provisioning_state: The provisioning state, which only appears in the response.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[dict] sku: The virtual machine scale set sku.
        :param pulumi.Input[dict] tags: Resource tags
        :param pulumi.Input[dict] upgrade_policy: The upgrade policy.
        :param pulumi.Input[dict] virtual_machine_profile: The virtual machine profile.

        The **sku** object supports the following:

          * `capacity` (`pulumi.Input[float]`) - Specifies the number of virtual machines in the scale set.
          * `name` (`pulumi.Input[str]`) - The sku name.
          * `tier` (`pulumi.Input[str]`) - Specifies the tier of virtual machines in a scale set.<br /><br /> Possible Values:<br /><br /> **Standard**<br /><br /> **Basic**

        The **upgrade_policy** object supports the following:

          * `mode` (`pulumi.Input[str]`) - Specifies the mode of an upgrade to virtual machines in the scale set.<br /><br /> Possible values are:<br /><br /> **Manual** - You  control the application of updates to virtual machines in the scale set. You do this by using the manualUpgrade action.<br /><br /> **Automatic** - All virtual machines in the scale set are  automatically updated at the same time.

        The **virtual_machine_profile** object supports the following:

          * `extension_profile` (`pulumi.Input[dict]`) - The virtual machine scale set extension profile.
            * `extensions` (`pulumi.Input[list]`) - The virtual machine scale set child extension resources.
              * `auto_upgrade_minor_version` (`pulumi.Input[bool]`) - Whether the extension handler should be automatically upgraded across minor versions.
              * `id` (`pulumi.Input[str]`) - Resource Id
              * `name` (`pulumi.Input[str]`) - The name of the extension.
              * `protected_settings` (`pulumi.Input[dict]`) - Json formatted protected settings for the extension.
              * `publisher` (`pulumi.Input[str]`) - The name of the extension handler publisher.
              * `settings` (`pulumi.Input[dict]`) - Json formatted public settings for the extension.
              * `type` (`pulumi.Input[str]`) - The type of the extension handler.
              * `type_handler_version` (`pulumi.Input[str]`) - The type version of the extension handler.

          * `network_profile` (`pulumi.Input[dict]`) - The virtual machine scale set network profile.
            * `network_interface_configurations` (`pulumi.Input[list]`) - The list of network configurations.
              * `id` (`pulumi.Input[str]`) - Resource Id
              * `ip_configurations` (`pulumi.Input[list]`) - The virtual machine scale set IP Configuration.
                * `id` (`pulumi.Input[str]`) - Resource Id
                * `load_balancer_backend_address_pools` (`pulumi.Input[list]`) - The load balancer backend address pools.
                  * `id` (`pulumi.Input[str]`) - Resource Id

                * `load_balancer_inbound_nat_pools` (`pulumi.Input[list]`) - The load balancer inbound nat pools.
                * `name` (`pulumi.Input[str]`) - The IP configuration name.
                * `subnet` (`pulumi.Input[dict]`) - The subnet.
                  * `id` (`pulumi.Input[str]`) - The ARM resource id in the form of /subscriptions/{SubscriptionId}/resourceGroups/{ResourceGroupName}/...

              * `name` (`pulumi.Input[str]`) - The network configuration name.
              * `primary` (`pulumi.Input[bool]`) - Whether this is a primary NIC on a virtual machine.

          * `os_profile` (`pulumi.Input[dict]`) - The virtual machine scale set OS profile.
            * `admin_password` (`pulumi.Input[str]`) - Specifies the password of the administrator account. <br><br> **Minimum-length (Windows):** 8 characters <br><br> **Minimum-length (Linux):** 6 characters <br><br> **Max-length (Windows):** 123 characters <br><br> **Max-length (Linux):** 72 characters <br><br> **Complexity requirements:** 3 out of 4 conditions below need to be fulfilled <br> Has lower characters <br>Has upper characters <br> Has a digit <br> Has a special character (Regex match [\W_]) <br><br> **Disallowed values:** "abc@123", "P@$$w0rd", "P@ssw0rd", "P@ssword123", "Pa$$word", "pass@word1", "Password!", "Password1", "Password22", "iloveyou!" <br><br> For resetting the password, see [How to reset the Remote Desktop service or its login password in a Windows VM](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-reset-rdp?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> For resetting root password, see [Manage users, SSH, and check or repair disks on Azure Linux VMs using the VMAccess Extension](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-using-vmaccess-extension?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json#reset-root-password)
            * `admin_username` (`pulumi.Input[str]`) - Specifies the name of the administrator account. <br><br> **Windows-only restriction:** Cannot end in "." <br><br> **Disallowed values:** "administrator", "admin", "user", "user1", "test", "user2", "test1", "user3", "admin1", "1", "123", "a", "actuser", "adm", "admin2", "aspnet", "backup", "console", "david", "guest", "john", "owner", "root", "server", "sql", "support", "support_388945a0", "sys", "test2", "test3", "user4", "user5". <br><br> **Minimum-length (Linux):** 1  character <br><br> **Max-length (Linux):** 64 characters <br><br> **Max-length (Windows):** 20 characters  <br><br><li> For root access to the Linux VM, see [Using root privileges on Linux virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-use-root-privileges?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json)<br><li> For a list of built-in system users on Linux that should not be used in this field, see [Selecting User Names for Linux on Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-usernames?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json)
            * `computer_name_prefix` (`pulumi.Input[str]`) - Specifies the computer name prefix for all of the virtual machines in the scale set. Computer name prefixes must be 1 to 15 characters long.
            * `custom_data` (`pulumi.Input[str]`) - A base-64 encoded string of custom data.
            * `linux_configuration` (`pulumi.Input[dict]`) - The Linux Configuration of the OS profile.
              * `disable_password_authentication` (`pulumi.Input[bool]`) - Specifies whether password authentication should be disabled.
              * `ssh` (`pulumi.Input[dict]`) - Specifies the ssh key configuration for a Linux OS.
                * `public_keys` (`pulumi.Input[list]`) - The list of SSH public keys used to authenticate with linux based VMs.
                  * `key_data` (`pulumi.Input[str]`) - SSH public key certificate used to authenticate with the VM through ssh. The key needs to be at least 2048-bit and in ssh-rsa format. <br><br> For creating ssh keys, see [Create SSH keys on Linux and Mac for Linux VMs in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-mac-create-ssh-keys?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json).
                  * `path` (`pulumi.Input[str]`) - Specifies the full path on the created VM where ssh public key is stored. If the file already exists, the specified key is appended to the file. Example: /home/user/.ssh/authorized_keys

            * `secrets` (`pulumi.Input[list]`) - The List of certificates for addition to the VM.
              * `source_vault` (`pulumi.Input[dict]`) - The relative URL of the Key Vault containing all of the certificates in VaultCertificates.
              * `vault_certificates` (`pulumi.Input[list]`) - The list of key vault references in SourceVault which contain certificates.
                * `certificate_store` (`pulumi.Input[str]`) - For Windows VMs, specifies the certificate store on the Virtual Machine to which the certificate should be added. The specified certificate store is implicitly in the LocalMachine account. <br><br>For Linux VMs, the certificate file is placed under the /var/lib/waagent directory, with the file name &lt;UppercaseThumbprint&gt;.crt for the X509 certificate file and &lt;UppercaseThumbprint&gt;.prv for private key. Both of these files are .pem formatted.
                * `certificate_url` (`pulumi.Input[str]`) - This is the URL of a certificate that has been uploaded to Key Vault as a secret. For adding a secret to the Key Vault, see [Add a key or secret to the key vault](https://docs.microsoft.com/azure/key-vault/key-vault-get-started/#add). In this case, your certificate needs to be It is the Base64 encoding of the following JSON Object which is encoded in UTF-8: <br><br> {<br>  "data":"<Base64-encoded-certificate>",<br>  "dataType":"pfx",<br>  "password":"<pfx-file-password>"<br>}

            * `windows_configuration` (`pulumi.Input[dict]`) - The Windows Configuration of the OS profile.
              * `additional_unattend_content` (`pulumi.Input[list]`) - Specifies additional base-64 encoded XML formatted information that can be included in the Unattend.xml file, which is used by Windows Setup.
                * `component_name` (`pulumi.Input[str]`) - The component name. Currently, the only allowable value is Microsoft-Windows-Shell-Setup.
                * `content` (`pulumi.Input[str]`) - Specifies the XML formatted content that is added to the unattend.xml file for the specified path and component. The XML must be less than 4KB and must include the root element for the setting or feature that is being inserted.
                * `pass_name` (`pulumi.Input[str]`) - The pass name. Currently, the only allowable value is OobeSystem.
                * `setting_name` (`pulumi.Input[str]`) - Specifies the name of the setting to which the content applies. Possible values are: FirstLogonCommands and AutoLogon.

              * `enable_automatic_updates` (`pulumi.Input[bool]`) - Indicates whether virtual machine is enabled for automatic updates.
              * `provision_vm_agent` (`pulumi.Input[bool]`) - Indicates whether virtual machine agent should be provisioned on the virtual machine. <br><br> When this property is not specified in the request body, default behavior is to set it to true.  This will ensure that VM Agent is installed on the VM so that extensions can be added to the VM later.
              * `time_zone` (`pulumi.Input[str]`) - Specifies the time zone of the virtual machine. e.g. "Pacific Standard Time"
              * `win_rm` (`pulumi.Input[dict]`) - Specifies the Windows Remote Management listeners. This enables remote Windows PowerShell.
                * `listeners` (`pulumi.Input[list]`) - The list of Windows Remote Management listeners
                  * `certificate_url` (`pulumi.Input[str]`) - This is the URL of a certificate that has been uploaded to Key Vault as a secret. For adding a secret to the Key Vault, see [Add a key or secret to the key vault](https://docs.microsoft.com/azure/key-vault/key-vault-get-started/#add). In this case, your certificate needs to be It is the Base64 encoding of the following JSON Object which is encoded in UTF-8: <br><br> {<br>  "data":"<Base64-encoded-certificate>",<br>  "dataType":"pfx",<br>  "password":"<pfx-file-password>"<br>}
                  * `protocol` (`pulumi.Input[str]`) - Specifies the protocol of listener. <br><br> Possible values are: <br>**http** <br><br> **https**

          * `storage_profile` (`pulumi.Input[dict]`) - The virtual machine scale set storage profile.
            * `image_reference` (`pulumi.Input[dict]`) - The image reference.
              * `offer` (`pulumi.Input[str]`) - Specifies the offer of the platform image or marketplace image used to create the virtual machine.
              * `publisher` (`pulumi.Input[str]`) - The image publisher.
              * `sku` (`pulumi.Input[str]`) - The image SKU.
              * `version` (`pulumi.Input[str]`) - Specifies the version of the platform image or marketplace image used to create the virtual machine. The allowed formats are Major.Minor.Build or 'latest'. Major, Minor, and Build are decimal numbers. Specify 'latest' to use the latest version of an image available at deploy time. Even if you use 'latest', the VM image will not automatically update after deploy time even if a new version becomes available.

            * `os_disk` (`pulumi.Input[dict]`) - The OS disk.
              * `caching` (`pulumi.Input[str]`) - Specifies the caching requirements. <br><br> Possible values are: <br><br> **None** <br><br> **ReadOnly** <br><br> **ReadWrite** <br><br> Default: **None for Standard storage. ReadOnly for Premium storage**
              * `create_option` (`pulumi.Input[str]`) - Specifies how the virtual machines in the scale set should be created.<br><br> The only allowed value is: **FromImage** \u2013 This value is used when you are using an image to create the virtual machine. If you are using a platform image, you also use the imageReference element described above. If you are using a marketplace image, you  also use the plan element previously described.
              * `image` (`pulumi.Input[dict]`) - The Source User Image VirtualHardDisk. This VirtualHardDisk will be copied before using it to attach to the Virtual Machine. If SourceImage is provided, the destination VirtualHardDisk should not exist.
                * `uri` (`pulumi.Input[str]`) - Specifies the virtual hard disk's uri.

              * `name` (`pulumi.Input[str]`) - The disk name.
              * `os_type` (`pulumi.Input[str]`) - This property allows you to specify the type of the OS that is included in the disk if creating a VM from user-image or a specialized VHD. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**
              * `vhd_containers` (`pulumi.Input[list]`) - The list of virtual hard disk container uris.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['over_provision'] = over_provision
            __props__['provisioning_state'] = provisioning_state
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['upgrade_policy'] = upgrade_policy
            __props__['virtual_machine_profile'] = virtual_machine_profile
            __props__['properties'] = None
            __props__['type'] = None
        super(VirtualMachineScaleSet, __self__).__init__(
            'azurerm:compute/v20150615:VirtualMachineScaleSet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing VirtualMachineScaleSet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return VirtualMachineScaleSet(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
