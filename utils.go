package norenapigo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/golang/glog"
)

type Time struct {
	time.Time
}

// API endpoints
const (
	URILogin              string = "QuickAuth"
	URIUserSessionRenew   string = "QuickAuth"
	URIUserProfile        string = "UserDetails"
	URILogout             string = "Logout"
	URIGetOrderBook       string = "SingleOrdHist"
	URIPlaceOrder         string = "PlaceOrder"
	URIModifyOrder        string = "ModifyOrder"
	URICancelOrder        string = "CancelOrder"
	URIGetHoldings        string = "Holdings"
	URIGetPositions       string = "PositionBook"
	URIGetTradeBook       string = "TradeBook"
	URILTP                string = "GetQuotes"
	URITPSeries           string = "TPSeries"
	URIEODTPSeries        string = "EODChartData"
	URISearchScript       string = "SearchScrip"
	URISingleOrderHistory string = "SingleOrdHist"
	URISecurityInfo       string = "GetSecurityInfo"
	URIOrderMargin        string = "GetOrderMargin"
	URIPlaceGTTOrder      string = "PlaceGTTOrder"
	URIPlaceOCOOrder      string = "PlaceOCOOrder"
	URICancelGTTOrder     string = "CancelGTTOrder"
	URIGetPendingGTTOrder string = "GetPendingGTTOrder"
	URIGetEnabledGTTs     string = "GetEnabledGTTs"
)

func structToMap(obj interface{}, tagName string) map[string]interface{} {
	var values reflect.Value
	switch obj.(type) {
	case OrderParams:
		{
			con := obj.(OrderParams)
			values = reflect.ValueOf(&con).Elem()
		}
	case GetOrderParams:
		{
			con := obj.(GetOrderParams)
			values = reflect.ValueOf(&con).Elem()
		}

	case ModifyOrderParams:
		{
			con := obj.(ModifyOrderParams)
			values = reflect.ValueOf(&con).Elem()
		}
	case LTPParams:
		{
			con := obj.(LTPParams)
			values = reflect.ValueOf(&con).Elem()
		}
	case SecurityInfoParams:
		{
			con := obj.(SecurityInfoParams)
			values = reflect.ValueOf(&con).Elem()
		}
	case TSPriceParam:
		{
			con := obj.(TSPriceParam)
			values = reflect.ValueOf(&con).Elem()
		}
	case DailyTSParam:
		{
			con := obj.(DailyTSParam)
			values = reflect.ValueOf(&con).Elem()
		}
	case ConvertPositionParams:
		{
			con := obj.(ConvertPositionParams)
			values = reflect.ValueOf(&con).Elem()
		}
	}

	tags := reflect.TypeOf(obj)
	params := make(map[string]interface{})
	for i := 0; i < values.NumField(); i++ {
		params[tags.Field(i).Tag.Get(tagName)] = values.Field(i).Interface()
	}

	return params
}

func getIpAndMac() (string, string, string, error) {

	//----------------------
	// Get the local machine IP address
	//----------------------

	var localIp, currentNetworkHardwareName string

	localIp, err := getLocalIP()

	if err != nil {
		return "", "", "", err
	}

	// get all the system's or local machine's network interfaces

	interfaces, _ := net.Interfaces()
	for _, interf := range interfaces {

		if addrs, err := interf.Addrs(); err == nil {
			for _, addr := range addrs {

				// only interested in the name with current IP address
				if strings.Contains(addr.String(), localIp) {
					currentNetworkHardwareName = interf.Name
				}
			}
		}
	}

	// extract the hardware information base on the interface name
	// capture above
	netInterface, err := net.InterfaceByName(currentNetworkHardwareName)

	if err != nil {
		return "", "", "", err
	}

	macAddress := netInterface.HardwareAddr

	// verify if the MAC address can be parsed properly
	_, err = net.ParseMAC(macAddress.String())

	if err != nil {
		return "", "", "", err
	}

	publicIp, err := getPublicIp()
	if err != nil {
		return "", "", "", err
	}

	return localIp, publicIp, macAddress.String(), nil

}

func getLocalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("please check your network connection")
}

func getPublicIp() (string, error) {
	resp, err := http.Get("https://myexternalip.com/raw")
	if err != nil {
		return "", err
	}

	content, _ := ioutil.ReadAll(resp.Body)
	err = resp.Body.Close()
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func GetTime(timeString string) string {
	layout := "02-01-2006 15:04:05"
	time_location, _ := time.LoadLocation("Asia/Kolkata")
	t, err := time.ParseInLocation(layout, timeString, time_location)
	// t, err := time.Parse(layout, timeString)
	if err != nil {
		glog.Fatal(err)
	}
	return fmt.Sprintf("%d", t.Unix())
}

func GetDate(timestampStr string) (string, error) {
	// Define the layout of the timestamp string
	layout := "02-01-2006"

	// Parse the timestamp string into a time.Time object
	timestamp, err := time.Parse(layout, timestampStr)
	if err != nil {
		return "", fmt.Errorf("error parsing timestamp: %v", err)
	}

	// Convert the time.Time object to epoch time (in seconds)
	epochTime := timestamp.Unix()

	// Convert the epoch time to a string
	epochTimeStr := fmt.Sprintf("%d", epochTime)

	return epochTimeStr, nil
}
func GetTodayAndLastWeekEpoch() (int64, int64) {
	// Get today's date
	today := time.Now()

	// Subtract 7 days to get the date from one week ago
	lastWeek := today.AddDate(0, 0, -7)

	// Convert dates to epoch time (in seconds)
	todayEpoch := today.Unix()
	lastWeekEpoch := lastWeek.Unix()

	return todayEpoch, lastWeekEpoch
}
