package completions

type AutoComplete struct {
	Name        string
	Description string
	Commands    []*Command
	Groups      []*Group
	Flags       []*Flag
}

type Command struct {
	Name        string
	Description string
	Flags       []*Flag
	URI         string
}

type Group struct {
	Name        string
	Description string
	Groups      []*Group
	Commands    []*Command
	Flags       []*Flag
	URI         string
}

type Flag struct {
	Name        string
	Description string
}
