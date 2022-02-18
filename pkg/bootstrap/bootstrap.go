package bootstrap

import (
	"io"
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
	baseFileName := os.Getenv("CLUSTER_ID") +"-" +  os.Getenv("SLICE_NAME")
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
		ccdFileName := "slice-" + os.Getenv("SLICE_NAME")
		files := map[string]string{
			"ovpnConfigFile": openVpnConfFileName,
			"pkiCACertFile":"pki/" + "ca.crt",
			"pkiDhPemFile":"pki/" + "dh.pem",
			"pkiTAKeyFile": "pki/" + baseFileName +"-ta.key",
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
	var sourcefile,destfile *os.File
	var err error

	sourcefile, err = os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	//check if the destination file already exists , if not create it
	present,err := exists(dest)
	if err != nil {
		return err
	}
	if !present{
		destfile, err := os.Create(dest)
		if err != nil {
			return err
		}
		defer destfile.Close()
	} else {
		destfile ,err := os.Open(dest)
		if err != nil {
			return err
		}
		defer destfile.Close()
	}
	
	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			_ = os.Chmod(dest, sourceinfo.Mode())
		}

	}
	return nil
} 