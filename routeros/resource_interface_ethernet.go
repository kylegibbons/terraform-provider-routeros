package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ResourceInterfaceEthernet https://help.mikrotik.com/docs/display/ROS/Ethernet
func ResourceInterfaceEthernet() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/ethernet"),
		MetaId:           PropId(Name),

		"advertise": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Advertised speed and duplex modes for Ethernet interfaces over twisted pair, only applies when " +
				"auto-negotiation is enabled. Advertising higher speeds than the actual interface supported speed will " +
				"have no effect, multiple options are allowed.",
			ValidateFunc: validation.StringInSlice([]string{
				"10M-full", "10M-half", "100M-full", "100M-half", "1000M-full",
				"1000M-half", "2500M-full", "5000M-full", "10000M-full"}, false),
		},
		KeyArp: PropArpRw,
		"auto_negotiation": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: `When enabled, the interface "advertises" its maximum capabilities to achieve the best connection possible.
	* Note1: Auto-negotiation should not be disabled on one end only, otherwise Ethernet Interfaces may not work properly.
	* Note2: Gigabit Ethernet and NBASE-T Ethernet links cannot work with auto-negotiation disabled.`,
		},
		"bandwidth": {
			Type:     schema.TypeString,
			Optional: true,
			Description: `Sets max rx/tx bandwidth in kbps that will be handled by an interface. TX limit is supported on " + 
				"all Atheros switch-chip ports. RX limit is supported only on Atheros8327/QCA8337 switch-chip ports.`,
		},
		"cable_settings": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  `Changes the cable length setting (only applicable to NS DP83815/6 cards).`,
			ValidateFunc: validation.StringInSlice([]string{"default", "short", "standard"}, false),
		},
		"combo_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "When auto mode is selected, the port that was first connected will establish the link. In case " +
				"this link fails, the other port will try to establish a new link. If both ports are connected at the " +
				"same time (e.g. after reboot), the priority will be the SFP/SFP+ port. When sfp mode is selected, the" +
				" interface will only work through SFP/SFP+ cage. When copper mode is selected, the interface will only " +
				"work through RJ45 Ethernet port.",
			ValidateFunc: validation.StringInSlice([]string{"auto", "copper", "sfp"}, false),
		},
		KeyComment: PropCommentRw,
		"diable_running_check": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Disable running check. If this value is set to 'no', the router automatically detects whether " +
				"the NIC is connected with a device in the network or not. Default value is 'yes' because older NICs do not " +
				"support it. (only applicable to x86)",
		},
		"tx_flow_control": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "When set to on, the port will generate pause frames to the upstream device to temporarily stop " +
				"the packet transmission. Pause frames are only generated when some routers output interface is congested " +
				"and packets cannot be transmitted anymore. auto is the same as on except when auto-negotiation=yes flow " +
				"control status is resolved by taking into account what other end advertises.",
			ValidateFunc: validation.StringInSlice([]string{"on", "off", "auto"}, false),
		},
		"rx_flow_control": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "When set to on, the port will process received pause frames and suspend transmission if required. " +
				"auto is the same as on except when auto-negotiation=yes flow control status is resolved by taking into " +
				"account what other end advertises.",
			ValidateFunc: validation.StringInSlice([]string{"on", "off", "auto"}, false),
		},
		"full_duplex": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Defines whether the transmission of data appears in two directions simultaneously, only applies " +
				"when auto-negotiation is disabled.",
		},
		KeyL2Mtu: {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Layer2 Maximum transmission unit. Read more: https://wiki.mikrotik.com/wiki/Maximum_Transmission_Unit_on_RouterBoards",
			ValidateFunc: validation.IntBetween(0, 65535),
		},
		"mac_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Media Access Control number of an interface.",
			ValidateFunc: validation.IsMACAddress,
		},
		"master_port": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Outdated property, more details about this property can be found in the Master-port page.",
		},
		"mdix_enable": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "Whether the MDI/X auto cross over cable correction feature is enabled for the port (Hardware " +
				"specific, e.g. ether1 on RB500 can be set to yes/no. Fixed to 'yes' on other hardware.)",
		},
		"mtu": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Layer3 Maximum transmission unit.",
			ValidateFunc: validation.IntBetween(0, 65535),
		},
		KeyName: PropNameForceNewRw,
		"poe_out": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Poe Out settings. Read more: https://wiki.mikrotik.com/wiki/Manual:PoE-Out",
			ValidateFunc: validation.StringInSlice([]string{"auto-on", "forced-on", "off"}, false),
		},
		"poe_priority": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Poe Out settings. Read more: https://wiki.mikrotik.com/wiki/Manual:PoE-Out",
			ValidateFunc: validation.IntBetween(0, 99),
		},
		"sfp_shutdown-temperature": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "The temperature in Celsius at which the interface will be temporarily turned off due to too high " +
				"detected SFP module temperature (introduced v6.48). The default value for SFP/SFP+/SFP28 interfaces is 95, " +
				"and for QSFP+/QSFP28 interfaces 80 (introduced v7.6).",
		},
		"speed": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Sets interface data transmission speed which takes effect only when auto-negotiation is disabled.",
			ValidateFunc: validation.StringInSlice([]string{"10Mbps", "10Gbps", "100Mbps", "1Gbps"}, false),
		},

		// READ ONLY PROPERTIES
		"orig_mac_address": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Original Media Access Control number of an interface.",
		},
		KeyRunning: PropRunningRo,
		"slave": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether interface is configured as a slave of another interface (for example Bonding)",
		},
		"switch": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "ID to which switch chip interface belongs to.",
		},

		// Some properties are not implemented, see: https://help.mikrotik.com/docs/display/ROS/Ethernet
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
