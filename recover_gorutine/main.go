package main

import (
	"log"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)

		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or(append(channels[3:], orDone)...):
			}
		}
	}()
	return orDone
}

type startGorutineFn func(
	done <-chan interface{}, pulseInterval time.Duration,
) (heartbeat <-chan interface{})

func newSteward(
	timeout time.Duration,
	startGorutine startGorutineFn,
) startGorutineFn {
	return func(
		done <-chan interface{}, pulseInterval time.Duration,
	) <-chan interface{} {
		heartbeat := make(chan interface{})
		go func() {
			defer close(heartbeat)

			var wardDone chan interface{}
			var wardHeartbeat <-chan interface{}

			startWard := func() {
				wardDone = make(chan interface{})
				wardHeartbeat = startGorutine(or(wardDone, done), timeout/2)
			}
			startWard()
			pulse := time.Tick(pulseInterval)

		monitorL:
			for {
				timeoutSignal := time.After(timeout)
				for {
					select {
					case <-pulse:
						select {
						case heartbeat <- struct{}{}:
						default:
						}
					case <-wardHeartbeat:
						continue monitorL
					case <-timeoutSignal:
						log.Println("ward unhealthy; restarting")
						close(wardDone)
						startWard()
						continue monitorL
					case <-done:
						return
					}
				}
			}
		}()

		return heartbeat
	}
}

func main() {
	doWork := func(done <-chan interface{}, _ time.Duration) <-chan interface{} {
		log.Println("ward: Hello, I am irresponsible!")
		go func() {
			<-done
			log.Println("ward: I am halting")
		}()
		return nil
	}

	doWorkWithSteaward := newSteward(4*time.Second, doWork)

	done := make(chan interface{})
	time.AfterFunc(9*time.Second, func() {
		log.Println("halting stward and wad")
		close(done)
	})

	for range doWorkWithSteaward(done, 4*time.Second) {
	}
	log.Println("done")
}
