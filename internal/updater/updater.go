package updater

type Updater interface {
	UpdateAndUpgrade() error
	ServiceFileContent() string
}
