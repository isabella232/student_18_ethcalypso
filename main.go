package main

import (
	"io"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/dedis/cothority"
	"github.com/dedis/cothority/byzcoin/bcadmin/lib"
	"github.com/dedis/cothority/darc"
	"github.com/dedis/onet/log"
	"github.com/dedis/onet/network"
	"github.com/dedis/student_18_ethcalypso/calypso"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	cli "gopkg.in/urfave/cli.v1"
)

var cliApp = cli.NewApp()

var cmds = cli.Commands{
	{
		Name:    "LTSID",
		Usage:   "create a ledger",
		Aliases: []string{"c"},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "roster, r",
				Usage: "the roster of the cothority that will host the ledger",
			},
			cli.StringFlag{
				Name:  "PrivateKey, pk",
				Usage: "The privatekey that you will use to interact with ethereum",
			},
			cli.StringFlag{
				Name:  "Calypso",
				Usage: "The Calypso contract you will interact with",
			},
		},
		Action: LTSID,
	},
	{
		Name:    "AddWrite",
		Usage:   "Add write request and returns the address",
		Aliases: []string{"s"},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "LTSID",
				Usage: "the ByzCoin config to use",
			},
			cli.StringFlag{
				Name:  "X",
				Usage: "The X variable of a writeRequest",
			},
			cli.StringFlag{
				Name:  "roster, r",
				Usage: "the roster of the cothority that will host the ledger",
			},
			cli.StringFlag{
				Name:  "PrivateKey, pk",
				Usage: "The privatekey that you will use to interact with ethereum",
			},
			cli.StringFlag{
				Name:  "Calypso",
				Usage: "The Calypso contract you will interact with",
			},
		},
		Action: AddWrite,
	},
	{
		Name:    "AddRead",
		Usage:   "Add Read request and returns the address and the secret",
		Aliases: []string{"s"},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "wr",
				Usage: "Address of the write request",
			},
			cli.StringFlag{
				Name:  "roster, r",
				Usage: "the roster of the cothority that will host the ledger",
			},
			cli.StringFlag{
				Name:  "PrivateKey, pk",
				Usage: "The privatekey that you will use to interact with ethereum",
			},
			cli.StringFlag{
				Name:  "Calypso",
				Usage: "The Calypso contract you will interact with",
			},
		},
		Action: AddRead,
	},
	{
		Name:    "DecryptKey",
		Usage:   "Decrypts the key",
		Aliases: []string{"s"},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "wr",
				Usage: "The address of the write request",
			},
			cli.StringFlag{
				Name:  "rr",
				Usage: "The address of the read request",
			},
			cli.StringFlag{
				Name:  "roster, r",
				Usage: "the roster of the cothority that will host the ledger",
			},
		},
		Action: DecryptKey,
	},
}

func init() {
	network.RegisterMessages(&darc.Darc{}, &darc.Identity{}, &darc.Signer{})
}

func LTSID(c *cli.Context) error {
	fn := c.String("roster")
	r, err := lib.ReadRoster(fn)
	if err != nil {
		return err
	}
	roster := r.NewRosterWithRoot(r.List[0])
	pk := c.String("pk")
	privateKey, e := crypto.HexToECDSA(pk)
	if e != nil {
		log.Fatal(e)
	}
	cal := c.String("Calypso")
	calAddr := common.HexToAddress(cal)
	calypsoClient := calypso.NewClient(*roster, privateKey)
	LTS, e := calypsoClient.CreateLTS(calAddr)
	if e != nil {
		fmt.Println("Can't get LTS")
		fmt.Println(e)
		return e
	}
	hexLTSID := hex.EncodeToString(LTS.LTSID)
	fmt.Println("LTSID: ", hexLTSID)
	fmt.Println("X: ", LTS.X.String())
	return nil

}

func init() {
	cliApp.Name = "bcadmin"
	cliApp.Usage = "Create eth whatever"
	cliApp.Version = "0.1"
	cliApp.Commands = cmds
	cliApp.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "debug, d",
			Value: 0,
			Usage: "debug-level: 1 for terse, 5 for maximal",
		},
	}
	cliApp.Before = func(c *cli.Context) error {
		log.SetDebugVisible(c.Int("debug"))
		lib.ConfigPath = c.String("config")
		return nil
	}
}

func DecryptKey(c *cli.Context) {
	/*
		wr := c.String("wr")
		rr := c.String("rr")
		wrAddr := common.HexToAddress(wr)
		rrAddr := common.HexToAddress(rr)
		fn := c.String("roster")
		r, err := lib.ReadRoster(fn)
		if err != nil {
			log.Fatal(err)
		}
		roster := r.NewRosterWithRoot(r.List[0])
		calypsoClient := calypso.NewClient(*roster)
		dk := &calypso.DecryptKey{
			Write: wrAddr,
			Read:  rrAddr,
		}
		dkr, e := calypsoClient.DecryptKey(dk)
		if e != nil {
			log.Fatal(e)
		}
		fmt.Println("Xhat: ", dkr.XhatEnc.String())
		fmt.Println("Cs", dkr.Cs)
		fmt.Println("X is ", dkr.X.String())
	*/
}

func AddRead(c *cli.Context) {
	wr := c.String("wr")
	wrAddr := common.HexToAddress(wr)
	fn := c.String("roster")
	r, err := lib.ReadRoster(fn)
	if err != nil {
		log.Fatal(err)
	}
	roster := r.NewRosterWithRoot(r.List[0])
	fmt.Println("Roster ", roster)
	pk := c.String("pk")
	privateKey, e := crypto.HexToECDSA(pk)
	if e != nil {
		log.Fatal(e)
	}
	cal := c.String("Calypso")
	calAddr := common.HexToAddress(cal)
	calypsoClient := calypso.NewClient(*roster, privateKey)
	secret, rAddr, e := calypsoClient.AddRead(wrAddr, calAddr)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println("Your secret is: ", secret)
	fmt.Println("Read address is ", rAddr.Hex())
}

func AddWrite(c *cli.Context) {
	ltsid := c.String("LTSID")
	X := c.String("X")
	xHex, e := hex.DecodeString(X)
	if e != nil {
		log.Fatal(e)
	}
	pk := c.String("pk")
	privateKey, e := crypto.HexToECDSA(pk)
	if e != nil {
		log.Fatal(e)
	}
	cal := c.String("Calypso")
	calAddr := common.HexToAddress(cal)
	point := cothority.Suite.Point()
	e = point.UnmarshalBinary(xHex)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println("LTSID: ", ltsid)
	fmt.Println("Value of X: ", X)
	fn := c.String("roster")
	r, err := lib.ReadRoster(fn)
	if err != nil {
		log.Fatal(err)
	}
	roster := r.NewRosterWithRoot(r.List[0])
	fmt.Println("Roster ", roster)
	calypsoClient := calypso.NewClient(*roster, privateKey)
	id, e := hex.DecodeString(ltsid)
	if e != nil {
		log.Fatal(e)
	}
	wrAddr, e := calypsoClient.AddWrite(id, point, []byte("Sabrina"), calAddr)
	fmt.Println("Address of write is: ", wrAddr.Hex())
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	http.ListenAndServe(":8080", nil)
	http.HandleFunc("/Bjorn", IndexWebsite)
}

var tpl *template.Template

func IndexWebsite(wr http.ResponseWriter, req *http.Request) {
	io.
}
