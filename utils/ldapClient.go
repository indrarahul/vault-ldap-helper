package utils

import (
	"crypto/tls"
	"fmt"

	"gopkg.in/ldap.v2"
)

//LDAPClient s
type LDAPClient struct {
	Conn *ldap.Conn

	Hosts       []string
	Port        int
	Base        string
	GroupFilter string
}

//Connect function for connecting to the ldap
func (lc *LDAPClient) Connect() error {
	if lc.Conn == nil {
		var ldapConn *ldap.Conn
		var connErr, tlsErr error

		for _, each := range lc.Hosts {
			ldapURL := fmt.Sprintf("%s:%d", each, lc.Port)

			ldapConn, connErr = ldap.Dial("tcp", ldapURL)
			if connErr != nil {
				continue
			}

			tlsErr = ldapConn.StartTLS(&tls.Config{InsecureSkipVerify: true})
			if tlsErr == nil {
				break
			}
		}

		if connErr != nil {
			return connErr
		}

		if tlsErr != nil {
			return tlsErr
		}

		lc.Conn = ldapConn
	}

	return nil
}

//Close function for closing the ldap connection
func (lc *LDAPClient) Close() {
	if lc.Conn != nil {
		lc.Conn.Close()
		lc.Conn = nil
	}
}

//GetUsers function for getting the groups a user is associated with
func (lc *LDAPClient) GetUsers() (*ldap.SearchResult, error) {

	err := lc.Connect()
	if err != nil {
		return nil, err
	}

	searchReq := ldap.NewSearchRequest(
		lc.Base,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		lc.GroupFilter,
		[]string{},
		nil,
	)

	req, err := lc.Conn.Search(searchReq)
	if err != nil {
		return nil, err
	}

	return req, nil
}
