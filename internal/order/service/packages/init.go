package packages

func Init() map[string]Package {
	packageTypes := make(map[string]Package)
	packageTypes["box"] = &Box{}
	packageTypes["packet"] = &Packet{}
	packageTypes["wrap"] = &Wrap{}
	return packageTypes
}
