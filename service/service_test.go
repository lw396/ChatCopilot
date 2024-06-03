package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/lw396/WeComCopilot/internal/repository/sqlite"
)

func TestUnmarshalImage(t *testing.T) {
	service := New()
	data := &sqlite.MessageContent{
		MsgContent: `wxid_dl7i57grkf5q12:
		<msg>
			<img aeskey="38d1392f289a48863d433293db41b924" encryver="1"
				cdnthumbaeskey="38d1392f289a48863d433293db41b924:"
				cdnthumburl="3057020100044b30490201000204a1b9e69d02032f5853020431743070020464048430042466393534313230342d353931302d346463372d383536622d633466303735396330396134020401292a010201000405004c543f00"
				cdnthumblength="7042" cdnthumbheight="90" cdnthumbwidth="120" cdnmidheight="0"
				cdnmidwidth="0" cdnhdheight="0" cdnhdwidth="0"
				cdnmidimgurl="3057020100044b30490201000204a1b9e69d02032f5853020431743070020464048430042466393534313230342d353931302d346463372d383536622d633466303735396330396134020401292a010201000405004c543f00"
				length="27227"
				cdnbigimgurl="3057020100044b30490201000204a1b9e69d02032f5853020431743070020464048430042466393534313230342d353931302d346463372d383536622d633466303735396330396134020401292a010201000405004c543f00"
				hdlength="4898674" md5="4d38359fa55446c544a8a9e34d167370" hevc_mid_size="27227" />
			<platform_signature></platform_signature>
			<imgdatahash></imgdatahash>
		</msg>`,
		MesDes: true,
	}
	result, err := service.HandleImage(context.Background(), data, true)
	if err != nil {
		t.Error(err)
	}

	t.Log(result)
}

func TestUnmarshalEmoji(t *testing.T) {
	service := New()
	data := &sqlite.MessageContent{
		MsgContent: `wxid_t99xk0w3nbe122:
		<msg>
			<emoji fromusername="wxid_t99xk0w3nbe122" tousername="22820114318@chatroom" type="2" idbuffer="media:0_0"
				md5="5e784554e79f8b94ad3a81465f397dff" len="4750361" productid="" androidmd5="5e784554e79f8b94ad3a81465f397dff"
				androidlen="4750361" s60v3md5="5e784554e79f8b94ad3a81465f397dff" s60v3len="4750361"
				s60v5md5="5e784554e79f8b94ad3a81465f397dff" s60v5len="4750361"
				cdnurl="http://wxapp.tc.qq.com/262/20304/stodownload?m=5e784554e79f8b94ad3a81465f397dff&amp;filekey=30350201010421301f020201060402535a04105e784554e79f8b94ad3a81465f397dff0203487c19040d00000004627466730000000132&amp;hy=SZ&amp;storeid=2630fc28900014929000000000000010600004f50535a20c2788096386f424&amp;bizid=1023"
				designerid="" thumburl=""
				encrypturl="http://wxapp.tc.qq.com/262/20304/stodownload?m=481c3a9a7fcf52ac229364afc1c66557&amp;filekey=30350201010421301f02020106040253480410481c3a9a7fcf52ac229364afc1c665570203487c20040d00000004627466730000000132&amp;hy=SH&amp;storeid=2630fc28a0000183b000000000000010600004f50534819165b40b65442a8a&amp;bizid=1023"
				aeskey="495459a67b77b63b38c2b686fd809f2b"
				externurl="http://wxapp.tc.qq.com/262/20304/stodownload?m=18094c89e122ed0fa46c9058f0841c17&amp;filekey=30350201010421301f0202010604025348041018094c89e122ed0fa46c9058f0841c170203032700040d00000004627466730000000132&amp;hy=SH&amp;storeid=2630fc28a00079f21000000000000010600004f5053480976fb40b65505be0&amp;bizid=1023"
				externmd5="1e66b7ea16a12277bc6ee350c428e2a0" width="290" height="290" tpurl="" tpauthkey="" attachedtext=""
				attachedtextcolor="" lensid="" emojiattr="" linkid="" desc=""></emoji>
			<gameext type="0" content="0"></gameext>
		</msg>`,
		MesDes: false,
	}
	result, err := service.HandleSticker(context.Background(), data, true)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)
}

func TestUnmarshalVideo(t *testing.T) {
	service := New()
	data := &sqlite.MessageContent{
		MsgContent: `wxid_t99xk0w3nbe122:
		<?xml version="1.0"?>
		<msg>
			<videomsg aeskey="71d50c579fbaebc25b267de3ff583966"
				cdnvideourl="3057020100044b30490201000204dc45dde602032f591902040da90c790204665d2978042439386366333530302d356639312d343863382d616164322d6534313734306331663564640204052400040201000405004c511d00"
				cdnthumbaeskey="71d50c579fbaebc25b267de3ff583966"
				cdnthumburl="3057020100044b30490201000204dc45dde602032f591902040da90c790204665d2978042439386366333530302d356639312d343863382d616164322d6534313734306331663564640204052400040201000405004c511d00"
				length="817956" playlength="4" cdnthumblength="6746" cdnthumbwidth="224"
				cdnthumbheight="298" fromusername="wxid_t99xk0w3nbe122"
				md5="443d7917ad367b35c289669bf4f0e789" newmd5="5a51c9abc61c8c606a5e8e580182fbea"
				isplaceholder="0" rawmd5="47130090203fce25bf8df70ea98343a8" rawlength="9926637"
				cdnrawvideourl="3057020100044b30490201000204dc45dde602032f591902040da90c790204665d2977042432393065633063352d373663352d343338652d613231312d346561326230393366343234020405a400040201000405004c4c6d00"
				cdnrawvideoaeskey="19931fd31753220a853409ac0a7bacc6" overwritenewmsgid="0"
				originsourcemd5="38ee59e8d1b22eddaaef4ebf54c9a3cc" isad="0" />
		</msg>`,
		MesDes: true,
	}
	result, err := service.HandleVideo(context.Background(), data, true)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)
}
