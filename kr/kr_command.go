package main

import (
	"encoding/base64"
	"os"
	"strconv"

	"github.com/kryptco/kr"
	"github.com/kryptco/kr/krdclient"
	"github.com/urfave/cli"

	"net/mail"
)

func setTeamNameCommand(c *cli.Context) (err error) {
	_, err = krdclient.RequestMe()
	if err != nil {
		PrintFatal(os.Stderr, kr.Red("Kryptonite ▶ "+err.Error()))
		return
	}

	name := c.String("name")
	if name == "" {
		PrintFatal(os.Stderr, "--name flag required")
	}
	kr.SetTeamName(name)
	return
}

func createInviteCommand(c *cli.Context) (err error) {

	emails := c.StringSlice("emails")
	domain := c.String("domain")
	if domain != "" {

		domain := c.String("domain")
		if domain == "" {
			PrintFatal(os.Stderr, "Supply a valid email domain")
			return
		}
		_, err := mail.ParseAddress("x@" + domain)
		if err != nil {
			PrintFatal(os.Stderr, "Email domain "+domain+" is invalid.")
			return err
		}

		kr.InviteDomain(domain)

	} else if len(emails) > 0 {
		var verifed_addresses []string

		for _, email := range emails {
			address, err := mail.ParseAddress(email)
			if err != nil {
				PrintFatal(os.Stderr, "Email "+email+" is invalid.")
				return err
			}
			verifed_addresses = append(verifed_addresses, address.Address)
		}

		kr.InviteEmails(verifed_addresses)

	} else {
		PrintFatal(os.Stderr, "--emails or --domain are required")
	}

	return
}

func closeInvitationsCommand(c *cli.Context) (err error) {
	kr.CancelInvite()
	return
}

func setPolicyCommand(c *cli.Context) (err error) {
	var window *int64
	if c.String("window") != "" {
		windowInt, err := strconv.ParseInt(c.String("window"), 10, 64)
		if err != nil {
			PrintFatal(os.Stderr, "Could not parse window as integer seconds: "+err.Error())
		}
		window = &windowInt
	}
	kr.SetApprovalWindow(window)
	return
}

func getMembersCommand(c *cli.Context) (err error) {
	var query *string
	if c.String("query") != "" {
		queryStr := c.String("query")
		query = &queryStr
	}
	kr.GetMembers(query, c.Bool("ssh"), c.Bool("pgp"))
	return
}

func addAdminCommand(c *cli.Context) (err error) {
	email := c.String("email")
	if email == "" {
		PrintFatal(os.Stderr, "--email required")
	}
	kr.AddAdmin(email)
	return
}

func removeAdminCommand(c *cli.Context) (err error) {
	email := c.String("email")
	if email == "" {
		PrintFatal(os.Stderr, "--email required")
	}
	kr.RemoveAdmin(email)
	return
}

func getAdminsCommand(c *cli.Context) (err error) {
	kr.GetAdmins()
	return
}

func pinHostKeyCommand(c *cli.Context) (err error) {
	if c.String("public-key") == "" {
		kr.PinKnownHostKeys(c.String("host"), c.Bool("update-from-server"))
		return
	}
	pk, err := base64.StdEncoding.DecodeString(c.String("public-key"))
	if err != nil {
		PrintFatal(os.Stderr, "error decoding public-key, make sure it is base64 encoded without the key type prefix (i.e. no 'ssh-rsa' or 'ssh-ed25519') "+err.Error())
	}
	kr.PinHostKey(c.String("host"), pk)
	return
}

func unpinHostKeyCommand(c *cli.Context) (err error) {
	pk, err := base64.StdEncoding.DecodeString(c.String("public-key"))
	if err != nil {
		PrintFatal(os.Stderr, "error decoding public-key, make sure it is base64 encoded without the key type prefix (i.e. no 'ssh-rsa' or 'ssh-ed25519') "+err.Error())
	}
	kr.UnpinHostKey(c.String("host"), pk)
	return
}

func listPinnedKeysCommand(c *cli.Context) (err error) {
	host := c.String("host")
	if host == "" {
		kr.GetAllPinnedHostKeys()
		return
	}
	kr.GetPinnedHostKeys(host, c.Bool("search"))
	return
}

func enableLoggingCommand(c *cli.Context) (err error) {
	kr.EnableLogging()
	return
}

func logsCommand(c *cli.Context) (err error) {
	kr.UpdateTeamLogs()
	return
}

func teamBillingCommand(c *cli.Context) (err error) {
	kr.OpenBilling()
	return
}
