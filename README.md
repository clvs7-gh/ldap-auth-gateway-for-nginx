# ldap-auth-gateway-for-nginx

## What's this?
A LDAP auth gateway for nginx.

Example config file for nginx is included. See example directory.

This software is based on auth gateway I developed privately for lab.
Some of the features are removed from it (because they are very private one).

I've confirmed that this gateway works with Samba AD or OpenLDAP.
Other LDAP(S) servers are not tested, but it may be works with. If not works with them, please fix gateway's source codes. This is CC0, so there is no support.

## Supported features
- Login with credentials stored in LDAP server, including Active Directory.
- Cookie-based session.
- Pass username to gateway-protected servers.

## How to build
1. Run `make dep-install`.
2. Let's build. Run `make`.

## Author
clvs7

## License
CC0 (Public Domain)

## Contribute
If you make PR for this repository, please license it with CC0.
