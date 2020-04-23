# システム間認証

## システムトークンの生成

### 準備

- git で使うプロトコルを変更する  
  `$ git config --global url."git@bitbucket.org:".insteadOf "https://bitbucket.org/"`
- キー生成コマンドをインストール  
  `$ go get -u bitbucket.org/linksportsinc/go-tools`  
  `$ go install go install bitbucket.org/linksportsinc/go-tools/system-auth/cmd/teamub-systemkey-gen`

### 生成

- `$ CLIENT_SYSTEM_NAME=xxx SECRET=xxx $GOPATH/bin/teamub-systemkey-gen`

## システムトークンの検証/Claims の取得

```
import systemAuth "bitbucket.org/linksportsinc/go-tools/system-auth"

...

func veirfy(toekn string) bool {
	secret := os.Getenv("SECRET")
	auth := systemAuth.NewApp(secret)

	claims, err := auth.VerifyToken(token)
	return err == nil
}
...

```
