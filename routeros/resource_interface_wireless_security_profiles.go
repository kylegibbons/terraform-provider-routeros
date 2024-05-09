package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ResourceInterfaceWirelessSecurityProfiles https://help.mikrotik.com/docs/display/ROS/Wireless+Interface#WirelessInterface-SecurityProfiles
func ResourceInterfaceWirelessSecurityProfiles() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/wireless/security-profiles"),
		MetaId:           PropId(Name),

		// BASIC PROPERTIES
		"mode": {
			Type:     schema.TypeString,
			Required: true,
			Description: `Encryption mode for the security profile.
	* none - Encryption is not used. Encrypted frames are not accepted.
	* static-keys-required - WEP mode. Do not accept and do not send unencrypted frames. Station in static-keys-required mode will not connect to an Access Point in static-keys-optional mode.
	* static-keys-optional - WEP mode. Support encryption and decryption, but allow also to receive and send unencrypted frames. Device will send unencrypted frames if encryption algorithm is specified as none. Station in static-keys-optional mode will not connect to an Access Point in static-keys-required mode. See also: static-sta-private-algo, static-transmit-key.
	* dynamic-keys - WPA mode.`,
			ValidateFunc: validation.StringInSlice([]string{
				"none",
				"static-keys-optional",
				"static-keys-required",
				"dynamic-keys",
			}, false),
		},
		KeyName: PropNameForceNewRw,

		"default": {
			Type:     schema.TypeString,
			Computed: true,
		},

		// WPA PSK PROPERTIES

		"authentication_types": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"wpa-psk", "wpa2-psk", "wpa-eap", "wpa2-eap"}, false),
			},
			Required: true,
			MinItems: 1,
			Description: "Set of supported authentication types, multiple values can be selected. Access Point " +
				"will advertise supported authentication types, and client will connect to Access Point only if " +
				"it supports any of the advertised authentication types.",
		},
		"disable_pmkid": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: `Whether to include PMKID into the EAPOL frame sent out by the Access Point. Disabling PMKID can cause compatibility issues with devices that use the PMKID to connect to an Access Point.
	* yes - removes PMKID from EAPOL frames (improves security, reduces compatibility).
	* no - includes PMKID into EAPOL frames (reduces security, improves compatibility).
	This property only has effect on Access Points.`,
		},
		"unicast_ciphers": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"tkip", "aes-ccm"}, false),
			},
			Optional: true,
			Computed: true,
			Description: "Access Point advertises that it supports specified ciphers, multiple values can be selected. " +
				"Client attempts connection only to Access Points that supports at least one of the specified ciphers. " +
				"One of the ciphers will be used to encrypt unicast frames that are sent between Access Point and Station.",
		},
		"group_ciphers": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"tkip", "aes-ccm"}, false),
			},
			Optional: true,
			Computed: true,
			Description: `Access Point advertises one of these ciphers, multiple values can be selected. Access Point uses it to encrypt all broadcast and multicast frames. Client attempts connection only to Access Points that use one of the specified group ciphers.
	* tkip - Temporal Key Integrity Protocol - encryption protocol, compatible with legacy WEP equipment, but enhanced to correct some of the WEP flaws.
	* aes-ccm - more secure WPA encryption protocol, based on the reliable AES (Advanced Encryption Standard). Networks free of WEP legacy should use only this cipher.
			`,
		},
		"group_key_update": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "5m",
			Description: "Controls how often Access Point updates the group key. This key is used to encrypt all broadcast " +
				"and multicast frames. property only has effect for Access Points.",
		},
		"wpa_pre_shared_key": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "WPA pre-shared key mode requires all devices in a BSS to have common secret key. Value of this key " +
				"can be an arbitrary text. Commonly referred to as the network password for WPA mode. property only has effect " +
				"when wpa-psk is added to authentication-types.",
		},
		"wpa2_pre_shared_key": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "WPA2 pre-shared key mode requires all devices in a BSS to have common secret key. Value of this key " +
				"can be an arbitrary text. Commonly referred to as the network password for WPA2 mode. property only has effect " +
				"when wpa2-psk is added to authentication-types. Some properties are not implemented, ",
		},

		// WPA EAP PROPERTIES

		"eap_methods": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"eap-tls", "eap-ttls-mschapv2", "passthrough", "peap"}, false),
			},
			Computed: true,
			Optional: true,
			Description: `Allowed types of authentication methods, multiple values can be selected. This property only has effect on Access Points.
	* eap-tls - Use built-in EAP TLS authentication. Both client and server certificates are supported. See description of tls-mode and tls-certificate properties.
	* eap-ttls-mschapv2 - Use EAP-TTLS with MS-CHAPv2 authentication.
	* passthrough - Access Point will relay authentication process to the RADIUS server.
	* peap - Use Protected EAP authentication.`,
		},
		"supplicant_identity": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: `EAP identity that is sent by client at the beginning of EAP authentication. This value is " + 
				"used as a value for User-Name attribute in RADIUS messages sent by RADIUS EAP accounting and RADIUS " + 
				"EAP pass-through authentication.`,
		},
		"mschapv2_username": {
			Type:     schema.TypeString,
			Optional: true,
			Description: `Username to use for authentication when eap-ttls-mschapv2 authentication method is being used. " + 
				"This property only has effect on Stations.`,
		},
		"mschapv2_password": {
			Type:     schema.TypeString,
			Optional: true,
			Description: `Password to use for authentication when eap-ttls-mschapv2 authentication method is being used. " + 
				"This property only has effect on Stations.`,
		},
		"tls_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: `This property has effect only when eap-methods contains eap-tls.
			* verify-certificate - Require remote device to have valid certificate. Check that it is signed by known certificate authority. No additional identity verification is done. Certificate may include information about time period during which it is valid. If router has incorrect time and date, it may reject valid certificate because router's clock is outside that period. See also the Certificates configuration.
			* dont-verify-certificate - Do not check certificate of the remote device. Access Point will not require client to provide certificate.
			* no-certificates - Do not use certificates. TLS session is established using 2048 bit anonymous Diffie-Hellman key exchange.
			* verify-certificate-with-crl - Same as verify-certificate but also checks if the certificate is valid by checking the Certificate Revocation List.`,
			ValidateFunc: validation.StringInSlice([]string{"verify-certificate", "dont-verify-certificate", "no-certificates", "verify-certificate-with-crl"}, false),
		},
		"tls_certificate": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: `Access Point always needs a certificate when configured when tls-mode is set to verify-certificate, or is set to " + 
				"dont-verify-certificate. Client needs a certificate only if Access Point is configured with tls-mode set to " + 
				"verify-certificate. In this case client needs a valid certificate that is signed by a CA known to the Access Point. " + 
				"This property only has effect when tls-mode is not set to no-certificates and eap-methods contains eap-tls.`,
		},
		"management_protection": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"management_protection_key": {
			Type:     schema.TypeString,
			Optional: true,
		},

		// RADIUS PROPERTIES
		"radius_mac_authentication": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: `This property affects the way how Access Point processes clients that are not found in the Access List.
			* no - allow or reject client authentication based on the value of default-authentication property of the Wireless interface.
			* yes - Query RADIUS server using MAC address of client as user name. With this setting the value of default-authentication has no effect.`,
		},
		"radius_mac_accounting": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"radius_eap_accounting": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"radius_called_format": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"interim_update": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "When RADIUS accounting is used, Access Point periodically sends accounting information updates to the " +
				"RADIUS server. This property specifies default update interval that can be overridden by the RADIUS server using " +
				"Acct-Interim-Interval attribute.",
		},
		"radius_mac_format": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "Controls how MAC address of the client is encoded by Access Point in the User-Name attribute of the MAC " +
				"authentication and MAC accounting RADIUS requests.",
			ValidateFunc: validation.StringInSlice([]string{
				"XX:XX:XX:XX:XX:XX", "XXXX:XXXX:XXXX", "XXXXXX:XXXXXX",
				"XX-XX-XX-XX-XX-XX", "XXXXXX-XXXXXX", "XXXXXXXXXXXX",
				"XX XX XX XX XX XX"}, false),
		},
		"radius_mac_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: `By default Access Point uses an empty password, when sending Access-Request during MAC authentication. " +
				"When this property is set to as-username-and-password, Access Point will use the same value for User-Password " +
				"attribute as for the User-Name attribute.`,
			ValidateFunc: validation.StringInSlice([]string{"as-username", "as-username-and-password"}, false),
		},
		"radius_mac_caching": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: `If this value is set to time interval, the Access Point will cache RADIUS MAC authentication responses " +
				"for specified time, and will not contact RADIUS server if matching cache entry already exists. Value disabled " +
				"will disable cache, Access Point will always contact RADIUS server.`,
		},

		// WEP PROPERTIES
		"static_key_0": {
			Type:     schema.TypeString,
			Optional: true,
			Description: `Hexadecimal representation of the key. Length of key must be appropriate for selected algorithm. " +
				"See the Statically configured WEP keys section.`,
		},
		"static_key_1": {
			Type:     schema.TypeString,
			Optional: true,
			Description: `Hexadecimal representation of the key. Length of key must be appropriate for selected algorithm. " +
				"See the Statically configured WEP keys section.`,
		},
		"static_key_2": {
			Type:     schema.TypeString,
			Optional: true,
			Description: `Hexadecimal representation of the key. Length of key must be appropriate for selected algorithm. " +
				"See the Statically configured WEP keys section.`,
		},
		"static_key_3": {
			Type:     schema.TypeString,
			Optional: true,
			Description: `Hexadecimal representation of the key. Length of key must be appropriate for selected algorithm. " +
				"See the Statically configured WEP keys section.`,
		},
		"static_algo_0": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			Description:  "Encryption algorithm to use with the corresponding key.",
			ValidateFunc: validation.StringInSlice([]string{"none", "40bit-wep", "104bit-wep", "tkip", "aes-ccm"}, false),
		},
		"static_algo_1": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			Description:  "Encryption algorithm to use with the corresponding key.",
			ValidateFunc: validation.StringInSlice([]string{"none", "40bit-wep", "104bit-wep", "tkip", "aes-ccm"}, false),
		},
		"static_algo_2": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			Description:  "Encryption algorithm to use with the corresponding key.",
			ValidateFunc: validation.StringInSlice([]string{"none", "40bit-wep", "104bit-wep", "tkip", "aes-ccm"}, false),
		},
		"static_algo_3": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			Description:  "Encryption algorithm to use with the corresponding key.",
			ValidateFunc: validation.StringInSlice([]string{"none", "40bit-wep", "104bit-wep", "tkip", "aes-ccm"}, false),
		},
		"static_transmit_key": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "Access Point will use the specified key to encrypt frames for clients that do not use private key. " +
				"Access Point will also use this key to encrypt broadcast and multicast frames. Client will use the specified " +
				"key to encrypt frames if static-sta-private-algo is set to none. If corresponding static-algo-N property has " +
				"value set to none, then frame will be sent unencrypted (when mode is set to static-keys-optional) or will not " +
				"be sent at all (when mode is set to static-keys-required).",
			ValidateFunc: validation.StringInSlice([]string{"key-0", "key-1", "key-2", "key-3"}, false),
		},
		"static_sta_private_key": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Length of key must be appropriate for selected algorithm, see the Statically configured WEP keys section. " +
				"This property is used only on Stations. Access Point uses corresponding key either from private-key property, " +
				"or from Mikrotik-Wireless-Enc-Key attribute.",
		},
		"static_sta_private_algo": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "Encryption algorithm to use with station private key. Value none disables use of the private key. " +
				"This property is only used on Stations. Access Point has to get corresponding value either from private-algo " +
				"property, or from Mikrotik-Wireless-Enc-Algo attribute. Station private key replaces key 0 for unicast frames. " +
				"Station will not use private key to decrypt broadcast frames.",
			ValidateFunc: validation.StringInSlice([]string{"none", "40bit-wep", "104bit-wep", "tkip", "aes-ccm"}, false),
		},

		// Some properties are not implemented, see: https://help.mikrotik.com/docs/display/ROS/Wireless+Interface#WirelessInterface-SecurityProfiles
	}

	return &schema.Resource{
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
