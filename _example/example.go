package main

import . "Api_hh"

func main() {
	hh := HH{}

	hh.View(hh.GetData("Ryazan", "it"))
	hh.View(hh.GetData("Moscow", "cook"))
}
