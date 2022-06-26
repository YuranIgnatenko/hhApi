package main

import . "hh_api"

func main() {
	hh := HH{}

	hh.View(hh.GetData("Ryazan", "it"))
	hh.View(hh.GetData("Moscow", "cook"))
}
