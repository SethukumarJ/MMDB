package models

// GeoData represents the structure of our IP geolocation data
type GeoData struct {
	GeoIsoCode   string `maxminddb:"geo_iso_code_2"`
	State        string `maxminddb:"geo_state"`
	City         string `maxminddb:"geo_city"`
	IspName      string `maxminddb:"isp_name"`
	ConnType     string `maxminddb:"conn_type"`
	VpnProxyType string `maxminddb:"vpn_proxy_type"`
}
