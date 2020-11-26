package routes

import "github.com/andynl/go-ongkir/controller"

func init() {
	Router.HandleFunc("/getOngkir", controller.GetOngkir)
}
