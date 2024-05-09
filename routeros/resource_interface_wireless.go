package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ResourceInterfaceWireless https://help.mikrotik.com/docs/display/ROS/Wireless+Interface
func ResourceInterfaceWireless() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/interface/wireless"),
		MetaId:           PropId(Name),

		KeyActualMtu: PropActualMtuRo,
		"adaptive_noise_immunity": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "This property is only effective for cards based on Atheros chipset.",
			ValidateFunc: validation.StringInSlice([]string{"ap-and-client-mode", "client-mode", "none"}, false),
		},
		"allow_sharedkey": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Allow WEP Shared Key clients to connect. Note that no authentication is done for these " +
				"clients (WEP Shared keys are not compared to anything) - they are just accepted at once (if access " +
				"list allows that)",
			ValidateFunc: validation.StringInSlice([]string{"yes", "no"}, false),
		},
		"ampdu_priorities": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 7),
			},
			Description: "Frame priorities for which AMPDU sending (aggregating frames and sending using block " +
				"acknowledgment) should get negotiated and used. Using AMPDUs will increase throughput, but may " +
				"increase latency, therefore, may not be desirable for real-time traffic (voice, video). Due to " +
				"this, by default AMPDUs are enabled only for best-effort traffic.",
		},
		"amsdu_limit": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Max AMSDU that device is allowed to prepare when negotiated. AMSDU aggregation may " +
				"significantly increase throughput especially for small frames, but may increase latency in " +
				"case of packet loss due to retransmission of aggregated frame. Sending and receiving AMSDUs " +
				"will also increase CPU usage.",
			ValidateFunc: validation.IntBetween(0, 8192),
		},
		"amsdu_threshold": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Max frame size to allow including in AMSDU.",
			ValidateFunc: validation.IntBetween(0, 8192),
		},
		"antenna_gain": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Antenna gain in dBi, used to calculate maximum transmit power according to country regulations.",
			// ValidateFunc: validation.FloatBetween(0, 4294967295),
		},
		"antenna_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Description: `Select antenna to use for transmitting and for receiving
				ant-a - use only 'a' antenna
				ant-b - use only 'b' antenna
				txa-rxb - use antenna 'a' for transmitting, antenna 'b' for receiving
				rxa-txb - use antenna 'b' for transmitting, antenna 'a' for receiving`,
			ValidateFunc: validation.StringInSlice([]string{"ant-a", "ant-b", "rxa-txb", "txa-rxb"}, false),
		},
		"area": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Identifies group of wireless networks. This value is announced by AP, " +
				"and can be matched in connect-list by area-prefix. This is a proprietary extension.",
		},
		KeyArp:        PropArpRw,
		KeyArpTimeout: PropArpTimeoutRw,
		"band": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Defines set of used data rates, channel frequencies and widths.",
			ValidateFunc: validation.StringInSlice([]string{"2ghz-b", "2ghz-b/g", "2ghz-b/g/n", "2ghz-onlyg", "2ghz-onlyn", "5ghz-a", "5ghz-a/n", "5ghz-onlyn", "5ghz-a/n/ac", "5ghz-onlyac", "5ghz-n/ac"}, false),
		},
		"basic_rates_a_g": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Similar to the basic-rates-b property, but used for 5ghz, 5ghz-10mhz, 5ghz-5mhz, 5ghz-turbo, " +
				"2.4ghz-b/g, 2.4ghz-onlyg, 2ghz-10mhz, 2ghz-5mhz and 2.4ghz-g-turbo bands.",
			ValidateFunc: validation.StringInSlice([]string{"12Mbps", "18Mbps", "24Mbps", "36Mbps", "48Mbps", "54Mbps", "6Mbps", "9Mbps"}, false),
		},
		"basic_rates_b": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "List of basic rates, used for 2.4ghz-b, 2.4ghz-b/g and 2.4ghz-onlyg bands. Client will connect " +
				"to AP only if it supports all basic rates announced by the AP. AP will establish WDS link only if it " +
				"supports all basic rates of the other AP. This property has effect only in AP modes, and when value of rate-set is configured.",
			ValidateFunc: validation.StringInSlice([]string{"11Mbps", "1Mbps", "2Mbps", "5.5Mbps"}, false),
		},
		"bridge_mode": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "enabled",
			Description:  "Allows to use station-bridge mode. Read more: https://wiki.mikrotik.com/wiki/Manual:Wireless_Station_Modes#Mode_station-bridge",
			ValidateFunc: validation.StringInSlice([]string{"disabled", "enabled"}, false),
		},
		// (integer | disabled; Default: disabled)	Time in microseconds which will be used to send data without stopping. Note that no other wireless cards in that network will be able to transmit data during burst-time microseconds. This setting is available only for AR5000, AR5001X, and AR5001X+ chipset based cards.
		// "burst-time": {
		// 	Type:     schema.TypeString,
		// 	Optional: true,
		// 	Description: "",
		// 	ValidateFunc: validation.StringInSlice([]string{"disabled", "enabled"}, false),
		// },
		"channel_width": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Use of extension channels (e.g. Ce, eC etc) allows additional 20MHz extension channels and if it " +
				"should be located below or above the control (main) channel. Extension channel allows 802.11n devices to " +
				"use up to 40MHz (802.11ac up to 160MHz) of spectrum in total thus increasing max throughput. Channel widths " +
				"with XX and XXXX extensions automatically scan for a less crowded control channel frequency based on the number " +
				"of concurrent devices running in every frequency and chooses the “C” - Control channel frequency automatically.",
			ValidateFunc: validation.StringInSlice([]string{"20mhz", "10mhz", "5mhz", "40mhz-turbo", "20/40mhz-Ce", "20/40mhz-eC", "20/40mhz-XX", "20/40/80mhz-Ceee", "20/40/80mhz-eCee", "20/40/80mhz-eeCe", "20/40/80mhz-eeeC", "20/40/80mhz-XXXX", "20/40/80/160mhz-Ceeeeeee", "20/40/80/160mhz-XXXXXXXX", "20/40/80/160mhz-eCeeeeee", "20/40/80/160mhz-eeCeeeee", "20/40/80/160mhz-eeeCeeee", "20/40/80/160mhz-eeeeCeee", "20/40/80/160mhz-eeeeeCee", "20/40/80/160mhz-eeeeeeCe", "20/40/80/160mhz-eeeeeeeC"}, false),
		},
		KeyComment: PropCommentRw,
		"compression": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Setting this property to yes will allow the use of the hardware compression. Wireless interface must have " +
				"support for hardware compression. Connections with devices that do not use compression will still work.",
			ValidateFunc: validation.StringInSlice([]string{"yes", "no"}, false),
		},
		"country": {
			Type:     schema.TypeString,
			Required: true,
			Description: "Limits available bands, frequencies and maximum transmit power for each frequency. Also specifies default " +
				"value of scan-list. Value no_country_set is an FCC compliant set of channels.",
		},
		"default_ap_tx_limit": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "This is the value of ap-tx-limit for clients that do not match any entry in the access-list. 0 means no limit.",
			// ValidateFunc: validation.FloatBetween(0, 4294967295),
		},
		"default_authentication": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "For AP mode, this is the value of authentication for clients that do not match any entry in the access-list. " +
				"For station mode, this is the value of connect for APs that do not match any entry in the connect-list",
			ValidateFunc: validation.StringInSlice([]string{"yes", "no"}, false),
		},
		"default_client_tx_limit": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "This is the value of client-tx-limit for clients that do not match any entry in the access-list. 0 means no limit",
			// ValidateFunc: validation.FloatBetween(0, 4294967295),
		},
		"default_forwarding": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			Description:  "This is the value of forwarding for clients that do not match any entry in the access-list",
			ValidateFunc: validation.StringInSlice([]string{"yes", "no"}, false),
		},
		"disable_running_check": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "When set to yes interface will always have running flag. If value is set to no', the router " +
				"determines whether the card is up and running - for AP one or more clients have to be registered to it, " +
				"for station, it should be connected to an AP.",
			ValidateFunc: validation.StringInSlice([]string{"yes", "no"}, false),
		},
		KeyDisabled: PropDisabledRw,
		"disconnect_timeout": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "This interval is measured from third sending failure on the lowest data rate. At this point 3 * " +
				"(hw-retries + 1) frame transmits on the lowest data rate had failed. During disconnect-timeout packet " +
				"transmission will be retried with on-fail-retry-time interval. If no frame can be transmitted successfully " +
				"during disconnect-timeout, the connection is closed, and this event is logged as \"extensive data loss\". " +
				"Successful frame transmission resets this timer.",
		},
		"distance": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "	How long to wait for confirmation of unicast frames (ACKs) before considering transmission " +
				"unsuccessful, or in short ACK-Timeout. Distance value has these behaviors: Dynamic - causes AP to detect " +
				"and use the smallest timeout that works with all connected clients. Indoor - uses the default ACK timeout " +
				"value that the hardware chip manufacturer has set. Number - uses the input value in formula: ACK-timeout = " +
				"((distance * 1000) + 299) / 300 us; Acknowledgments are not used in Nstreme/NV2 protocols.",
		},
		"frame_lifetime": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Discard frames that have been queued for sending longer than frame-lifetime. By default, when " +
				"value of this property is 0, frames are discarded only after connection is closed.",
			// ValidateFunc: validation.FloatBetween(0, 4294967295),
		},
		"frequency": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Channel frequency value in MHz on which AP will operate. Allowed values depend on the selected band, " +
				"and are restricted by country setting and wireless card capabilities. This setting has no effect if interface " +
				"is in any of station modes, or in wds-slave mode, or if DFS is active. Note: If using mode \"superchannel\", any " +
				"frequency supported by the card will be accepted, but on the RouterOS client, any non-standard frequency must be " +
				"configured in the scan-list, otherwise it will not be scanning in non-standard range. In Winbox, scanlist " +
				"frequencies are in bold, any other frequency means the clients will need scan-list configured.",
			// ValidateFunc: validation.FloatBetween(0, 4294967295),
		},
		"frequency_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Three frequency modes are available: regulatory-domain - Limit available channels and maximum transmit " +
				"power for each channel according to the value of country manual-txpower - Same as above, but do not limit maximum " +
				"transmit power. superchannel - Conformance Testing Mode. Allow all channels supported by the card. List of available " +
				"channels for each band can be seen in /interface wireless info allowed-channels. This mode allows you to test " +
				"wireless channels outside the default scan-list and/or regulatory domain. This mode should only be used in controlled " +
				"environments, or if you have special permission to use it in your region. Before v4.3 this was called Custom Frequency " +
				"Upgrade, or Superchannel. Since RouterOS v4.3 this mode is available without special key upgrades to all installations.",
			ValidateFunc: validation.StringInSlice([]string{"regulatory-domain", "manual-txpower", "superchannel"}, false),
		},
		"frequency_offset": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Allows to specify offset if the used wireless card operates at a different frequency than is shown in RouterOS, " +
				"in case a frequency converter is used in the card. So if your card works at 4000MHz but RouterOS shows 5000MHz, set " +
				"offset to 1000MHz and it will be displayed correctly. The value is in MHz and can be positive or negative.",
			ValidateFunc: validation.IntBetween(-2147483648, 2147483647),
		},
		"guard_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Whether to allow use of short guard interval (refer to 802.11n MCS specification to see how this may affect " +
				"throughput). \"any\" will use either short or long, depending on data rate, \"long\" will use long.",
			ValidateFunc: validation.StringInSlice([]string{"any", "long"}, false),
		},
		"hide_ssid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "yes - AP does not include SSID in the beacon frames, and does not reply to probe requests that have " +
				"broadcast SSID. no - AP includes SSID in the beacon frames, and replies to probe requests that have broadcast SSID. " +
				"This property has an effect only in AP mode. Setting it to yes can remove this network from the list of wireless " +
				"networks that are shown by some client software. Changing this setting does not improve the security of the wireless " +
				"network, because SSID is included in other frames sent by the AP.",
			ValidateFunc: validation.StringInSlice([]string{"no", "yes"}, false),
		},
		//
		// ht-basic-mcs (list of (mcs-0 | mcs-1 | mcs-2 | mcs-3 | mcs-4 | mcs-5 | mcs-6 | mcs-7 | mcs-8 | mcs-9 | mcs-10 | mcs-11 | mcs-12 | mcs-13 |
		// mcs-14 | mcs-15 | mcs-16 | mcs-17 | mcs-18 | mcs-19 | mcs-20 | mcs-21 | mcs-22 | mcs-23); Default: mcs-0; mcs-1; mcs-2; mcs-3; mcs-4; mcs-5;
		// mcs-6; mcs-7)	Modulation and Coding Schemes that every connecting client must support. Refer to 802.11n for MCS specification.
		//
		// ht-supported-mcs (list of (mcs-0 | mcs-1 | mcs-2 | mcs-3 | mcs-4 | mcs-5 | mcs-6 | mcs-7 | mcs-8 | mcs-9 | mcs-10 | mcs-11 | mcs-12
		// | mcs-13 | mcs-14 | mcs-15 | mcs-16 | mcs-17 | mcs-18 | mcs-19 | mcs-20 | mcs-21 | mcs-22 | mcs-23); Default: mcs-0; mcs-1; mcs-2;
		// mcs-3; mcs-4; mcs-5; mcs-6; mcs-7; mcs-8; mcs-9; mcs-10; mcs-11; mcs-12; mcs-13; mcs-14; mcs-15; mcs-16; mcs-17; mcs-18; mcs-19; mcs-20;
		// mcs-21; mcs-22; mcs-23)	Modulation and Coding Schemes that this device advertises as supported. Refer to 802.11n for MCS specification.
		//
		"hw_fragmentation_threshold": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Specifies maximum fragment size in bytes when transmitted over the wireless medium. 802.11 standard packet (MSDU in 802.11 " +
				"terminologies) fragmentation allows packets to be fragmented before transmitting over a wireless medium to increase the probability " +
				"of successful transmission (only fragments that did not transmit correctly are retransmitted). Note that transmission of a fragmented " +
				"packet is less efficient than transmitting unfragmented packet because of protocol overhead and increased resource usage at both - " +
				"transmitting and receiving party.",
			ValidateFunc: validation.IntBetween(256, 3000),
		},
		"hw_protection_mode": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Frame protection support property read more: https://wiki.mikrotik.com/wiki/Manual:Interface/Wireless#Frame_protection_support_.28RTS.2FCTS.29",
			ValidateFunc: validation.StringInSlice([]string{"cts-to-self", "none", "rts-cts"}, false),
		},
		"hw_protection_threshold": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "Frame protection support property read more: https://wiki.mikrotik.com/wiki/Manual:Interface/Wireless#Frame_protection_support_.28RTS.2FCTS.29",
			ValidateFunc: validation.IntBetween(0, 65535),
		},
		"hw_retries": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Number of times sending frame is retried without considering it a transmission failure. Data-rate is decreased upon " +
				"failure and the frame is sent again. Three sequential failures on the lowest supported rate suspend transmission to this destination " +
				"for the duration of on-fail-retry-time. After that, the frame is sent again. The frame is being retransmitted until transmission " +
				"success, or until the client is disconnected after disconnect-timeout. The frame can be discarded during this time if frame-lifetime is exceeded.",
			ValidateFunc: validation.IntBetween(0, 15),
		},
		"installation": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Adjusts scan-list to use indoor, outdoor or all frequencies for the country that is set.",
			ValidateFunc: validation.StringInSlice([]string{"any", "indoor", "outdoor"}, false),
		},
		"interworking_profile": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validation.StringInSlice([]string{"enabled", "disabled"}, false),
		},
		"keepalive_frames": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "Applies only if wireless interface is in mode=ap-bridge. If a client has not communicated for around 20 seconds, AP sends a " +
				"\"keepalive-frame\". Note, disabling the feature can lead to \"ghost\" clients in registration-table.",
			ValidateFunc: validation.StringInSlice([]string{"enabled", "disabled"}, false),
		},
		"l2mtu": {
			Type:         schema.TypeInt,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validation.IntBetween(0, 65536),
		},
		"mac_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validation.IsMACAddress,
		},
		"master_interface": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "Name of wireless interface that has virtual-ap capability. Virtual AP interface will only work if master " +
				"interface is in ap-bridge, bridge, station or wds-slave mode. This property is only for virtual AP interfaces.",
		},
		"max_station_count": {
			Type:         schema.TypeInt,
			Optional:     true,
			Computed:     true,
			Description:  "Maximum number of associated clients. WDS links also count toward this limit.",
			ValidateFunc: validation.IntBetween(1, 2007),
		},
		"mode": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			Description:  "Selection between different station and access point (AP) modes. https://wiki.mikrotik.com/wiki/Manual:Wireless_Station_Modes",
			ValidateFunc: validation.StringInSlice([]string{"station", "station-wds", "ap-bridge", "bridge", "alignment-only", "nstreme-dual-slave", "wds-slave", "station-pseudobridge", "station-pseudobridge-clone", "station-bridge"}, false),
		},
		KeyMtu: PropMtuRw(),
		"multicast_buffering": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "For a client that has power saving, buffer multicast packets until next beacon time. A client should " +
				"wake up to receive a beacon, by receiving beacon it sees that there are multicast packets pending, and it should " +
				"wait for multicast packets to be sent.",
			ValidateFunc: validation.StringInSlice([]string{"enabled", "disabled"}, false),
		},
		"multicast_helper": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "When set to full, multicast packets will be sent with a unicast destination MAC address, " +
				"resolving multicast problem on the wireless link. This option should be enabled only on the access point, " +
				"clients should be configured in station-bridge mode. Available starting from v5.15. disabled - disables the " +
				"helper and sends multicast packets with multicast destination MAC addresses dhcp - dhcp packet mac addresses " +
				"are changed to unicast mac addresses prior to sending them out full - all multicast packet mac address are " +
				"changed to unicast mac addresses prior to sending them out default - default choice that currently is set to " +
				"disabled. Value can be changed in future releases.",
			ValidateFunc: validation.StringInSlice([]string{"default", "disabled", "full"}, false),
		},
		KeyName: PropNameForceNewRw,
		"noise_floor_threshold": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "For advanced use only, as it can badly affect the performance of the interface. It is possible " +
				"to manually set noise floor threshold value. By default, it is dynamically calculated. This property also " +
				"affects received signal strength. This property is only effective on non-AC chips.",
			ValidateFunc: validation.IntBetween(-128, 127),
		},
		"nv2_cell_radius": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Setting affects the size of contention time slot that AP allocates for clients to initiate " +
				"connection and also size of time slots used for estimating distance to client. When setting is too small, " +
				"clients that are farther away may have trouble connecting and/or disconnect with \"ranging timeout\" error. " +
				"Although during normal operation the effect of this setting should be negligible, in order to maintain maximum " +
				"performance, it is advised to not increase this setting if not necessary, so AP is not reserving time that " +
				"is actually never used, but instead allocates it for actual data transfer. on AP: distance to farthest client " +
				"in km on station: no effect",
			ValidateFunc: validation.IntBetween(10, 200),
		},
		"nv2_noise_floor_offset": {
			Type:         schema.TypeInt,
			Optional:     true,
			ValidateFunc: validation.IntBetween(0, 20),
		},
		"nv2_preshared_key": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"nv2_qos": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Sets the packet priority mechanism, firstly data from high priority queue is sent, then lower " +
				"queue priority data until 0 queue priority is reached. When link is full with high priority queue data, lower " +
				"priority data is not sent. Use it very carefully, setting works on AP frame-priority - manual setting that " +
				"can be tuned with Mangle rules. default - default setting where small packets receive priority for best latency",
			ValidateFunc: validation.StringInSlice([]string{"default", "frame-priority"}, false),
		},
		"nv2_queue_count": {
			Type:         schema.TypeInt,
			Optional:     true,
			ValidateFunc: validation.IntBetween(2, 8),
		},
		"nv2_security": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice([]string{"disabled", "enabled"}, false),
		},
		"on_fail_retry_time": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "After third sending failure on the lowest data rate, wait for specified time interval before retrying.",
		},
		"periodic_calibration": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Setting default enables periodic calibration if info default-periodic-calibration property is enabled. " +
				"Value of that property depends on the type of wireless card. This property is only effective for cards based on Atheros chipset.",
			ValidateFunc: validation.StringInSlice([]string{"default", "disabled", "enabled"}, false),
		},
		"periodic_calibration_interval": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "This property is only effective for cards based on Atheros chipset.",
			ValidateFunc: validation.IntBetween(1, 10000),
		},
		"preamble_mode": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Short preamble mode is an option of 802.11b standard that reduces per-frame overhead.",
			ValidateFunc: validation.StringInSlice([]string{"both", "long", "short"}, false),
		},
		"prism_cardtype": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Specify type of the installed Prism wireless card.",
			ValidateFunc: validation.StringInSlice([]string{"100mW", "200mW", "30mW"}, false),
		},
		"proprietary_extensions": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "	RouterOS includes proprietary information in an information element of management frames. " +
				"This parameter controls how this information is included.",
			ValidateFunc: validation.StringInSlice([]string{"post-2.9.25", "post-2.9.25"}, false),
		},
		"radio_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "Descriptive name of the device, that is shown in registration table entries on the remote devices. " +
				"This is a proprietary extension.",
		},
		"rate_selection": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Starting from v5.9 default value is advanced since legacy mode was inefficient.",
			ValidateFunc: validation.StringInSlice([]string{"advanced", "legacy"}, false),
		},
		"rate_set": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Two options are available: default - default basic and supported rate sets are used. " +
				"Values from basic-rates and supported-rates parameters have no effect. configured - use values " +
				"from basic-rates, supported-rates, basic-mcs, mcs. Read more: " +
				"https://wiki.mikrotik.com/wiki/Manual:Interface/Wireless#Basic_and_MCS_Rate_table",
			ValidateFunc: validation.StringInSlice([]string{"configured", "default"}, false),
		},
		"rx_chains": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Which antennas to use for receive. In current MikroTik routers, both RX and TX chain must " +
				"be enabled, for the chain to be enabled.",
			ValidateFunc: validation.IntBetween(0, 3),
		},
		"scan_list": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Comma separated list of frequencies and frequency ranges | default. Since v6.35 (wireless-rep) type also support range:step option",
		},
		"security_profile": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of profile from security-profiles",
		},
		"secondary_channel": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Specifies secondary channel, required to enable 80+80MHz transmission. To disable 80+80MHz functionality, " +
				"set secondary-channel to \"\" or unset the value via CLI/GUI.",
		},
		"ssid": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "SSID (service set identifier) is a name that identifies wireless network.",
			ValidateFunc: validation.StringLenBetween(0, 32),
		},
		"skip_dfs_channels": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "These values are used to skip all DFS channels or specifically skip DFS CAC " +
				"channels in range 5600-5650MHz which detection could go up to 10min.",
			ValidateFunc: validation.StringInSlice([]string{"disabled", "all", "10min-cac"}, false),
		},
		"station_bridge_clone_mac": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "This property has effect only in the station-pseudobridge-clone mode. Use this " +
				"MAC address when connection to AP. If this value is 00:00:00:00:00:00, station will initially " +
				"use MAC address of the wireless interface. As soon as packet with MAC address of another device " +
				"needs to be transmitted, station will reconnect to AP using that address.",
			ValidateFunc: validation.IsMACAddress,
		},
		"station_roaming": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "Station Roaming feature is available only for 802.11 wireless protocol and only for station modes. Read more: " +
				"https://wiki.mikrotik.com/wiki/Manual:Wireless_Station_Roaming",
			ValidateFunc: validation.StringInSlice([]string{"disabled", "enabled"}, false),
		},
		"supported_rates_a_g": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"6Mbps", "9Mbps", "12Mbps", "18Mbps", "24Mbps", "36Mbps", "48Mbps", "54Mbps"}, false),
			},
			Description: "List of supported rates, used for all bands except 2ghz-b.",
		},
		"supported_rates_b": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"11Mbps", "1Mbps", "2Mbps", "5.5Mbps"}, false),
			},
			Description: "List of supported rates, used for 2ghz-b, 2ghz-b/g and 2ghz-b/g/n bands. Two devices " +
				"will communicate only using rates that are supported by both devices. This property has effect " +
				"only when value of rate-set is configured.",
		},
		"tdma_period_size": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Specifies TDMA period in milliseconds. It could help on the longer distance links, it " +
				"could slightly increase bandwidth, while latency is increased too.",
			ValidateFunc: validation.IntBetween(1, 10),
		},
		"tx_chains": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Which antennas to use for transmitting. In current MikroTik routers, both RX and TX " +
				"chain must be enabled, for the chain to be enabled.",
			ValidateFunc: validation.IntBetween(0, 3),
		},
		"tx_power": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "For 802.11ac wireless interface it's total power but for 802.11a/b/g/n it's power per chain.",
			ValidateFunc: validation.IntBetween(-30, 40),
		},
		"tx_power_mode": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Sets up tx-power mode for wireless card",
			ValidateFunc: validation.StringInSlice([]string{"default", "card-rates", "all-rates-fixed", "manual-table"}, false),
		},
		//
		// update-stats-interval (; Default: )	How often to request update of signals strength and ccq values from clients.
		// Access to registration-table also triggers update of these values.
		// This is proprietary extension.
		//
		"vht_basic_mcs": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Modulation and Coding Schemes that every connecting client must support. Refer to 802.11ac for MCS specification.",
			ValidateFunc: validation.StringInSlice([]string{"none", "MCS 0-7", "MCS 0-8", "MCS 0-9"}, false),
		},
		"vht_supported_mcs": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Modulation and Coding Schemes that this device advertises as supported. Refer to 802.11ac for MCS specification.",
			ValidateFunc: validation.StringInSlice([]string{"none", "MCS 0-7", "MCS 0-8", "MCS 0-9"}, false),
		},
		//
		// wds-cost-range (start [-end] integer[0..4294967295]; Default: 50-150)
		// Bridge port cost of WDS links are automatically adjusted, depending on measured link throughput. Port cost is recalculated and adjusted every 5 seconds if it has changed by more than 10%, or if more than 20 seconds have passed since the last adjustment. Setting this property to 0 disables automatic cost adjustment. Automatic adjustment does not work for WDS links that are manually configured as a bridge port.
		//
		"wds_default_bridge": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "When WDS link is established and status of the wds interface becomes running, it will be added as a bridge port to the " +
				"bridge interface specified by this property. When WDS link is lost, wds interface is removed from the bridge. If wds interface " +
				"is already included in a bridge setup when WDS link becomes active, it will not be added to bridge specified by , and will (needs editing)",
		},
		"wds_default_cost": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Initial bridge port cost of the WDS links.",
			// ValidateFunc: validation.FloatBetween(0, 4294967295),
		},
		"wds_ignore_ssid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			Description: "By default, WDS link between two APs can be created only when they work on the same frequency and have the same " +
				"SSID value. If this property is set to yes, then SSID of the remote AP will not be checked. This property has no effect " +
				"on connections from clients in station-wds mode. It also does not work if wds-mode is static-mesh or dynamic-mesh.",
			ValidateFunc: validation.StringInSlice([]string{"yes", "no"}, false),
		},
		"wds_mode": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			Description:  "Controls how WDS links with other devices (APs and clients in station-wds mode) are established.",
			ValidateFunc: validation.StringInSlice([]string{"disabled", "dynamic", "dynamic-mesh", "static", "static-mesh"}, false),
		},
		"wireless_protocol": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Specifies protocol used on wireless interface.",
			ValidateFunc: validation.StringInSlice([]string{"802.11", "any", "nstreme", "nv2", "nv2-nstreme", "nv2-nstreme-802.11", "unspecified"}, false),
		},
		"wmm_support": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			Description:  "Specifies whether to enable WMM.  Only applies to bands B and G. Other bands will have it enabled regardless of this setting",
			ValidateFunc: validation.StringInSlice([]string{"disabled", "enabled", "required"}, false),
		},
		"wps_mode": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			Description:  "https://wiki.mikrotik.com/wiki/Manual:Interface/Wireless#WPS_Server",
			ValidateFunc: validation.StringInSlice([]string{"disabled", "push-button", "push-button-virtual-only"}, false),
		},
		"interface_type": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"update_stats_interval": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"wds_cost_range": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"vlan_mode": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"running": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"vlan_id": {
			Type:     schema.TypeInt,
			Computed: true,
		},

		// Some properties are not implemented, see: https://help.mikrotik.com/docs/display/ROS/Wireless+Interface
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
