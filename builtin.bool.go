package dot

type Bool bool

func (b Bool) String() string {
	if b {
		return "true"
	}
	return "false"
}
