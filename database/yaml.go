package database

type configYAML struct {
	Database struct {
		Driver string `yaml:"driver"`
		Port   string `yaml:"port"`
		User   string `yaml:"user"`
		Pass   string `yaml:"pass"`
		Name   string `yaml:"name"`
		Host   string `yaml:"host"`
	}
}

type settings struct {
	config configYAML
}

func (s *settings) DriverDB() string {
	return s.config.Database.Driver
}

func (s *settings) PortDB() string {
	return s.config.Database.Port
}

func (s *settings) UserDB() string {
	return s.config.Database.User
}

func (s *settings) PassDB() string {
	return s.config.Database.Pass
}

func (s *settings) NameDB() string {
	return s.config.Database.Name
}

func (s *settings) HostDB() string {
	return s.config.Database.Host
}
