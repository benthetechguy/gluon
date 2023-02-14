package command

import (
	"bytes"
	"github.com/ProtonMail/gluon/imap/parser"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParser_SubscribeCommand(t *testing.T) {
	input := toIMAPLine(`tag SUBSCRIBE INBOX`)
	s := parser.NewScanner(bytes.NewReader(input))
	p := NewParser(s)

	expected := Command{Tag: "tag", Payload: &SubscribeCommand{
		Mailbox: "INBOX",
	}}

	cmd, err := p.Parse()
	require.NoError(t, err)
	require.Equal(t, expected, cmd)
	require.Equal(t, "subscribe", p.LastParsedCommand())
	require.Equal(t, "tag", p.LastParsedTag())
}
