package libmacouflage

import (
        "fmt"
        "net"
        "strings"
        "encoding/json"
)

var OuiDb []Oui

type Oui struct {
        VendorPrefix string     `json:"vendor_prefix"`
        Vendor string           `json:"vendor_name"`
}

type NoVendorError struct {
	msg string
}

func (e NoVendorError) Error() string {
	return e.msg
}

func init() {
	OuiData, err := Asset("data/ouis.json")
	if err != nil {
		fmt.Println(err)
		return
	}
        err = json.Unmarshal(OuiData, &OuiDb)
        if err != nil {
                fmt.Println(err)
                return
        }
}

func ValidateMac(mac string) (err error) {
        _, err = net.ParseMAC(mac)
        return
}

func FindVendorByMac(mac string) (vendor Oui, err error) {
        err = ValidateMac(mac)
        if err != nil {
                return
        }
        for _, oui := range OuiDb {
                if(strings.EqualFold(oui.VendorPrefix, mac[:8])) {
                        vendor = oui
                        return
                }
        }
        msg := fmt.Sprintf("No vendor found in OuiDb for vendor prefix: %s", mac[:8])
	err = NoVendorError{msg}
        return
}
