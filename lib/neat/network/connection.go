package network

type Connection struct {
	In         int
	Out        int
	Weight     float64
	Enabled    bool
	Innovation int
}

func NewConnection(in int, out int, weight float64, inno int) Connection {
	return Connection{in, out, weight, true, inno}
}

func (c Connection) GetKey() Key {
	return Key{c.In, c.Out}
}

func (c *Connection) Disable() {
	c.Enabled = false
}
