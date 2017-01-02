package emailcheck

import (
	"net"
	"regexp"
	"time"
)

var EMAIL_REGEX = regexp.MustCompile("[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?")

func CheckRegex(email string) bool {
	return EMAIL_REGEX.Match([]byte(email))
}

func CheckRecords(email string) (bool, error) {
	records, err := getRecords(email)

	if err != nil {
		return false, err
	}

	return len(records) != 0, nil
}

func CheckConnectivity(email string) (bool, error) {
	records, err := getRecords(email)

	if err != nil {
		return false, err
	}

	d := &net.Dialer{Timeout: 1 * time.Second}

	for _, record := range records {
		host := record.Host

		if host[len(host)-1] == '.' {
			host = host[0 : len(host)-1]
		}

		//most home IPS will have this port blocked
		//but server providers usually keep it open
		conn, err := d.Dial("tcp", host+":25")

		if err == nil {
			conn.Close()

			return true, nil
		}
	}

	return false, nil
}

func Check(email string) (bool, error) {
	if !CheckRegex(email) {
		return false, nil
	}

	return CheckConnectivity(email)
}
