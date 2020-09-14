package memory

import "time"

// HostNotIdle indicates that user requests are being responded to and background file updates
// should be paused to avoid affecting the player experience by using process time.
func HostNotIdle() { notIdle <- signal{} }

const WaitUntilIdleTimeout = 200 * time.Millisecond

// WaitUntilIdle stops the calling goroutine for the WaitUntilIdleTimeout duration if HostNotIdle
// has been called recently. This function is used to slow down the expensive computer player move
// calculation while much shorter player requests are being done.
func WaitUntilIdle() {

}

const idleTimeSeconds = 1

var notIdle = make(chan signal, 8)

func continuouslySignalIdle() <-chan signal {
	idle := make(chan signal)
	timer := time.NewTimer(time.Second * idleTimeSeconds)
	go func() {
		for {
			select {
			case <-timer.C:
				idle <- signal{}
			case <-notIdle:
				if timer.Stop() == false {
					<-timer.C
				}
			}
			timer.Reset(time.Second * IdleTimeSeconds)
		}
	}()
	return idle
}
