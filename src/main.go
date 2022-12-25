package main

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

func backupAwsConfig(homedir string) {

	tm := time.Now().Format("20060102150405")
	println(tm)

	backUpDIr := filepath.Join(homedir, ".go-get-sts-token-backup")

	if _, dirrErr := os.Stat(backUpDIr); os.IsNotExist(dirrErr) {
		if mkdirErr := os.Mkdir(backUpDIr, os.ModePerm); mkdirErr != nil {
			log.Fatal(mkdirErr)
		}
	}

	// backup config
	config_file := filepath.Join(homedir, ".aws", "config")
	config_r, err := os.Open(config_file)
	if err != nil {
		log.Fatal(err)
	}
	defer config_r.Close()

	var cfb_name string = "config-" + tm
	config_file_backup := filepath.Join(backUpDIr, cfb_name)
	config_w, err := os.Create(config_file_backup)
	if err != nil {
		log.Fatal(err)
	}
	defer config_w.Close()
	config_w.ReadFrom(config_r)

	// backup credentials
	credentials_file := filepath.Join(homedir, ".aws", "credentials")
	credentials_r, err := os.Open(credentials_file)
	if err != nil {
		log.Fatal(err)
	}
	defer credentials_r.Close()

	var crdfb_name string = "credentials-" + tm
	credentials_file_backup := filepath.Join(backUpDIr, crdfb_name)
	credentials_w, err := os.Create(credentials_file_backup)
	if err != nil {
		log.Fatal(err)
	}
	defer credentials_w.Close()
	credentials_w.ReadFrom(credentials_r)

}

func main() {

	homedir, hd_err := os.UserHomeDir()
	if hd_err != nil {
		log.Fatal("Error looking up homedir: %s", hd_err)
	}

	// aws_dir := filepath.Join(homedir, "/.aws/credentials")

	// aws_credentials, err := ini.Load(aws_dir)
	// if err != nil {
	// 	log.Fatal("AWS Credentials file nt loaded")
	// }
	backupAwsConfig(homedir)
}
