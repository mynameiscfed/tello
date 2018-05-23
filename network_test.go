// tello project network_test.go

// Copyright (C) 2018  Steve Merrony

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.package memory

package tello

import (
	"log"
	"testing"
	"time"
)

// use go test -count=1 to bypass test caching

func TestControlConnectDisconnect(t *testing.T) {
	conn, err := ControlConnectDefault()
	if err != nil {
		log.Fatalf("CCD failed with error %v", err)
	}
	log.Println("Connected to Tello")
	stop := make(chan bool, 2)
	go ControlResponseListener(conn, stop)
	time.Sleep(10 * time.Second)
	stop <- true
	ControlDisconnect(conn)
	log.Println("Disconnected normally from Tello")
}
