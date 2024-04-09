package packages

const (
	BoxPackage    = "box"
	PacketPackage = "packet"
	WrapPackage   = "wrap"
)

func Init() map[string]*Package {
	return map[string]*Package{
		BoxPackage:    {Price: 20, MaxWeight: 30},
		PacketPackage: {Price: 5, MaxWeight: 10},
		WrapPackage:   {Price: 1, MaxWeight: 0},
	}
}
