package funfacts

/**
  Implementer funfacts-funksjon:
    GetFunFacts(about string) []string
      hvor about kan ha en av tre testverdier, -
        sun, luna eller terra

  Sett inn alle Funfucts i en struktur
  type FunFacts struct {
      Sun []string
      Luna []string
      Terra []string
  }
*/

type FunFacts struct {
	Sun   []string
	Luna  []string
	Terra []string
}

func GetFunFacts(about string) []string {
	var funFacts FunFacts
	funFacts.Sun = []string{
		"The sun is a star that is located at the center of the solar system.",
		"It is estimated to be about 4.6 billion years old.",
		"The sun makes up about 99.86% of the total mass of the solar system.",
		"The temperature at the sun's core is estimated to be about 15 million Kelvin.",
	}
	funFacts.Luna = []string{
		"The moon is the Earth's only natural satellite.",
		"The moon is approximately 238,855 miles away from Earth.",
		"The moon is approximately 1/6th the size of the Earth.",
		"The moon is believed to have formed around 4.5 billion years ago.",
	}
	funFacts.Terra = []string{
		"Earth is the third planet from the sun.",
		"Earth has a diameter of about 7,926 miles.",
		"Earth is the only known planet to have liquid water on its surface.",
		"Earth is estimated to be about 4.54 billion years old.",
	}

	switch about {
	case "sun":
		return funFacts.Sun
	case "luna":
		return funFacts.Luna
	case "terra":
		return funFacts.Terra
	default:
		return []string{}
	}
}
