package main

import "github.com/nggenius/nggame/apps/tools/obj_gen/parser"

func main() {
	parser.ParseFromXml("player.xml", "object.tpl", "./parser/", "../../../gameobject/entity/player.go")
	parser.ParseFromXml("scene.xml", "object.tpl", "./parser/", "../../../gameobject/entity/scene.go")
}
