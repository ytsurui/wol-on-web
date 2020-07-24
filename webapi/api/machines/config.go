package machines

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type apiconfig struct {
	HttpPort int           `json:"httpport"`
	ReadOnly bool          `json:"readonly"`
	Machines []MachineInfo `json:"machines"`
}

var apiConfigData apiconfig
var maxID int

func ReadFile(fileName string) (bytes []byte, err error) {
	bytes, err = ioutil.ReadFile(fileName)
	if err != nil {
		bytes = nil
		return
	}
	return
}

func WriteFile(fileName string, writedata []byte) error {
	if apiConfigData.ReadOnly {
		err := errors.New("[Warning] Read Only Config")
		return err
	}

	err := ioutil.WriteFile(fileName, writedata, os.ModePerm)
	if err != nil {
		return (err)
	}
	return nil
}

func InitConfig() {
	var jsonFile string
	jsonFile = "config.json"

	configData, err := ReadFile(jsonFile)
	if err != nil {
		fmt.Println("[InitConfig] Config File Not Found")
		apiConfigData.ReadOnly = false
		apiConfigData.Machines = make([]MachineInfo, 0)
		return
	}

	err = json.Unmarshal(configData, &apiConfigData)
	if err != nil {
		fmt.Println("[InitConfig] Config JSON decode error, err=", err)
		apiConfigData.ReadOnly = true
		apiConfigData.Machines = make([]MachineInfo, 0)
		return
	}

	maxID = len(apiConfigData.Machines)
	for _, v := range apiConfigData.Machines {
		if v.ID > maxID {
			maxID = v.ID
		}
	}
}

func writeConfigData() (errInfo error) {
	var jsonFile string
	jsonFile = "config.json"

	writeBytes, err := json.MarshalIndent(apiConfigData, "", "  ")
	if err != nil {
		return err
	}

	err = WriteFile(jsonFile, writeBytes)
	if err != nil {
		return (err)
	}

	return nil
}

func getMachineListData() []MachineInfo {
	return (apiConfigData.Machines)
}

func getMachineItemData(id int) (machineData *MachineInfo, errInfo error) {
	for _, v := range apiConfigData.Machines {
		if v.ID == id {
			return &v, nil
		}
	}

	return nil, errors.New("Machine Not Found")
}

func writeMachineItemData(id int, data MachineInfo) (errInfo error) {
	if apiConfigData.ReadOnly {
		errInfo = errors.New("[Warning] Read Only Config")
		return
	}

	if id == 0 {
		maxID++
		data.ID = maxID
		apiConfigData.Machines = append(apiConfigData.Machines, data)
	} else {
		for i, v := range apiConfigData.Machines {
			if v.ID == id {
				data.ID = id
				apiConfigData.Machines[i] = data
			}
		}
	}

	err := writeConfigData()
	return err
}

func deleteMachineItemData(id int) (errInfo error) {
	if apiConfigData.ReadOnly {
		errInfo = errors.New("[Warning] Read Only Config")
		return
	}

	for i, v := range apiConfigData.Machines {
		if v.ID == id {
			// Delete
			newMachines := make([]MachineInfo, 0)
			newMachines = append(apiConfigData.Machines[:i], apiConfigData.Machines[i+1:]...)

			apiConfigData.Machines = newMachines
			err := writeConfigData()
			return err
		}
	}

	err := writeConfigData()
	return err
}

func GetHttpPortNum() int {
	if apiConfigData.HttpPort <= 0 || apiConfigData.HttpPort > 65535 {
		return (80)
	}
	return apiConfigData.HttpPort
}
