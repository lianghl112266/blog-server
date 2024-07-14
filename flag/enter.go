package flag

import (
	"flag"
)

type Option struct {
	DB bool
}

func Parse() Option {
	db := flag.Bool("db", false, "init db")
	flag.Parse()
	return Option{
		DB: *db,
	}
}

func IsWebStop(option Option) bool {
	return option.DB
}
func SwitchOption(option Option) {
	if option.DB {
		MakeMigrations()
	}
}
