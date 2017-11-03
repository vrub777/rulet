package Models

type ViewFirstCategory struct {
	Id                      int
	Name                    string
	Order                   int
	CountRequest            int
	CountSecondLevel        int
	IcoFileName             string
	IcoFullPath             string
	ListSecondLavelCategory []*ViewSecondCategory
}
