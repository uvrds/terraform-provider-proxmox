package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
	log "github.com/sirupsen/logrus"
	"github.com/terraform-provider-proxmox/pkg/provider"
	"os"
)

var pack = "main"

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

//func main() {
//	//cookie, err := model.Auth()
//	//if err != nil {
//	//	log.WithFields(log.Fields{
//	//		"package":  pack,
//	//		"function": "Req",
//	//		"error":    err,
//	//		"data":     nil,
//	//	}).Fatal("Response", err)
//	//}
//	//
//	//fmt.Println(cookie)
//	//
//		BaseURL := "https://192.168.122.54:8006/api2/json"
//		Username := "root@pam"
//		Password := "asdqz123"
//		t := client.NewClient(BaseURL, Username, Password)
//		tr := &http.Transport{
//			MaxIdleConns:       10,
//			IdleConnTimeout:    30 * time.Second,
//			DisableCompression: true,
//			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
//		}
//		t.Client = &http.Client{Transport: tr}
//		err := t.Authenticate()
//		//err = t.GetStatus("pve", "101")
//		err = t.CreateLxc("pve")
//		fmt.Println(err)
//	}

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return provider.Provider()
		},
	})
}
