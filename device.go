package uasurfer

import (
	"strings"
)

var TabletDevices = []string{
	"tab", 					//2.08%
	"pad",					//1.66%
	" nexus 7 ", 				//0.08%
	" nexus 9 ", 				//0.00%
	" nexus 10 ", 				//0.03%
	" xoom ", 				//0.00%
	" a101 ",				//0.00%
	" a1-810 ",				//0.00%
	" a1-811 ",				//0.00%
	" a1-841 ",				//0.00%
	" a1-850 ",				//0.03%
	" a210 ",				//0.01%
	" a211 ",				//0.00%
	" a3-a20 ",				//0.00%
	" a3-a40 ",				//0.00%
	" a500 ",				//0.00%
	" a700 ",				//0.01%
	" b1-711 ",				//0.00%
	" b1-730",				//0.00%
	" b1-a71 ",				//0.01%
	" b3-a20 ",				//0.00%
	" b3-a30 ",				//0.00%
	" b1-710 ",				//0.00%
	" 8063 ",				//0.00%
	" 9003x ",				//0.00%
	" 9005x ",				//0.00%
	" 9010x ",				//0.01%
	" i216x ",				//0.00%
	" i213 ",				//0.00%
	" p310x ",				//0.00%
	" p320x ",				//0.00%
	" x96 ",				//0.02%
	" archos 101",				//0.02%
	" ap-110 ",				//0.00%
	" k007 ",				//0.00%
	" k00e ",				//0.01%
	" k013 ",				//0.03%
	" asus_l001 ",				//0.00%
	" p008 ",				//0.00%
	" a727 ",				//0.00%
	" bq7008g ",				//0.00%
	" bq-7021g ",				//0.00%
	" nb74 ",				//0.01%
	" nb75 ",				//0.01%
	" nb751 ",				//0.01%
	" nb85 ",				//0.01%
	" cavion_10_3grq ",			//0.00%
	" cw-hi8-super ",			//0.00%
	" u55gt",				//0.00%
	" u65gt",				//0.01%
	" venue 8 ",				//0.01%
	" venue 7 ",				//0.00%
	" ht7070mg ",				//0.01%
	" tt7071mg",				//0.00%
	" ps1060mg",				//0.00% Digma Plane 1601 3G
	" eny m8s",				//0.00%
	" squad ",				//0.00%
	"fly flylife connect", 			//0.05%
	" p771a ",				//0.01%
	" fdr-a01l ",				//0.00%
	" tz60 ",				//0.03%
	" tz70 ",				//0.00%
	" lenovoa3300", 			//0.11%
	" lenovo a3500", 			//0.06%
	" lenovo a5500", 			//0.05%
	" lenovo a7600-h", 			//0.05%
	" lenovo a7600-f", 			//0.02%
	" lenovo b6000",			//0.04%
	" lenovo b8000",			//0.03%
	" lenovo s5000",			//0.01%
	" lenovo s6000",			//0.07%
	" lenovo t2-x30",			//0.05%
	" lenovo tb3-710",			//0.14%
	" lenovo tb3-730", 			//0.02%
	" lenovo tb3-850",			//0.02%
	" lenovo yt3-850",			//0.01%
	" lenovo yt3-x50",			//0.01%
	" lenovo tb-x103f",			//0.02%
	" mz608 ",				//0.01%
	" next761 ",				//0.01%
	" nomi_c0700",				//0.08%
	" nomi c10103",				//0.01%
	" oysters T72",				//0.01%
	" pmp3670b ",				//0.00%
	" pmp3970b ",				//0.00%
	" pmp5670c ",				//0.00%
	" pmp5870c ",				//0.00%
	" pmp7100d3g_quad ",			//0.00%
	" pmp7280",				//0.01%
	" pmt3057_3g ",				//0.00%
	" pmt3111_wi ",				//0.00%
	" pmt3147 3g ",				//0.00%
	" pmt3331_3g ",				//0.00%
	" pmt3341_3g ",				//0.00%
	" pmt3787_3g ",				//0.00%
	" pmt3797_3g ",				//0.00%
	" pmt5001",				//0.00%
	" pmt5011_3g ",				//0.00%
	" pmt5777_3g ",				//0.00%
	" pmt5887",				//0.00%
	" rk29sdk ",				//0.00%
	" rk30sdk ",				//0.00%
	" gt-p1000 ",				//0.00%
	" gt-p3100 ",				//0.08%
	" gt-p3110 ",				//0.01%
	" gt-p3113 ",				//0.00%
	" gt-p5100 ",				//0.06%
	" gt-p5110 ",				//0.02%
	" gt-p5113 ",				//0.00%
	" gt-p5200 ",				//0.03%
	" gt-p5210 ",				//0.02%
	" gt-p6200 ",				//0.00%
	" gt-p6800 ",				//0.00%
	" gt-p7300 ",				//0.00%
	" gt-p7500 ",				//0.00%
	" gt-p7510 ",				//0.00%
	" gt-n5100 ",				//0.01%
	" gt-n5110 ",				//0.00%
	" gt-n8000 ",				//0.04%
	" gt-n8010 ",				//0.00%
	" gt-n8013 ",				//0.00%
	" sm-t110 ",				//0.06%
	" sm-t111 ",				//0.03%
	" sm-t113 ",				//0.02%
	" sm-t116 ",				//0.07%
	" sm-t210 ",				//0.02%
	" sm-t211 ",				//0.02%
	" sm-t211 ",				//0.00%
	" sm-t230 ",				//0.02%
	" sm-t230nu ",				//0.00%
	" sm-t231 ",				//0.05%
	" sm-t235 ",				//0.01%
	" sm-t239m ",				//0.00%
	" sm-t280 ",				//0.01%
	" sm-t285 ",				//0.04%
	" sm-t310 ",				//0.03%
	" sm-t311 ",				//0.01%
	" sm-t320 ",				//0.00%
	" sm-t321 ",				//0.00%
	" sm-t325 ",				//0.00%
	" sm-t331 ",				//0.01%
	" sm-t350 ",				//0.01%
	" sm-t355 ",				//0.03%
	" sm-t520 ",				//0.00%
	" sm-t525 ",				//0.00%
	" sm-t530 ",				//0.02%
	" sm-t530nu ",				//0.00%
	" sm-t531 ",				//0.05%
	" sm-t535 ",				//0.00%
	" sm-t550 ",				//0.01%
	" sm-t555 ",				//0.04%
	" sm-t560 ",				//0.05%
	" sm-t560nu ",				//0.00%
	" sm-t561m ",				//0.00%
	" sm-t580 ",				//0.01%
	" sm-t585 ",				//0.06%
	" sm-t700 ",				//0.00%
	" sm-t705 ",				//0.01%
	" sm-t710 ",				//0.00%
	" sm-t715 ",				//0.00%
	" sm-t719 ",				//0.01%
	" sm-t800 ",				//0.01%
	" sm-t805 ",				//0.01%
	" sm-t805m ",				//0.00%
	" sm-t810 ",				//0.00%
	" sm-t813 ",				//0.00%
	" sm-t815 ",				//0.01%
	" sm-t819 ",				//0.01%
	" sm-t820 ",				//0.00%
	" sm-t900 ",				//0.00%
	" sm-t561 ",				//0.16%
	" sm-p600 ",				//0.01%
	" sm-p601 ",				//0.01%
	" sm-p605 ",				//0.00%
	" sm-p901 ",				//0.00%
	" sm-p905 ",				//0.00%
	" sgp312 ",				//0.00%
	" sgp321 ",				//0.01%
	" sgp511 ",				//0.00%
	" sgp512 ",				//0.00%
	" sgp521 ",				//0.01%
	" sgp621 ",				//0.00%
	" sgp771 ",				//0.00%
	" d101 ",				//0.00%
	" m102 ",				//0.00%
}

