package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API("goa simple sample", func() {
	// APIのタイトル
	Title("tikasan/goa-simple-sample")
	// APIの説明
	Description("goaのサンプルです")
	// 作成者へのコンタクト情報
	Contact(func() {
		Name("pei")
		Email("satak47cpc@gmail.com")
		URL("./issues")
	})
	// APIのライセンス
	License(func() {
		Name("MIT")
		URL("./LICENSE")
	})
	// APIのドキュメント
	Docs(func() {
		Description("wiki")
		URL("./wiki")
	})
	// ホストの設定
	// Host("localhost:8080")
	// // 対応しているプロトコル定義、httpかhttpsまたはその両方
	// Scheme("http", "https")
	// // 全てのエンドポイントのベースパス
	// // /usersというエンドポイントがあったら、/api/v1/usersとなる
	// BasePath("/api/v1")

	Server("app", func() {
		Host("localhost", func() {
			URI("http://localhost:8080/api/v3")
		})
	})

	cors.Origin("http://localhost:8080/swagger", func() {
		//クライアントに公開された1つ以上のヘッダー
		cors.Headers("X-Time")
		// 1つまたは複数の許可されたHTTPメソッド
		cors.Methods("GET", "POST", "PUT", "DELETE")
		//プリフライト要求応答をキャッシュする時間
		cors.MaxAge(600)
		// Access-Control-Allow-Credentialsヘッダーを設定する
		cors.Credentials()
	})
})
