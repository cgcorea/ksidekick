package kannel

import (
	"fmt"
	"strconv"
)

// Username set header X-Kannel-Username (username in SMS Push)
func Username(value string) func(r *Request) {
	return func(r *Request) {
		r.Header.Set("X-Kannel-Username", value)
	}
}

// Password set header X-Kannel-password (Password in SMS Push)
func Password(value string) func(r *Request) {
	return func(r *Request) {
		r.Header.Set("X-Kannel-password", value)
	}
}

// From set header X-Kannel-From (from in SMS Push)
func From(value string) func(r *Request) {
	return func(r *Request) {
		r.Header.Set("X-Kannel-From", value)
	}
}

// To set header X-Kannel-To (to in SMS Push)
func To(value string) func(r *Request) {
	return func(r *Request) {
		r.Header.Set("X-Kannel-To", value)
	}
}

// Charset set charset as in Content-Type: text/html; charset=ISO-8859-1
func (r *Request) Charset(value string) func(r *Request) {
	return func(r *Request) {
		contentType := fmt.Sprintf("text/html; charset=%v", value)
		r.Header.Set("Content-Type", contentType)
	}
}

// UDH set header X-Kannel-UDH (udh in SMS Push)
func UDH(value string) func(r *Request) {
	return func(r *Request) {
		r.Header.Set("X-Kannel-UDH", value)
	}
}

// SMSC Optional virtual smsc-id from which the message is supposed to have
// arrived. This is used for routing purposes, if any denied or preferred SMS
// centers are set up in SMS center configuration.
func SMSC(value string) func(r *Request) {
	return func(r *Request) {
		r.Header.Set("X-Kannel-SMSC", value)
	}
}

// MClass Sets the Message Class in DCS field. Accepts values between 0 and
// 3, for Message Class 0 to 3, A value of 0 sends the message directly to
// display, 1 sends to mobile, 2 to SIM and 3 to SIM toolkit.
func MClass(value int) func(r *Request) {
	return func(r *Request) {
		r.Header.Set("X-Kannel-MClass", strconv.Itoa(value))
	}
}

// MWI Optional. Sets Message Waiting Indicator bits in DCS field. If given,
// the message will be encoded as a Message Waiting Indicator. The accepted
// values are 0,1,2 and 3 for activating the voice, fax, email and other
// indicator, or 4,5,6,7 for deactivating, respectively. This option excludes
// the flash option
func MWI(value int) func(r *Request) {
	return func(r *Request) {
		r.Header.Set("X-Kannel-MWI", strconv.Itoa(value))
	}
}

// Compress Optional. Sets the Compression bit in DCS Field.
func (r *Request) Compress(value int) func(r *Request) {
	return func(r *Request) {
		r.Header.Set("X-Kannel-Compress", strconv.Itoa(value))
	}
}

// Coding Optional. Sets the coding scheme bits in DCS field. Accepts values
// 0 to 2, for 7bit, 8bit or UCS-2. If unset, defaults to 7 bits unless a udh is
// defined, which sets coding to 8bits.
//
// If unset, defaults to 0 (7 bits) if Content-Type is text/plain,  text/html or
// text/vnd.wap.wml. On application/octet-stream, defaults to 8 bits (1). All
// other Content-Type values are rejected.
func Coding(value int) func(r *Request) {
	return func(r *Request) {
		r.Header.Set("X-Kannel-Coding", strconv.Itoa(value))
	}
}

// Validity Optional (minutes). If given, Kannel will inform SMS Center that
// it should only try to send the message for this many minutes. If the
// destination mobile is off other situation that it cannot receive the sms, the
// smsc discards the message. Note: you must have your Kannel box time
// synchronized with the SMS Center.
func Validity(value int) func(r *Request) {
	return func(r *Request) {
		r.Header.Set("X-Kannel-Validity", strconv.Itoa(value))
	}
}

// Deferred Optional (minutes). If given, the SMS center will postpone the
// message to be delivered at now plus this many minutes. Note: you must have
// your Kannel box time synchronized with the SMS Center.
func Deferred(value int) func(r *Request) {
	return func(r *Request) {
		r.Header.Set("X-Kannel-Deferred", strconv.Itoa(value))
	}
}

// DLRMask Optional (bit mask). Request for delivery reports with the state
// of the sent message. The value is a bit mask composed of:
//   1: Delivered to phone
//   2: Non-Delivered to Phone
//   4: Queued on SMSC
//   8: Delivered to SMSC
//  16: Non-Delivered to SMSC
// Must set dlr-url on sendsms-user group or use the dlr-url CGI variable.
func DLRMask(value int) func(r *Request) {
	return func(r *Request) {
		r.Header.Set("X-Kannel-DLR-Mask", strconv.Itoa(value))
	}
}

// DLRUrl Optional. If dlr-mask is given, this is the url to be fetched.
// (Must be url-encoded)
func DLRUrl(value string) func(r *Request) {
	return func(r *Request) {
		r.Header.Set("X-Kannel-DLR-Url", value)
	}
}

// Account Optional. Account name or number to carry forward for billing
// purposes. This field is logged as ACT in the log file so it allows you to do
// some accounting on it if your front end uses the same username for all
// services but wants to distinguish them in the log. In the case of a HTTP SMSC
// type the account name is prepended with the service-name (username) and a
// colon (:) and forwarded to the next instance of Kannel. This allows
// hierarchical accounting.
func Account(value string) func(r *Request) {
	return func(r *Request) {
		r.Header.Set("X-Kannel-Account", value)
	}
}

// PID Optional. Sets the PID value. (See ETSI Documentation). Ex: SIM
// Toolkit messages would use something like &pid=127&coding=1&alt-dcs=1&mclass=3
func PID(value int) func(r *Request) {
	return func(r *Request) {
		r.Header.Set("X-Kannel-PID", strconv.Itoa(value))
	}
}

// AltDCS Optional. If unset, Kannel uses the alt-dcs defined on smsc
// configuration, or 0X per default. If equals to 1, uses FX. If equals to 0,
// force 0X.
func AltDCS(value int) func(r *Request) {
	return func(r *Request) {
		r.Header.Set("X-Kannel-Alt-DCS", strconv.Itoa(value))
	}
}

// BInfo Optional. Billing identifier/information proxy field used to pass
// arbitrary billing transaction IDs or information to the specific SMSC
// modules. For EMI2 this is encapsulated into the XSer 0c field, for SMPP this
// is encapsulated into the service_type of the submit_sm PDU.
func BInfo(value string) func(r *Request) {
	return func(r *Request) {
		r.Header.Set("X-Kannel-BInfo", value)
	}
}

// RPI Optional. Sets the Return Path Indicator (RPI) value.
// (See ETSI Documentation).
func RPI(value int) func(r *Request) {
	return func(r *Request) {
		r.Header.Set("X-Kannel-RPI", strconv.Itoa(value))
	}
}

// Priority Optional. Sets the Priority value (range 0-3 is allowed).
// (Defaults to 0, which is the lowest priority).
func Priority(value int) func(r *Request) {
	return func(r *Request) {
		r.Header.Set("X-Kannel-Priority", strconv.Itoa(value))
	}
}

// Metadata SMPP TLVs
func Metadata(value string) func(r *Request) {
	return func(r *Request) {
		r.Header.Set("X-Kannel-Meta-Data", value)
	}
}

func (r *Request) Set(option func(r *Request)) {
	option(r)
}
