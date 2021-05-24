package main

func selectDifficulty() string {
	for {
		cw.clear()
		cw.print("Welome to Hydra Slayer CLI!\n")
		cw.println("Select your difficulty: ")
		cw.println("  1: Easy")
		cw.println("  2: Medium")
		cw.println("  3: Hard")
		cw.println("  4: Chaotic (totally random dungeon)")
		cw.print("> ")
		cw.flush()
		input := cw.read()
		switch input {
		case "1": return "easy"
		case "2": return "medium"
		case "3": return "hard"
		case "4": return "chaotic"
		case "exit":
			abortGame = true
			return "easy"
		}
	}
}
