/*

    Proxy Bridge is a simple Golang program that acts as a proxy server, 
    forwarding incoming requests from clients to a target server. The program listens for 
    incoming connections on a specified port, and then forwards these connections to the 
    target server using the TCP protocol.

    Author    : Wildy Sheverando
    Contact   : hai@wildyrando.com
    Git Repo  : https://github.com/wildyrando/Proxy-Bridge.git

    This project licensed under GNU Public License V3
    Link: GNU Public License V3

*/

package main

import (
    "io"
    "log"
    "net"
    "time"
)

func main() {
    // Configuration
    host := "127.0.0.1"
    port := "22"
    listen := "8880"

    // Listen new bridged port
    ln, err := net.Listen("tcp", ":"+listen)
    if err != nil {
        log.Fatal(err)
    }
    defer ln.Close() // Closing defer

    // Log for starting information
    log.Printf("Server started on port: %s\n", listen)
    log.Printf("Redirecting requests to: %s at port %s\n", host, port)

    // looping
    for {
        // Conn handle
        conn, err := ln.Accept()
        if err != nil {
            log.Println(err)
            continue
        }

        // Log Output
        log.Printf("Connection received from %s:%s\n", conn.RemoteAddr().Network(), conn.RemoteAddr().String())

        // create a func with c net connection to handle proxy
        go func(c net.Conn) {
            // Dial Target in tcp with time out
            dconn, err := net.DialTimeout("tcp", host+":"+port, 5*time.Second) // change to 5 seconds
            if err != nil {
                log.Println(err)
                return
            }
            
            // Return HTTP Response Switching Protocols to client
            _, err = c.Write([]byte("HTTP/1.1 101 Switching Protocols\r\nUpgrade: websocket\r\nConnection: Upgrade\r\n\r\n"))
            if err != nil {
                log.Println(err)
                return
            }

            // Copy client request to target (Forward)
            go func() {
                if _, err := io.Copy(dconn, c); err != nil {
                    log.Printf("Error copying data from client to destination server: %v\n", err)
                }
            }()

            // Copy Return from target to client
            if _, err := io.Copy(c, dconn); err != nil {
                log.Printf("Error copying data from destination server to client: %v\n", err)
            }

            // Log output
            log.Printf("Connection terminated for %s:%s\n", conn.RemoteAddr().Network(), conn.RemoteAddr().String())
        }(conn)
    }
}
