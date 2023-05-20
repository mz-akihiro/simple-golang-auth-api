package logout

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	cookie := &http.Cookie{}

	// 名前(token)を設定
	cookie.Name = "token"
	// jwtトークンを空文字で上書きする
	cookie.Value = ""
	// 有効期限もすぐに消えるよう設定
	cookie.Expires = time.Now()
	// パス設定
	cookie.Path = "/"
	// ドメイン設定
	cookie.Domain = os.Getenv("API_DOMAIN")
	// セキュア設定falseへ
	/*cookie.Secure = true*/
	//
	// クッキー登録と以下同文
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode

	http.SetCookie(w, cookie)
	fmt.Fprintln(w, "Delete cookie")

}
