package entity

type TblAmi struct {
	Id               int
	Name             string
	ImageId          string
	MinecraftEdition string
	MinecraftVersion string
	ModLoader        string
	ModLoaderVersion string
	MinimumRam       int
	MinimumVcpu      int
}
