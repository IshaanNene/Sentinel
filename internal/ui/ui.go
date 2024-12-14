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

			output := fmt.Sprintf(
				"[cyan]HACKER TERMINAL - SYSTEM MONITOR[white]\n\n"+
				"[yellow]CPU USAGE:[white] %.2f%% (CORES: %d)\n"+
				"[yellow]MEMORY USAGE:[white] %.2f%% (%s / %s)\n"+
				"[yellow]DISK USAGE:[white] %.2f%% (%s / %s)\n\n",
				cpuUsage, cpuCount, memUsage, utils.FormatBytes(memUsed), utils.FormatBytes(memTotal), diskUsage, utils.FormatBytes(diskUsed), utils.FormatBytes(diskTotal),
			)

			output += "[yellow]NETWORK STATS:[white]\n"
			for _, stat := range netStats {
				output += fmt.Sprintf("  [green]INTERFACE:[white] %s | [blue]SENT:[white] %s | [blue]RECEIVED:[white] %s\n",
					stat.Name, utils.FormatBytes(stat.BytesSent), utils.FormatBytes(stat.BytesRecv))
			}

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
