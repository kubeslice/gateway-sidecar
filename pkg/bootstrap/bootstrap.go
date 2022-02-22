package bootstrap

import (
	"io/ioutil"
	"os"
	"bitbucket.org/realtimeai/kubeslice-gw-sidecar/pkg/logger"
)

const (
	SERVER string = "SERVER"
	CLIENT string = "CLIENT" 
)

type GatewayPod struct{
	mode 	  string
	mountPath string
	secretMountPath string
	logger    *logger.Logger
}

func NewGatewayPod(mode string,mountPath,secretMountPath string ,logger *logger.Logger) (*GatewayPod){
	return &GatewayPod{
		mode: mode,
		mountPath: mountPath,
		secretMountPath: secretMountPath,
		logger: logger,
	}
}

// Process() creates a directory structure as required by openvpn pods
func (gw *GatewayPod) Process() error {
	baseFileName := os.Getenv("CLUSTER_ID") +"-" +  os.Getenv("SLICE_NAME") + "-1.vpn.aveshasystems.com"
	if gw.mode == SERVER {
		//create two directories named pki and ccd in /mountPath (eg: /config/pki) if not exists
		present,err := exists(gw.mountPath+"/pki")
		if err != nil {
			return err
		} 
		if !present{
			err = os.Mkdir(gw.mountPath+"/pki",0755)
			if err != nil {
				return err
			}
		}

		present,err = exists(gw.mountPath+"/ccd")
		if err != nil {
			return err
		} 
		if !present {
			err = os.Mkdir(gw.mountPath+"/ccd",0755)
			if err != nil {
				return err
			}
		}

		// create sub-directories "issued" and "private" in "/mountPath/pki"
		present,err = exists(gw.mountPath+"/pki/"+"issued")
		if err != nil {
			return err
		} 
		if !present {
			err = os.Mkdir(gw.mountPath+"/pki/"+"issued",0755)
			if err != nil {
				return err
			}
		}
		
		present,err = exists(gw.mountPath+"/pki/"+"private")
		if err != nil {
			return err
		} 
		if !present {
			err = os.Mkdir(gw.mountPath+"/pki/"+"private",0755)
			if err != nil {
				return err
			}
		}
		//Copy these files from /secretMountPath/* to /mountPath/*
		openVpnConfFileName := "openvpn.conf"
		crtFileName := baseFileName + ".crt"
		keyFileName := baseFileName + ".key"
		takeyFileName := baseFileName + "-ta.key"
		ccdFileName := "slice-" + os.Getenv("SLICE_NAME")
		files := map[string]string{
			"ovpnConfigFile": openVpnConfFileName,
			"pkiCACertFile":"pki/" + "ca.crt",
			"pkiDhPemFile":"pki/" + "dh.pem",
			"pkiTAKeyFile": "pki/" + takeyFileName,
			"pkiIssuedCertFile": "pki/issued/" + crtFileName,
			"pkiPrivateKeyFile": "pki/private/" + keyFileName,
			"ccdFile" : "ccd/" + ccdFileName,
		}
		for source,dest := range files{
			sourceFile := gw.secretMountPath + "/" + source
			destinationFile := gw.mountPath +"/" + dest
			err = CopyFile(sourceFile,destinationFile)
			if err != nil {
				return err
			}
		}
		//create the ovpn_env.sh file in /config directory

		err = writeFile("config/ovpn_env.sh")
		if err != nil {
			return err
		}

	}else {
		//Mount files for client
	}
	return nil
}

func exists(path string) (bool,error) {
	_,err := os.Stat(path)
	if err == nil {
		return true,nil
	}
	if os.IsNotExist(err){
		return false,nil
	}
	return false,err
}
func CopyFile(source string, dest string) (error) {
	bytesRead, err := ioutil.ReadFile(source)
	if err != nil {
		return err
    }

    err = ioutil.WriteFile(dest, bytesRead, 0755)
	if err != nil {
		return err
	}
	return err
}

func writeFile(source string) error {
	f, err := os.Create(source)

    if err != nil {
        return err
    }

	contents := `declare -x OVPN_DEFROUTE=0
declare -x OVPN_DEVICE=tun
declare -x OVPN_NAT=0`
	_, err = f.WriteString(contents)

    if err != nil {
        return err
    }

    defer f.Close()

	return err

}