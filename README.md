# Sentinel

**Sentinel** is a real-time system monitoring tool built with Go. It provides a terminal-based user interface (TUI) to monitor vital system metrics such as CPU usage, memory consumption, disk space, and network activity. **Sentinel** offers live updates and detailed insights into your system’s health, making it a perfect tool for developers, sysadmins, and anyone interested in tracking their system performance in real-time.

## Features

- **Real-Time Monitoring**: Get live updates on CPU, memory, disk, and network usage.
- **Terminal User Interface (TUI)**: Clean, dynamic, and intuitive interface for quick insights.
- **Cross-Platform**: Works on Linux, macOS, and Windows (via WSL).
- **Lightweight and Fast**: Built with Go, ensuring minimal resource usage while providing accurate, up-to-date metrics.
- **Customizable**: Easy to extend and modify for additional system metrics.

## Technologies Used

- **Go**: The primary language for building the application.
- **tview**: A package for building TUI applications in Go.
- **gopsutil**: A Go library for fetching system metrics like CPU, memory, and disk usage.
- **Terminal UI**: Responsive, dynamic terminal interface for real-time data visualization.

## Installation

### Prerequisites

- Go 1.20 or later installed on your system.

### Steps to Install

1. Clone the repository:
   ```bash
   git clone https://github.com/<your-username>/sentinel.git
   cd sentinel
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run main.go
   ```
Once the application runs, you'll see a terminal-based UI displaying live system metrics.

## Usage

- **CPU**: Displays real-time CPU usage and load.
- **Memory**: Shows the current memory usage, including used and available memory.
- **Disk**: Monitors disk space usage and free space.
- **Network**: Provides network activity stats, including bytes sent and received.

## Screenshots
<img width="388" alt="Screenshot 2024-12-14 at 10 10 11 AM" src="https://github.com/user-attachments/assets/ae1e8888-bcda-40ee-a70a-2594cc6a5ceb" />
## Future Enhancements
- **Alerting System**: Notify users when system metrics exceed predefined thresholds.
- **Export Feature**: Export system stats to CSV or JSON format.
- **Support for More Metrics**: Add support for additional metrics like GPU usage, temperature, etc.
- **Docker Support**: Containerize the application for easier deployment and usage in Docker environments.

## Contributing
I welcome contributions! If you'd like to contribute, please follow the steps below:

1. Fork the repository.
2. Create a new branch for your changes (`git checkout -b feature-branch`).
3. Make your changes and commit them (`git commit -am 'Add new feature'`).
4. Push to your fork (`git push origin feature-branch`).
5. Submit a pull request.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
