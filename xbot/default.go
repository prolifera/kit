package xbot

var defaultCli = &Client{}

func Default() *Client {
	return defaultCli
}

func InitDefault(token string) error {
	var err error
	defaultCli, err = NewClientByToken(token)
	if err != nil {
		return err
	}
	return nil
}
