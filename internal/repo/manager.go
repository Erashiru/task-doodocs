package repo

type RepoI interface {
}

func New(storagePath string) (RepoI, error) {
	return &RepoI{}
}
