package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"

	"gopkg.in/ldap.v2"
)

func ConnectLDAP() (*ldap.Conn, error) {
	config := LoadConfig()
	tlsConfig := tls.Config{}

	var l *ldap.Conn

	if config.LDAP.IsTLS {
        tlsConfig.ServerName = config.LDAP.ServerHost
		if config.LDAP.CACertFilePath != "" {
			caCert, err := ioutil.ReadFile(config.LDAP.CACertFilePath)
			if err != nil {
				return nil, errors.Wrap(err, "Failed to load CACert")
			}
			caCertPool := x509.NewCertPool()
			caCertPool.AppendCertsFromPEM(caCert)
			tlsConfig.RootCAs = caCertPool
		}
		_l, err := ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", config.LDAP.ServerHost, config.LDAP.ServerPort), &tlsConfig)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to dial to LDAPS server")
		}
		l = _l
	} else {
		_l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", config.LDAP.ServerHost, config.LDAP.ServerPort))
		if err != nil {
			return nil, errors.Wrap(err, "Failed to dial to LDAP server")
		}
		l = _l
	}

	err := l.Bind(config.LDAP.BindDN, config.LDAP.BindPassword)
	if err != nil {
		l.Close()
		return nil, errors.Wrap(err, "Failed to bind binding user")
	}

	return l, nil
}

func AuthLDAP(l *ldap.Conn, username string, password string) error {
	config := LoadConfig()

	searchRequest := ldap.NewSearchRequest(
		config.LDAP.BaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf(config.LDAP.SearchFilter, username),
		[]string{"dn"}, nil,
	)
	result, err := l.Search(searchRequest)
	if err != nil {
		return errors.Wrap(err, "Failed to search")
	}
	count := len(result.Entries)
	if count > 1 {
		return errors.New("Unexpected user")
	}
	if count <= 0 {
		return errors.New("No such user")
	}
	userDN := result.Entries[0].DN

	err = l.Bind(userDN, password)
	if err != nil {
		return errors.Wrap(err, "Failed to bind normal user")
	}
	err = l.Bind(config.LDAP.BindDN, config.LDAP.BindPassword)
	if err != nil {
		return errors.Wrap(err, "Failed to re-bind")
	}

	return nil
}
