package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addProxyPass)
	addProxyPass.Flags().IntP("port", "p", 80, "Especifica el puerto (default 80)")
}

var addProxyPass = &cobra.Command{
	Use:   "proxy <HOST>",
	Short: "Agrega un proxy_pass",
	Args: func(cmd *cobra.Command, args []string) error {
		if (len(args)) < 2 {
			return errors.New("Necesito el DOMINIO y la IP de tu app")
		}
		if len(args) > 2 {
			return errors.New("Demasiados argumentos, para la mano")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("agregando proxy")
		host := args[0]
		ip := args[1]
		port, _ := cmd.Flags().GetInt("port")
		createTemplate(host, ip, port)
	},
}

func createTemplate(host string, ip string, port int) {
	absPath, _ := filepath.Abs("./templates/proxy_pass.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	read, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	output := bytes.ReplaceAll(read, []byte("{HOST}"), []byte(host))
	output = bytes.ReplaceAll(output, []byte("{IP}"), []byte(ip))
	output = bytes.ReplaceAll(output, []byte("{PORT}"), []byte(strconv.Itoa(port)))

	if err = ioutil.WriteFile("./outputs/proxy_pass", output, 0666); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println("proxy_pass generado para " + host)
}
