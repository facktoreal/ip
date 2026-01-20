[![Deploy to Appengine](https://github.com/facktoreal/ip/actions/workflows/master.yml/badge.svg)](https://github.com/facktoreal/ip/actions/workflows/master.yml)

# Public IP service

This service provides a simple way to retrieve your public IP address. It is useful for applications that need to know your current public IP address, such as VPN clients, remote access tools, and more.

Available free of charge at [clean-ip.nw.r.appspot.com](https://clean-ip.nw.r.appspot.com/)

Default response is JSON with IP address.

TODO:
- [ ] Add support for IPv6 vs IPv4
- [ ] Add support for other formats (html, xml)
- [ ] Add support for ISP detection
- [ ] Add support for VPN detection
- [ ] Add support for proxy detection
- [ ] Add support for TOR detection
- [ ] Add support for GeoIP location

### Install on Appengine

Google offer a free tier for Appengine. If your application can fit into F1 instance, you have 28 hours per day for free.

Required GCP API's:
-  App Engine Admin API (if you deploy applicaation from Github Actions) [console.cloud.google.com/apis/library/appengine.googleapis.com?project=PROJECT_NAME](https://console.cloud.google.com/apis/library/appengine.googleapis.com?project=PROJECT_NAME)

Service account need a permission to deploy application to Appengine:
- App Engine Deployer (if you deploy applicaation from Github Actions) [console.cloud.google.com/iam-admin/serviceaccounts/project?project=PROJECT_NAME](https://console.cloud.google.com/iam-admin/serviceaccounts/project?project=PROJECT_NAME)
- Service Account User [console.cloud.google.com/iam-admin/serviceaccounts/project?project=PROJECT_NAME](https://console.cloud.google.com/iam-admin/serviceaccounts/project?project=PROJECT_NAME)
- 