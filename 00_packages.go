package main

import (
	"log"
	"os/exec"
	"os/user"
	"time"
)

func main() {

	err := checkUser("openshift")
	if err != nil {
		log.Println("User not found, firing routine...")
		func() {
			userCreation := exec.Command("sudo", "useradd", "-m", "openshift")
			err := userCreation.Run()
			if err != nil {
				panic(err)
			}
		}()
	}

	basic := []string{"nginx", "httpd", "buildah", "podman"}
	start := time.Now()
	for _, v := range basic {
		instPkg(v)
	}

	duration := time.Since(start)

	log.Printf("Install complete, ET: %v ms", duration.Milliseconds())

}

func instPkg(pkgName string) error {

	inst := exec.Command("sudo", "yum", "install", "-y", pkgName)
	log.Printf("Installing package: %v", pkgName)
	err := inst.Run()

	if err != nil {
		log.Fatalf("Unable to install package %v , exiting...", pkgName)
	}
	return err
}

func checkUser(u string) error {
	_, err := user.Lookup(u)
	return err
}
