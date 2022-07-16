package factory

type musket struct {
	gun
}

func newMusket() iGun {
	return &musket{
		gun: gun{
			name:  "Musket Gun",
			power: 2,
		},
	}
}
