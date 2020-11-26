package cli

func validImpactOptions() impactOptions {
	return impactOptions{
		RootDir: "RootDir",
		Pattern: "Pattern",
		Files:   []string{"File_1", "File_2", "File_3"},
	}
}
