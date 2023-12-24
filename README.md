# Proxy Bridge
Proxy Bridge is a simple Golang program that acts as a proxy server, forwarding incoming requests from clients to a target server. The program listens for incoming connections on a specified port, and then forwards these connections to the target server using the TCP protocol.

## Installation
1. Clone the repository:
  - git clone https://github.com/wildyrando/proxy-bridge.git
2. Run the program:
  - go run main.go
 
By default, the program listens for incoming connections on port 8880 and forwards them to a server running on `127.0.0.1:22`. You can change these values by modifying the `host`, `port`, and `listen` variables in the code.

## License
This program is licensed under the GNU Public License V3. See the [LICENSE](LICENSE) file for more information.
