package config

// SetupFromTOML produces a default configuration
func SetupFromTOML(configFilePath string) error {
	if err := setupDominionConfigFromTOML(configFilePath); err != nil {
		return err
	}
	if err := setupServicesConfigFromTOML(configFilePath); err != nil {
		return err
	}
	return nil
}
