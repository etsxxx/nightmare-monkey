package nightmare

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/mitchellh/go-ps"
)

type KillInput struct {
	Desc string `json:"desc"`
}

type KilledProcess struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	PID  int    `json:"pid"`
	PPID int    `json:"ppid"`
	//Registerd int64 `json:""`
}

func pickupProcess() (process ps.Process, err error) {
	processes, err := ps.Processes()
	if err != nil {
		return nil, err
	}

	for i := 0; i < 5; i++ {
		process = processes[ran.Intn(len(processes))]
		if process.Pid() <= 1 {
			continue
		} else if (process.Pid() == os.Getpid()) || process.PPid() == os.Getpid() {
			continue
		}
		return process, nil
	}

	return nil, fmt.Errorf("cannot find any process to kill")
}

func (monkey *NightmareMonkey) killRandomProcess(desc string) (process *KilledProcess, err error) {
	target, err := pickupProcess()
	if err != nil {
		return nil, err
	}

	process = &KilledProcess{
		Desc: desc,
		Name: target.Executable(),
		PID:  target.Pid(),
		PPID: target.PPid(),
	}
	if !monkey.Dryrun {
		ps, err := os.FindProcess(target.Pid())
		if err != nil {
			return process, err
		}

		err = ps.Signal(os.Kill)
		if err != nil {
			return process, err
		}
	}

	return process, nil
}

func (monkey *NightmareMonkey) kill(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	//now := time.Now()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		monkey.Logger.Printf("error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "something went wrong",
		})
		return
	}

	input := KillInput{}
	err = json.Unmarshal(body, &input)
	if err != nil {
		monkey.Logger.Printf("error: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "format error",
		})
		return
	}

	if monkey.isSleeping() {
		// TODO: change error code
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Nightmare Monkey is sleeping now",
		})
		return
	}

	killed, err := monkey.killRandomProcess(input.Desc)
	if err != nil {
		if killed != nil {
			monkey.Logger.Printf("error: attempt to kill %s(PID:%d PPID:%d) and got error: %s", killed.Name, killed.PID, killed.PPID, err)
		} else {
			monkey.Logger.Printf("error: %s", err)
		}

		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	if monkey.Dryrun {
		monkey.Logger.Printf("(Dryrun) Kill: %s(PID:%d PPID:%d) - %s", killed.Name, killed.PID, killed.PPID, killed.Desc)
	} else {
		monkey.Logger.Printf("Kill: %s(PID:%d PPID:%d) - %s", killed.Name, killed.PID, killed.PPID, killed.Desc)
	}
	w.WriteHeader(http.StatusOK)
	encoder.Encode(killed)
	return
}
