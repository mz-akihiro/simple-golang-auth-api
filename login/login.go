package login

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"simple-golang-auth-api/db"
	"simple-golang-auth-api/model"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	dbCnt := db.Newdb()

	defer db.CloseDB(dbCnt)

	//リクエスト用の構造体をmodelから引っ張ってきて用意
	var user model.User
	//用意した構造体にリクエストのjsonをデコード
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid JSON format")
		return
	}
	fmt.Println(user)

	//email passwordにちゃんと値が入ってるか確認
	if user.Email == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Email and password are required")
		return
	}

	//userテーブル内のemailカラムに同じ値がないか確認。あればcountが1以上になる。
	//今回はログインなので、emailが登録されていることを確認。
	var count int
	if err := dbCnt.QueryRow("SELECT COUNT(*) FROM user WHERE email = ?", user.Email).Scan(&count); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Database error")
		return
	}
	if count == 0 {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintln(w, "Email not registered")
		return
	}

	//emailの値と同じ行のpasswordカラムから値を取得
	var password string
	row := dbCnt.QueryRow("SELECT password FROM user WHERE email = ?", user.Email)
	if err := row.Scan(&password); err != nil {
		w.WriteHeader(http.StatusInternalServerError) //どのエラー使えばいい？
		fmt.Fprintln(w, "Database error")
		return
	}

	//クライアントから送られたパスワードと登録済みのパスワードが一緒かどうか確認
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "different password")
		return
	}

	//tokenにidを埋め込んでおくために取り出す。
	var idToken int
	row2 := dbCnt.QueryRow("SELECT id FROM user WHERE email = ?", user.Email)
	if err := row2.Scan(&password); err != nil {
		w.WriteHeader(http.StatusInternalServerError) //どのエラー使えばいい？
		fmt.Fprintln(w, "Database error")
		return
	}

	//トークン作成
	// パスワードが一致する場合はJWTトークンを発行
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": idToken,
		"exp":     time.Now().Add(time.Hour * 12).Unix(), // jwtトークンの有効期限(12時間)
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Failed to create token")
		return
	}

	cookie := &http.Cookie{}
	// クッキーの名前(token)を設定
	cookie.Name = "token"
	// クッキーの値としてjwtトークンを代入
	cookie.Value = tokenString
	// クッキーの有効期限(24h)を設定
	cookie.Expires = time.Now().Add(24 * time.Hour)
	// クッキーが送信されるパスを設定
	cookie.Path = "/"
	// ドメインは環境変数に設定されている変数を使用
	// RFC2109・RFC6265でドメインにipアドレスを指定することが
	// 禁止されているので注意。
	cookie.Domain = os.Getenv("API_DOMAIN")
	// クッキーのセキュア設定。trueだとHTTPS通信のみで
	// クッキーが提供されるようになる。
	// なのでここは一旦falseにする。(設定しないとデフォルトでfalseになる)
	/*cookie.Secure = true*/
	//
	// HttpOnly属性をtrueにすることで
	// クライアント側のJavaScript等がクッキーにアクセスできなくなる。
	cookie.HttpOnly = true
	// 全てのリクエストにクッキーが付与される
	cookie.SameSite = http.SameSiteNoneMode

	http.SetCookie(w, cookie)
	fmt.Fprintln(w, "Cookie sent")
}
