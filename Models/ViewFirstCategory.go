package Models

type ViewFirstCategory struct {
	Id                      int
	Name                    string
	CountRequest            int
	IcoFileName             string
	IcoFullPath             string
	ListSecondLavelCategory []*ViewSecondCategory
}
