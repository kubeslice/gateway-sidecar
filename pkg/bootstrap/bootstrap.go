package bootstrap

import (
	"os"
	"io"
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
	baseFileName := os.Getenv("CLUSTER_ID") + os.Getenv("SLICE_NAME")
	if gw.mode == SERVER {
		//create two directories named pki and ccd in /mountPath (eg: /config/pki)
		err := os.Mkdir(gw.mountPath+"/pki",0755)
		if err != nil {
			return err
		}
		err = os.Mkdir(gw.mountPath+"/ccd",0755)
		if err != nil {
			return err
		}

		//create two files named openvpn.conf and copy the contents from source file in secretMountPath
		files:= []string{"openvpn.conf"}

		for _,file := range files{
			sourceFile := gw.secretMountPath + "/" + file
			destinationFile := gw.mountPath + "/" + file
			err = CopyFile(sourceFile,destinationFile)
			if err != nil {
				return err
			}
		}

		// create sub-directories "issued" and "private"
		err = os.Mkdir(gw.mountPath+"/pki/"+"issued",0755)
		if err != nil {
			return err
		}
		err = os.Mkdir(gw.mountPath+"/pki/"+"private",0755)
		if err != nil {
			return err
		}
		//TODO:Discuss on the last fileName
		//Copy these 3 files from /secretMountPath to /mountPath/pki/
		files = []string{"ca.crt","dh.pem",baseFileName+"-ta"+".key"}
		for _,file := range files{
			sourceFile := gw.secretMountPath + "/" + file
			destinationFile := gw.mountPath + "/pki/" + file
			err = CopyFile(sourceFile,destinationFile)
			if err != nil {
				return err
			}
		}

		//copy the .crt file in /mountPath/pki/issued
		crtFileName := baseFileName + ".crt"
		sourceFile := gw.secretMountPath + "/" + crtFileName
		destinationFile := gw.mountPath + "/pki/issued/" + crtFileName
		err = CopyFile(sourceFile,destinationFile)
		if err != nil {
			return err
		}

		//copy the .crt file in /mountPath/pki/private
		keyFileName := baseFileName + ".key"
		sourceFile = gw.secretMountPath + "/" + keyFileName
		destinationFile = gw.mountPath + "/pki/private/" + keyFileName
		err = CopyFile(sourceFile,destinationFile)
		if err != nil {
			return err
		}
		//TODO: copy file in /ccd directory
	}
	return nil
}

func CopyFile(source string, dest string) (error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			_ = os.Chmod(dest, sourceinfo.Mode())
		}

	}
	return nil
} 