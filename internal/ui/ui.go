package ui

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
	"goscope/internal/monitor"
	"goscope/pkg/utils"
)

func StartUI() {
	app := tview.NewApplication()
	textView := tview.NewTextView().SetDynamicColors(true).SetTextAlign(tview.AlignLeft)

	go func() {
		for {
			cpuUsage, _ := monitor.GetCPUUsage()
			cpuCount, _ := monitor.GetCPUCount()
			memUsage, memUsed, _ := monitor.GetMemoryUsage()
			memTotal, _ := monitor.GetMemoryTotal()
			diskUsage, diskUsed, _ := monitor.GetDiskUsage()
			diskTotal, _ := monitor.GetDiskTotal()
			netStats, _ := monitor.GetNetworkStats()
			swapTotal, _ := monitor.GetMemorySwapTotal()
			swapUsed, _ := monitor.GetMemorySwapUsed()
			diskFree, _ := monitor.GetDiskFree()
			memFree, _ := monitor.GetMemoryFree()
			cpuIdle, _ := monitor.GetCPUIdle()
			diskInodesUsed, _ := monitor.GetDiskInodesUsed()
			diskInodesFree, _ := monitor.GetDiskInodesFree()
			cpuFrequencyInfo, _ := monitor.GetCPUFrequency()
			cpuFrequency := cpuFrequencyInfo[len(cpuFrequencyInfo)-1]
			cpuModelName := cpuFrequency.ModelName
			cpuMhz := cpuFrequency.Mhz
			diskReadBytes, _ := monitor.GetDiskReadBytes()
			diskWriteBytes, _ := monitor.GetDiskWriteBytes()

			output := fmt.Sprintf(
				"[cyan]HACKER TERMINAL - SYSTEM MONITOR[white]\n\n"+
				"[yellow]CPU USAGE:[white] %.2f%% (CORES: %d)\n"+
				"[yellow]MEMORY USAGE:[white] %.2f%% (%s / %s)\n"+
				"[yellow]DISK USAGE:[white] %.2f%% (%s / %s)\n"+
				"[yellow]SWAP USAGE:[white] %.2f%% (%s / %s)\n"+
				"[yellow]DISK FREE:[white] %s\n"+
				"[yellow]MEMORY FREE:[white] %s\n"+
				"[yellow]CPU IDLE:[white] %.2f%%\n"+
				"[yellow]DISK INODES:[white] Used: %s, Free: %s\n"+
				"[yellow]NETWORK STATS:[white]\n",
				cpuUsage, cpuCount, memUsage, utils.FormatBytes(memUsed), utils.FormatBytes(memTotal),
					diskUsage, utils.FormatBytes(diskUsed), utils.FormatBytes(diskTotal),
					(float64(swapUsed)/float64(swapTotal))*100, utils.FormatBytes(swapUsed), utils.FormatBytes(swapTotal),
					utils.FormatBytes(diskFree), utils.FormatBytes(memFree),
					cpuIdle, utils.FormatBytes(diskInodesUsed), utils.FormatBytes(diskInodesFree),
			)

			for _, stat := range netStats {
				output += fmt.Sprintf("  [green]INTERFACE:[white] %s | [blue]SENT:[white] %s | [blue]RECEIVED:[white] %s\n",
					stat.Name, utils.FormatBytes(stat.BytesSent), utils.FormatBytes(stat.BytesRecv))
			}

			output += fmt.Sprintf(
				"[yellow]DISK READ:[white] %s\n"+
				"[yellow]DISK WRITE:[white] %s\n"+
				"[yellow]CPU FREQUENCY:[white] %s (%.2f MHz)\n\n",
				utils.FormatBytes(diskReadBytes), utils.FormatBytes(diskWriteBytes),
				cpuModelName, cpuMhz,
			)

			output += "\n[red]UPDATING EVERY SECOND... STAY ALERT!\n"

			app.QueueUpdateDraw(func() {
				textView.SetText(output)
			})
			time.Sleep(1 * time.Second)
		}
	}()

	if err := app.SetRoot(textView, true).Run(); err != nil {
		panic(err)
	}
}