func SearchStrings(a []string, x string) bool {

	for _, v := range a {
		if strings.Contains(x, v) {
			return true
		}
	}
	return false

}

func (u *UserAgent) evalDevice(ua string) {
	switch {

	case u.OS.Platform == PlatformWindows || u.OS.Platform == PlatformMac || u.OS.Name == OSChromeOS:
		if strings.Contains(ua, "mobile") || strings.Contains(ua, "touch") {
			u.DeviceType = DeviceTablet // windows rt, linux haxor tablets
			return
		}
		u.DeviceType = DeviceComputer

	case u.OS.Platform == PlatformiPad || u.OS.Platform == PlatformiPod || strings.Contains(ua, "tablet") || strings.Contains(ua, "kindle/") || strings.Contains(ua, "playbook"):
		u.DeviceType = DeviceTablet

	case u.OS.Platform == PlatformiPhone || u.OS.Platform == PlatformBlackberry || strings.Contains(ua, "phone"):
		u.DeviceType = DevicePhone

	// long list of smarttv and tv dongle identifiers
	case strings.Contains(ua, "tv") || strings.Contains(ua, "crkey") || strings.Contains(ua, "googletv") || strings.Contains(ua, "aftb") || strings.Contains(ua, "adt-") || strings.Contains(ua, "roku") || strings.Contains(ua, "viera") || strings.Contains(ua, "aquos") || strings.Contains(ua, "dtv") || strings.Contains(ua, "appletv") || strings.Contains(ua, "smarttv") || strings.Contains(ua, "tuner") || strings.Contains(ua, "smart-tv") || strings.Contains(ua, "hbbtv") || strings.Contains(ua, "netcast") || strings.Contains(ua, "vizio"):
		u.DeviceType = DeviceTV

	case u.OS.Name == OSAndroid:
		if SearchStrings(TabletDevices, ua) {
			u.DeviceType = DeviceTablet
			return
		}

		u.DeviceType = DevicePhone // default to phone

	case u.OS.Platform == PlatformPlaystation || u.OS.Platform == PlatformXbox || u.OS.Platform == PlatformNintendo:
		u.DeviceType = DeviceConsole

	case strings.Contains(ua, "glass") || strings.Contains(ua, "watch") || strings.Contains(ua, "sm-v"):
		u.DeviceType = DeviceWearable

	// specifically above "mobile" string check as Kindle Fire tablets report as "mobile"
	case u.Browser.Name == BrowserSilk || u.OS.Name == OSKindle && !strings.Contains(ua, "sd4930ur"):
		u.DeviceType = DeviceTablet

	case strings.Contains(ua, "mobile") || strings.Contains(ua, "touch") || strings.Contains(ua, " mobi") || strings.Contains(ua, "webos"): //anything "mobile"/"touch" that didn't get captured as tablet, console or wearable is presumed a phone
		u.DeviceType = DevicePhone

	case u.OS.Name == OSLinux: // linux goes last since it's in so many other device types (tvs, wearables, android-based stuff)
		u.DeviceType = DeviceComputer

	default:
		u.DeviceType = DeviceUnknown
	}
}
