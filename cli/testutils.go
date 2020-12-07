package cli

func validImpactOptions() ImpactOptions {
	return ImpactOptions{
		RootDir:     "RootDir",
		Pattern:     "Pattern",
		Files:       []string{"File_1", "File_2", "File_3"},
		Credentials: "user:pwd",
	}
}
