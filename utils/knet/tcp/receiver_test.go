// expect：be sure to finish!
// author：KercyLAN
// create at：2020-3-5 10:29

package tcp

import "testing"

func TestReceiver_Accept(t *testing.T) {
	tcpReceiver := NewReceiver(":1998")
	if err := tcpReceiver.Accept(); err != nil {
		t.Log("application exit.")
	}
}
