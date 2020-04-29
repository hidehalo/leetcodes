package main

import "testing"

func Test_toGoatLatin(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// {
		// 	"case1",
		// 	args{"I speak Goat Latin"},
		// 	"Imaa peaksmaaa oatGmaaaa atinLmaaaaa",
		// },
		// {
		// 	"case2",
		// 	args{"The quick brown fox jumped over the lazy dog"},
		// 	"heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa",
		// },
		// {
		// 	"case3",
		// 	args{"HZ sg L"},
		// 	"ZHmaa gsmaaa Lmaaaa",
		// },
		{
			"case4",
			args{"YS I tH FQcP eWnE jHM aN xGI kGn Slyo Usgg XO AxP vvWK kmH S WJ olmh uor mvE Fngj BqF FNvl wOW CnD tq i IBUe iD H swN s PsfA hs pHd YYY"},
			"SYmaa Imaaa Htmaaaa QcPFmaaaaa eWnEmaaaaaa HMjmaaaaaaa aNmaaaaaaaa GIxmaaaaaaaaa Gnkmaaaaaaaaaa lyoSmaaaaaaaaaaa Usggmaaaaaaaaaaaa OXmaaaaaaaaaaaaa AxPmaaaaaaaaaaaaaa vWKvmaaaaaaaaaaaaaaa mHkmaaaaaaaaaaaaaaaa Smaaaaaaaaaaaaaaaaa JWmaaaaaaaaaaaaaaaaaa olmhmaaaaaaaaaaaaaaaaaaa uormaaaaaaaaaaaaaaaaaaaa vEmmaaaaaaaaaaaaaaaaaaaaa ngjFmaaaaaaaaaaaaaaaaaaaaaa qFBmaaaaaaaaaaaaaaaaaaaaaaa NvlFmaaaaaaaaaaaaaaaaaaaaaaaa OWwmaaaaaaaaaaaaaaaaaaaaaaaaa nDCmaaaaaaaaaaaaaaaaaaaaaaaaaa qtmaaaaaaaaaaaaaaaaaaaaaaaaaaa imaaaaaaaaaaaaaaaaaaaaaaaaaaaa IBUemaaaaaaaaaaaaaaaaaaaaaaaaaaaaa iDmaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa Hmaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa wNsmaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa smaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa sfAPmaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa shmaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa Hdpmaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa YYYmaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toGoatLatin(tt.args.S); got != tt.want {
				t.Errorf("toGoatLatin() = %v, want %v", got, tt.want)
			}
		})
	}
}
