package scan

import (
	"fmt"
	"path/filepath"
	"toolscan/core/model"
	"toolscan/core/utils"
)

// ---------------------------------------------------------------------------------------
// NMAP INTERACTION
// ---------------------------------------------------------------------------------------
type ArpScan model.Scan

// Constructor for ArpScan
func NewArpScan(name, target, folder, file, arpArgs string) *ArpScan {
	// Create a Scan
	s := &ArpScan{
		Name:   name,
		Target: target,
		Status: model.NOT_STARTED,
	}
	// Construct output path and create if it doesn't exist
	s.Outfolder = filepath.Join(utils.Config.Outfolder, utils.CleanPath(target), folder)
	s.Outfile = filepath.Join(s.Outfolder, utils.CleanPath(file))
	utils.EnsureDir(s.Outfolder)
	// Construct command
	s.Cmd = s.constructCmd(arpArgs)
	return s
}

func (s *ArpScan) constructCmd(args string) string {
	return fmt.Sprintf("arp-scan %s %s -oA %s", args, s.Target, s.Outfile)
}
