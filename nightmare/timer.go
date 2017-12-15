package nightmare

import "time"

func (monkey *NightmareMonkey) Nightmare() {
	go func() {
		for {
			next := ran.Int63n(int64(monkey.Interval)) + 1
			monkey.Logger.Printf("Next: %d seconds after", next)
			t := time.NewTicker(time.Duration(next) * time.Second)

			select {
			case <-t.C:
				t.Stop()
				if monkey.isSleeping() {
					monkey.Logger.Printf("Nightmare Monkey is sleeping with good dream.")
				} else {
					killed, err := monkey.killRandomProcess("by monkey")
					if err != nil {
						if killed != nil {
							monkey.Logger.Printf("error: attempt to kill %s(PID:%d PPID:%d) and got error: %s", killed.Name, killed.PID, killed.PPID, err)
						} else {
							monkey.Logger.Printf("error: %s", err)
						}
					} else {
						if monkey.Dryrun {
							monkey.Logger.Printf("(Dryrun) Kill: %s(PID:%d PPID:%d) - %s", killed.Name, killed.PID, killed.PPID, killed.Desc)
						} else {
							monkey.Logger.Printf("Kill: %s(PID:%d PPID:%d) - %s", killed.Name, killed.PID, killed.PPID, killed.Desc)
						}
					}
				}
			}

		}
	}()
}
