package cli

func validImpactOptions() ImpactOptions {
	return ImpactOptions{
		Files:       []string{"File_1", "File_2", "File_3"},
		RootDir:     "RootDir",
		Pattern:     "Pattern",
		Credentials: "user:pwd",
		Output:      "output.json",
		ListStates:  false,
	}
}
